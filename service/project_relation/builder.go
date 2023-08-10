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

package project_relation

import (
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type CreateProjectRelationInstancesReq struct {
	apiReq *core.ApiReq
}
type CreateProjectRelationInstancesReqBody struct {
	RelationRuleID string `json:"relation_rule_id"`

	Instances []*RelationBindInstance `json:"instances"`
}

type CreateProjectRelationInstancesResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type CreateProjectRelationInstancesReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateProjectRelationInstancesReqBuilder() *CreateProjectRelationInstancesReqBuilder {
	builder := &CreateProjectRelationInstancesReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &CreateProjectRelationInstancesReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *CreateProjectRelationInstancesReqBuilder) AccessUser(userKey string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *CreateProjectRelationInstancesReqBuilder) UUID(uuid string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) ProjectKey(projectKey string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) WorkItemID(workItemID int64) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) RelationRuleID(relationRuleID string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Body.(*CreateProjectRelationInstancesReqBody).RelationRuleID = relationRuleID
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) Instances(instances []*RelationBindInstance) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Body.(*CreateProjectRelationInstancesReqBody).Instances = instances
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) Build() *CreateProjectRelationInstancesReq {
	req := &CreateProjectRelationInstancesReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteProjectRelationInstanceReq struct {
	apiReq *core.ApiReq
}
type DeleteProjectRelationInstanceReqBody struct {
	RelationRuleID string `json:"relation_rule_id"`

	RelationWorkItemID int64 `json:"relation_work_item_id"`
}

type DeleteProjectRelationInstanceResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type DeleteProjectRelationInstanceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewDeleteProjectRelationInstanceReqBuilder() *DeleteProjectRelationInstanceReqBuilder {
	builder := &DeleteProjectRelationInstanceReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &DeleteProjectRelationInstanceReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *DeleteProjectRelationInstanceReqBuilder) AccessUser(userKey string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *DeleteProjectRelationInstanceReqBuilder) UUID(uuid string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *DeleteProjectRelationInstanceReqBuilder) ProjectKey(projectKey string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *DeleteProjectRelationInstanceReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *DeleteProjectRelationInstanceReqBuilder) WorkItemID(workItemID int64) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *DeleteProjectRelationInstanceReqBuilder) RelationRuleID(relationRuleID string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*DeleteProjectRelationInstanceReqBody).RelationRuleID = relationRuleID
	return builder
}
func (builder *DeleteProjectRelationInstanceReqBuilder) RelationWorkItemID(relationWorkItemID int64) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*DeleteProjectRelationInstanceReqBody).RelationWorkItemID = relationWorkItemID
	return builder
}
func (builder *DeleteProjectRelationInstanceReqBuilder) Build() *DeleteProjectRelationInstanceReq {
	req := &DeleteProjectRelationInstanceReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectRelationReq struct {
	apiReq *core.ApiReq
}
type QueryProjectRelationReqBody struct {
	RemoteProjects []string `json:"remote_projects"`
}

type QueryProjectRelationResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*ProjectRelationRule `json:"data"`
}

type QueryProjectRelationReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryProjectRelationReqBuilder() *QueryProjectRelationReqBuilder {
	builder := &QueryProjectRelationReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &QueryProjectRelationReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryProjectRelationReqBuilder) AccessUser(userKey string) *QueryProjectRelationReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryProjectRelationReqBuilder) UUID(uuid string) *QueryProjectRelationReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *QueryProjectRelationReqBuilder) ProjectKey(projectKey string) *QueryProjectRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryProjectRelationReqBuilder) RemoteProjects(remoteProjects []string) *QueryProjectRelationReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationReqBody).RemoteProjects = remoteProjects
	return builder
}
func (builder *QueryProjectRelationReqBuilder) Build() *QueryProjectRelationReq {
	req := &QueryProjectRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectRelationInstanceReq struct {
	apiReq *core.ApiReq
}
type QueryProjectRelationInstanceReqBody struct {
	RelationRuleID string `json:"relation_rule_id"`

	RelationWorkItemID int64 `json:"relation_work_item_id"`

	RelationWorkItemTypeKey string `json:"relation_work_item_type_key"`

	RelationProjectKey string `json:"relation_project_key"`
}

type QueryProjectRelationInstanceResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*RelationInstance `json:"data"`
}

type QueryProjectRelationInstanceReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryProjectRelationInstanceReqBuilder() *QueryProjectRelationInstanceReqBuilder {
	builder := &QueryProjectRelationInstanceReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &QueryProjectRelationInstanceReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryProjectRelationInstanceReqBuilder) AccessUser(userKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryProjectRelationInstanceReqBuilder) UUID(uuid string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) ProjectKey(projectKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) WorkItemID(workItemID int64) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) RelationRuleID(relationRuleID string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationRuleID = relationRuleID
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) RelationWorkItemID(relationWorkItemID int64) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationWorkItemID = relationWorkItemID
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) RelationWorkItemTypeKey(relationWorkItemTypeKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationWorkItemTypeKey = relationWorkItemTypeKey
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) RelationProjectKey(relationProjectKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationProjectKey = relationProjectKey
	return builder
}
func (builder *QueryProjectRelationInstanceReqBuilder) Build() *QueryProjectRelationInstanceReq {
	req := &QueryProjectRelationInstanceReq{}
	req.apiReq = builder.apiReq
	return req
}
