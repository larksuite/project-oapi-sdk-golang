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

package project

import (
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"

	"github.com/larksuite/project-oapi-sdk-golang/service/workitem"
)

type GetProjectDetailReq struct {
	apiReq *core.APIReq
}
type GetProjectDetailReqBody struct {
	ProjectKeys []string `json:"project_keys"`

	UserKey string `json:"user_key"`

	SimpleNames []string `json:"simple_names"`

	TenantGroupID int64 `json:"tenant_group_id"`
}

type GetProjectDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data map[string]*Project `json:"data"`
}

type GetProjectDetailReqBuilder struct {
	apiReq *core.APIReq
	body   *GetProjectDetailReqBody
}

func NewGetProjectDetailReqBuilder() *GetProjectDetailReqBuilder {
	builder := &GetProjectDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &GetProjectDetailReqBody{}
	return builder
}
func (builder *GetProjectDetailReqBuilder) ProjectKeys(projectKeys []string) *GetProjectDetailReqBuilder {
	builder.body.ProjectKeys = projectKeys
	return builder
}
func (builder *GetProjectDetailReqBuilder) UserKey(userKey string) *GetProjectDetailReqBuilder {
	builder.body.UserKey = userKey
	return builder
}
func (builder *GetProjectDetailReqBuilder) SimpleNames(simpleNames []string) *GetProjectDetailReqBuilder {
	builder.body.SimpleNames = simpleNames
	return builder
}
func (builder *GetProjectDetailReqBuilder) TenantGroupID(tenantGroupID int64) *GetProjectDetailReqBuilder {
	builder.body.TenantGroupID = tenantGroupID
	return builder
}
func (builder *GetProjectDetailReqBuilder) Build() *GetProjectDetailReq {
	req := &GetProjectDetailReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type ListProjectReq struct {
	apiReq *core.APIReq
}
type ListProjectReqBody struct {
	UserKey string `json:"user_key"`

	TenantGroupID int64 `json:"tenant_group_id"`

	AssetKey string `json:"asset_key"`

	Order []string `json:"order"`
}

type ListProjectResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []string `json:"data"`
}

type ListProjectReqBuilder struct {
	apiReq *core.APIReq
	body   *ListProjectReqBody
}

func NewListProjectReqBuilder() *ListProjectReqBuilder {
	builder := &ListProjectReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &ListProjectReqBody{}
	return builder
}
func (builder *ListProjectReqBuilder) UserKey(userKey string) *ListProjectReqBuilder {
	builder.body.UserKey = userKey
	return builder
}
func (builder *ListProjectReqBuilder) TenantGroupID(tenantGroupID int64) *ListProjectReqBuilder {
	builder.body.TenantGroupID = tenantGroupID
	return builder
}
func (builder *ListProjectReqBuilder) AssetKey(assetKey string) *ListProjectReqBuilder {
	builder.body.AssetKey = assetKey
	return builder
}
func (builder *ListProjectReqBuilder) Order(order []string) *ListProjectReqBuilder {
	builder.body.Order = order
	return builder
}
func (builder *ListProjectReqBuilder) Build() *ListProjectReq {
	req := &ListProjectReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type ListProjectBusinessReq struct {
	apiReq *core.APIReq
}

type ListProjectBusinessResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Business `json:"data"`
}

type ListProjectBusinessReqBuilder struct {
	apiReq *core.APIReq
}

func NewListProjectBusinessReqBuilder() *ListProjectBusinessReqBuilder {
	builder := &ListProjectBusinessReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *ListProjectBusinessReqBuilder) ProjectKey(projectKey string) *ListProjectBusinessReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *ListProjectBusinessReqBuilder) Build() *ListProjectBusinessReq {
	req := &ListProjectBusinessReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListProjectTeamReq struct {
	apiReq *core.APIReq
}

type ListProjectTeamResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Team `json:"data"`
}

type ListProjectTeamReqBuilder struct {
	apiReq *core.APIReq
}

func NewListProjectTeamReqBuilder() *ListProjectTeamReqBuilder {
	builder := &ListProjectTeamReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *ListProjectTeamReqBuilder) ProjectKey(projectKey string) *ListProjectTeamReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *ListProjectTeamReqBuilder) Build() *ListProjectTeamReq {
	req := &ListProjectTeamReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListProjectWorkItemTypeReq struct {
	apiReq *core.APIReq
}

type ListProjectWorkItemTypeResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*workitem.WorkItemKeyType `json:"data"`
}

type ListProjectWorkItemTypeReqBuilder struct {
	apiReq *core.APIReq
}

func NewListProjectWorkItemTypeReqBuilder() *ListProjectWorkItemTypeReqBuilder {
	builder := &ListProjectWorkItemTypeReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *ListProjectWorkItemTypeReqBuilder) ProjectKey(projectKey string) *ListProjectWorkItemTypeReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *ListProjectWorkItemTypeReqBuilder) Build() *ListProjectWorkItemTypeReq {
	req := &ListProjectWorkItemTypeReq{}
	req.apiReq = builder.apiReq
	return req
}
