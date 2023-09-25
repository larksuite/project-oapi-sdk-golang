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

package user

import (
	"context"
	"fmt"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

const APIPathQueryUserDetail = "/open_api/user/query"

const APIPathSearchUser = "/open_api/user/search"

func NewService(config *core.Config) *UserService {
	a := &UserService{config: config}
	return a
}

type UserService struct {
	config *core.Config
}

// 获取用户详情
func (a *UserService) QueryUserDetail(ctx context.Context, req *QueryUserDetailReq, options ...core.RequestOptionFunc) (*QueryUserDetailResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryUserDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryUserDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryUserDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryUserDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 模糊查询指定空间的用户列表
func (a *UserService) SearchUser(ctx context.Context, req *SearchUserReq, options ...core.RequestOptionFunc) (*SearchUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathSearchUser
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SearchUserResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchUser] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
