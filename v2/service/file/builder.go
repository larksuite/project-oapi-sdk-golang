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
	"encoding/json"
	"fmt"
	"io"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

// 上传文件（表单）接口参数构造器
type UploadFileByFormReqBuilder struct {
	apiReq *core.APIReq
	body   *core.FormData
}

// 上传文件（表单）接口参数
type UploadFileByFormReq struct {
	apiReq *core.APIReq
}

// 上传文件（表单）接口响应
type UploadFileByFormResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *UploadFileByFormResult `json:"data"`
}

func NewUploadFileByFormReqBuilder() *UploadFileByFormReqBuilder {
	builder := &UploadFileByFormReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	builder.body = core.NewFormdata()
	return builder
}

func (builder *UploadFileByFormReqBuilder) ProjectKey(projectKey string) *UploadFileByFormReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *UploadFileByFormReqBuilder) ResourceType(resourceType string) *UploadFileByFormReqBuilder {
	builder.apiReq.PathParams.Set("resource_type", resourceType)
	return builder
}

func (builder *UploadFileByFormReqBuilder) File(file io.Reader) *UploadFileByFormReqBuilder {
	builder.body.AddFile("unknown-file", file)
	return builder
}

func (builder *UploadFileByFormReqBuilder) FileWithFileName(fileName string, file io.Reader) *UploadFileByFormReqBuilder {
	builder.body.AddFile(fileName, file)
	return builder
}

func (builder *UploadFileByFormReqBuilder) FileMimeType(mimeType string) *UploadFileByFormReqBuilder {
	builder.body.SetMimeType(mimeType)
	return builder
}

func (builder *UploadFileByFormReqBuilder) FieldMap(fieldMap map[string]string) *UploadFileByFormReqBuilder {
	bs, _ := json.Marshal(fieldMap)
	builder.body.AddField("field_map", string(bs))
	return builder
}

func (builder *UploadFileByFormReqBuilder) Build() *UploadFileByFormReq {
	req := &UploadFileByFormReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

// 下载文件接口参数构造器
type DownloadFileReqBuilder struct {
	apiReq *core.APIReq
}

// 下载文件接口参数
type DownloadFileReq struct {
	apiReq *core.APIReq
}

// 下载文件接口响应
type DownloadFileResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	// File 为下载响应的网络流；成功时非空，使用完毕后【必须调用 File.Close()】以释放连接。
	File     io.ReadCloser `json:"-"`
	FileName string        `json:"-"`
	MimeType string        `json:"-"`
}

func NewDownloadFileReqBuilder() *DownloadFileReqBuilder {
	builder := &DownloadFileReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	return builder
}

func (builder *DownloadFileReqBuilder) ProjectKey(projectKey string) *DownloadFileReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *DownloadFileReqBuilder) FileToken(fileToken string) *DownloadFileReqBuilder {
	builder.apiReq.PathParams.Set("file_token", fileToken)
	return builder
}

func (builder *DownloadFileReqBuilder) Build() *DownloadFileReq {
	req := &DownloadFileReq{}
	req.apiReq = builder.apiReq
	return req
}

// 分片初始化接口参数
type UploadPreProcessReqBody struct {
	Size       int64             `json:"size"`
	FileName   string            `json:"file_name"`
	FieldMap   map[string]string `json:"field_map,omitempty"`
	MimeType   string            `json:"mime_type,omitempty"`
	ResourceID string            `json:"resource_id,omitempty"`
}

// 分片初始化接口参数构造器
type UploadPreProcessReqBuilder struct {
	apiReq *core.APIReq
	body   *UploadPreProcessReqBody
}

// 分片初始化接口参数
type UploadPreProcessReq struct {
	apiReq *core.APIReq
}

// 分片初始化接口响应
type UploadPreProcessResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *UploadPreProcessResult `json:"data"`
}

func NewUploadPreProcessReqBuilder() *UploadPreProcessReqBuilder {
	builder := &UploadPreProcessReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	builder.body = &UploadPreProcessReqBody{}
	return builder
}

func (builder *UploadPreProcessReqBuilder) ProjectKey(projectKey string) *UploadPreProcessReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *UploadPreProcessReqBuilder) ResourceType(resourceType string) *UploadPreProcessReqBuilder {
	builder.apiReq.PathParams.Set("resource_type", resourceType)
	return builder
}

func (builder *UploadPreProcessReqBuilder) Size(size int64) *UploadPreProcessReqBuilder {
	builder.body.Size = size
	return builder
}

func (builder *UploadPreProcessReqBuilder) FileName(fileName string) *UploadPreProcessReqBuilder {
	builder.body.FileName = fileName
	return builder
}

func (builder *UploadPreProcessReqBuilder) FieldMap(fieldMap map[string]string) *UploadPreProcessReqBuilder {
	builder.body.FieldMap = fieldMap
	return builder
}

func (builder *UploadPreProcessReqBuilder) MimeType(mimeType string) *UploadPreProcessReqBuilder {
	builder.body.MimeType = mimeType
	return builder
}

func (builder *UploadPreProcessReqBuilder) ResourceID(resourceID string) *UploadPreProcessReqBuilder {
	builder.body.ResourceID = resourceID
	return builder
}

func (builder *UploadPreProcessReqBuilder) Build() *UploadPreProcessReq {
	req := &UploadPreProcessReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

// 上传分片接口参数构造器
type UploadPartReqBuilder struct {
	apiReq  *core.APIReq
	body    io.Reader
	bodyLen int64
}

// 上传分片接口参数
type UploadPartReq struct {
	apiReq  *core.APIReq
	body    io.Reader
	bodyLen int64
}

// 上传分片接口响应
type UploadPartResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

func NewUploadPartReqBuilder() *UploadPartReqBuilder {
	builder := &UploadPartReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *UploadPartReqBuilder) ProjectKey(projectKey string) *UploadPartReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *UploadPartReqBuilder) ResourceID(resourceID string) *UploadPartReqBuilder {
	builder.apiReq.PathParams.Set("resource_id", resourceID)
	return builder
}

// 分片序号，从 0 开始
func (builder *UploadPartReqBuilder) PartNumber(partNumber int64) *UploadPartReqBuilder {
	builder.apiReq.PathParams.Set("part_number", fmt.Sprint(partNumber))
	return builder
}

// 分片二进制的 MD5，必填
func (builder *UploadPartReqBuilder) Md5(md5 string) *UploadPartReqBuilder {
	builder.apiReq.QueryParams.Set("md5", md5)
	return builder
}

// 分片二进制内容（流式）：file 为该分片的可读流，size 为该分片字节数
func (builder *UploadPartReqBuilder) File(file io.Reader, size int64) *UploadPartReqBuilder {
	builder.body = file
	builder.bodyLen = size
	return builder
}

func (builder *UploadPartReqBuilder) Build() *UploadPartReq {
	req := &UploadPartReq{}
	req.apiReq = builder.apiReq
	req.body = builder.body
	req.bodyLen = builder.bodyLen
	return req
}

// 完成分片上传接口参数构造器
type UploadPartFinishReqBuilder struct {
	apiReq *core.APIReq
}

// 完成分片上传接口参数
type UploadPartFinishReq struct {
	apiReq *core.APIReq
}

// 完成分片上传接口响应
type UploadPartFinishResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *UploadPartFinishResult `json:"data"`
}

func NewUploadPartFinishReqBuilder() *UploadPartFinishReqBuilder {
	builder := &UploadPartFinishReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	return builder
}

func (builder *UploadPartFinishReqBuilder) ProjectKey(projectKey string) *UploadPartFinishReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *UploadPartFinishReqBuilder) ResourceID(resourceID string) *UploadPartFinishReqBuilder {
	builder.apiReq.PathParams.Set("resource_id", resourceID)
	return builder
}

func (builder *UploadPartFinishReqBuilder) Build() *UploadPartFinishReq {
	req := &UploadPartFinishReq{}
	req.apiReq = builder.apiReq
	return req
}

// 终止分片上传接口参数构造器
type UploadPartAbortReqBuilder struct {
	apiReq *core.APIReq
}

// 终止分片上传接口参数
type UploadPartAbortReq struct {
	apiReq *core.APIReq
}

// 终止分片上传接口响应
type UploadPartAbortResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

func NewUploadPartAbortReqBuilder() *UploadPartAbortReqBuilder {
	builder := &UploadPartAbortReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	return builder
}

func (builder *UploadPartAbortReqBuilder) ProjectKey(projectKey string) *UploadPartAbortReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *UploadPartAbortReqBuilder) ResourceID(resourceID string) *UploadPartAbortReqBuilder {
	builder.apiReq.PathParams.Set("resource_id", resourceID)
	return builder
}

func (builder *UploadPartAbortReqBuilder) Build() *UploadPartAbortReq {
	req := &UploadPartAbortReq{}
	req.apiReq = builder.apiReq
	return req
}

// 分片上传（聚合）接口参数构造器
type MultipartUploadReqBuilder struct {
	projectKey   string
	resourceType string
	fileName     string
	mimeType     string
	fieldMap     map[string]string
	resourceID   string
	content      io.ReaderAt // 支持并发 ReadAt 的随机访问源（*os.File / bytes.NewReader 均可）
	size         int64
	concurrency  int
}

// 分片上传（聚合）接口参数
type MultipartUploadReq struct {
	projectKey   string
	resourceType string
	fileName     string
	mimeType     string
	fieldMap     map[string]string
	resourceID   string
	content      io.ReaderAt
	size         int64
	concurrency  int
}

func NewMultipartUploadReqBuilder() *MultipartUploadReqBuilder {
	return &MultipartUploadReqBuilder{}
}

func (builder *MultipartUploadReqBuilder) ProjectKey(projectKey string) *MultipartUploadReqBuilder {
	builder.projectKey = projectKey
	return builder
}

func (builder *MultipartUploadReqBuilder) ResourceType(resourceType string) *MultipartUploadReqBuilder {
	builder.resourceType = resourceType
	return builder
}

func (builder *MultipartUploadReqBuilder) FileName(fileName string) *MultipartUploadReqBuilder {
	builder.fileName = fileName
	return builder
}

func (builder *MultipartUploadReqBuilder) MimeType(mimeType string) *MultipartUploadReqBuilder {
	builder.mimeType = mimeType
	return builder
}

func (builder *MultipartUploadReqBuilder) FieldMap(fieldMap map[string]string) *MultipartUploadReqBuilder {
	builder.fieldMap = fieldMap
	return builder
}

// 断点续传可选：与初始化传入一致的 resource_id
func (builder *MultipartUploadReqBuilder) ResourceID(resourceID string) *MultipartUploadReqBuilder {
	builder.resourceID = resourceID
	return builder
}

// 完整文件内容（流式）：file 为支持随机读的源（*os.File / bytes.NewReader 均可），size 为文件总字节数
func (builder *MultipartUploadReqBuilder) File(file io.ReaderAt, size int64) *MultipartUploadReqBuilder {
	builder.content = file
	builder.size = size
	return builder
}

// 并发上传的最大并发数，<=0 时用默认值
func (builder *MultipartUploadReqBuilder) Concurrency(concurrency int) *MultipartUploadReqBuilder {
	builder.concurrency = concurrency
	return builder
}

func (builder *MultipartUploadReqBuilder) Build() *MultipartUploadReq {
	req := &MultipartUploadReq{
		projectKey:   builder.projectKey,
		resourceType: builder.resourceType,
		fileName:     builder.fileName,
		mimeType:     builder.mimeType,
		fieldMap:     builder.fieldMap,
		resourceID:   builder.resourceID,
		content:      builder.content,
		size:         builder.size,
		concurrency:  builder.concurrency,
	}
	return req
}
