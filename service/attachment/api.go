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

package attachment

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const (
	// 添加附件接口path
	ApiPathUploadAttachment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/file/upload"

	// 下载附件接口path
	ApiPathDownloadAttachment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/file/download"

	// 附件上传接口path
	ApiPathSpecialUploadAttachment = "/open_api/:project_key/file/upload"
)

func NewService(config *core.Config) *AttachmentService {
	a := &AttachmentService{config: config}
	return a
}

type AttachmentService struct {
	config *core.Config
}

// 添加附件
func (a *AttachmentService) UploadAttachment(ctx context.Context, req *UploadAttachmentReq, options ...core.RequestOptionFunc) (*UploadAttachmentResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUploadAttachment
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadAttachment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UploadAttachmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadAttachment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 下载附件
func (a *AttachmentService) DownloadAttachment(ctx context.Context, req *DownloadAttachmentReq, options ...core.RequestOptionFunc) (*DownloadAttachmentResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathDownloadAttachment
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DownloadAttachment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DownloadAttachmentResp{APIResp: apiResp}
	// 如果是下载，则设置响应结果
	if apiResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(apiResp.RawBody)
		resp.FileName = core.FileNameByHeader(apiResp.Header)
		return resp, err
	}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DownloadAttachment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 附件上传
func (a *AttachmentService) SpecialUploadAttachment(ctx context.Context, req *SpecialUploadAttachmentReq, options ...core.RequestOptionFunc) (*SpecialUploadAttachmentResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathSpecialUploadAttachment
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SpecialUploadAttachment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SpecialUploadAttachmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SpecialUploadAttachment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
