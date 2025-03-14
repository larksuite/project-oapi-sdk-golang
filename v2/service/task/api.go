package task

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_SearchSubtask = "/open_api/work_item/subtask/search"


func NewService(config *core.Config) *TaskService {
	a := &TaskService{config: config}
	return a
}

type TaskService struct {
	config *core.Config
}


/*
 * @name: openapi获取指定的子任务列表（跨空间）
 * @desc: openapi获取指定的子任务列表（跨空间）
 */
func (a *TaskService) SearchSubtask(ctx context.Context, req *SearchSubtaskReq, options ...core.RequestOptionFunc) (*SearchSubtaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchSubtask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchSubtask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchSubtaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchSubtask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

