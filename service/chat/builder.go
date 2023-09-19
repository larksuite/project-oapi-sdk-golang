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

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

type BotJoinChatReq struct {
	apiReq *core.APIReq
}
type BotJoinChatReqBody struct {
	WorkItemTypeKey string `json:"work_item_type_key"`

	AppIDs []string `json:"app_ids"`
}

type BotJoinChatResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *BotJoinChatInfo `json:"data"`
}

type BotJoinChatReqBuilder struct {
	apiReq *core.APIReq
	body   *BotJoinChatReqBody
}

func NewBotJoinChatReqBuilder() *BotJoinChatReqBuilder {
	builder := &BotJoinChatReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &BotJoinChatReqBody{}
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
	builder.body.WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *BotJoinChatReqBuilder) AppIDs(appIDs []string) *BotJoinChatReqBuilder {
	builder.body.AppIDs = appIDs
	return builder
}
func (builder *BotJoinChatReqBuilder) Build() *BotJoinChatReq {
	req := &BotJoinChatReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
