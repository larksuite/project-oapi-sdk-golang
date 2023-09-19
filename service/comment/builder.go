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

package comment

import (
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"

	"code.byted.org/bits/project-oapi-sdk-golang/service/common"
)

type CreateCommentReq struct {
	apiReq *core.APIReq
}
type CreateCommentReqBody struct {
	Content string `json:"content"`
}

type CreateCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data int64 `json:"data"`
}

type CreateCommentReqBuilder struct {
	apiReq *core.APIReq
	body   *CreateCommentReqBody
}

func NewCreateCommentReqBuilder() *CreateCommentReqBuilder {
	builder := &CreateCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &CreateCommentReqBody{}
	return builder
}
func (builder *CreateCommentReqBuilder) ProjectKey(projectKey string) *CreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *CreateCommentReqBuilder) WorkItemID(workItemID int64) *CreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *CreateCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *CreateCommentReqBuilder) Content(content string) *CreateCommentReqBuilder {
	builder.body.Content = content
	return builder
}
func (builder *CreateCommentReqBuilder) Build() *CreateCommentReq {
	req := &CreateCommentReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type DeleteCommentReq struct {
	apiReq *core.APIReq
}

type DeleteCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteCommentReqBuilder() *DeleteCommentReqBuilder {
	builder := &DeleteCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *DeleteCommentReqBuilder) ProjectKey(projectKey string) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *DeleteCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *DeleteCommentReqBuilder) WorkItemID(workItemID int64) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *DeleteCommentReqBuilder) CommentID(commentID int64) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("comment_id", fmt.Sprint(commentID))
	return builder
}
func (builder *DeleteCommentReqBuilder) Build() *DeleteCommentReq {
	req := &DeleteCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryCommentsReq struct {
	apiReq *core.APIReq
}

type QueryCommentsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Comment `json:"data"`

	Pagination *common.Pagination `json:"pagination"`
}

type QueryCommentsReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryCommentsReqBuilder() *QueryCommentsReqBuilder {
	builder := &QueryCommentsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *QueryCommentsReqBuilder) ProjectKey(projectKey string) *QueryCommentsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryCommentsReqBuilder) WorkItemID(workItemID int64) *QueryCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *QueryCommentsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *QueryCommentsReqBuilder) PageSize(pageSize int64) *QueryCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}
func (builder *QueryCommentsReqBuilder) PageNum(pageNum int64) *QueryCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_num", fmt.Sprint(pageNum))
	return builder
}
func (builder *QueryCommentsReqBuilder) Build() *QueryCommentsReq {
	req := &QueryCommentsReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateCommentReq struct {
	apiReq *core.APIReq
}
type UpdateCommentReqBody struct {
	Content string `json:"content"`
}

type UpdateCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateCommentReqBuilder struct {
	apiReq *core.APIReq
	body   *UpdateCommentReqBody
}

func NewUpdateCommentReqBuilder() *UpdateCommentReqBuilder {
	builder := &UpdateCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	builder.body = &UpdateCommentReqBody{}
	return builder
}
func (builder *UpdateCommentReqBuilder) ProjectKey(projectKey string) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *UpdateCommentReqBuilder) WorkItemID(workItemID int64) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *UpdateCommentReqBuilder) CommentID(commentID int64) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("comment_id", fmt.Sprint(commentID))
	return builder
}
func (builder *UpdateCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}
func (builder *UpdateCommentReqBuilder) Content(content string) *UpdateCommentReqBuilder {
	builder.body.Content = content
	return builder
}
func (builder *UpdateCommentReqBuilder) Build() *UpdateCommentReq {
	req := &UpdateCommentReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
