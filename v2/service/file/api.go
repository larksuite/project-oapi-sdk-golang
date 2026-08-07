/*
 * Copyright (c) 2023 Lark Technologies Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package file

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

const (
	// 上传文件（表单）接口path
	ApiPathUploadFileByForm = "/open_api/:project_key/file/stream/resource/:resource_type/upload_form"

	// 下载文件接口path
	ApiPathDownloadFile = "/open_api/:project_key/file/stream/download/:file_token"

	// 分片初始化接口path
	ApiPathUploadPreProcess = "/open_api/:project_key/file/resource/:resource_type/upload/preprocess"

	// 上传分片接口path
	ApiPathUploadPart = "/open_api/:project_key/file/stream/upload/multipart/:resource_id/:part_number"

	// 完成分片上传接口path
	ApiPathUploadPartFinish = "/open_api/:project_key/file/upload/multipart/finish/:resource_id"

	// 终止分片上传接口path
	ApiPathUploadPartAbort = "/open_api/:project_key/file/upload/multipart/abort/:resource_id"
)

// 分片上传（聚合）默认并发数
const defaultMultipartConcurrency = 5

// 表单上传的文件大小上限（100MB）；超过则走分片上传
const maxFormUploadSize = 100 * 1024 * 1024

func NewService(config *core.Config) *FileService {
	a := &FileService{config: config}
	return a
}

type FileService struct {
	config *core.Config
}

// 上传文件（表单）
func (a *FileService) UploadFileByForm(ctx context.Context, req *UploadFileByFormReq, options ...core.RequestOptionFunc) (*UploadFileByFormResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUploadFileByForm
	apiReq.HttpMethod = http.MethodPost
	// 流式编码 multipart body（io.Pipe），Content-Length 预算已知，走自建链路避免全量缓冲
	contentType, body, bodyLen := req.form.streamBody()
	apiResp, err := a.doStreamRequest(ctx, apiReq.HttpMethod, apiReq.ApiPath,
		apiReq.PathParams, apiReq.QueryParams, contentType, body, bodyLen, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadFileByForm] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UploadFileByFormResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalFileBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadFileByForm] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 下载文件。注意：成功时 DownloadFileResp.File 为网络流，调用方【必须 Close】。
func (a *FileService) DownloadFile(ctx context.Context, req *DownloadFileReq, options ...core.RequestOptionFunc) (*DownloadFileResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathDownloadFile
	apiReq.HttpMethod = http.MethodGet
	httpResp, err := a.doStreamDownload(ctx, apiReq.HttpMethod, apiReq.ApiPath, apiReq.PathParams, apiReq.QueryParams, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DownloadFile] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 成功：直接以响应流作为文件内容，不做全量缓冲（由调用方 Close）
	if httpResp.StatusCode == http.StatusOK {
		resp := &DownloadFileResp{APIResp: &core.APIResp{StatusCode: httpResp.StatusCode, Header: httpResp.Header}}
		resp.File = httpResp.Body
		resp.FileName = core.FileNameByHeader(httpResp.Header)
		resp.MimeType = core.MimeTypeByHeader(httpResp.Header)
		return resp, nil
	}
	// 失败：错误体为小 JSON，读入后解析为 CodeError（并 Close 释放连接）
	defer httpResp.Body.Close()
	raw, err := io.ReadAll(httpResp.Body)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DownloadFile] fail to read error body, error: %v", err.Error()))
		return nil, err
	}
	resp := &DownloadFileResp{APIResp: &core.APIResp{StatusCode: httpResp.StatusCode, Header: httpResp.Header, RawBody: raw}}
	err = resp.APIResp.JSONUnmarshalFileBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DownloadFile] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// 分片初始化
func (a *FileService) UploadPreProcess(ctx context.Context, req *UploadPreProcessReq, options ...core.RequestOptionFunc) (*UploadPreProcessResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUploadPreProcess
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPreProcess] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UploadPreProcessResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalFileBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPreProcess] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 上传分片
func (a *FileService) UploadPart(ctx context.Context, req *UploadPartReq, options ...core.RequestOptionFunc) (*UploadPartResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUploadPart
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := a.doStreamRequest(ctx, apiReq.HttpMethod, apiReq.ApiPath,
		apiReq.PathParams, apiReq.QueryParams, octetStream, req.body, req.bodyLen, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPart] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UploadPartResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalFileBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPart] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 完成分片上传
func (a *FileService) UploadPartFinish(ctx context.Context, req *UploadPartFinishReq, options ...core.RequestOptionFunc) (*UploadPartFinishResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUploadPartFinish
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPartFinish] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UploadPartFinishResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalFileBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPartFinish] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 终止分片上传
func (a *FileService) UploadPartAbort(ctx context.Context, req *UploadPartAbortReq, options ...core.RequestOptionFunc) (*UploadPartAbortResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUploadPartAbort
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPartAbort] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UploadPartAbortResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalFileBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadPartAbort] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 分片上传（聚合）：数据准备 -> 并发上传 -> 完成；遇错中止并返回统一报错。
// 成功时返回「完成分片上传」的响应体。
func (a *FileService) MultipartUpload(ctx context.Context, req *MultipartUploadReq, options ...core.RequestOptionFunc) (*UploadPartFinishResp, error) {
	// 文件大小 <= 100MB 时走表单上传，> 100MB 时走分片上传
	if req.size <= maxFormUploadSize {
		formBuilder := NewUploadFileByFormReqBuilder().
			ProjectKey(req.projectKey).
			ResourceType(req.resourceType)
		if req.fileName != "" {
			formBuilder.FileWithFileName(req.fileName, io.NewSectionReader(req.content, 0, req.size)).FileSize(req.size)
		} else {
			formBuilder.File(io.NewSectionReader(req.content, 0, req.size)).FileSize(req.size)
		}
		if req.mimeType != "" {
			formBuilder.FileMimeType(req.mimeType)
		}
		if req.fieldMap != nil {
			formBuilder.FieldMap(req.fieldMap)
		}
		formResp, err := a.UploadFileByForm(ctx, formBuilder.Build(), options...)
		if err != nil {
			a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to upload by form, error: %v", err.Error()))
			return nil, err
		}
		if !formResp.Success() {
			err = fmt.Errorf("[MultipartUpload] upload by form error:[code=%v, msg=%v]", formResp.CodeError.ErrCode, formResp.CodeError.ErrMsg)
			a.config.Logger.Error(ctx, err.Error())
			return nil, err
		}
		// 转换：桥接为 *UploadPartFinishResp，保证返回类型与本方法声明一致
		formFinishResp := &UploadPartFinishResp{
			APIResp:   formResp.APIResp,
			CodeError: formResp.CodeError,
		}
		if formResp.Data != nil {
			formFinishResp.Data = &UploadPartFinishResult{FileToken: formResp.Data.FileToken}
		}
		return formFinishResp, nil
	}

	// 1. 数据准备：分片初始化
	preBuilder := NewUploadPreProcessReqBuilder().
		ProjectKey(req.projectKey).
		ResourceType(req.resourceType).
		Size(req.size).
		FileName(req.fileName).
		MimeType(req.mimeType)
	if req.fieldMap != nil {
		preBuilder.FieldMap(req.fieldMap)
	}
	if req.resourceID != "" {
		preBuilder.ResourceID(req.resourceID)
	}
	preResp, err := a.UploadPreProcess(ctx, preBuilder.Build(), options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to preprocess, error: %v", err.Error()))
		return nil, err
	}
	if !preResp.Success() {
		err = fmt.Errorf("[MultipartUpload] preprocess error:[code=%v, msg=%v]", preResp.CodeError.ErrCode, preResp.CodeError.ErrMsg)
		a.config.Logger.Error(ctx, err.Error())
		return nil, err
	}
	if preResp.Data == nil {
		err = fmt.Errorf("[MultipartUpload] empty preprocess data")
		a.config.Logger.Error(ctx, err.Error())
		return nil, err
	}

	resourceID := preResp.Data.ResourceID
	var partSize int64
	var partCount int64
	var need []int64
	if preResp.Data.Multipart != nil {
		partSize = preResp.Data.Multipart.PartSize
		partCount = preResp.Data.Multipart.PartCount
		need = preResp.Data.Multipart.Need
	}

	concurrency := req.concurrency
	if concurrency <= 0 {
		concurrency = defaultMultipartConcurrency
	}

	// 2. 并发上传分片，遇错中止
	if err := a.uploadParts(ctx, req.projectKey, resourceID, partSize, partCount, need, req.content, req.size, concurrency, options...); err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to upload parts, error: %v", err.Error()))
		a.abortMultipart(ctx, req.projectKey, resourceID, options...)
		return nil, err
	}

	// 3. 完成响应：完成分片上传
	finishResp, err := a.UploadPartFinish(ctx,
		NewUploadPartFinishReqBuilder().
			ProjectKey(req.projectKey).
			ResourceID(resourceID).
			Build(),
		options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to finish upload, error: %v", err.Error()))
		a.abortMultipart(ctx, req.projectKey, resourceID, options...)
		return nil, err
	}
	if !finishResp.Success() {
		err = fmt.Errorf("[MultipartUpload] finish error:[code=%v, msg=%v]", finishResp.CodeError.ErrCode, finishResp.CodeError.ErrMsg)
		a.config.Logger.Error(ctx, err.Error())
		a.abortMultipart(ctx, req.projectKey, resourceID, options...)
		return nil, err
	}
	return finishResp, nil
}

// uploadParts 受限并发上传所有待上传分片，任一分片失败即取消其余并返回首个错误。
// file 为支持并发随机读的源，total 为文件总字节数；每片通过 io.SectionReader 流式读取，避免整片入内存。
func (a *FileService) uploadParts(ctx context.Context, projectKey, resourceID string, partSize int64, partCount int64, need []int64, file io.ReaderAt, total int64, concurrency int, options ...core.RequestOptionFunc) error {
	if len(need) == 0 {
		return nil
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error
	setErr := func(e error) {
		once.Do(func() {
			firstErr = e
			cancel()
		})
	}

	for _, partNumber := range need {
		if ctx.Err() != nil { // 已出错，停止派发
			break
		}
		sem <- struct{}{}
		wg.Add(1)
		go func(pn int64) {
			defer wg.Done()
			defer func() { <-sem }()
			if ctx.Err() != nil {
				return
			}
			// 切分该分片区间：最后一个分片取剩余全部内容（可超过 partSize），其余分片为 partSize
			start := pn * partSize
			if start > total {
				start = total
			}
			var end int64
			if pn == partCount-1 {
				end = total
			} else {
				end = start + partSize
				if end > total {
					end = total
				}
			}
			length := end - start
			// 第一遍：流式计算 md5（不缓冲整片）
			h := md5.New()
			if _, err := io.Copy(h, io.NewSectionReader(file, start, length)); err != nil {
				a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to md5 part %v, error: %v", pn, err.Error()))
				setErr(err)
				return
			}
			// 第二遍：以新的 SectionReader 作为请求体流式上传
			resp, err := a.UploadPart(ctx,
				NewUploadPartReqBuilder().
					ProjectKey(projectKey).
					ResourceID(resourceID).
					PartNumber(pn).
					Md5(hex.EncodeToString(h.Sum(nil))).
					File(io.NewSectionReader(file, start, length)).
					FileSize(length).
					Build(),
				options...)
			if err != nil {
				a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to upload part %v, error: %v", pn, err.Error()))
				setErr(err)
				return
			}
			if !resp.Success() {
				err = fmt.Errorf("[MultipartUpload] upload part %v error:[code=%v, msg=%v]", pn, resp.CodeError.ErrCode, resp.CodeError.ErrMsg)
				a.config.Logger.Error(ctx, err.Error())
				setErr(err)
				return
			}
		}(partNumber)
	}
	wg.Wait()
	return firstErr
}

// abortMultipart 中止分片上传，best-effort，失败仅记日志。
func (a *FileService) abortMultipart(ctx context.Context, projectKey, resourceID string, options ...core.RequestOptionFunc) {
	resp, err := a.UploadPartAbort(ctx,
		NewUploadPartAbortReqBuilder().
			ProjectKey(projectKey).
			ResourceID(resourceID).
			Build(),
		options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to abort upload, resource_id: %v, error: %v", resourceID, err.Error()))
		return
	}
	if !resp.Success() {
		a.config.Logger.Error(ctx, fmt.Sprintf("[MultipartUpload] fail to abort upload, resource_id: %v, error:[code=%v, msg=%v]", resourceID, resp.CodeError.ErrCode, resp.CodeError.ErrMsg))
	}
}
