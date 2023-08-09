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
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"

	"code.byted.org/bits/project-oapi-sdk-golang/service/common"

	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem/v2"
)

type CreateFixViewReq struct {
	apiReq *core.ApiReq
}
type CreateFixViewReqBody struct {
	WorkItemIDList []int64 `json:"work_item_id_list"`

	Name string `json:"name"`
}

type CreateFixViewResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *FixView `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type CreateFixViewReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateFixViewReqBuilder() *CreateFixViewReqBuilder {
	builder := &CreateFixViewReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &CreateFixViewReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *CreateFixViewReqBuilder) AccessUser(userKey string) *CreateFixViewReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *CreateFixViewReqBuilder) UUID(uuid string) *CreateFixViewReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	builder.apiReq.Body.(*CreateFixViewReqBody).WorkItemIDList = workItemIDList
	return builder
}
func (builder *CreateFixViewReqBuilder) Name(name string) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).Name = name
	return builder
}
func (builder *CreateFixViewReqBuilder) Build() *CreateFixViewReq {
	req := &CreateFixViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteFixViewReq struct {
	apiReq *core.ApiReq
}

type DeleteFixViewResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type DeleteFixViewReqBuilder struct {
	apiReq *core.ApiReq
}

func NewDeleteFixViewReqBuilder() *DeleteFixViewReqBuilder {
	builder := &DeleteFixViewReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *DeleteFixViewReqBuilder) AccessUser(userKey string) *DeleteFixViewReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *DeleteFixViewReqBuilder) UUID(uuid string) *DeleteFixViewReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	apiReq *core.ApiReq
}
type QueryWorkItemDetailsByViewIDReqBody struct {
	PageSize int64 `json:"page_size"`

	PageNum int64 `json:"page_num"`

	Expand *workitem.Expand `json:"expand"`
}

type QueryWorkItemDetailsByViewIDResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*workitem.WorkItemInfo `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type QueryWorkItemDetailsByViewIDReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryWorkItemDetailsByViewIDReqBuilder() *QueryWorkItemDetailsByViewIDReqBuilder {
	builder := &QueryWorkItemDetailsByViewIDReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &QueryWorkItemDetailsByViewIDReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) AccessUser(userKey string) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) UUID(uuid string) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	builder.apiReq.Body.(*QueryWorkItemDetailsByViewIDReqBody).PageSize = pageSize
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) PageNum(pageNum int64) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Body.(*QueryWorkItemDetailsByViewIDReqBody).PageNum = pageNum
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) Expand(expand *workitem.Expand) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Body.(*QueryWorkItemDetailsByViewIDReqBody).Expand = expand
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) Build() *QueryWorkItemDetailsByViewIDReq {
	req := &QueryWorkItemDetailsByViewIDReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateFixViewReq struct {
	apiReq *core.ApiReq
}
type UpdateFixViewReqBody struct {
	AddWorkItemIDs []int64 `json:"add_work_item_ids"`

	RemoveWorkItemIDs []int64 `json:"remove_work_item_ids"`
}

type UpdateFixViewResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *FixView `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type UpdateFixViewReqBuilder struct {
	apiReq *core.ApiReq
}

func NewUpdateFixViewReqBuilder() *UpdateFixViewReqBuilder {
	builder := &UpdateFixViewReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &UpdateFixViewReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *UpdateFixViewReqBuilder) AccessUser(userKey string) *UpdateFixViewReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *UpdateFixViewReqBuilder) UUID(uuid string) *UpdateFixViewReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	builder.apiReq.Body.(*UpdateFixViewReqBody).AddWorkItemIDs = addWorkItemIDs
	return builder
}
func (builder *UpdateFixViewReqBuilder) RemoveWorkItemIDs(removeWorkItemIDs []int64) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).RemoveWorkItemIDs = removeWorkItemIDs
	return builder
}
func (builder *UpdateFixViewReqBuilder) Build() *UpdateFixViewReq {
	req := &UpdateFixViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type ViewListReq struct {
	apiReq *core.ApiReq
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
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*ViewConf `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type ViewListReqBuilder struct {
	apiReq *core.ApiReq
}

func NewViewListReqBuilder() *ViewListReqBuilder {
	builder := &ViewListReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &ViewListReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *ViewListReqBuilder) AccessUser(userKey string) *ViewListReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *ViewListReqBuilder) UUID(uuid string) *ViewListReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *ViewListReqBuilder) ProjectKey(projectKey string) *ViewListReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *ViewListReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *ViewListReqBuilder) ViewIDs(viewIDs []string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).ViewIDs = viewIDs
	return builder
}
func (builder *ViewListReqBuilder) CreatedBy(createdBy string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).CreatedBy = createdBy
	return builder
}
func (builder *ViewListReqBuilder) CreatedAt(createdAt *workitem.TimeInterval) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).CreatedAt = createdAt
	return builder
}
func (builder *ViewListReqBuilder) PageSize(pageSize int64) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).PageSize = pageSize
	return builder
}
func (builder *ViewListReqBuilder) PageNum(pageNum int64) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).PageNum = pageNum
	return builder
}
func (builder *ViewListReqBuilder) Build() *ViewListReq {
	req := &ViewListReq{}
	req.apiReq = builder.apiReq
	return req
}

type WorkItemListReq struct {
	apiReq *core.ApiReq
}

type WorkItemListResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *FixView `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type WorkItemListReqBuilder struct {
	apiReq *core.ApiReq
}

func NewWorkItemListReqBuilder() *WorkItemListReqBuilder {
	builder := &WorkItemListReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *WorkItemListReqBuilder) AccessUser(userKey string) *WorkItemListReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *WorkItemListReqBuilder) UUID(uuid string) *WorkItemListReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
