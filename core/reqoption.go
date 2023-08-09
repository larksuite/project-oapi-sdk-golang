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

import "net/http"

type RequestOption struct {
	PluginAccessToken        string
	VirtualPluginAccessToken string
	UserPluginAccessToken    string
	Header                   http.Header
}

type RequestOptionFunc func(option *RequestOption)

func WithHeaders(header http.Header) RequestOptionFunc {
	return func(option *RequestOption) {
		option.Header = header
	}
}

func WithPluginAccessToken(pluginAccessToken string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.PluginAccessToken = pluginAccessToken
	}
}

func WithVirtualPluginAccessToken(virtualPluginAccessToken string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.VirtualPluginAccessToken = virtualPluginAccessToken
	}
}

func WithUserPluginAccessToken(userPluginAccessToken string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.UserPluginAccessToken = userPluginAccessToken
	}
}
