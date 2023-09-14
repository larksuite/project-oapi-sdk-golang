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
	body   *QueryUserDetailReqBody
}

func NewQueryUserDetailReqBuilder() *QueryUserDetailReqBuilder {
	builder := &QueryUserDetailReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &QueryUserDetailReqBody{}
	return builder
}
func (builder *QueryUserDetailReqBuilder) UserKeys(userKeys []string) *QueryUserDetailReqBuilder {
	builder.body.UserKeys = userKeys
	return builder
}
func (builder *QueryUserDetailReqBuilder) Usernames(usernames []string) *QueryUserDetailReqBuilder {
	builder.body.Usernames = usernames
	return builder
}
func (builder *QueryUserDetailReqBuilder) OutIDs(outIDs []string) *QueryUserDetailReqBuilder {
	builder.body.OutIDs = outIDs
	return builder
}
func (builder *QueryUserDetailReqBuilder) Emails(emails []string) *QueryUserDetailReqBuilder {
	builder.body.Emails = emails
	return builder
}
func (builder *QueryUserDetailReqBuilder) Build() *QueryUserDetailReq {
	req := &QueryUserDetailReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
