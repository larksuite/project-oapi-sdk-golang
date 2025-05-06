package project

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_BatchQueryProjectInfo = "/open_api/projects/detail"

const APIPath_QueryProjects = "/open_api/projects"


func NewService(config *core.Config) *ProjectService {
	a := &ProjectService{config: config}
	return a
}

type ProjectService struct {
	config *core.Config
}


/*
 * @name:OAPIBatchQueryProjectInfo
 * @desc:OpenAPI，获取空间详情
 */
func (a *ProjectService) BatchQueryProjectInfo(ctx context.Context, req *BatchQueryProjectInfoReq, options ...core.RequestOptionFunc) (*BatchQueryProjectInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_BatchQueryProjectInfo
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryProjectInfo] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchQueryProjectInfoResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryProjectInfo] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:OAPIQueryProjects
 * @desc:OpenAPI，获取空间列表
 */
func (a *ProjectService) QueryProjects(ctx context.Context, req *QueryProjectsReq, options ...core.RequestOptionFunc) (*QueryProjectsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryProjects
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjects] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryProjectsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjects] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

