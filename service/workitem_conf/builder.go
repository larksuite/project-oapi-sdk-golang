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

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type CreateTemplateDetailReq struct {
	apiReq *core.APIReq
}
type CreateTemplateDetailReqBody struct {
	ProjectKey string `json:"project_key"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	TemplateName string `json:"template_name"`

	CopyTemplateID int64 `json:"copy_template_id"`
}

type CreateTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data int64 `json:"data"`
}

type CreateTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
	body   *CreateTemplateDetailReqBody
}

func NewCreateTemplateDetailReqBuilder() *CreateTemplateDetailReqBuilder {
	builder := &CreateTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateTemplateDetailReqBody{}
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) ProjectKey(projectKey string) *CreateTemplateDetailReqBuilder {
	builder.body.ProjectKey = projectKey
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateTemplateDetailReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) TemplateName(templateName string) *CreateTemplateDetailReqBuilder {
	builder.body.TemplateName = templateName
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) CopyTemplateID(copyTemplateID int64) *CreateTemplateDetailReqBuilder {
	builder.body.CopyTemplateID = copyTemplateID
	return builder
}
func (builder *CreateTemplateDetailReqBuilder) Build() *CreateTemplateDetailReq {
	req := &CreateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type DeleteTemplateDetailReq struct {
	apiReq *core.APIReq
}

type DeleteTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteTemplateDetailReqBuilder() *DeleteTemplateDetailReqBuilder {
	builder := &DeleteTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
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
	apiReq *core.APIReq
}

type QueryTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *TemplateDetail `json:"data"`
}

type QueryTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryTemplateDetailReqBuilder() *QueryTemplateDetailReqBuilder {
	builder := &QueryTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
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

type QueryWbsTemplateConfReq struct {
	apiReq *core.APIReq
}
type QueryWbsTemplateConfReqBody struct {
	TemplateKey string `json:"template_key"`
}

type QueryWbsTemplateConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *WbsTemplate `json:"data"`
}

type QueryWbsTemplateConfReqBuilder struct {
	apiReq *core.APIReq
	body   *QueryWbsTemplateConfReqBody
}

func NewQueryWbsTemplateConfReqBuilder() *QueryWbsTemplateConfReqBuilder {
	builder := &QueryWbsTemplateConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &QueryWbsTemplateConfReqBody{}
	return builder
}
func (builder *QueryWbsTemplateConfReqBuilder) ProjectKey(projectKey string) *QueryWbsTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryWbsTemplateConfReqBuilder) TemplateKey(templateKey string) *QueryWbsTemplateConfReqBuilder {
	builder.body.TemplateKey = templateKey
	return builder
}
func (builder *QueryWbsTemplateConfReqBuilder) Build() *QueryWbsTemplateConfReq {
	req := &QueryWbsTemplateConfReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type QueryWorkItemTemplatesReq struct {
	apiReq *core.APIReq
}

type QueryWorkItemTemplatesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*TemplateConf `json:"data"`
}

type QueryWorkItemTemplatesReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWorkItemTemplatesReqBuilder() *QueryWorkItemTemplatesReqBuilder {
	builder := &QueryWorkItemTemplatesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
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
	apiReq *core.APIReq
}
type UpdateTemplateDetailReqBody struct {
	ProjectKey string `json:"project_key"`

	TemplateID int64 `json:"template_id"`

	WorkflowConfs []*WorkflowConfInfo `json:"workflow_confs"`

	StateFlowConfs []*StateFlowConfInfo `json:"state_flow_confs"`
}

type UpdateTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
	body   *UpdateTemplateDetailReqBody
}

func NewUpdateTemplateDetailReqBuilder() *UpdateTemplateDetailReqBuilder {
	builder := &UpdateTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateTemplateDetailReqBody{}
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) ProjectKey(projectKey string) *UpdateTemplateDetailReqBuilder {
	builder.body.ProjectKey = projectKey
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) TemplateID(templateID int64) *UpdateTemplateDetailReqBuilder {
	builder.body.TemplateID = templateID
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) WorkflowConfs(workflowConfs []*WorkflowConfInfo) *UpdateTemplateDetailReqBuilder {
	builder.body.WorkflowConfs = workflowConfs
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) StateFlowConfs(stateFlowConfs []*StateFlowConfInfo) *UpdateTemplateDetailReqBuilder {
	builder.body.StateFlowConfs = stateFlowConfs
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) Build() *UpdateTemplateDetailReq {
	req := &UpdateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
