package search

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_OAPICreateConditionView = "/open_api/view/v1/create_condition_view"

const APIPath_OAPIUpdateConditionView = "/open_api/view/v1/update_condition_view"


func NewService(config *core.Config) *SearchService {
	a := &SearchService{config: config}
	return a
}

type SearchService struct {
	config *core.Config
}


/*
 * @name: openapi创建条件视图
 * @desc: openapi创建条件视图
 */
func (a *SearchService) OAPICreateConditionView(ctx context.Context, req *OAPICreateConditionViewReq, options ...core.RequestOptionFunc) (*OAPICreateConditionViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateConditionView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateConditionView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateConditionViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateConditionView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi更新条件视图
 * @desc: openapi更新条件视图
 */
func (a *SearchService) OAPIUpdateConditionView(ctx context.Context, req *OAPIUpdateConditionViewReq, options ...core.RequestOptionFunc) (*OAPIUpdateConditionViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateConditionView
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateConditionView] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateConditionViewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateConditionView] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

