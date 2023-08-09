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
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"

	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem/v2"
)

type GetProjectDetailReq struct {
	apiReq *core.ApiReq
}
type GetProjectDetailReqBody struct {
	ProjectKeys []string `json:"project_keys"`

	UserKey string `json:"user_key"`

	SimpleNames []string `json:"simple_names"`

	TenantGroupID int64 `json:"tenant_group_id"`
}

type GetProjectDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data map[string]*Project `json:"data"`
}

type GetProjectDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewGetProjectDetailReqBuilder() *GetProjectDetailReqBuilder {
	builder := &GetProjectDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &GetProjectDetailReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *GetProjectDetailReqBuilder) AccessUser(userKey string) *GetProjectDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *GetProjectDetailReqBuilder) UUID(uuid string) *GetProjectDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *GetProjectDetailReqBuilder) ProjectKeys(projectKeys []string) *GetProjectDetailReqBuilder {
	builder.apiReq.Body.(*GetProjectDetailReqBody).ProjectKeys = projectKeys
	return builder
}
func (builder *GetProjectDetailReqBuilder) UserKey(userKey string) *GetProjectDetailReqBuilder {
	builder.apiReq.Body.(*GetProjectDetailReqBody).UserKey = userKey
	return builder
}
func (builder *GetProjectDetailReqBuilder) SimpleNames(simpleNames []string) *GetProjectDetailReqBuilder {
	builder.apiReq.Body.(*GetProjectDetailReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *GetProjectDetailReqBuilder) TenantGroupID(tenantGroupID int64) *GetProjectDetailReqBuilder {
	builder.apiReq.Body.(*GetProjectDetailReqBody).TenantGroupID = tenantGroupID
	return builder
}
func (builder *GetProjectDetailReqBuilder) Build() *GetProjectDetailReq {
	req := &GetProjectDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListProjectReq struct {
	apiReq *core.ApiReq
}
type ListProjectReqBody struct {
	UserKey string `json:"user_key"`

	TenantGroupID int64 `json:"tenant_group_id"`

	AssetKey string `json:"asset_key"`
}

type ListProjectResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []string `json:"data"`
}

type ListProjectReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListProjectReqBuilder() *ListProjectReqBuilder {
	builder := &ListProjectReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &ListProjectReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *ListProjectReqBuilder) AccessUser(userKey string) *ListProjectReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *ListProjectReqBuilder) UUID(uuid string) *ListProjectReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *ListProjectReqBuilder) UserKey(userKey string) *ListProjectReqBuilder {
	builder.apiReq.Body.(*ListProjectReqBody).UserKey = userKey
	return builder
}
func (builder *ListProjectReqBuilder) TenantGroupID(tenantGroupID int64) *ListProjectReqBuilder {
	builder.apiReq.Body.(*ListProjectReqBody).TenantGroupID = tenantGroupID
	return builder
}
func (builder *ListProjectReqBuilder) AssetKey(assetKey string) *ListProjectReqBuilder {
	builder.apiReq.Body.(*ListProjectReqBody).AssetKey = assetKey
	return builder
}
func (builder *ListProjectReqBuilder) Build() *ListProjectReq {
	req := &ListProjectReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListProjectBusinessReq struct {
	apiReq *core.ApiReq
}

type ListProjectBusinessResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*Business `json:"data"`
}

type ListProjectBusinessReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListProjectBusinessReqBuilder() *ListProjectBusinessReqBuilder {
	builder := &ListProjectBusinessReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *ListProjectBusinessReqBuilder) AccessUser(userKey string) *ListProjectBusinessReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *ListProjectBusinessReqBuilder) UUID(uuid string) *ListProjectBusinessReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	apiReq *core.ApiReq
}

type ListProjectTeamResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*Team `json:"data"`
}

type ListProjectTeamReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListProjectTeamReqBuilder() *ListProjectTeamReqBuilder {
	builder := &ListProjectTeamReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *ListProjectTeamReqBuilder) AccessUser(userKey string) *ListProjectTeamReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *ListProjectTeamReqBuilder) UUID(uuid string) *ListProjectTeamReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	apiReq *core.ApiReq
}

type ListProjectWorkItemTypeResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*workitem.WorkItemKeyType `json:"data"`
}

type ListProjectWorkItemTypeReqBuilder struct {
	apiReq *core.ApiReq
}

func NewListProjectWorkItemTypeReqBuilder() *ListProjectWorkItemTypeReqBuilder {
	builder := &ListProjectWorkItemTypeReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *ListProjectWorkItemTypeReqBuilder) AccessUser(userKey string) *ListProjectWorkItemTypeReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *ListProjectWorkItemTypeReqBuilder) UUID(uuid string) *ListProjectWorkItemTypeReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
