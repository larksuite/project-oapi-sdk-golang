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
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"

	"code.byted.org/bits/project-oapi-sdk-golang/service/common"

	"code.byted.org/bits/project-oapi-sdk-golang/service/field"

	"code.byted.org/bits/project-oapi-sdk-golang/service/user"
)

type AbortWorkItemReq struct {
	apiReq *core.ApiReq
}
type AbortWorkItemReqBody struct {
	IsAborted bool `json:"is_aborted"`

	Reason string `json:"reason"`
}

type AbortWorkItemResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type AbortWorkItemReqBuilder struct {
	apiReq *core.ApiReq
	body   *AbortWorkItemReqBody
}

func NewAbortWorkItemReqBuilder() *AbortWorkItemReqBuilder {
	builder := &AbortWorkItemReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &AbortWorkItemReqBody{}
	return builder
}
func (builder *AbortWorkItemReqBuilder) ProjectKey(projectKey string) *AbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *AbortWorkItemReqBuilder) WorkItemID(workItemID int64) *AbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *AbortWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *AbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *AbortWorkItemReqBuilder) IsAborted(isAborted bool) *AbortWorkItemReqBuilder {
	builder.body.IsAborted = isAborted
	return builder
}
func (builder *AbortWorkItemReqBuilder) Reason(reason string) *AbortWorkItemReqBuilder {
	builder.body.Reason = reason
	return builder
}
func (builder *AbortWorkItemReqBuilder) Build() *AbortWorkItemReq {
	req := &AbortWorkItemReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type CompositiveSearchReq struct {
	apiReq *core.ApiReq
}
type CompositiveSearchReqBody struct {
	ProjectKeys []string `json:"project_keys"`

	QueryType string `json:"query_type"`

	Query string `json:"query"`

	QuerySubType []string `json:"query_sub_type"`

	PageSize int64 `json:"page_size"`

	PageNum int64 `json:"page_num"`

	SimpleNames []string `json:"simple_names"`
}

type CompositiveSearchResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*CompInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type CompositiveSearchReqBuilder struct {
	apiReq *core.ApiReq
	body   *CompositiveSearchReqBody
}

func NewCompositiveSearchReqBuilder() *CompositiveSearchReqBuilder {
	builder := &CompositiveSearchReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CompositiveSearchReqBody{}
	return builder
}
func (builder *CompositiveSearchReqBuilder) ProjectKeys(projectKeys []string) *CompositiveSearchReqBuilder {
	builder.body.ProjectKeys = projectKeys
	return builder
}
func (builder *CompositiveSearchReqBuilder) QueryType(queryType string) *CompositiveSearchReqBuilder {
	builder.body.QueryType = queryType
	return builder
}
func (builder *CompositiveSearchReqBuilder) Query(query string) *CompositiveSearchReqBuilder {
	builder.body.Query = query
	return builder
}
func (builder *CompositiveSearchReqBuilder) QuerySubType(querySubType []string) *CompositiveSearchReqBuilder {
	builder.body.QuerySubType = querySubType
	return builder
}
func (builder *CompositiveSearchReqBuilder) PageSize(pageSize int64) *CompositiveSearchReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *CompositiveSearchReqBuilder) PageNum(pageNum int64) *CompositiveSearchReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *CompositiveSearchReqBuilder) SimpleNames(simpleNames []string) *CompositiveSearchReqBuilder {
	builder.body.SimpleNames = simpleNames
	return builder
}
func (builder *CompositiveSearchReqBuilder) Build() *CompositiveSearchReq {
	req := &CompositiveSearchReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type CreateWorkItemReq struct {
	apiReq *core.ApiReq
}
type CreateWorkItemReqBody struct {
	WorkItemTypeKey string `json:"work_item_type_key"`

	FieldValuePairs []*field.FieldValuePair `json:"field_value_pairs"`

	TemplateID int64 `json:"template_id"`

	Name string `json:"name"`
}

type CreateWorkItemResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data int64 `json:"data"`
}

type CreateWorkItemReqBuilder struct {
	apiReq *core.ApiReq
	body   *CreateWorkItemReqBody
}

func NewCreateWorkItemReqBuilder() *CreateWorkItemReqBuilder {
	builder := &CreateWorkItemReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateWorkItemReqBody{}
	return builder
}
func (builder *CreateWorkItemReqBuilder) ProjectKey(projectKey string) *CreateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *CreateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateWorkItemReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *CreateWorkItemReqBuilder) FieldValuePairs(fieldValuePairs []*field.FieldValuePair) *CreateWorkItemReqBuilder {
	builder.body.FieldValuePairs = fieldValuePairs
	return builder
}
func (builder *CreateWorkItemReqBuilder) TemplateID(templateID int64) *CreateWorkItemReqBuilder {
	builder.body.TemplateID = templateID
	return builder
}
func (builder *CreateWorkItemReqBuilder) Name(name string) *CreateWorkItemReqBuilder {
	builder.body.Name = name
	return builder
}
func (builder *CreateWorkItemReqBuilder) Build() *CreateWorkItemReq {
	req := &CreateWorkItemReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type CreateWorkItemRelationReq struct {
	apiReq *core.ApiReq
}
type CreateWorkItemRelationReqBody struct {
	ProjectKey string `json:"project_key"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	Name string `json:"name"`

	RelationDetails []*RelationDetail `json:"relation_details"`
}

type CreateWorkItemRelationResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *CreateWorkItemRelationData `json:"data"`
}

type CreateWorkItemRelationReqBuilder struct {
	apiReq *core.ApiReq
	body   *CreateWorkItemRelationReqBody
}

func NewCreateWorkItemRelationReqBuilder() *CreateWorkItemRelationReqBuilder {
	builder := &CreateWorkItemRelationReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateWorkItemRelationReqBody{}
	return builder
}
func (builder *CreateWorkItemRelationReqBuilder) ProjectKey(projectKey string) *CreateWorkItemRelationReqBuilder {
	builder.body.ProjectKey = projectKey
	return builder
}
func (builder *CreateWorkItemRelationReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateWorkItemRelationReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *CreateWorkItemRelationReqBuilder) Name(name string) *CreateWorkItemRelationReqBuilder {
	builder.body.Name = name
	return builder
}
func (builder *CreateWorkItemRelationReqBuilder) RelationDetails(relationDetails []*RelationDetail) *CreateWorkItemRelationReqBuilder {
	builder.body.RelationDetails = relationDetails
	return builder
}
func (builder *CreateWorkItemRelationReqBuilder) Build() *CreateWorkItemRelationReq {
	req := &CreateWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type DeleteWorkItemReq struct {
	apiReq *core.ApiReq
}

type DeleteWorkItemResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type DeleteWorkItemReqBuilder struct {
	apiReq *core.ApiReq
}

func NewDeleteWorkItemReqBuilder() *DeleteWorkItemReqBuilder {
	builder := &DeleteWorkItemReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *DeleteWorkItemReqBuilder) ProjectKey(projectKey string) *DeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *DeleteWorkItemReqBuilder) WorkItemID(workItemID int64) *DeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *DeleteWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *DeleteWorkItemReqBuilder) Build() *DeleteWorkItemReq {
	req := &DeleteWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteWorkItemRelationReq struct {
	apiReq *core.ApiReq
}
type DeleteWorkItemRelationReqBody struct {
	RelationID string `json:"relation_id"`

	ProjectKey string `json:"project_key"`
}

type DeleteWorkItemRelationResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data string `json:"data"`
}

type DeleteWorkItemRelationReqBuilder struct {
	apiReq *core.ApiReq
	body   *DeleteWorkItemRelationReqBody
}

func NewDeleteWorkItemRelationReqBuilder() *DeleteWorkItemRelationReqBuilder {
	builder := &DeleteWorkItemRelationReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &DeleteWorkItemRelationReqBody{}
	return builder
}
func (builder *DeleteWorkItemRelationReqBuilder) RelationID(relationID string) *DeleteWorkItemRelationReqBuilder {
	builder.body.RelationID = relationID
	return builder
}
func (builder *DeleteWorkItemRelationReqBuilder) ProjectKey(projectKey string) *DeleteWorkItemRelationReqBuilder {
	builder.body.ProjectKey = projectKey
	return builder
}
func (builder *DeleteWorkItemRelationReqBuilder) Build() *DeleteWorkItemRelationReq {
	req := &DeleteWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type FilterReq struct {
	apiReq *core.ApiReq
}
type FilterReqBody struct {
	WorkItemName string `json:"work_item_name"`

	UserKeys []string `json:"user_keys"`

	WorkItemIDs []int64 `json:"work_item_ids"`

	WorkItemTypeKeys []string `json:"work_item_type_keys"`

	CreatedAt *TimeInterval `json:"created_at"`

	UpdatedAt *TimeInterval `json:"updated_at"`

	SubStages []string `json:"sub_stages"`

	Businesses []string `json:"businesses"`

	Priorities []string `json:"priorities"`

	Tags []string `json:"tags"`

	PageNum int64 `json:"page_num"`

	PageSize int64 `json:"page_size"`

	WorkItemStatus []*WorkItemStatus `json:"work_item_status"`

	Expand *Expand `json:"expand"`
}

type FilterResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*WorkItemInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type FilterReqBuilder struct {
	apiReq *core.ApiReq
	body   *FilterReqBody
}

func NewFilterReqBuilder() *FilterReqBuilder {
	builder := &FilterReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &FilterReqBody{}
	return builder
}
func (builder *FilterReqBuilder) ProjectKey(projectKey string) *FilterReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *FilterReqBuilder) WorkItemName(workItemName string) *FilterReqBuilder {
	builder.body.WorkItemName = workItemName
	return builder
}
func (builder *FilterReqBuilder) UserKeys(userKeys []string) *FilterReqBuilder {
	builder.body.UserKeys = userKeys
	return builder
}
func (builder *FilterReqBuilder) WorkItemIDs(workItemIDs []int64) *FilterReqBuilder {
	builder.body.WorkItemIDs = workItemIDs
	return builder
}
func (builder *FilterReqBuilder) WorkItemTypeKeys(workItemTypeKeys []string) *FilterReqBuilder {
	builder.body.WorkItemTypeKeys = workItemTypeKeys
	return builder
}
func (builder *FilterReqBuilder) CreatedAt(createdAt *TimeInterval) *FilterReqBuilder {
	builder.body.CreatedAt = createdAt
	return builder
}
func (builder *FilterReqBuilder) UpdatedAt(updatedAt *TimeInterval) *FilterReqBuilder {
	builder.body.UpdatedAt = updatedAt
	return builder
}
func (builder *FilterReqBuilder) SubStages(subStages []string) *FilterReqBuilder {
	builder.body.SubStages = subStages
	return builder
}
func (builder *FilterReqBuilder) Businesses(businesses []string) *FilterReqBuilder {
	builder.body.Businesses = businesses
	return builder
}
func (builder *FilterReqBuilder) Priorities(priorities []string) *FilterReqBuilder {
	builder.body.Priorities = priorities
	return builder
}
func (builder *FilterReqBuilder) Tags(tags []string) *FilterReqBuilder {
	builder.body.Tags = tags
	return builder
}
func (builder *FilterReqBuilder) PageNum(pageNum int64) *FilterReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *FilterReqBuilder) PageSize(pageSize int64) *FilterReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *FilterReqBuilder) WorkItemStatus(workItemStatus []*WorkItemStatus) *FilterReqBuilder {
	builder.body.WorkItemStatus = workItemStatus
	return builder
}
func (builder *FilterReqBuilder) Expand(expand *Expand) *FilterReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *FilterReqBuilder) Build() *FilterReq {
	req := &FilterReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type FilterAcrossProjectReq struct {
	apiReq *core.ApiReq
}
type FilterAcrossProjectReqBody struct {
	ProjectKeys []string `json:"project_keys"`

	SearchUser *SearchUser `json:"search_user"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	CreatedAt *TimeInterval `json:"created_at"`

	UpdatedAt *TimeInterval `json:"updated_at"`

	WorkItemStatus []*WorkItemStatus `json:"work_item_status"`

	WorkItemName string `json:"work_item_name"`

	PageNum int64 `json:"page_num"`

	PageSize int64 `json:"page_size"`

	TenantGroupID int64 `json:"tenant_group_id"`

	WorkItemIDs []int64 `json:"work_item_ids"`

	Businesses []string `json:"businesses"`

	Priorities []string `json:"priorities"`

	Tags []string `json:"tags"`

	SimpleNames []string `json:"simple_names"`

	TemplateIDs []int64 `json:"template_ids"`

	Expand *Expand `json:"expand"`
}

type FilterAcrossProjectResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*WorkItemInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type FilterAcrossProjectReqBuilder struct {
	apiReq *core.ApiReq
	body   *FilterAcrossProjectReqBody
}

func NewFilterAcrossProjectReqBuilder() *FilterAcrossProjectReqBuilder {
	builder := &FilterAcrossProjectReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &FilterAcrossProjectReqBody{}
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) ProjectKeys(projectKeys []string) *FilterAcrossProjectReqBuilder {
	builder.body.ProjectKeys = projectKeys
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) SearchUser(searchUser *SearchUser) *FilterAcrossProjectReqBuilder {
	builder.body.SearchUser = searchUser
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemTypeKey(workItemTypeKey string) *FilterAcrossProjectReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) CreatedAt(createdAt *TimeInterval) *FilterAcrossProjectReqBuilder {
	builder.body.CreatedAt = createdAt
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) UpdatedAt(updatedAt *TimeInterval) *FilterAcrossProjectReqBuilder {
	builder.body.UpdatedAt = updatedAt
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemStatus(workItemStatus []*WorkItemStatus) *FilterAcrossProjectReqBuilder {
	builder.body.WorkItemStatus = workItemStatus
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemName(workItemName string) *FilterAcrossProjectReqBuilder {
	builder.body.WorkItemName = workItemName
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) PageNum(pageNum int64) *FilterAcrossProjectReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) PageSize(pageSize int64) *FilterAcrossProjectReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) TenantGroupID(tenantGroupID int64) *FilterAcrossProjectReqBuilder {
	builder.body.TenantGroupID = tenantGroupID
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemIDs(workItemIDs []int64) *FilterAcrossProjectReqBuilder {
	builder.body.WorkItemIDs = workItemIDs
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Businesses(businesses []string) *FilterAcrossProjectReqBuilder {
	builder.body.Businesses = businesses
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Priorities(priorities []string) *FilterAcrossProjectReqBuilder {
	builder.body.Priorities = priorities
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Tags(tags []string) *FilterAcrossProjectReqBuilder {
	builder.body.Tags = tags
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) SimpleNames(simpleNames []string) *FilterAcrossProjectReqBuilder {
	builder.body.SimpleNames = simpleNames
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) TemplateIDs(templateIDs []int64) *FilterAcrossProjectReqBuilder {
	builder.body.TemplateIDs = templateIDs
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Expand(expand *Expand) *FilterAcrossProjectReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Build() *FilterAcrossProjectReq {
	req := &FilterAcrossProjectReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type GetMetaReq struct {
	apiReq *core.ApiReq
}

type GetMetaResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*FieldConf `json:"data"`
}

type GetMetaReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetMetaReqBuilder() *GetMetaReqBuilder {
	builder := &GetMetaReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *GetMetaReqBuilder) ProjectKey(projectKey string) *GetMetaReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *GetMetaReqBuilder) WorkItemTypeKey(workItemTypeKey string) *GetMetaReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *GetMetaReqBuilder) Build() *GetMetaReq {
	req := &GetMetaReq{}
	req.apiReq = builder.apiReq
	return req
}

type NodeOperateReq struct {
	apiReq *core.ApiReq
}
type NodeOperateReqBody struct {
	Action string `json:"action"`

	RollbackReason string `json:"rollback_reason"`

	NodeOwners []string `json:"node_owners"`

	NodeSchedule *Schedule `json:"node_schedule"`

	Schedules []*Schedule `json:"schedules"`

	Fields []*field.FieldValuePair `json:"fields"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`
}

type NodeOperateResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type NodeOperateReqBuilder struct {
	apiReq *core.ApiReq
	body   *NodeOperateReqBody
}

func NewNodeOperateReqBuilder() *NodeOperateReqBuilder {
	builder := &NodeOperateReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &NodeOperateReqBody{}
	return builder
}
func (builder *NodeOperateReqBuilder) ProjectKey(projectKey string) *NodeOperateReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *NodeOperateReqBuilder) WorkItemID(workItemID int64) *NodeOperateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *NodeOperateReqBuilder) NodeID(nodeID string) *NodeOperateReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}
func (builder *NodeOperateReqBuilder) WorkItemTypeKey(workItemTypeKey string) *NodeOperateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *NodeOperateReqBuilder) Action(action string) *NodeOperateReqBuilder {
	builder.body.Action = action
	return builder
}
func (builder *NodeOperateReqBuilder) RollbackReason(rollbackReason string) *NodeOperateReqBuilder {
	builder.body.RollbackReason = rollbackReason
	return builder
}
func (builder *NodeOperateReqBuilder) NodeOwners(nodeOwners []string) *NodeOperateReqBuilder {
	builder.body.NodeOwners = nodeOwners
	return builder
}
func (builder *NodeOperateReqBuilder) NodeSchedule(nodeSchedule *Schedule) *NodeOperateReqBuilder {
	builder.body.NodeSchedule = nodeSchedule
	return builder
}
func (builder *NodeOperateReqBuilder) Schedules(schedules []*Schedule) *NodeOperateReqBuilder {
	builder.body.Schedules = schedules
	return builder
}
func (builder *NodeOperateReqBuilder) Fields(fields []*field.FieldValuePair) *NodeOperateReqBuilder {
	builder.body.Fields = fields
	return builder
}
func (builder *NodeOperateReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *NodeOperateReqBuilder {
	builder.body.RoleAssignee = roleAssignee
	return builder
}
func (builder *NodeOperateReqBuilder) Build() *NodeOperateReq {
	req := &NodeOperateReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type NodeStateChangeReq struct {
	apiReq *core.ApiReq
}
type NodeStateChangeReqBody struct {
	TransitionID int64 `json:"transition_id"`

	RoleOwners []*user.RoleOwner `json:"role_owners"`

	Fields []*field.FieldValuePair `json:"fields"`
}

type NodeStateChangeResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type NodeStateChangeReqBuilder struct {
	apiReq *core.ApiReq
	body   *NodeStateChangeReqBody
}

func NewNodeStateChangeReqBuilder() *NodeStateChangeReqBuilder {
	builder := &NodeStateChangeReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &NodeStateChangeReqBody{}
	return builder
}
func (builder *NodeStateChangeReqBuilder) ProjectKey(projectKey string) *NodeStateChangeReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *NodeStateChangeReqBuilder) WorkItemID(workItemID int64) *NodeStateChangeReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *NodeStateChangeReqBuilder) WorkItemTypeKey(workItemTypeKey string) *NodeStateChangeReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *NodeStateChangeReqBuilder) TransitionID(transitionID int64) *NodeStateChangeReqBuilder {
	builder.body.TransitionID = transitionID
	return builder
}
func (builder *NodeStateChangeReqBuilder) RoleOwners(roleOwners []*user.RoleOwner) *NodeStateChangeReqBuilder {
	builder.body.RoleOwners = roleOwners
	return builder
}
func (builder *NodeStateChangeReqBuilder) Fields(fields []*field.FieldValuePair) *NodeStateChangeReqBuilder {
	builder.body.Fields = fields
	return builder
}
func (builder *NodeStateChangeReqBuilder) Build() *NodeStateChangeReq {
	req := &NodeStateChangeReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type NodeUpdateReq struct {
	apiReq *core.ApiReq
}
type NodeUpdateReqBody struct {
	NodeOwners []string `json:"node_owners"`

	NodeSchedule *Schedule `json:"node_schedule"`

	Schedules []*Schedule `json:"schedules"`

	Fields []*field.FieldValuePair `json:"fields"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`
}

type NodeUpdateResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type NodeUpdateReqBuilder struct {
	apiReq *core.ApiReq
	body   *NodeUpdateReqBody
}

func NewNodeUpdateReqBuilder() *NodeUpdateReqBuilder {
	builder := &NodeUpdateReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &NodeUpdateReqBody{}
	return builder
}
func (builder *NodeUpdateReqBuilder) ProjectKey(projectKey string) *NodeUpdateReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *NodeUpdateReqBuilder) WorkItemID(workItemID int64) *NodeUpdateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *NodeUpdateReqBuilder) NodeID(nodeID string) *NodeUpdateReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}
func (builder *NodeUpdateReqBuilder) WorkItemTypeKey(workItemTypeKey string) *NodeUpdateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *NodeUpdateReqBuilder) NodeOwners(nodeOwners []string) *NodeUpdateReqBuilder {
	builder.body.NodeOwners = nodeOwners
	return builder
}
func (builder *NodeUpdateReqBuilder) NodeSchedule(nodeSchedule *Schedule) *NodeUpdateReqBuilder {
	builder.body.NodeSchedule = nodeSchedule
	return builder
}
func (builder *NodeUpdateReqBuilder) Schedules(schedules []*Schedule) *NodeUpdateReqBuilder {
	builder.body.Schedules = schedules
	return builder
}
func (builder *NodeUpdateReqBuilder) Fields(fields []*field.FieldValuePair) *NodeUpdateReqBuilder {
	builder.body.Fields = fields
	return builder
}
func (builder *NodeUpdateReqBuilder) RoleAssignee(roleAssignee []*user.RoleOwner) *NodeUpdateReqBuilder {
	builder.body.RoleAssignee = roleAssignee
	return builder
}
func (builder *NodeUpdateReqBuilder) Build() *NodeUpdateReq {
	req := &NodeUpdateReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type QueryWorkItemDetailReq struct {
	apiReq *core.ApiReq
}
type QueryWorkItemDetailReqBody struct {
	WorkItemIDs []int64 `json:"work_item_ids"`

	Fields []string `json:"fields"`

	Expand *Expand `json:"expand"`
}

type QueryWorkItemDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*WorkItemInfo `json:"data"`
}

type QueryWorkItemDetailReqBuilder struct {
	apiReq *core.ApiReq
	body   *QueryWorkItemDetailReqBody
}

func NewQueryWorkItemDetailReqBuilder() *QueryWorkItemDetailReqBuilder {
	builder := &QueryWorkItemDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &QueryWorkItemDetailReqBody{}
	return builder
}
func (builder *QueryWorkItemDetailReqBuilder) ProjectKey(projectKey string) *QueryWorkItemDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryWorkItemDetailReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryWorkItemDetailReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *QueryWorkItemDetailReqBuilder) WorkItemIDs(workItemIDs []int64) *QueryWorkItemDetailReqBuilder {
	builder.body.WorkItemIDs = workItemIDs
	return builder
}
func (builder *QueryWorkItemDetailReqBuilder) Fields(fields []string) *QueryWorkItemDetailReqBuilder {
	builder.body.Fields = fields
	return builder
}
func (builder *QueryWorkItemDetailReqBuilder) Expand(expand *Expand) *QueryWorkItemDetailReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *QueryWorkItemDetailReqBuilder) Build() *QueryWorkItemDetailReq {
	req := &QueryWorkItemDetailReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type QueryWorkItemRelationReq struct {
	apiReq *core.ApiReq
}

type QueryWorkItemRelationResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*WorkItemRelation `json:"data"`
}

type QueryWorkItemRelationReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryWorkItemRelationReqBuilder() *QueryWorkItemRelationReqBuilder {
	builder := &QueryWorkItemRelationReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *QueryWorkItemRelationReqBuilder) ProjectKey(projectKey string) *QueryWorkItemRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryWorkItemRelationReqBuilder) Build() *QueryWorkItemRelationReq {
	req := &QueryWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkflowReq struct {
	apiReq *core.ApiReq
}
type QueryWorkflowReqBody struct {
	Fields []string `json:"fields"`

	FlowType int64 `json:"flow_type"`

	Expand *Expand `json:"expand"`
}

type QueryWorkflowResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *NodesConnections `json:"data"`
}

type QueryWorkflowReqBuilder struct {
	apiReq *core.ApiReq
	body   *QueryWorkflowReqBody
}

func NewQueryWorkflowReqBuilder() *QueryWorkflowReqBuilder {
	builder := &QueryWorkflowReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &QueryWorkflowReqBody{}
	return builder
}
func (builder *QueryWorkflowReqBuilder) ProjectKey(projectKey string) *QueryWorkflowReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryWorkflowReqBuilder) WorkItemID(workItemID int64) *QueryWorkflowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *QueryWorkflowReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryWorkflowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *QueryWorkflowReqBuilder) Fields(fields []string) *QueryWorkflowReqBuilder {
	builder.body.Fields = fields
	return builder
}
func (builder *QueryWorkflowReqBuilder) FlowType(flowType int64) *QueryWorkflowReqBuilder {
	builder.body.FlowType = flowType
	return builder
}
func (builder *QueryWorkflowReqBuilder) Expand(expand *Expand) *QueryWorkflowReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *QueryWorkflowReqBuilder) Build() *QueryWorkflowReq {
	req := &QueryWorkflowReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type SearchByParamsReq struct {
	apiReq *core.ApiReq
}
type SearchByParamsReqBody struct {
	SearchGroup *SearchGroup `json:"search_group"`

	PageNum int64 `json:"page_num"`

	PageSize int64 `json:"page_size"`

	Fields []string `json:"fields"`

	Expand *Expand `json:"expand"`
}

type SearchByParamsResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*WorkItemInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type SearchByParamsReqBuilder struct {
	apiReq *core.ApiReq
	body   *SearchByParamsReqBody
}

func NewSearchByParamsReqBuilder() *SearchByParamsReqBuilder {
	builder := &SearchByParamsReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &SearchByParamsReqBody{}
	return builder
}
func (builder *SearchByParamsReqBuilder) ProjectKey(projectKey string) *SearchByParamsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *SearchByParamsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *SearchByParamsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *SearchByParamsReqBuilder) SearchGroup(searchGroup *SearchGroup) *SearchByParamsReqBuilder {
	builder.body.SearchGroup = searchGroup
	return builder
}
func (builder *SearchByParamsReqBuilder) PageNum(pageNum int64) *SearchByParamsReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *SearchByParamsReqBuilder) PageSize(pageSize int64) *SearchByParamsReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *SearchByParamsReqBuilder) Fields(fields []string) *SearchByParamsReqBuilder {
	builder.body.Fields = fields
	return builder
}
func (builder *SearchByParamsReqBuilder) Expand(expand *Expand) *SearchByParamsReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *SearchByParamsReqBuilder) Build() *SearchByParamsReq {
	req := &SearchByParamsReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type SearchByRelationReq struct {
	apiReq *core.ApiReq
}
type SearchByRelationReqBody struct {
	RelationWorkItemTypeKey string `json:"relation_work_item_type_key"`

	RelationKey string `json:"relation_key"`

	PageNum int64 `json:"page_num"`

	PageSize int64 `json:"page_size"`

	RelationType int32 `json:"relation_type"`

	Expand *Expand `json:"expand"`
}

type SearchByRelationResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*WorkItemInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type SearchByRelationReqBuilder struct {
	apiReq *core.ApiReq
	body   *SearchByRelationReqBody
}

func NewSearchByRelationReqBuilder() *SearchByRelationReqBuilder {
	builder := &SearchByRelationReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &SearchByRelationReqBody{}
	return builder
}
func (builder *SearchByRelationReqBuilder) ProjectKey(projectKey string) *SearchByRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *SearchByRelationReqBuilder) WorkItemTypeKey(workItemTypeKey string) *SearchByRelationReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *SearchByRelationReqBuilder) WorkItemID(workItemID int64) *SearchByRelationReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *SearchByRelationReqBuilder) RelationWorkItemTypeKey(relationWorkItemTypeKey string) *SearchByRelationReqBuilder {
	builder.body.RelationWorkItemTypeKey = relationWorkItemTypeKey
	return builder
}
func (builder *SearchByRelationReqBuilder) RelationKey(relationKey string) *SearchByRelationReqBuilder {
	builder.body.RelationKey = relationKey
	return builder
}
func (builder *SearchByRelationReqBuilder) PageNum(pageNum int64) *SearchByRelationReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *SearchByRelationReqBuilder) PageSize(pageSize int64) *SearchByRelationReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *SearchByRelationReqBuilder) RelationType(relationType int32) *SearchByRelationReqBuilder {
	builder.body.RelationType = relationType
	return builder
}
func (builder *SearchByRelationReqBuilder) Expand(expand *Expand) *SearchByRelationReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *SearchByRelationReqBuilder) Build() *SearchByRelationReq {
	req := &SearchByRelationReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type UpdateMultiSignalReq struct {
	apiReq *core.ApiReq
}
type UpdateMultiSignalReqBody struct {
	FieldKey string `json:"field_key"`

	FieldAlias string `json:"field_alias"`

	Details []*field.MultiSignalDetail `json:"details"`

	UpdateType string `json:"update_type"`
}

type UpdateMultiSignalResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *field.MultiSignal `json:"data"`
}

type UpdateMultiSignalReqBuilder struct {
	apiReq *core.ApiReq
	body   *UpdateMultiSignalReqBody
}

func NewUpdateMultiSignalReqBuilder() *UpdateMultiSignalReqBuilder {
	builder := &UpdateMultiSignalReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateMultiSignalReqBody{}
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) ProjectKey(projectKey string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) WorkItemID(workItemID int64) *UpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) FieldKey(fieldKey string) *UpdateMultiSignalReqBuilder {
	builder.body.FieldKey = fieldKey
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) FieldAlias(fieldAlias string) *UpdateMultiSignalReqBuilder {
	builder.body.FieldAlias = fieldAlias
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) Details(details []*field.MultiSignalDetail) *UpdateMultiSignalReqBuilder {
	builder.body.Details = details
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) UpdateType(updateType string) *UpdateMultiSignalReqBuilder {
	builder.body.UpdateType = updateType
	return builder
}
func (builder *UpdateMultiSignalReqBuilder) Build() *UpdateMultiSignalReq {
	req := &UpdateMultiSignalReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type UpdateWorkItemReq struct {
	apiReq *core.ApiReq
}
type UpdateWorkItemReqBody struct {
	UpdateFields []*field.FieldValuePair `json:"update_fields"`
}

type UpdateWorkItemResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type UpdateWorkItemReqBuilder struct {
	apiReq *core.ApiReq
	body   *UpdateWorkItemReqBody
}

func NewUpdateWorkItemReqBuilder() *UpdateWorkItemReqBuilder {
	builder := &UpdateWorkItemReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateWorkItemReqBody{}
	return builder
}
func (builder *UpdateWorkItemReqBuilder) ProjectKey(projectKey string) *UpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *UpdateWorkItemReqBuilder) WorkItemID(workItemID int64) *UpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *UpdateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *UpdateWorkItemReqBuilder) UpdateFields(updateFields []*field.FieldValuePair) *UpdateWorkItemReqBuilder {
	builder.body.UpdateFields = updateFields
	return builder
}
func (builder *UpdateWorkItemReqBuilder) Build() *UpdateWorkItemReq {
	req := &UpdateWorkItemReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type UpdateWorkItemRelationReq struct {
	apiReq *core.ApiReq
}
type UpdateWorkItemRelationReqBody struct {
	RelationID string `json:"relation_id"`

	ProjectKey string `json:"project_key"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	Name string `json:"name"`

	RelationDetails []*RelationDetail `json:"relation_details"`
}

type UpdateWorkItemRelationResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data string `json:"data"`
}

type UpdateWorkItemRelationReqBuilder struct {
	apiReq *core.ApiReq
	body   *UpdateWorkItemRelationReqBody
}

func NewUpdateWorkItemRelationReqBuilder() *UpdateWorkItemRelationReqBuilder {
	builder := &UpdateWorkItemRelationReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateWorkItemRelationReqBody{}
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) RelationID(relationID string) *UpdateWorkItemRelationReqBuilder {
	builder.body.RelationID = relationID
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) ProjectKey(projectKey string) *UpdateWorkItemRelationReqBuilder {
	builder.body.ProjectKey = projectKey
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkItemRelationReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) Name(name string) *UpdateWorkItemRelationReqBuilder {
	builder.body.Name = name
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) RelationDetails(relationDetails []*RelationDetail) *UpdateWorkItemRelationReqBuilder {
	builder.body.RelationDetails = relationDetails
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) Build() *UpdateWorkItemRelationReq {
	req := &UpdateWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type WbsViewReq struct {
	apiReq *core.ApiReq
}
type WbsViewReqBody struct {
	Expand *Expand `json:"expand"`
}

type WbsViewResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *WbsViewResponse `json:"data"`
}

type WbsViewReqBuilder struct {
	apiReq *core.ApiReq
	body   *WbsViewReqBody
}

func NewWbsViewReqBuilder() *WbsViewReqBuilder {
	builder := &WbsViewReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &WbsViewReqBody{}
	return builder
}
func (builder *WbsViewReqBuilder) ProjectKey(projectKey string) *WbsViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *WbsViewReqBuilder) WorkItemID(workItemID int64) *WbsViewReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *WbsViewReqBuilder) WorkItemTypeKey(workItemTypeKey string) *WbsViewReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *WbsViewReqBuilder) Expand(expand *Expand) *WbsViewReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *WbsViewReqBuilder) Build() *WbsViewReq {
	req := &WbsViewReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
