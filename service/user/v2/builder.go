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

package user

import (
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type QueryUserDetailReq struct {
	apiReq *core.ApiReq
}
type QueryUserDetailReqBody struct {
	UserKeys []string `json:"user_keys"`

	Usernames []string `json:"usernames"`

	OutIDs []string `json:"out_ids"`

	Emails []string `json:"emails"`
}

type QueryUserDetailResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data []*UserBasicInfo `json:"data"`
}

type QueryUserDetailReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryUserDetailReqBuilder() *QueryUserDetailReqBuilder {
	builder := &QueryUserDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &QueryUserDetailReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *QueryUserDetailReqBuilder) AccessUser(userKey string) *QueryUserDetailReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *QueryUserDetailReqBuilder) UUID(uuid string) *QueryUserDetailReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *QueryUserDetailReqBuilder) UserKeys(userKeys []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).UserKeys = userKeys
	return builder
}
func (builder *QueryUserDetailReqBuilder) Usernames(usernames []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).Usernames = usernames
	return builder
}
func (builder *QueryUserDetailReqBuilder) OutIDs(outIDs []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).OutIDs = outIDs
	return builder
}
func (builder *QueryUserDetailReqBuilder) Emails(emails []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).Emails = emails
	return builder
}
func (builder *QueryUserDetailReqBuilder) Build() *QueryUserDetailReq {
	req := &QueryUserDetailReq{}
	req.apiReq = builder.apiReq
	return req
}
