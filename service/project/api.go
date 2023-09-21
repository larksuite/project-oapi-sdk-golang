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

const APIPathGetProjectDetail = "/open_api/projects/detail"

const APIPathListProject = "/open_api/projects"

const APIPathListProjectBusiness = "/open_api/:project_key/business/all"

const APIPathListProjectTeam = "/open_api/:project_key/teams/all"

const APIPathListProjectWorkItemType = "/open_api/:project_key/work_item/all-types"

func NewService(config *core.Config) *ProjectService {
	a := &ProjectService{config: config}
	return a
}

type ProjectService struct {
	config *core.Config
}

// 获取空间详情
func (a *ProjectService) GetProjectDetail(ctx context.Context, req *GetProjectDetailReq, options ...core.RequestOptionFunc) (*GetProjectDetailResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathGetProjectDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetProjectDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetProjectDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetProjectDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间列表
func (a *ProjectService) ListProject(ctx context.Context, req *ListProjectReq, options ...core.RequestOptionFunc) (*ListProjectResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathListProject
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProject] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListProjectResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProject] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间下业务线详情
func (a *ProjectService) ListProjectBusiness(ctx context.Context, req *ListProjectBusinessReq, options ...core.RequestOptionFunc) (*ListProjectBusinessResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathListProjectBusiness
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectBusiness] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListProjectBusinessResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectBusiness] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间下团队人员
func (a *ProjectService) ListProjectTeam(ctx context.Context, req *ListProjectTeamReq, options ...core.RequestOptionFunc) (*ListProjectTeamResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathListProjectTeam
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListProjectTeamResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间下工作项类型
func (a *ProjectService) ListProjectWorkItemType(ctx context.Context, req *ListProjectWorkItemTypeReq, options ...core.RequestOptionFunc) (*ListProjectWorkItemTypeResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathListProjectWorkItemType
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectWorkItemType] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListProjectWorkItemTypeResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectWorkItemType] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
