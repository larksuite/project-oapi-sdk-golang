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
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type WbsCollaborationPublishReq struct {
	apiReq *core.APIReq
}

type WbsCollaborationPublishReqBody struct {
	WorkItemID int64  `json:"work_item_id"`
	DraftID    string `json:"draft_id"`
}

type WbsCollaborationPublishReqBuilder struct {
	apiReq *core.APIReq
	body   *WbsCollaborationPublishReqBody
}

func NewWbsCollaborationPublishReqBuilder() *WbsCollaborationPublishReqBuilder {
	builder := &WbsCollaborationPublishReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &WbsCollaborationPublishReqBody{}
	return builder
}

func (builder *WbsCollaborationPublishReqBuilder) ProjectKey(projectKey string) *WbsCollaborationPublishReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *WbsCollaborationPublishReqBuilder) WorkItemID(workItemID int64) *WbsCollaborationPublishReqBuilder {
	builder.body.WorkItemID = workItemID
	return builder
}

func (builder *WbsCollaborationPublishReqBuilder) DraftID(draftID string) *WbsCollaborationPublishReqBuilder {
	builder.body.DraftID = draftID
	return builder
}

func (builder *WbsCollaborationPublishReqBuilder) Build() *WbsCollaborationPublishReq {
	req := &WbsCollaborationPublishReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type WbsCollaborationPublishResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *WbsViewDraftRespData `json:"data"`
}

type SwitchBackToWbsViewDraftReq struct {
	apiReq *core.APIReq
}

type SwitchBackToWbsViewDraftReqBody struct {
	WorkItemID int64  `json:"work_item_id"`
	DraftID    string `json:"draft_id"`
}

type SwitchBackToWbsViewDraftReqBuilder struct {
	apiReq *core.APIReq
	body   *SwitchBackToWbsViewDraftReqBody
}

func NewSwitchBackToWbsViewDraftReqBuilder() *SwitchBackToWbsViewDraftReqBuilder {
	builder := &SwitchBackToWbsViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &SwitchBackToWbsViewDraftReqBody{}
	return builder
}

func (builder *SwitchBackToWbsViewDraftReqBuilder) ProjectKey(projectKey string) *SwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *SwitchBackToWbsViewDraftReqBuilder) WorkItemID(workItemID int64) *SwitchBackToWbsViewDraftReqBuilder {
	builder.body.WorkItemID = workItemID
	return builder
}

func (builder *SwitchBackToWbsViewDraftReqBuilder) DraftID(draftID string) *SwitchBackToWbsViewDraftReqBuilder {
	builder.body.DraftID = draftID
	return builder
}

func (builder *SwitchBackToWbsViewDraftReqBuilder) Build() *SwitchBackToWbsViewDraftReq {
	req := &SwitchBackToWbsViewDraftReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type SwitchBackToWbsViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *WbsViewDraftRespData `json:"data"`
}
