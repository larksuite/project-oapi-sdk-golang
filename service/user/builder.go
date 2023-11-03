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
	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type QueryUserDetailReq struct {
	apiReq *core.APIReq
}
type QueryUserDetailReqBody struct {
	UserKeys []string `json:"user_keys"`

	Usernames []string `json:"usernames"`

	OutIDs []string `json:"out_ids"`

	Emails []string `json:"emails"`
}

type QueryUserDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*UserBasicInfo `json:"data"`
}

type QueryUserDetailReqBuilder struct {
	apiReq *core.APIReq
	body   *QueryUserDetailReqBody
}

func NewQueryUserDetailReqBuilder() *QueryUserDetailReqBuilder {
	builder := &QueryUserDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
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

type SearchUserReq struct {
	apiReq *core.APIReq
}
type SearchUserReqBody struct {
	Query string `json:"query"`

	ProjectKey string `json:"project_key"`
}

type SearchUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*UserBasicInfo `json:"data"`
}

type SearchUserReqBuilder struct {
	apiReq *core.APIReq
	body   *SearchUserReqBody
}

func NewSearchUserReqBuilder() *SearchUserReqBuilder {
	builder := &SearchUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &SearchUserReqBody{}
	return builder
}
func (builder *SearchUserReqBuilder) Query(query string) *SearchUserReqBuilder {
	builder.body.Query = query
	return builder
}
func (builder *SearchUserReqBuilder) ProjectKey(projectKey string) *SearchUserReqBuilder {
	builder.body.ProjectKey = projectKey
	return builder
}
func (builder *SearchUserReqBuilder) Build() *SearchUserReq {
	req := &SearchUserReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
