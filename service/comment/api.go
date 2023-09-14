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
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPath_CreateComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/create"

const APIPath_DeleteComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/:comment_id"

const APIPath_QueryComments = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comments"

const APIPath_UpdateComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/:comment_id"

func NewService(config *core.Config) *CommentService {
	a := &CommentService{config: config}
	return a
}

type CommentService struct {
	config *core.Config
}

/*
 *   添加评论
 */
func (a *CommentService) CreateComment(ctx context.Context, req *CreateCommentReq, options ...core.RequestOptionFunc) (*CreateCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateComment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateComment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   删除评论
 */
func (a *CommentService) DeleteComment(ctx context.Context, req *DeleteCommentReq, options ...core.RequestOptionFunc) (*DeleteCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteComment
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteComment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   查询评论
 */
func (a *CommentService) QueryComments(ctx context.Context, req *QueryCommentsReq, options ...core.RequestOptionFunc) (*QueryCommentsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryComments
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryComments] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryCommentsResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryComments] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *   更新评论
 */
func (a *CommentService) UpdateComment(ctx context.Context, req *UpdateCommentReq, options ...core.RequestOptionFunc) (*UpdateCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateComment
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateComment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
