package view

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_CreateConditionView = "/open_api/view/v1/create_condition_view"

const APIPath_CreateFixView = "/open_api/:project_key/:work_item_type_key/fix_view"

const APIPath_DeleteFixView = "/open_api/:project_key/fix_view/:view_id"

const APIPath_QuerySpaceUIAggFields = "/open_api/view_search/space_agg_fields/query_ui"

const APIPath_QueryWorkItemDetailsByViewID = "/open_api/:project_key/view/:view_id"

const APIPath_UpdateConditionView = "/open_api/view/v1/update_condition_view"

const APIPath_UpdateFixView = "/open_api/:project_key/:work_item_type_key/fix_view/:view_id"

const APIPath_ViewList = "/open_api/:project_key/view_conf/list"

const APIPath_WorkItemList = "/open_api/:project_key/fix_view/:view_id"


func NewService(config *core.Config) *ViewService {
	a := &ViewService{config: config}
	return a
}

type ViewService struct {
	config *core.Config
}


/*
 * @name: openapi创建条件视图
 * @desc: openapi创建条件视图
 */
func (a *ViewService) CreateConditionView(ctx context.Context, req *CreateConditionViewReq, options ...core.RequestOptionFunc) (*CreateConditionViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateConditionView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateConditionView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateConditionViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateConditionView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi创建固定视图
 * @desc: openapi创建固定视图
 */
func (a *ViewService) CreateFixView(ctx context.Context, req *CreateFixViewReq, options ...core.RequestOptionFunc) (*CreateFixViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateFixView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateFixView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFixViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateFixView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi删除固定视图
 * @desc: openapi删除固定视图
 */
func (a *ViewService) DeleteFixView(ctx context.Context, req *DeleteFixViewReq, options ...core.RequestOptionFunc) (*DeleteFixViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteFixView
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFixView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFixViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFixView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取空间级别聚合字段
 * @desc: openapi获取空间级别聚合字段
 */
func (a *ViewService) QuerySpaceUIAggFields(ctx context.Context, req *QuerySpaceUIAggFieldsReq, options ...core.RequestOptionFunc) (*QuerySpaceUIAggFieldsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QuerySpaceUIAggFields
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QuerySpaceUIAggFields] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpaceUIAggFieldsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QuerySpaceUIAggFields] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取视图下工作项列表（全景视图）
 * @desc: openapi获取视图下工作项列表（全景视图）
 */
func (a *ViewService) QueryWorkItemDetailsByViewID(ctx context.Context, req *QueryWorkItemDetailsByViewIDReq, options ...core.RequestOptionFunc) (*QueryWorkItemDetailsByViewIDResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemDetailsByViewID
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemDetailsByViewID] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemDetailsByViewIDResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemDetailsByViewID] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi更新条件视图
 * @desc: openapi更新条件视图
 */
func (a *ViewService) UpdateConditionView(ctx context.Context, req *UpdateConditionViewReq, options ...core.RequestOptionFunc) (*UpdateConditionViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateConditionView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateConditionView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateConditionViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateConditionView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi更新固定视图
 * @desc: openapi更新固定视图
 */
func (a *ViewService) UpdateFixView(ctx context.Context, req *UpdateFixViewReq, options ...core.RequestOptionFunc) (*UpdateFixViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateFixView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFixView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateFixViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFixView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取视图列表
 * @desc: openapi获取视图列表
 */
func (a *ViewService) ViewList(ctx context.Context, req *ViewListReq, options ...core.RequestOptionFunc) (*ViewListResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ViewList
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ViewList] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ViewListResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ViewList] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取视图下工作项列表
 * @desc: openapi获取视图下工作项列表
 */
func (a *ViewService) WorkItemList(ctx context.Context, req *WorkItemListReq, options ...core.RequestOptionFunc) (*WorkItemListResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_WorkItemList
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WorkItemList] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &WorkItemListResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WorkItemList] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

