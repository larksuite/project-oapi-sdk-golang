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

package project_relation

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

const APIPathCreateProjectRelationInstances = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/batch_bind"

const APIPathDeleteProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id"

const APIPathQueryProjectRelation = "/open_api/:project_key/relation/rules"

const APIPathQueryProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/work_item_list"

func NewService(config *core.Config) *ProjectRelationService {
	a := &ProjectRelationService{config: config}
	return a
}

type ProjectRelationService struct {
	config *core.Config
}

// 通过空间关联绑定关联工作项
func (a *ProjectRelationService) CreateProjectRelationInstances(ctx context.Context, req *CreateProjectRelationInstancesReq, options ...core.RequestOptionFunc) (*CreateProjectRelationInstancesResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathCreateProjectRelationInstances
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateProjectRelationInstances] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateProjectRelationInstancesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateProjectRelationInstances] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 通过空间关联解绑关联工作项
func (a *ProjectRelationService) DeleteProjectRelationInstance(ctx context.Context, req *DeleteProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*DeleteProjectRelationInstanceResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathDeleteProjectRelationInstance
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteProjectRelationInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间关联规则列表
func (a *ProjectRelationService) QueryProjectRelation(ctx context.Context, req *QueryProjectRelationReq, options ...core.RequestOptionFunc) (*QueryProjectRelationResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryProjectRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryProjectRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间关联下的关联工作项列表
func (a *ProjectRelationService) QueryProjectRelationInstance(ctx context.Context, req *QueryProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*QueryProjectRelationInstanceResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryProjectRelationInstance
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryProjectRelationInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
