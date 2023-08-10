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

import "code.byted.org/bits/project-oapi-sdk-golang/core"

type GetUserPluginTokenReq struct {
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

type GetUserPluginTokenResp struct {
	*core.ApiResp `json:"-"`
	Error         *TokenErr        `json:"error"`
	Data          *UserPluginToken `json:"data"`
}

type RefreshTokenReq struct {
	apiReq *core.ApiReq
}

type RefreshTokenReqBuilder struct {
	apiReq *core.ApiReq
}

func NewRefreshTokenReqBuilder() *RefreshTokenReqBuilder {
	builder := &RefreshTokenReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &RefreshTokenReqBody{},
	}
	return builder
}

type RefreshTokenReqBody struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    int    `json:"type,omitempty"`
}

// 刷新token
func (builder *RefreshTokenReqBuilder) RefreshToken(refreshToken string) *RefreshTokenReqBuilder {
	builder.apiReq.Body.(*RefreshTokenReqBody).RefreshToken = refreshToken
	return builder
}

// 要刷新的token类型，目前固定填1
func (builder *RefreshTokenReqBuilder) TokenType(tokenType int) *RefreshTokenReqBuilder {
	builder.apiReq.Body.(*RefreshTokenReqBody).TokenType = tokenType
	return builder
}

func (builder *RefreshTokenReqBuilder) Build() *RefreshTokenReq {
	req := &RefreshTokenReq{}
	req.apiReq = builder.apiReq
	return req
}

type RefreshTokenResp struct {
	*core.ApiResp `json:"-"`
	Error         *TokenErr     `json:"error"`
	Data          *RefreshToken `json:"data"`
}
