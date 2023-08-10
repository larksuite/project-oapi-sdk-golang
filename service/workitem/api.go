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

package workitem

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPath_AbortWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/abort"

const APIPath_CompositiveSearch = "/open_api/compositive_search"

const APIPath_CreateWorkItem = "/open_api/:project_key/work_item/create"

const APIPath_CreateWorkItemRelation = "/open_api/work_item/relation/create"

const APIPath_DeleteWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id"

const APIPath_DeleteWorkItemRelation = "/open_api/work_item/relation/delete"

const APIPath_Filter = "/open_api/:project_key/work_item/filter"

const APIPath_FilterAcrossProject = "/open_api/work_items/filter_across_project"

const APIPath_GetMeta = "/open_api/:project_key/work_item/:work_item_type_key/meta"

const APIPath_NodeOperate = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/:node_id/operate"

const APIPath_NodeStateChange = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/state_change"

const APIPath_NodeUpdate = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/:node_id"

const APIPath_QueryWorkItemDetail = "/open_api/:project_key/work_item/:work_item_type_key/query"

const APIPath_QueryWorkItemRelation = "/open_api/:project_key/work_item/relation"

const APIPath_QueryWorkflow = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/query"

const APIPath_SearchByParams = "/open_api/:project_key/work_item/:work_item_type_key/search/params"

const APIPath_SearchByRelation = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/search_by_relation"

const APIPath_UpdateMultiSignal = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/update/multi_signal"

const APIPath_UpdateWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id"

const APIPath_UpdateWorkItemRelation = "/open_api/work_item/relation/update"

const APIPath_WbsView = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/wbs_view"

func NewService(config *core.Config) *WorkItemService {
	a := &WorkItemService{config: config}
	return a
}

type WorkItemService struct {
	config *core.Config
}

/*
 *
 */
func (a *WorkItemService) AbortWorkItem(ctx context.Context, req *AbortWorkItemReq, options ...core.RequestOptionFunc) (*AbortWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_AbortWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[AbortWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &AbortWorkItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[AbortWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) CompositiveSearch(ctx context.Context, req *CompositiveSearchReq, options ...core.RequestOptionFunc) (*CompositiveSearchResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CompositiveSearch
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CompositiveSearch] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CompositiveSearchResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CompositiveSearch] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) CreateWorkItem(ctx context.Context, req *CreateWorkItemReq, options ...core.RequestOptionFunc) (*CreateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateWorkItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) CreateWorkItemRelation(ctx context.Context, req *CreateWorkItemRelationReq, options ...core.RequestOptionFunc) (*CreateWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateWorkItemRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateWorkItemRelationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) DeleteWorkItem(ctx context.Context, req *DeleteWorkItemReq, options ...core.RequestOptionFunc) (*DeleteWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteWorkItem
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteWorkItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) DeleteWorkItemRelation(ctx context.Context, req *DeleteWorkItemRelationReq, options ...core.RequestOptionFunc) (*DeleteWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteWorkItemRelation
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteWorkItemRelationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) Filter(ctx context.Context, req *FilterReq, options ...core.RequestOptionFunc) (*FilterResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_Filter
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[Filter] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &FilterResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[Filter] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) FilterAcrossProject(ctx context.Context, req *FilterAcrossProjectReq, options ...core.RequestOptionFunc) (*FilterAcrossProjectResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_FilterAcrossProject
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[FilterAcrossProject] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &FilterAcrossProjectResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[FilterAcrossProject] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) GetMeta(ctx context.Context, req *GetMetaReq, options ...core.RequestOptionFunc) (*GetMetaResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetMeta
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetMeta] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMetaResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetMeta] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) NodeOperate(ctx context.Context, req *NodeOperateReq, options ...core.RequestOptionFunc) (*NodeOperateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_NodeOperate
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[NodeOperate] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &NodeOperateResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[NodeOperate] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) NodeStateChange(ctx context.Context, req *NodeStateChangeReq, options ...core.RequestOptionFunc) (*NodeStateChangeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_NodeStateChange
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[NodeStateChange] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &NodeStateChangeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[NodeStateChange] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) NodeUpdate(ctx context.Context, req *NodeUpdateReq, options ...core.RequestOptionFunc) (*NodeUpdateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_NodeUpdate
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[NodeUpdate] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &NodeUpdateResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[NodeUpdate] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) QueryWorkItemDetail(ctx context.Context, req *QueryWorkItemDetailReq, options ...core.RequestOptionFunc) (*QueryWorkItemDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemDetailResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) QueryWorkItemRelation(ctx context.Context, req *QueryWorkItemRelationReq, options ...core.RequestOptionFunc) (*QueryWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemRelation
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemRelationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) QueryWorkflow(ctx context.Context, req *QueryWorkflowReq, options ...core.RequestOptionFunc) (*QueryWorkflowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkflow
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkflow] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkflowResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkflow] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) SearchByParams(ctx context.Context, req *SearchByParamsReq, options ...core.RequestOptionFunc) (*SearchByParamsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchByParams
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchByParams] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchByParamsResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchByParams] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) SearchByRelation(ctx context.Context, req *SearchByRelationReq, options ...core.RequestOptionFunc) (*SearchByRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchByRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchByRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchByRelationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchByRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) UpdateMultiSignal(ctx context.Context, req *UpdateMultiSignalReq, options ...core.RequestOptionFunc) (*UpdateMultiSignalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateMultiSignal
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateMultiSignal] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateMultiSignalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateMultiSignal] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) UpdateWorkItem(ctx context.Context, req *UpdateWorkItemReq, options ...core.RequestOptionFunc) (*UpdateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkItemResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) UpdateWorkItemRelation(ctx context.Context, req *UpdateWorkItemRelationReq, options ...core.RequestOptionFunc) (*UpdateWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkItemRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkItemRelationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) WbsView(ctx context.Context, req *WbsViewReq, options ...core.RequestOptionFunc) (*WbsViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_WbsView
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WbsView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &WbsViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WbsView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
