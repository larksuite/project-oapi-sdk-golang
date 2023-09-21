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

package plugin

import (
	"code.byted.org/bits/project-oapi-sdk-golang/core"
	"context"
	"fmt"
	"net/http"
)

const (
	// 获取plugin_token
	ApiPathGetPluginToken     = "/open_api/authen/plugin_token"
	ApiPathGetUserPluginToken = "/open_api/authen/user_plugin_token"
	ApiPathRefreshToken       = "/open_api/authen/refresh_token"
)

func NewService(config *core.Config) *PluginService {
	a := &PluginService{config: config}
	return a
}

type PluginService struct {
	config *core.Config
}

// 获取plugin_token（0为plugin_token，1为虚拟plugin_token）
func (a *PluginService) GetPluginToken(ctx context.Context, pluginType int, options ...core.RequestOptionFunc) (*core.GetAccessTokenResp, error) {
	apiReq := &core.APIReq{
		HttpMethod: http.MethodPost,
		ApiPath:    ApiPathGetPluginToken,
		Body: &core.GetAccessTokenReq{
			PluginId:     a.config.AppID,
			PluginSecret: a.config.AppSecret,
			Type:         pluginType,
		},
		SkipAuth: true,
	}
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetPluginToken] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &core.GetAccessTokenResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetPluginToken] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 获取user_plugin_token
func (a *PluginService) GetUserPluginToken(ctx context.Context, code string, options ...core.RequestOptionFunc) (*GetUserPluginTokenResp, error) {
	apiReq := &core.APIReq{
		HttpMethod: http.MethodPost,
		ApiPath:    ApiPathGetUserPluginToken,
		Body: &GetUserPluginTokenReq{
			Code:      code,
			GrantType: "authorization_code",
		},
	}
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetUserPluginToken] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetUserPluginTokenResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetUserPluginToken] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

// 刷新Token（目前刷新token仅支持刷新user_plugin_token）
func (a *PluginService) RefreshToken(ctx context.Context, req *RefreshTokenReq, options ...core.RequestOptionFunc) (*RefreshTokenResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathRefreshToken
	apiReq.HttpMethod = http.MethodPost
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[RefreshToken] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &RefreshTokenResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[RefreshToken] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
