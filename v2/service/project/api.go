package project

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_ListProjectTeam = "/open_api/:project_key/teams/all"

const APIPath_OAPIBatchQueryProjectInfo = "/open_api/projects/detail"

const APIPath_OAPICreateComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/create"

const APIPath_OAPIDeleteComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/:comment_id"

const APIPath_OAPIListComments = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comments"

const APIPath_OAPIPageGetWorkItemOpRecord = "/open_api/op_record/work_item/list"

const APIPath_OAPIQueryProjects = "/open_api/projects"

const APIPath_OAPIUpdateComment = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/comment/:comment_id"


func NewService(config *core.Config) *ProjectService {
	a := &ProjectService{config: config}
	return a
}

type ProjectService struct {
	config *core.Config
}


/*
 * open_api下沉接口
@name: 获取空间下团队人员
 * @desc: open_api下沉
 */
func (a *ProjectService) ListProjectTeam(ctx context.Context, req *ListProjectTeamReq, options ...core.RequestOptionFunc) (*ListProjectTeamResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListProjectTeam
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListProjectTeamResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:OAPIBatchQueryProjectInfo
 * @desc:OpenAPI，获取空间详情
 */
func (a *ProjectService) OAPIBatchQueryProjectInfo(ctx context.Context, req *OAPIBatchQueryProjectInfoReq, options ...core.RequestOptionFunc) (*OAPIBatchQueryProjectInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIBatchQueryProjectInfo
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryProjectInfo] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIBatchQueryProjectInfoResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryProjectInfo] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * comment openapi接口
@name:APICreateCommentForOpen
 * @desc:OpenAPI，创建评论接口
 */
func (a *ProjectService) OAPICreateComment(ctx context.Context, req *OAPICreateCommentReq, options ...core.RequestOptionFunc) (*OAPICreateCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateComment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateComment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateCommentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:APIDeleteCommentForOpen
 * @desc:OpenAPI，删除评论接口
 */
func (a *ProjectService) OAPIDeleteComment(ctx context.Context, req *OAPIDeleteCommentReq, options ...core.RequestOptionFunc) (*OAPIDeleteCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteComment
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteComment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteCommentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:APIListCommentsForOpen
 * @desc:OpenAPI，分页查询评论接口
 */
func (a *ProjectService) OAPIListComments(ctx context.Context, req *OAPIListCommentsReq, options ...core.RequestOptionFunc) (*OAPIListCommentsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIListComments
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIListComments] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIListCommentsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIListComments] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:OAPIPageGetWorkItemOpRecord
 * @desc:OpenAPI，查询操作记录详情
 */
func (a *ProjectService) OAPIPageGetWorkItemOpRecord(ctx context.Context, req *OAPIPageGetWorkItemOpRecordReq, options ...core.RequestOptionFunc) (*OAPIPageGetWorkItemOpRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIPageGetWorkItemOpRecord
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPageGetWorkItemOpRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIPageGetWorkItemOpRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPageGetWorkItemOpRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:OAPIQueryProjects
 * @desc:OpenAPI，获取空间列表
 */
func (a *ProjectService) OAPIQueryProjects(ctx context.Context, req *OAPIQueryProjectsReq, options ...core.RequestOptionFunc) (*OAPIQueryProjectsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryProjects
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjects] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryProjectsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjects] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:APIUpdateCommentForOpen
 * @desc:OpenAPI，更新评论接口
 */
func (a *ProjectService) OAPIUpdateComment(ctx context.Context, req *OAPIUpdateCommentReq, options ...core.RequestOptionFunc) (*OAPIUpdateCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateComment
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateComment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateCommentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateComment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

