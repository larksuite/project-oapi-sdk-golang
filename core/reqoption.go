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
	RequestId    string
	Header       http.Header
	FileUpload   bool
	FileDownload bool
}

type RequestOptionFunc func(option *RequestOption)

func WithRequestId(requestId string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.RequestId = requestId
	}
}

func WithHeaders(header http.Header) RequestOptionFunc {
	return func(option *RequestOption) {
		for k, valueList := range header {
			for _, value := range valueList {
				option.Header.Add(k, value)
			}
		}
	}
}

func WithAccessToken(accessToken string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.Header.Set(HTTPHeaderAccessToken, accessToken)
	}
}

func WithUserKey(userKey string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.Header.Set(HTTPHeaderUserKey, userKey)
	}
}

func WithIdemUUID(uuid string) RequestOptionFunc {
	return func(option *RequestOption) {
		option.Header.Set(HTTPHeaderIdemUUID, uuid)
	}
}

func WithFileUpload() RequestOptionFunc {
	return func(option *RequestOption) {
		option.FileUpload = true
	}
}

func WithFileDownload() RequestOptionFunc {
	return func(option *RequestOption) {
		option.FileDownload = true
	}
}
