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

package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var tokenManager TokenManager = TokenManager{cache: cache}

type TokenManager struct {
	cache Cache
}

func (m *TokenManager) getPluginAccessToken(ctx context.Context, config *Config) (string, error) {
	token, err := m.get(ctx, pluginAccessTokenKey(config.AppId))
	if err != nil {
		return "", err
	}

	if token == "" {
		token, err = m.getPluginAccessTokenThenCache(ctx, config)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return token, nil
}

func (m *TokenManager) getVirtualPluginAccessToken(ctx context.Context, config *Config) (string, error) {
	token, err := m.get(ctx, virtualPluginAccessTokenKey(config.AppId))
	if err != nil {
		return "", err
	}

	if token == "" {
		token, err = m.getVitualPluginAccessTokenThenCache(ctx, config)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return token, nil
}

func (m *TokenManager) set(ctx context.Context, key, value string, ttl time.Duration) error {
	return m.cache.Set(ctx, key, value, ttl)
}

func (m *TokenManager) get(ctx context.Context, tokenKey string) (string, error) {
	token, err := m.cache.Get(ctx, tokenKey)
	return token, err
}

type GetAccessTokenReq struct {
	PluginId     string `json:"plugin_id"`
	PluginSecret string `json:"plugin_secret"`
	Type         int    `json:"type"`
}

type AccessToken struct {
	ExpireTime int    `json:"expire_time"`
	Token      string `json:"token"`
}
type GetAccessTokenResp struct {
	*ApiResp `json:"-"`
	Error    *AccessTokenErr `json:"error"`
	Data     *AccessToken    `json:"data"`
}

func (t *GetAccessTokenResp) Success() bool {
	return t.Error.Code == 0
}

type MarketplaceAppAccessTokenReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppTicket string `json:"app_ticket"`
}

type MarketplaceTenantAccessTokenReq struct {
	AppAccessToken string `json:"app_access_token"`
	TenantKey      string `json:"tenant_key"`
}

func pluginAccessTokenKey(appID string) string {
	return fmt.Sprintf("%s-%s", pluginAccessTokenKeyPrefix, appID)
}

func virtualPluginAccessTokenKey(appID string) string {
	return fmt.Sprintf("%s-%s", virtualPluginAccessTokenKeyPrefix, appID)
}
func (m *TokenManager) getPluginAccessTokenThenCache(ctx context.Context, config *Config) (string, error) {
	rawResp, err := Request(ctx, &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    PluginAccessTokenInternalUrlPath,
		Body: &GetAccessTokenReq{
			PluginId:     config.AppId,
			PluginSecret: config.AppSecret,
			Type:         0,
		},
		SkipAuth: true,
	}, config)

	if err != nil {
		return "", err
	}
	resp := &GetAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return "", err
	}
	if resp.Error.Code != 0 {
		config.Logger.Warn(ctx, fmt.Sprintf("get pluginAccessToken err:%v", Prettify(resp)))
		return "", resp.Error
	}
	expire := time.Duration(resp.Data.ExpireTime)*time.Second - expiryDelta
	err = m.set(ctx, pluginAccessTokenKey(config.AppId), resp.Data.Token, expire)
	if err != nil {
		config.Logger.Warn(ctx, fmt.Sprintf("pluginAccessToken save cache, err:%v", err))
	}
	return resp.Data.Token, err
}

func (m *TokenManager) getVitualPluginAccessTokenThenCache(ctx context.Context, config *Config) (string, error) {
	rawResp, err := Request(ctx, &ApiReq{
		HttpMethod: http.MethodPost,
		ApiPath:    PluginAccessTokenInternalUrlPath,
		Body: &GetAccessTokenReq{
			PluginId:     config.AppId,
			PluginSecret: config.AppSecret,
			Type:         1,
		},
		SkipAuth: true,
	}, config)

	if err != nil {
		return "", err
	}
	resp := &GetAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return "", err
	}
	if resp.Error.Code != 0 {
		config.Logger.Warn(ctx, fmt.Sprintf("get vitualPluginAccessToken cache, err:%v", Prettify(resp)))
		return "", resp.Error
	}
	expire := time.Duration(resp.Data.ExpireTime)*time.Second - expiryDelta
	err = m.cache.Set(ctx, virtualPluginAccessTokenKey(config.AppId), resp.Data.Token, expire)
	if err != nil {
		config.Logger.Warn(ctx, fmt.Sprintf("vitualPluginAccessToken save cache, err:%v", err))
	}
	return resp.Data.Token, err
}
