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

package task

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPathCreateSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/task"

const APIPathDeleteSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/task/:task_id"

const APIPathModifySubtask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/subtask/modify"

const APIPathSearchSubtask = "/open_api/work_item/subtask/search"

const APIPathTaskDetail = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/task"

const APIPathUpdateSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/:node_id/task/:task_id"

func NewService(config *core.Config) *TaskService {
	a := &TaskService{config: config}
	return a
}

type TaskService struct {
	config *core.Config
}

/*
 *   创建子任务
 */
func (a *TaskService) CreateSubTask(ctx context.Context, req *CreateSubTaskReq, options ...core.RequestOptionFunc) (*CreateSubTaskResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathCreateSubTask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   删除子任务
 */
func (a *TaskService) DeleteSubTask(ctx context.Context, req *DeleteSubTaskReq, options ...core.RequestOptionFunc) (*DeleteSubTaskResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathDeleteSubTask
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   子任务完成/回滚
 */
func (a *TaskService) ModifySubtask(ctx context.Context, req *ModifySubtaskReq, options ...core.RequestOptionFunc) (*ModifySubtaskResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathModifySubtask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ModifySubtask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ModifySubtaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ModifySubtask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   获取指定的子任务列表（跨空间）
 */
func (a *TaskService) SearchSubtask(ctx context.Context, req *SearchSubtaskReq, options ...core.RequestOptionFunc) (*SearchSubtaskResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathSearchSubtask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchSubtask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SearchSubtaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchSubtask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   获取子任务详情
 */
func (a *TaskService) TaskDetail(ctx context.Context, req *TaskDetailReq, options ...core.RequestOptionFunc) (*TaskDetailResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathTaskDetail
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[TaskDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &TaskDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[TaskDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   更新子任务
 */
func (a *TaskService) UpdateSubTask(ctx context.Context, req *UpdateSubTaskReq, options ...core.RequestOptionFunc) (*UpdateSubTaskResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathUpdateSubTask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
