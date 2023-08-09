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
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type QueryRoleConfDetailsReq struct {
	apiReq *core.ApiReq
}

type QueryRoleConfDetailsResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*RoleConfDetail `json:"data"`
}

type QueryRoleConfDetailsReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryRoleConfDetailsReqBuilder() *QueryRoleConfDetailsReqBuilder {
	builder := &QueryRoleConfDetailsReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryRoleConfDetailsReqBuilder) AccessUser(userKey string) *QueryRoleConfDetailsReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryRoleConfDetailsReqBuilder) UUID(uuid string) *QueryRoleConfDetailsReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
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
