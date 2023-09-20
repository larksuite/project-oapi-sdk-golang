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
	"time"
)

const defaultContentType = contentTypeJson + "; charset=utf-8"
const userAgentHeader = "User-Agent"

const (
	HTTPHeaderUserKey      = "X-USER-KEY"
	HTTPHeaderAccessToken  = "X-PLUGIN-TOKEN"
	HTTPHeaderIdemUUID     = "X-IDEM-UUID"
	HTTPHeaderKeyRequestID = "X-Request-Id"
	HTTPHeaderKeyLogID     = "X-Tt-Logid"
	contentTypeHeader      = "Content-Type"
	contentTypeJson        = "application/json"
	customRequestId        = "Oapi-Sdk-Request-Id"
)

const (
	PluginAccessTokenInternalUrlPath string = "/open_api/authen/plugin_token"
)

type AccessTokenType int

const (
	AccessTokenTypePlugin        AccessTokenType = 0
	AccessTokenTypeVirtualPlugin AccessTokenType = 1
)

const (
	accessTokenKeyPrefix = "access_token"
)
const expiryDelta = 3 * time.Minute
