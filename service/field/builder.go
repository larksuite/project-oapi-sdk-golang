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

package field

import (
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type CreateFieldReq struct {
	apiReq *core.APIReq
}
type CreateFieldReqBody struct {
	FieldName string `json:"field_name"`

	FieldTypeKey string `json:"field_type_key"`

	ValueType int64 `json:"value_type"`

	ReferenceWorkItemTypeKey string `json:"reference_work_item_type_key"`

	ReferenceFieldKey string `json:"reference_field_key"`

	FieldValue interface{} `json:"field_value"`

	FreeAdd int64 `json:"free_add"`

	WorkItemRelationUUID string `json:"work_item_relation_uuid"`

	DefaultValue interface{} `json:"default_value"`

	FieldAlias string `json:"field_alias"`

	HelpDescription string `json:"help_description"`

	AuthorizedRoles []string `json:"authorized_roles"`

	IsMulti bool `json:"is_multi"`

	Format bool `json:"format"`
}

type CreateFieldResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data string `json:"data"`
}

type CreateFieldReqBuilder struct {
	apiReq *core.APIReq
	body   *CreateFieldReqBody
}

func NewCreateFieldReqBuilder() *CreateFieldReqBuilder {
	builder := &CreateFieldReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateFieldReqBody{}
	return builder
}
func (builder *CreateFieldReqBuilder) ProjectKey(projectKey string) *CreateFieldReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *CreateFieldReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateFieldReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *CreateFieldReqBuilder) FieldName(fieldName string) *CreateFieldReqBuilder {
	builder.body.FieldName = fieldName
	return builder
}
func (builder *CreateFieldReqBuilder) FieldTypeKey(fieldTypeKey string) *CreateFieldReqBuilder {
	builder.body.FieldTypeKey = fieldTypeKey
	return builder
}
func (builder *CreateFieldReqBuilder) ValueType(valueType int64) *CreateFieldReqBuilder {
	builder.body.ValueType = valueType
	return builder
}
func (builder *CreateFieldReqBuilder) ReferenceWorkItemTypeKey(referenceWorkItemTypeKey string) *CreateFieldReqBuilder {
	builder.body.ReferenceWorkItemTypeKey = referenceWorkItemTypeKey
	return builder
}
func (builder *CreateFieldReqBuilder) ReferenceFieldKey(referenceFieldKey string) *CreateFieldReqBuilder {
	builder.body.ReferenceFieldKey = referenceFieldKey
	return builder
}
func (builder *CreateFieldReqBuilder) FieldValue(fieldValue interface{}) *CreateFieldReqBuilder {
	builder.body.FieldValue = fieldValue
	return builder
}
func (builder *CreateFieldReqBuilder) FreeAdd(freeAdd int64) *CreateFieldReqBuilder {
	builder.body.FreeAdd = freeAdd
	return builder
}
func (builder *CreateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID string) *CreateFieldReqBuilder {
	builder.body.WorkItemRelationUUID = workItemRelationUUID
	return builder
}
func (builder *CreateFieldReqBuilder) DefaultValue(defaultValue interface{}) *CreateFieldReqBuilder {
	builder.body.DefaultValue = defaultValue
	return builder
}
func (builder *CreateFieldReqBuilder) FieldAlias(fieldAlias string) *CreateFieldReqBuilder {
	builder.body.FieldAlias = fieldAlias
	return builder
}
func (builder *CreateFieldReqBuilder) HelpDescription(helpDescription string) *CreateFieldReqBuilder {
	builder.body.HelpDescription = helpDescription
	return builder
}
func (builder *CreateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *CreateFieldReqBuilder {
	builder.body.AuthorizedRoles = authorizedRoles
	return builder
}
func (builder *CreateFieldReqBuilder) IsMulti(isMulti bool) *CreateFieldReqBuilder {
	builder.body.IsMulti = isMulti
	return builder
}
func (builder *CreateFieldReqBuilder) Format(format bool) *CreateFieldReqBuilder {
	builder.body.Format = format
	return builder
}
func (builder *CreateFieldReqBuilder) Build() *CreateFieldReq {
	req := &CreateFieldReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type QueryProjectFieldsReq struct {
	apiReq *core.APIReq
}
type QueryProjectFieldsReqBody struct {
	WorkItemTypeKey string `json:"work_item_type_key"`
}

type QueryProjectFieldsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*SimpleField `json:"data"`
}

type QueryProjectFieldsReqBuilder struct {
	apiReq *core.APIReq
	body   *QueryProjectFieldsReqBody
}

func NewQueryProjectFieldsReqBuilder() *QueryProjectFieldsReqBuilder {
	builder := &QueryProjectFieldsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &QueryProjectFieldsReqBody{}
	return builder
}
func (builder *QueryProjectFieldsReqBuilder) ProjectKey(projectKey string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryProjectFieldsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryProjectFieldsReqBuilder {
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *QueryProjectFieldsReqBuilder) Build() *QueryProjectFieldsReq {
	req := &QueryProjectFieldsReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type UpdateFieldReq struct {
	apiReq *core.APIReq
}
type UpdateFieldReqBody struct {
	FieldName string `json:"field_name"`

	FieldKey string `json:"field_key"`

	FieldValue interface{} `json:"field_value"`

	FreeAdd int64 `json:"free_add"`

	WorkItemRelationUUID string `json:"work_item_relation_uuid"`

	DefaultValue interface{} `json:"default_value"`

	FieldAlias string `json:"field_alias"`

	HelpDescription string `json:"help_description"`

	AuthorizedRoles []string `json:"authorized_roles"`
}

type UpdateFieldResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateFieldReqBuilder struct {
	apiReq *core.APIReq
	body   *UpdateFieldReqBody
}

func NewUpdateFieldReqBuilder() *UpdateFieldReqBuilder {
	builder := &UpdateFieldReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateFieldReqBody{}
	return builder
}
func (builder *UpdateFieldReqBuilder) ProjectKey(projectKey string) *UpdateFieldReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *UpdateFieldReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateFieldReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldName(fieldName string) *UpdateFieldReqBuilder {
	builder.body.FieldName = fieldName
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldKey(fieldKey string) *UpdateFieldReqBuilder {
	builder.body.FieldKey = fieldKey
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldValue(fieldValue interface{}) *UpdateFieldReqBuilder {
	builder.body.FieldValue = fieldValue
	return builder
}
func (builder *UpdateFieldReqBuilder) FreeAdd(freeAdd int64) *UpdateFieldReqBuilder {
	builder.body.FreeAdd = freeAdd
	return builder
}
func (builder *UpdateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID string) *UpdateFieldReqBuilder {
	builder.body.WorkItemRelationUUID = workItemRelationUUID
	return builder
}
func (builder *UpdateFieldReqBuilder) DefaultValue(defaultValue interface{}) *UpdateFieldReqBuilder {
	builder.body.DefaultValue = defaultValue
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldAlias(fieldAlias string) *UpdateFieldReqBuilder {
	builder.body.FieldAlias = fieldAlias
	return builder
}
func (builder *UpdateFieldReqBuilder) HelpDescription(helpDescription string) *UpdateFieldReqBuilder {
	builder.body.HelpDescription = helpDescription
	return builder
}
func (builder *UpdateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *UpdateFieldReqBuilder {
	builder.body.AuthorizedRoles = authorizedRoles
	return builder
}
func (builder *UpdateFieldReqBuilder) Build() *UpdateFieldReq {
	req := &UpdateFieldReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
