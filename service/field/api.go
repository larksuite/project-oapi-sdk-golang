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

package field

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPathCreateField = "/open_api/:project_key/field/:work_item_type_key/create"

const APIPathQueryProjectFields = "/open_api/:project_key/field/all"

const APIPathUpdateField = "/open_api/:project_key/field/:work_item_type_key"

func NewService(config *core.Config) *FieldService {
	a := &FieldService{config: config}
	return a
}

type FieldService struct {
	config *core.Config
}

// 创建自定义字段
func (a *FieldService) CreateField(ctx context.Context, req *CreateFieldReq, options ...core.RequestOptionFunc) (*CreateFieldResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathCreateField
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateField] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateFieldResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateField] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取空间字段
func (a *FieldService) QueryProjectFields(ctx context.Context, req *QueryProjectFieldsReq, options ...core.RequestOptionFunc) (*QueryProjectFieldsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryProjectFields
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectFields] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryProjectFieldsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectFields] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 更新自定义字段
func (a *FieldService) UpdateField(ctx context.Context, req *UpdateFieldReq, options ...core.RequestOptionFunc) (*UpdateFieldResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathUpdateField
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateField] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateFieldResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateField] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
