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

package role_conf

import (
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type QueryRoleConfDetailsReq struct {
	apiReq *core.APIReq
}

type QueryRoleConfDetailsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*RoleConfDetail `json:"data"`
}

type QueryRoleConfDetailsReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryRoleConfDetailsReqBuilder() *QueryRoleConfDetailsReqBuilder {
	builder := &QueryRoleConfDetailsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *QueryRoleConfDetailsReqBuilder) ProjectKey(projectKey string) *QueryRoleConfDetailsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryRoleConfDetailsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryRoleConfDetailsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *QueryRoleConfDetailsReqBuilder) Build() *QueryRoleConfDetailsReq {
	req := &QueryRoleConfDetailsReq{}
	req.apiReq = builder.apiReq
	return req
}
