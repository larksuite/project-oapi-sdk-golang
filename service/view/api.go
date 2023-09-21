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

package view

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPathCreateFixView = "/open_api/:project_key/:work_item_type_key/fix_view"

const APIPathDeleteFixView = "/open_api/:project_key/fix_view/:view_id"

const APIPathQueryWorkItemDetailsByViewID = "/open_api/:project_key/view/:view_id"

const APIPathUpdateFixView = "/open_api/:project_key/:work_item_type_key/fix_view/:view_id"

const APIPathViewList = "/open_api/:project_key/view_conf/list"

const APIPathWorkItemList = "/open_api/:project_key/fix_view/:view_id"

func NewService(config *core.Config) *ViewService {
	a := &ViewService{config: config}
	return a
}

type ViewService struct {
	config *core.Config
}

// 创建固定视图
func (a *ViewService) CreateFixView(ctx context.Context, req *CreateFixViewReq, options ...core.RequestOptionFunc) (*CreateFixViewResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathCreateFixView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateFixView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateFixViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateFixView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 删除固定视图
func (a *ViewService) DeleteFixView(ctx context.Context, req *DeleteFixViewReq, options ...core.RequestOptionFunc) (*DeleteFixViewResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathDeleteFixView
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFixView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteFixViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFixView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取视图下工作项列表（全景视图）
func (a *ViewService) QueryWorkItemDetailsByViewID(ctx context.Context, req *QueryWorkItemDetailsByViewIDReq, options ...core.RequestOptionFunc) (*QueryWorkItemDetailsByViewIDResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryWorkItemDetailsByViewID
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemDetailsByViewID] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryWorkItemDetailsByViewIDResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemDetailsByViewID] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 更新固定视图
func (a *ViewService) UpdateFixView(ctx context.Context, req *UpdateFixViewReq, options ...core.RequestOptionFunc) (*UpdateFixViewResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathUpdateFixView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFixView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateFixViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFixView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取视图列表
func (a *ViewService) ViewList(ctx context.Context, req *ViewListReq, options ...core.RequestOptionFunc) (*ViewListResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathViewList
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ViewList] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ViewListResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ViewList] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取视图下工作项列表
func (a *ViewService) WorkItemList(ctx context.Context, req *WorkItemListReq, options ...core.RequestOptionFunc) (*WorkItemListResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathWorkItemList
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WorkItemList] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &WorkItemListResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WorkItemList] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
