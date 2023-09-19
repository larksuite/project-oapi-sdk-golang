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
	"context"
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const (
	// 添加附件接口path
	ApiPath_UploadAttachement = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/file/upload"
)

func NewService(config *core.Config) *AttachmentService {
	a := &AttachmentService{config: config}
	return a
}

type AttachmentService struct {
	config *core.Config
}

// 添加附件
//
// - 官网API文档链接:https://bytedance.feishu.cn/docs/doccntEfMPoh8Qv3hDCshfRLEuY#tkxHhZ
func (a *AttachmentService) UploadAttachment(ctx context.Context, req *UploadAttachmentReq, options ...core.RequestOptionFunc) (*UploadAttachmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPath_UploadAttachement
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadAttachment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadAttachmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UploadAttachment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
