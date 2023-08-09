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

package project

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPath_GetProjectDetail = "/open_api/projects/detail"

const APIPath_ListProject = "/open_api/projects"

const APIPath_ListProjectBusiness = "/open_api/:project_key/business/all"

const APIPath_ListProjectTeam = "/open_api/:project_key/teams/all"

const APIPath_ListProjectWorkItemType = "/open_api/:project_key/work_item/all-types"

func NewService(config *core.Config) *ProjectService {
	a := &ProjectService{config: config}
	return a
}

type ProjectService struct {
	config *core.Config
}

/*
 *
 */
func (a *ProjectService) GetProjectDetail(ctx context.Context, req *GetProjectDetailReq, options ...core.RequestOptionFunc) (*GetProjectDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetProjectDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetProjectDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetProjectDetailResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetProjectDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *ProjectService) ListProject(ctx context.Context, req *ListProjectReq, options ...core.RequestOptionFunc) (*ListProjectResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListProject
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProject] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListProjectResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProject] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *ProjectService) ListProjectBusiness(ctx context.Context, req *ListProjectBusinessReq, options ...core.RequestOptionFunc) (*ListProjectBusinessResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListProjectBusiness
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectBusiness] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListProjectBusinessResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectBusiness] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
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
	resp := &ListProjectTeamResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *ProjectService) ListProjectWorkItemType(ctx context.Context, req *ListProjectWorkItemTypeReq, options ...core.RequestOptionFunc) (*ListProjectWorkItemTypeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListProjectWorkItemType
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectWorkItemType] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListProjectWorkItemTypeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectWorkItemType] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
