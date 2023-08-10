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

type TokenErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserPluginToken struct {
	ExpireTime             int    `json:"expire_time"`
	Token                  string `json:"token"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpireTime int    `json:"refresh_token_expire_time"`
	UserKey                string `json:"user_key"`
	TenantKey              string `json:"tenant_key"`
}

type RefreshToken struct {
	ExpireTime             int    `json:"expire_time"`
	Token                  string `json:"token"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpireTime int    `json:"refresh_token_expire_time"`
}
