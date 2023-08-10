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
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type CreateFieldReq struct {
	apiReq *core.ApiReq
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
	*core.ApiResp `json:"-"`
	core.CodeError
	Data string `json:"data"`
}

type CreateFieldReqBuilder struct {
	apiReq *core.ApiReq
}

func NewCreateFieldReqBuilder() *CreateFieldReqBuilder {
	builder := &CreateFieldReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &CreateFieldReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *CreateFieldReqBuilder) AccessUser(userKey string) *CreateFieldReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *CreateFieldReqBuilder) UUID(uuid string) *CreateFieldReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	builder.apiReq.Body.(*CreateFieldReqBody).FieldName = fieldName
	return builder
}
func (builder *CreateFieldReqBuilder) FieldTypeKey(fieldTypeKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldTypeKey = fieldTypeKey
	return builder
}
func (builder *CreateFieldReqBuilder) ValueType(valueType int64) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ValueType = valueType
	return builder
}
func (builder *CreateFieldReqBuilder) ReferenceWorkItemTypeKey(referenceWorkItemTypeKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ReferenceWorkItemTypeKey = referenceWorkItemTypeKey
	return builder
}
func (builder *CreateFieldReqBuilder) ReferenceFieldKey(referenceFieldKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ReferenceFieldKey = referenceFieldKey
	return builder
}
func (builder *CreateFieldReqBuilder) FieldValue(fieldValue interface{}) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldValue = fieldValue
	return builder
}
func (builder *CreateFieldReqBuilder) FreeAdd(freeAdd int64) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FreeAdd = freeAdd
	return builder
}
func (builder *CreateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).WorkItemRelationUUID = workItemRelationUUID
	return builder
}
func (builder *CreateFieldReqBuilder) DefaultValue(defaultValue interface{}) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).DefaultValue = defaultValue
	return builder
}
func (builder *CreateFieldReqBuilder) FieldAlias(fieldAlias string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldAlias = fieldAlias
	return builder
}
func (builder *CreateFieldReqBuilder) HelpDescription(helpDescription string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).HelpDescription = helpDescription
	return builder
}
func (builder *CreateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).AuthorizedRoles = authorizedRoles
	return builder
}
func (builder *CreateFieldReqBuilder) IsMulti(isMulti bool) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).IsMulti = isMulti
	return builder
}
func (builder *CreateFieldReqBuilder) Format(format bool) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).Format = format
	return builder
}
func (builder *CreateFieldReqBuilder) Build() *CreateFieldReq {
	req := &CreateFieldReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectFieldsReq struct {
	apiReq *core.ApiReq
}
type QueryProjectFieldsReqBody struct {
	WorkItemTypeKey string `json:"work_item_type_key"`
}

type QueryProjectFieldsResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*SimpleField `json:"data"`
}

type QueryProjectFieldsReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryProjectFieldsReqBuilder() *QueryProjectFieldsReqBuilder {
	builder := &QueryProjectFieldsReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &QueryProjectFieldsReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryProjectFieldsReqBuilder) AccessUser(userKey string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryProjectFieldsReqBuilder) UUID(uuid string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *QueryProjectFieldsReqBuilder) ProjectKey(projectKey string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryProjectFieldsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.Body.(*QueryProjectFieldsReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *QueryProjectFieldsReqBuilder) Build() *QueryProjectFieldsReq {
	req := &QueryProjectFieldsReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateFieldReq struct {
	apiReq *core.ApiReq
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
	*core.ApiResp `json:"-"`
	core.CodeError
}

type UpdateFieldReqBuilder struct {
	apiReq *core.ApiReq
}

func NewUpdateFieldReqBuilder() *UpdateFieldReqBuilder {
	builder := &UpdateFieldReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &UpdateFieldReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *UpdateFieldReqBuilder) AccessUser(userKey string) *UpdateFieldReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *UpdateFieldReqBuilder) UUID(uuid string) *UpdateFieldReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldName = fieldName
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldKey(fieldKey string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldKey = fieldKey
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldValue(fieldValue interface{}) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldValue = fieldValue
	return builder
}
func (builder *UpdateFieldReqBuilder) FreeAdd(freeAdd int64) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FreeAdd = freeAdd
	return builder
}
func (builder *UpdateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).WorkItemRelationUUID = workItemRelationUUID
	return builder
}
func (builder *UpdateFieldReqBuilder) DefaultValue(defaultValue interface{}) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).DefaultValue = defaultValue
	return builder
}
func (builder *UpdateFieldReqBuilder) FieldAlias(fieldAlias string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldAlias = fieldAlias
	return builder
}
func (builder *UpdateFieldReqBuilder) HelpDescription(helpDescription string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).HelpDescription = helpDescription
	return builder
}
func (builder *UpdateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).AuthorizedRoles = authorizedRoles
	return builder
}
func (builder *UpdateFieldReqBuilder) Build() *UpdateFieldReq {
	req := &UpdateFieldReq{}
	req.apiReq = builder.apiReq
	return req
}
