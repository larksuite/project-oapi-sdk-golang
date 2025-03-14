package open_api

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_OAPIBatchWebHook = "/open_api/control/batch/webhook"

const APIPath_OAPIQueryTaskResult = "/open_api/task_result"


func NewService(config *core.Config) *OpenAPIService {
	a := &OpenAPIService{config: config}
	return a
}

type OpenAPIService struct {
	config *core.Config
}


/*
 * ================================================= 插件运行测试使用的open api接口 ======================================================
@name: OAPIQueryTaskResult
 * @desc: 查询任务执行情况——主要针对批量任务处理
 */
func (a *OpenAPIService) OAPIBatchWebHook(ctx context.Context, req *OAPIBatchWebHookReq, options ...core.RequestOptionFunc) (*OAPIBatchWebHookResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIBatchWebHook
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchWebHook] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIBatchWebHookResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchWebHook] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryTaskResult
 * @desc: 查询任务执行情况——主要针对批量任务处理
 */
func (a *OpenAPIService) OAPIQueryTaskResult(ctx context.Context, req *OAPIQueryTaskResultReq, options ...core.RequestOptionFunc) (*OAPIQueryTaskResultResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryTaskResult
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryTaskResult] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryTaskResultResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryTaskResult] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

