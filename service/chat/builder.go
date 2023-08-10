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

package chat

import (
	"fmt"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type BotJoinChatReq struct {
	apiReq *core.ApiReq
}
type BotJoinChatReqBody struct {
	WorkItemTypeKey string `json:"work_item_type_key"`

	AppIDs []string `json:"app_ids"`
}

type BotJoinChatResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
	Data *BotJoinChatInfo `json:"data"`
}

type BotJoinChatReqBuilder struct {
	apiReq *core.ApiReq
}

func NewBotJoinChatReqBuilder() *BotJoinChatReqBuilder {
	builder := &BotJoinChatReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Header:      make(http.Header),
		Body:        &BotJoinChatReqBody{},
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *BotJoinChatReqBuilder) AccessUser(userKey string) *BotJoinChatReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *BotJoinChatReqBuilder) UUID(uuid string) *BotJoinChatReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}
func (builder *BotJoinChatReqBuilder) ProjectKey(projectKey string) *BotJoinChatReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *BotJoinChatReqBuilder) WorkItemID(workItemID int64) *BotJoinChatReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *BotJoinChatReqBuilder) WorkItemTypeKey(workItemTypeKey string) *BotJoinChatReqBuilder {
	builder.apiReq.Body.(*BotJoinChatReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *BotJoinChatReqBuilder) AppIDs(appIDs []string) *BotJoinChatReqBuilder {
	builder.apiReq.Body.(*BotJoinChatReqBody).AppIDs = appIDs
	return builder
}
func (builder *BotJoinChatReqBuilder) Build() *BotJoinChatReq {
	req := &BotJoinChatReq{}
	req.apiReq = builder.apiReq
	return req
}
