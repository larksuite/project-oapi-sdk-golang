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
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPath_BotJoinChat = "/open_api/:project_key/work_item/:work_item_id/bot_join_chat"

func NewService(config *core.Config) *ChatService {
	a := &ChatService{config: config}
	return a
}

type ChatService struct {
	config *core.Config
}

/*
 *   拉机器人入群
 */
func (a *ChatService) BotJoinChat(ctx context.Context, req *BotJoinChatReq, options ...core.RequestOptionFunc) (*BotJoinChatResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_BotJoinChat
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BotJoinChat] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &BotJoinChatResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BotJoinChat] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
