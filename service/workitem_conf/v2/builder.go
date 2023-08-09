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

package workitem_conf

import (
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type CreateTemplateDetailReq struct {
	apiReq *core.ApiReq
}
type CreateTemplateDetailReqBody struct {
	ProjectKey string `json:"project_key"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	TemplateName string `json:"template_name"`

	CopyTemplateID int64 `json:"copy_template_id"`
}

type CreateTemplateDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data int64 `json:"data"`
}

type CreateTemplateDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateTemplateDetailReqBuilder() *CreateTemplateDetailReqBuilder {
	builder := &CreateTemplateDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &CreateTemplateDetailReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *CreateTemplateDetailReqBuilder) AccessUser(userKey string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *CreateTemplateDetailReqBuilder) UUID(uuid string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) ProjectKey(projectKey string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).ProjectKey = projectKey
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) TemplateName(templateName string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).TemplateName = templateName
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) CopyTemplateID(copyTemplateID int64) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).CopyTemplateID = copyTemplateID
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) Build() *CreateTemplateDetailReq {
	req := &CreateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteTemplateDetailReq struct {
	apiReq *core.ApiReq
}

type DeleteTemplateDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type DeleteTemplateDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewDeleteTemplateDetailReqBuilder() *DeleteTemplateDetailReqBuilder {
	builder := &DeleteTemplateDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *DeleteTemplateDetailReqBuilder) AccessUser(userKey string) *DeleteTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *DeleteTemplateDetailReqBuilder) UUID(uuid string) *DeleteTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *DeleteTemplateDetailReqBuilder) ProjectKey(projectKey string) *DeleteTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *DeleteTemplateDetailReqBuilder) TemplateID(templateID int64) *DeleteTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("template_id", fmt.Sprint(templateID))
	return builder
}
func (builder *DeleteTemplateDetailReqBuilder) Build() *DeleteTemplateDetailReq {
	req := &DeleteTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryTemplateDetailReq struct {
	apiReq *core.ApiReq
}

type QueryTemplateDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *TemplateDetail `json:"data"`
}

type QueryTemplateDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryTemplateDetailReqBuilder() *QueryTemplateDetailReqBuilder {
	builder := &QueryTemplateDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryTemplateDetailReqBuilder) AccessUser(userKey string) *QueryTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryTemplateDetailReqBuilder) UUID(uuid string) *QueryTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *QueryTemplateDetailReqBuilder) ProjectKey(projectKey string) *QueryTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryTemplateDetailReqBuilder) TemplateID(templateID int64) *QueryTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("template_id", fmt.Sprint(templateID))
	return builder
}
func (builder *QueryTemplateDetailReqBuilder) Build() *QueryTemplateDetailReq {
	req := &QueryTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkItemTemplatesReq struct {
	apiReq *core.ApiReq
}

type QueryWorkItemTemplatesResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*TemplateConf `json:"data"`
}

type QueryWorkItemTemplatesReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryWorkItemTemplatesReqBuilder() *QueryWorkItemTemplatesReqBuilder {
	builder := &QueryWorkItemTemplatesReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryWorkItemTemplatesReqBuilder) AccessUser(userKey string) *QueryWorkItemTemplatesReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryWorkItemTemplatesReqBuilder) UUID(uuid string) *QueryWorkItemTemplatesReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *QueryWorkItemTemplatesReqBuilder) ProjectKey(projectKey string) *QueryWorkItemTemplatesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryWorkItemTemplatesReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryWorkItemTemplatesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *QueryWorkItemTemplatesReqBuilder) Build() *QueryWorkItemTemplatesReq {
	req := &QueryWorkItemTemplatesReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateTemplateDetailReq struct {
	apiReq *core.ApiReq
}
type UpdateTemplateDetailReqBody struct {
	ProjectKey string `json:"project_key"`

	TemplateID int64 `json:"template_id"`

	WorkflowConfs []*WorkflowConfInfo `json:"workflow_confs"`

	StateFlowConfs []*StateFlowConfInfo `json:"state_flow_confs"`
}

type UpdateTemplateDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

type UpdateTemplateDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewUpdateTemplateDetailReqBuilder() *UpdateTemplateDetailReqBuilder {
	builder := &UpdateTemplateDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &UpdateTemplateDetailReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *UpdateTemplateDetailReqBuilder) AccessUser(userKey string) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *UpdateTemplateDetailReqBuilder) UUID(uuid string) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) ProjectKey(projectKey string) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).ProjectKey = projectKey
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) TemplateID(templateID int64) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).TemplateID = templateID
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) WorkflowConfs(workflowConfs []*WorkflowConfInfo) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).WorkflowConfs = workflowConfs
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) StateFlowConfs(stateFlowConfs []*StateFlowConfInfo) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).StateFlowConfs = stateFlowConfs
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) Build() *UpdateTemplateDetailReq {
	req := &UpdateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}
