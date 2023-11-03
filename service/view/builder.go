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

package view

import (
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"

	"github.com/larksuite/project-oapi-sdk-golang/service/common"

	"github.com/larksuite/project-oapi-sdk-golang/service/workitem"
)

type CreateFixViewReq struct {
	apiReq *core.APIReq
}
type CreateFixViewReqBody struct {
	WorkItemIDList []int64 `json:"work_item_id_list"`

	Name string `json:"name"`

	CooperationMode int64 `json:"cooperation_mode"`

	CooperationUserKeys []string `json:"cooperation_user_keys"`

	CooperationTeamIDs []int64 `json:"cooperation_team_ids"`
}

type CreateFixViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *FixView `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type CreateFixViewReqBuilder struct {
	apiReq *core.APIReq
	body   *CreateFixViewReqBody
}

func NewCreateFixViewReqBuilder() *CreateFixViewReqBuilder {
	builder := &CreateFixViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateFixViewReqBody{}
	return builder
}
func (builder *CreateFixViewReqBuilder) ProjectKey(projectKey string) *CreateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *CreateFixViewReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *CreateFixViewReqBuilder) WorkItemIDList(workItemIDList []int64) *CreateFixViewReqBuilder {
	builder.body.WorkItemIDList = workItemIDList
	return builder
}
func (builder *CreateFixViewReqBuilder) Name(name string) *CreateFixViewReqBuilder {
	builder.body.Name = name
	return builder
}
func (builder *CreateFixViewReqBuilder) CooperationMode(cooperationMode int64) *CreateFixViewReqBuilder {
	builder.body.CooperationMode = cooperationMode
	return builder
}
func (builder *CreateFixViewReqBuilder) CooperationUserKeys(cooperationUserKeys []string) *CreateFixViewReqBuilder {
	builder.body.CooperationUserKeys = cooperationUserKeys
	return builder
}
func (builder *CreateFixViewReqBuilder) CooperationTeamIDs(cooperationTeamIDs []int64) *CreateFixViewReqBuilder {
	builder.body.CooperationTeamIDs = cooperationTeamIDs
	return builder
}
func (builder *CreateFixViewReqBuilder) Build() *CreateFixViewReq {
	req := &CreateFixViewReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type DeleteFixViewReq struct {
	apiReq *core.APIReq
}

type DeleteFixViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteFixViewReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteFixViewReqBuilder() *DeleteFixViewReqBuilder {
	builder := &DeleteFixViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *DeleteFixViewReqBuilder) ProjectKey(projectKey string) *DeleteFixViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *DeleteFixViewReqBuilder) ViewID(viewID string) *DeleteFixViewReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}
func (builder *DeleteFixViewReqBuilder) Build() *DeleteFixViewReq {
	req := &DeleteFixViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkItemDetailsByViewIDReq struct {
	apiReq *core.APIReq
}
type QueryWorkItemDetailsByViewIDReqBody struct {
	PageSize int64 `json:"page_size"`

	PageNum int64 `json:"page_num"`

	Expand *workitem.Expand `json:"expand"`
}

type QueryWorkItemDetailsByViewIDResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*workitem.WorkItemInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type QueryWorkItemDetailsByViewIDReqBuilder struct {
	apiReq *core.APIReq
	body   *QueryWorkItemDetailsByViewIDReqBody
}

func NewQueryWorkItemDetailsByViewIDReqBuilder() *QueryWorkItemDetailsByViewIDReqBuilder {
	builder := &QueryWorkItemDetailsByViewIDReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &QueryWorkItemDetailsByViewIDReqBody{}
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) ProjectKey(projectKey string) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) ViewID(viewID string) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) PageSize(pageSize int64) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) PageNum(pageNum int64) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) Expand(expand *workitem.Expand) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.body.Expand = expand
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) Build() *QueryWorkItemDetailsByViewIDReq {
	req := &QueryWorkItemDetailsByViewIDReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type UpdateFixViewReq struct {
	apiReq *core.APIReq
}
type UpdateFixViewReqBody struct {
	AddWorkItemIDs []int64 `json:"add_work_item_ids"`

	RemoveWorkItemIDs []int64 `json:"remove_work_item_ids"`

	CooperationMode int64 `json:"cooperation_mode"`

	CooperationUserKeys []string `json:"cooperation_user_keys"`

	CooperationTeamIDs []int64 `json:"cooperation_team_ids"`
}

type UpdateFixViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *FixView `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type UpdateFixViewReqBuilder struct {
	apiReq *core.APIReq
	body   *UpdateFixViewReqBody
}

func NewUpdateFixViewReqBuilder() *UpdateFixViewReqBuilder {
	builder := &UpdateFixViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateFixViewReqBody{}
	return builder
}
func (builder *UpdateFixViewReqBuilder) ProjectKey(projectKey string) *UpdateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *UpdateFixViewReqBuilder) ViewID(viewID string) *UpdateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}
func (builder *UpdateFixViewReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *UpdateFixViewReqBuilder) AddWorkItemIDs(addWorkItemIDs []int64) *UpdateFixViewReqBuilder {
	builder.body.AddWorkItemIDs = addWorkItemIDs
	return builder
}
func (builder *UpdateFixViewReqBuilder) RemoveWorkItemIDs(removeWorkItemIDs []int64) *UpdateFixViewReqBuilder {
	builder.body.RemoveWorkItemIDs = removeWorkItemIDs
	return builder
}
func (builder *UpdateFixViewReqBuilder) CooperationMode(cooperationMode int64) *UpdateFixViewReqBuilder {
	builder.body.CooperationMode = cooperationMode
	return builder
}
func (builder *UpdateFixViewReqBuilder) CooperationUserKeys(cooperationUserKeys []string) *UpdateFixViewReqBuilder {
	builder.body.CooperationUserKeys = cooperationUserKeys
	return builder
}
func (builder *UpdateFixViewReqBuilder) CooperationTeamIDs(cooperationTeamIDs []int64) *UpdateFixViewReqBuilder {
	builder.body.CooperationTeamIDs = cooperationTeamIDs
	return builder
}
func (builder *UpdateFixViewReqBuilder) Build() *UpdateFixViewReq {
	req := &UpdateFixViewReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type ViewListReq struct {
	apiReq *core.APIReq
}
type ViewListReqBody struct {
	WorkItemTypeKey string `json:"work_item_type_key"`

	ViewIDs []string `json:"view_ids"`

	CreatedBy string `json:"created_by"`

	CreatedAt *workitem.TimeInterval `json:"created_at"`

	PageSize int64 `json:"page_size"`

	PageNum int64 `json:"page_num"`
}

type ViewListResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*ViewConf `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type ViewListReqBuilder struct {
	apiReq *core.APIReq
	body   *ViewListReqBody
}

func NewViewListReqBuilder() *ViewListReqBuilder {
	builder := &ViewListReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &ViewListReqBody{}
	return builder
}
func (builder *ViewListReqBuilder) ProjectKey(projectKey string) *ViewListReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *ViewListReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ViewListReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *ViewListReqBuilder) ViewIDs(viewIDs []string) *ViewListReqBuilder {
	builder.body.ViewIDs = viewIDs
	return builder
}
func (builder *ViewListReqBuilder) CreatedBy(createdBy string) *ViewListReqBuilder {
	builder.body.CreatedBy = createdBy
	return builder
}
func (builder *ViewListReqBuilder) CreatedAt(createdAt *workitem.TimeInterval) *ViewListReqBuilder {
	builder.body.CreatedAt = createdAt
	return builder
}
func (builder *ViewListReqBuilder) PageSize(pageSize int64) *ViewListReqBuilder {
	builder.body.PageSize = pageSize
	return builder
}
func (builder *ViewListReqBuilder) PageNum(pageNum int64) *ViewListReqBuilder {
	builder.body.PageNum = pageNum
	return builder
}
func (builder *ViewListReqBuilder) Build() *ViewListReq {
	req := &ViewListReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type WorkItemListReq struct {
	apiReq *core.APIReq
}

type WorkItemListResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *FixView `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type WorkItemListReqBuilder struct {
	apiReq *core.APIReq
}

func NewWorkItemListReqBuilder() *WorkItemListReqBuilder {
	builder := &WorkItemListReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *WorkItemListReqBuilder) ProjectKey(projectKey string) *WorkItemListReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *WorkItemListReqBuilder) ViewID(viewID string) *WorkItemListReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}
func (builder *WorkItemListReqBuilder) PageSize(pageSize int64) *WorkItemListReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}
func (builder *WorkItemListReqBuilder) PageNum(pageNum int64) *WorkItemListReqBuilder {
	builder.apiReq.QueryParams.Set("page_num", fmt.Sprint(pageNum))
	return builder
}
func (builder *WorkItemListReqBuilder) Build() *WorkItemListReq {
	req := &WorkItemListReq{}
	req.apiReq = builder.apiReq
	return req
}
