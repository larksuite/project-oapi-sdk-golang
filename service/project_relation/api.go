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

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPath_CreateProjectRelationInstances = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/batch_bind"

const APIPath_DeleteProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id"

const APIPath_QueryProjectRelation = "/open_api/:project_key/relation/rules"

const APIPath_QueryProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/work_item_list"

func NewService(config *core.Config) *ProjectRelationService {
	a := &ProjectRelationService{config: config}
	return a
}

type ProjectRelationService struct {
	config *core.Config
}

/*
 *
 */
func (a *ProjectRelationService) CreateProjectRelationInstances(ctx context.Context, req *CreateProjectRelationInstancesReq, options ...core.RequestOptionFunc) (*CreateProjectRelationInstancesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateProjectRelationInstances
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateProjectRelationInstances] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateProjectRelationInstancesResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateProjectRelationInstances] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *ProjectRelationService) DeleteProjectRelationInstance(ctx context.Context, req *DeleteProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*DeleteProjectRelationInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteProjectRelationInstance
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteProjectRelationInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *ProjectRelationService) QueryProjectRelation(ctx context.Context, req *QueryProjectRelationReq, options ...core.RequestOptionFunc) (*QueryProjectRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryProjectRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryProjectRelationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *ProjectRelationService) QueryProjectRelationInstance(ctx context.Context, req *QueryProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*QueryProjectRelationInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryProjectRelationInstance
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryProjectRelationInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
