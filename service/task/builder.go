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
	"code.byted.org/bits/project-oapi-sdk-golang/service/field"
	"code.byted.org/bits/project-oapi-sdk-golang/service/user"
	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem"
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"

	"code.byted.org/bits/project-oapi-sdk-golang/service/common"
)

type CreateSubTaskReq struct {
	apiReq *core.ApiReq
}
type CreateSubTaskReqBody struct {
	NodeID string `json:"node_id"`

	Name string `json:"name"`

	AliasKey string `json:"alias_key"`

	Assignee []string `json:"assignee"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`

	Schedule *workitem.Schedule `json:"schedule"`

	Note string `json:"note"`
}

type CreateSubTaskResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data int64 `json:"data"`
}

type CreateSubTaskReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateSubTaskReqBuilder() *CreateSubTaskReqBuilder {
	builder := &CreateSubTaskReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &CreateSubTaskReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *CreateSubTaskReqBuilder) AccessUser(userKey string) *CreateSubTaskReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *CreateSubTaskReqBuilder) UUID(uuid string) *CreateSubTaskReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *CreateSubTaskReqBuilder) ProjectKey(projectKey string) *CreateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *CreateSubTaskReqBuilder) WorkItemID(workItemID int64) *CreateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *CreateSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *CreateSubTaskReqBuilder) NodeID(nodeID string) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).NodeID = nodeID
	return builder
}
func (builder *CreateSubTaskReqBuilder) Name(name string) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).Name = name
	return builder
}
func (builder *CreateSubTaskReqBuilder) AliasKey(aliasKey string) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).AliasKey = aliasKey
	return builder
}
func (builder *CreateSubTaskReqBuilder) Assignee(assignee []string) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).Assignee = assignee
	return builder
}
func (builder *CreateSubTaskReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *CreateSubTaskReqBuilder) Schedule(schedule *workitem.Schedule) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).Schedule = schedule
	return builder
}
func (builder *CreateSubTaskReqBuilder) Note(note string) *CreateSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateSubTaskReqBody).Note = note
	return builder
}
func (builder *CreateSubTaskReqBuilder) Build() *CreateSubTaskReq {
	req := &CreateSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteSubTaskReq struct {
	apiReq *core.ApiReq
}

type DeleteSubTaskResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type DeleteSubTaskReqBuilder struct {
	apiReq *core.ApiReq
}

func NewDeleteSubTaskReqBuilder() *DeleteSubTaskReqBuilder {
	builder := &DeleteSubTaskReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *DeleteSubTaskReqBuilder) AccessUser(userKey string) *DeleteSubTaskReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *DeleteSubTaskReqBuilder) UUID(uuid string) *DeleteSubTaskReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *DeleteSubTaskReqBuilder) ProjectKey(projectKey string) *DeleteSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *DeleteSubTaskReqBuilder) WorkItemID(workItemID int64) *DeleteSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *DeleteSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *DeleteSubTaskReqBuilder) TaskID(taskID int64) *DeleteSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("task_id", fmt.Sprint(taskID))
	return builder
}
func (builder *DeleteSubTaskReqBuilder) Build() *DeleteSubTaskReq {
	req := &DeleteSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type ModifySubtaskReq struct {
	apiReq *core.ApiReq
}
type ModifySubtaskReqBody struct {
	NodeID string `json:"node_id"`

	TaskID int64 `json:"task_id"`

	Action string `json:"action"`

	Assignee []string `json:"assignee"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`

	Schedules []*workitem.Schedule `json:"schedules"`

	Deliverable []*field.FieldValuePair `json:"deliverable"`

	Note string `json:"note"`
}

type ModifySubtaskResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type ModifySubtaskReqBuilder struct {
	apiReq *core.ApiReq
}

func NewModifySubtaskReqBuilder() *ModifySubtaskReqBuilder {
	builder := &ModifySubtaskReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &ModifySubtaskReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *ModifySubtaskReqBuilder) AccessUser(userKey string) *ModifySubtaskReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *ModifySubtaskReqBuilder) UUID(uuid string) *ModifySubtaskReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *ModifySubtaskReqBuilder) ProjectKey(projectKey string) *ModifySubtaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *ModifySubtaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ModifySubtaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *ModifySubtaskReqBuilder) WorkItemID(workItemID int64) *ModifySubtaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *ModifySubtaskReqBuilder) NodeID(nodeID string) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).NodeID = nodeID
	return builder
}
func (builder *ModifySubtaskReqBuilder) TaskID(taskID int64) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).TaskID = taskID
	return builder
}
func (builder *ModifySubtaskReqBuilder) Action(action string) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).Action = action
	return builder
}
func (builder *ModifySubtaskReqBuilder) Assignee(assignee []string) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).Assignee = assignee
	return builder
}
func (builder *ModifySubtaskReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *ModifySubtaskReqBuilder) Schedules(schedules []*workitem.Schedule) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).Schedules = schedules
	return builder
}
func (builder *ModifySubtaskReqBuilder) Deliverable(deliverable []*field.FieldValuePair) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).Deliverable = deliverable
	return builder
}
func (builder *ModifySubtaskReqBuilder) Note(note string) *ModifySubtaskReqBuilder {
	builder.apiReq.Body.(*ModifySubtaskReqBody).Note = note
	return builder
}
func (builder *ModifySubtaskReqBuilder) Build() *ModifySubtaskReq {
	req := &ModifySubtaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type SearchSubtaskReq struct {
	apiReq *core.ApiReq
}
type SearchSubtaskReqBody struct {
	ProjectKeys []string `json:"project_keys"`

	PageSize int64 `json:"page_size"`

	PageNum int64 `json:"page_num"`

	Name string `json:"name"`

	UserKeys []string `json:"user_keys"`

	Status int32 `json:"status"`

	CreatedAt *workitem.TimeInterval `json:"created_at"`

	SimpleNames []string `json:"simple_names"`
}

type SearchSubtaskResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Pagination *common.Pagination `json:"pagination"`

	Data []*SubDetail `json:"data"`
}

type SearchSubtaskReqBuilder struct {
	apiReq *core.ApiReq
}

func NewSearchSubtaskReqBuilder() *SearchSubtaskReqBuilder {
	builder := &SearchSubtaskReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &SearchSubtaskReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *SearchSubtaskReqBuilder) AccessUser(userKey string) *SearchSubtaskReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *SearchSubtaskReqBuilder) UUID(uuid string) *SearchSubtaskReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *SearchSubtaskReqBuilder) ProjectKeys(projectKeys []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).ProjectKeys = projectKeys
	return builder
}
func (builder *SearchSubtaskReqBuilder) PageSize(pageSize int64) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).PageSize = pageSize
	return builder
}
func (builder *SearchSubtaskReqBuilder) PageNum(pageNum int64) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).PageNum = pageNum
	return builder
}
func (builder *SearchSubtaskReqBuilder) Name(name string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).Name = name
	return builder
}
func (builder *SearchSubtaskReqBuilder) UserKeys(userKeys []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).UserKeys = userKeys
	return builder
}
func (builder *SearchSubtaskReqBuilder) Status(status int32) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).Status = status
	return builder
}
func (builder *SearchSubtaskReqBuilder) CreatedAt(createdAt *workitem.TimeInterval) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).CreatedAt = createdAt
	return builder
}
func (builder *SearchSubtaskReqBuilder) SimpleNames(simpleNames []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *SearchSubtaskReqBuilder) Build() *SearchSubtaskReq {
	req := &SearchSubtaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type TaskDetailReq struct {
	apiReq *core.ApiReq
}

type TaskDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*workitem.workitem `json:"data"`
}

type TaskDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewTaskDetailReqBuilder() *TaskDetailReqBuilder {
	builder := &TaskDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *TaskDetailReqBuilder) AccessUser(userKey string) *TaskDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *TaskDetailReqBuilder) UUID(uuid string) *TaskDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *TaskDetailReqBuilder) ProjectKey(projectKey string) *TaskDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *TaskDetailReqBuilder) WorkItemID(workItemID int64) *TaskDetailReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *TaskDetailReqBuilder) WorkItemTypeKey(workItemTypeKey string) *TaskDetailReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *TaskDetailReqBuilder) NodeID(nodeID string) *TaskDetailReqBuilder {
	builder.apiReq.QueryParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}
func (builder *TaskDetailReqBuilder) Build() *TaskDetailReq {
	req := &TaskDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateSubTaskReq struct {
	apiReq *core.ApiReq
}
type UpdateSubTaskReqBody struct {
	Name string `json:"name"`

	Assignee []string `json:"assignee"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`

	Schedule *workitem.Schedule `json:"schedule"`

	Note string `json:"note"`

	Deliverable []*field.FieldValuePair `json:"deliverable"`
}

type UpdateSubTaskResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type UpdateSubTaskReqBuilder struct {
	apiReq *core.ApiReq
}

func NewUpdateSubTaskReqBuilder() *UpdateSubTaskReqBuilder {
	builder := &UpdateSubTaskReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &UpdateSubTaskReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *UpdateSubTaskReqBuilder) AccessUser(userKey string) *UpdateSubTaskReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *UpdateSubTaskReqBuilder) UUID(uuid string) *UpdateSubTaskReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *UpdateSubTaskReqBuilder) ProjectKey(projectKey string) *UpdateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *UpdateSubTaskReqBuilder) WorkItemID(workItemID int64) *UpdateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *UpdateSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *UpdateSubTaskReqBuilder) NodeID(nodeID string) *UpdateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}
func (builder *UpdateSubTaskReqBuilder) TaskID(taskID int64) *UpdateSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("task_id", fmt.Sprint(taskID))
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Name(name string) *UpdateSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateSubTaskReqBody).Name = name
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Assignee(assignee []string) *UpdateSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateSubTaskReqBody).Assignee = assignee
	return builder
}
func (builder *UpdateSubTaskReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *UpdateSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateSubTaskReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Schedule(schedule *workitem.Schedule) *UpdateSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateSubTaskReqBody).Schedule = schedule
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Note(note string) *UpdateSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateSubTaskReqBody).Note = note
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Deliverable(deliverable []*field.FieldValuePair) *UpdateSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateSubTaskReqBody).Deliverable = deliverable
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Build() *UpdateSubTaskReq {
	req := &UpdateSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}
