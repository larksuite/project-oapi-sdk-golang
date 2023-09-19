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
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"

	"code.byted.org/bits/project-oapi-sdk-golang/service/common"

	"code.byted.org/bits/project-oapi-sdk-golang/service/field"

	"code.byted.org/bits/project-oapi-sdk-golang/service/user"

	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem"
)

type CreateSubTaskReq struct {
	apiReq *core.APIReq
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
	*core.APIResp `json:"-"`
	core.CodeError
	Data int64 `json:"data"`
}

type CreateSubTaskReqBuilder struct {
	apiReq *core.APIReq
	body   *CreateSubTaskReqBody
}

func NewCreateSubTaskReqBuilder() *CreateSubTaskReqBuilder {
	builder := &CreateSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateSubTaskReqBody{}
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
	builder.body.NodeID = nodeID
	return builder
}
func (builder *CreateSubTaskReqBuilder) Name(name string) *CreateSubTaskReqBuilder {
	builder.body.Name = name
	return builder
}
func (builder *CreateSubTaskReqBuilder) AliasKey(aliasKey string) *CreateSubTaskReqBuilder {
	builder.body.AliasKey = aliasKey
	return builder
}
func (builder *CreateSubTaskReqBuilder) Assignee(assignee []string) *CreateSubTaskReqBuilder {
	builder.body.Assignee = assignee
	return builder
}
func (builder *CreateSubTaskReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *CreateSubTaskReqBuilder {
	builder.body.RoleAssignee = roleAssignee
	return builder
}
func (builder *CreateSubTaskReqBuilder) Schedule(schedule *workitem.Schedule) *CreateSubTaskReqBuilder {
	builder.body.Schedule = schedule
	return builder
}
func (builder *CreateSubTaskReqBuilder) Note(note string) *CreateSubTaskReqBuilder {
	builder.body.Note = note
	return builder
}
func (builder *CreateSubTaskReqBuilder) Build() *CreateSubTaskReq {
	req := &CreateSubTaskReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type DeleteSubTaskReq struct {
	apiReq *core.APIReq
}

type DeleteSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteSubTaskReqBuilder() *DeleteSubTaskReqBuilder {
	builder := &DeleteSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
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
	apiReq *core.APIReq
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
	*core.APIResp `json:"-"`
	core.CodeError
}

type ModifySubtaskReqBuilder struct {
	apiReq *core.APIReq
	body   *ModifySubtaskReqBody
}

func NewModifySubtaskReqBuilder() *ModifySubtaskReqBuilder {
	builder := &ModifySubtaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &ModifySubtaskReqBody{}
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
	builder.body.NodeID = nodeID
	return builder
}
func (builder *ModifySubtaskReqBuilder) TaskID(taskID int64) *ModifySubtaskReqBuilder {
	builder.body.TaskID = taskID
	return builder
}
func (builder *ModifySubtaskReqBuilder) Action(action string) *ModifySubtaskReqBuilder {
	builder.body.Action = action
	return builder
}
func (builder *ModifySubtaskReqBuilder) Assignee(assignee []string) *ModifySubtaskReqBuilder {
	builder.body.Assignee = assignee
	return builder
}
func (builder *ModifySubtaskReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *ModifySubtaskReqBuilder {
	builder.body.RoleAssignee = roleAssignee
	return builder
}
func (builder *ModifySubtaskReqBuilder) Schedules(schedules []*workitem.Schedule) *ModifySubtaskReqBuilder {
	builder.body.Schedules = schedules
	return builder
}
func (builder *ModifySubtaskReqBuilder) Deliverable(deliverable []*field.FieldValuePair) *ModifySubtaskReqBuilder {
	builder.body.Deliverable = deliverable
	return builder
}
func (builder *ModifySubtaskReqBuilder) Note(note string) *ModifySubtaskReqBuilder {
	builder.body.Note = note
	return builder
}
func (builder *ModifySubtaskReqBuilder) Build() *ModifySubtaskReq {
	req := &ModifySubtaskReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type SearchSubtaskReq struct {
	apiReq *core.APIReq
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
	*core.APIResp `json:"-"`
	core.CodeError
	Pagination *common.Pagination `json:"pagination"`

	Data []*SubDetail `json:"data"`
}

type SearchSubtaskReqBuilder struct {
	apiReq *core.APIReq
	body   *SearchSubtaskReqBody
}

func NewSearchSubtaskReqBuilder() *SearchSubtaskReqBuilder {
	builder := &SearchSubtaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &SearchSubtaskReqBody{}
	return builder
}
func (builder *SearchSubtaskReqBuilder) ProjectKeys(projectKeys []string) *SearchSubtaskReqBuilder {
	builder.body.ProjectKeys = projectKeys
	return builder
}
func (builder *SearchSubtaskReqBuilder) PageSize(pageSize int64) *SearchSubtaskReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *SearchSubtaskReqBuilder) PageNum(pageNum int64) *SearchSubtaskReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *SearchSubtaskReqBuilder) Name(name string) *SearchSubtaskReqBuilder {
	builder.body.Name = name
	return builder
}
func (builder *SearchSubtaskReqBuilder) UserKeys(userKeys []string) *SearchSubtaskReqBuilder {
	builder.body.UserKeys = userKeys
	return builder
}
func (builder *SearchSubtaskReqBuilder) Status(status int32) *SearchSubtaskReqBuilder {
	builder.body.Status = status
	return builder
}
func (builder *SearchSubtaskReqBuilder) CreatedAt(createdAt *workitem.TimeInterval) *SearchSubtaskReqBuilder {
	builder.body.CreatedAt = createdAt
	return builder
}
func (builder *SearchSubtaskReqBuilder) SimpleNames(simpleNames []string) *SearchSubtaskReqBuilder {
	builder.body.SimpleNames = simpleNames
	return builder
}
func (builder *SearchSubtaskReqBuilder) Build() *SearchSubtaskReq {
	req := &SearchSubtaskReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type TaskDetailReq struct {
	apiReq *core.APIReq
}

type TaskDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*workitem.NodeTask `json:"data"`
}

type TaskDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewTaskDetailReqBuilder() *TaskDetailReqBuilder {
	builder := &TaskDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
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
	apiReq *core.APIReq
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
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateSubTaskReqBuilder struct {
	apiReq *core.APIReq
	body   *UpdateSubTaskReqBody
}

func NewUpdateSubTaskReqBuilder() *UpdateSubTaskReqBuilder {
	builder := &UpdateSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateSubTaskReqBody{}
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
	builder.body.Name = name
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Assignee(assignee []string) *UpdateSubTaskReqBuilder {
	builder.body.Assignee = assignee
	return builder
}
func (builder *UpdateSubTaskReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *UpdateSubTaskReqBuilder {
	builder.body.RoleAssignee = roleAssignee
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Schedule(schedule *workitem.Schedule) *UpdateSubTaskReqBuilder {
	builder.body.Schedule = schedule
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Note(note string) *UpdateSubTaskReqBuilder {
	builder.body.Note = note
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Deliverable(deliverable []*field.FieldValuePair) *UpdateSubTaskReqBuilder {
	builder.body.Deliverable = deliverable
	return builder
}
func (builder *UpdateSubTaskReqBuilder) Build() *UpdateSubTaskReq {
	req := &UpdateSubTaskReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
