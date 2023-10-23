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

package role_conf

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

const APIPathQueryRoleConfDetails = "/open_api/:project_key/flow_roles/:work_item_type_key"

func NewService(config *core.Config) *RoleConfService {
	a := &RoleConfService{config: config}
	return a
}

type RoleConfService struct {
	config *core.Config
}

// 获取流程角色配置详情
func (a *RoleConfService) QueryRoleConfDetails(ctx context.Context, req *QueryRoleConfDetailsReq, options ...core.RequestOptionFunc) (*QueryRoleConfDetailsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryRoleConfDetails
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryRoleConfDetails] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryRoleConfDetailsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryRoleConfDetails] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
