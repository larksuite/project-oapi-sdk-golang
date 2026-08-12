package comment

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_CreateComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/create"

const APIPath_CreateCommentNew = "/open_api/comment/create"

const APIPath_DeleteComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/:comment_id"

const APIPath_DeleteCommentNew = "/open_api/comment/delete"

const APIPath_ListComments = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comments"

const APIPath_QueryCommentNew = "/open_api/comment/query"

const APIPath_UpdateComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/:comment_id"

const APIPath_UpdateCommentNew = "/open_api/comment/update"


func NewService(config *core.Config) *CommentService {
	a := &CommentService{config: config}
	return a
}

type CommentService struct {
	config *core.Config
}


/*
 * comment openapi接口
@name:APICreateCommentForOpen
 * @desc:OpenAPI，创建评论接口
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
	resp := &CreateCommentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateComment
 * @desc: 创建评论
 */
func (a *CommentService) CreateCommentNew(ctx context.Context, req *CreateCommentNewReq, options ...core.RequestOptionFunc) (*CreateCommentNewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateCommentNew
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateCommentNew] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCommentNewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateCommentNew] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:APIDeleteCommentForOpen
 * @desc:OpenAPI，删除评论接口
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
	resp := &DeleteCommentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteComment
 * @desc: 删除评论
 */
func (a *CommentService) DeleteCommentNew(ctx context.Context, req *DeleteCommentNewReq, options ...core.RequestOptionFunc) (*DeleteCommentNewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteCommentNew
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteCommentNew] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCommentNewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteCommentNew] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:APIListCommentsForOpen
 * @desc:OpenAPI，分页查询评论接口
 */
func (a *CommentService) ListComments(ctx context.Context, req *ListCommentsReq, options ...core.RequestOptionFunc) (*ListCommentsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListComments
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListComments] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCommentsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListComments] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryComment
 * @desc: 查询评论
 */
func (a *CommentService) QueryCommentNew(ctx context.Context, req *QueryCommentNewReq, options ...core.RequestOptionFunc) (*QueryCommentNewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryCommentNew
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryCommentNew] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryCommentNewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryCommentNew] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:APIUpdateCommentForOpen
 * @desc:OpenAPI，更新评论接口
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
	resp := &UpdateCommentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateComment
 * @desc: 更新评论
 */
func (a *CommentService) UpdateCommentNew(ctx context.Context, req *UpdateCommentNewReq, options ...core.RequestOptionFunc) (*UpdateCommentNewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateCommentNew
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateCommentNew] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateCommentNewResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateCommentNew] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

