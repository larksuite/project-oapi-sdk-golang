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

package wbs_view_draft

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

const APIPathWbsCollaborationPublish = "/open_api/:project_key/wbs_view_draft/publish"

const APIPathSwitchBackToWbsViewDraft = "/open_api/:project_key/wbs_view_draft/switch"

func NewService(config *core.Config) *WbsViewDraftService {
	a := &WbsViewDraftService{config: config}
	return a
}

type WbsViewDraftService struct {
	config *core.Config
}

// WbsCollaborationPublish 计划表基于草稿发布
func (a *WbsViewDraftService) WbsCollaborationPublish(ctx context.Context, req *WbsCollaborationPublishReq, options ...core.RequestOptionFunc) (*WbsCollaborationPublishResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathWbsCollaborationPublish
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WbsCollaborationPublish] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &WbsCollaborationPublishResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WbsCollaborationPublish] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// SwitchBackToWbsViewDraft 审批拒绝切回草稿
func (a *WbsViewDraftService) SwitchBackToWbsViewDraft(ctx context.Context, req *SwitchBackToWbsViewDraftReq, options ...core.RequestOptionFunc) (*SwitchBackToWbsViewDraftResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathSwitchBackToWbsViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SwitchBackToWbsViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SwitchBackToWbsViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SwitchBackToWbsViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
