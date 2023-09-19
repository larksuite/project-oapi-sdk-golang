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

package workitem_conf

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPath_CreateTemplateDetail = "/open_api/template/v2/create_template"

const APIPath_DeleteTemplateDetail = "/open_api/template/v2/delete_template/:project_key/:template_id"

const APIPath_QueryTemplateDetail = "/open_api/:project_key/template_detail/:template_id"

const APIPath_QueryWorkItemTemplates = "/open_api/:project_key/template_list/:work_item_type_key"

const APIPath_UpdateTemplateDetail = "/open_api/template/v2/update_template"

func NewService(config *core.Config) *WorkItemConfService {
	a := &WorkItemConfService{config: config}
	return a
}

type WorkItemConfService struct {
	config *core.Config
}

/*
 *   新增流程类型配置
 */
func (a *WorkItemConfService) CreateTemplateDetail(ctx context.Context, req *CreateTemplateDetailReq, options ...core.RequestOptionFunc) (*CreateTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateTemplateDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   删除流程类型配置
 */
func (a *WorkItemConfService) DeleteTemplateDetail(ctx context.Context, req *DeleteTemplateDetailReq, options ...core.RequestOptionFunc) (*DeleteTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteTemplateDetail
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   获取流程类型配置详情
 */
func (a *WorkItemConfService) QueryTemplateDetail(ctx context.Context, req *QueryTemplateDetailReq, options ...core.RequestOptionFunc) (*QueryTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryTemplateDetail
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   获取工作项下的流程类型
 */
func (a *WorkItemConfService) QueryWorkItemTemplates(ctx context.Context, req *QueryWorkItemTemplatesReq, options ...core.RequestOptionFunc) (*QueryWorkItemTemplatesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemTemplates
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemTemplates] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemTemplatesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemTemplates] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   更新流程类型配置
 */
func (a *WorkItemConfService) UpdateTemplateDetail(ctx context.Context, req *UpdateTemplateDetailReq, options ...core.RequestOptionFunc) (*UpdateTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateTemplateDetail
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
