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

package file

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

// 自建链路使用的头名与常量。
// core 内同名常量（userAgentHeader/contentTypeHeader）未导出，此处以字面量声明，避免改 core。
const (
	headerUserAgent   = "User-Agent"
	headerContentType = "Content-Type"
	userAgentValue    = "oapi-sdk-go-file-stream"
	octetStream       = "application/octet-stream"
)

// buildURL 复刻 core.ReqTranslator.pathRebuild + BaseUrl 前缀 + query 编码。
func (a *FileService) buildURL(apiPath string, pathParams core.PathParams, query core.QueryParams) (string, error) {
	var segs []string
	for _, p := range strings.Split(apiPath, "/") {
		if strings.HasPrefix(p, ":") {
			name := p[1:]
			v := pathParams.Get(name)
			if v == "" {
				return "", fmt.Errorf("http path:%s, name:%s, value is empty", apiPath, name)
			}
			segs = append(segs, url.PathEscape(v))
			continue
		}
		segs = append(segs, p)
	}
	newPath := strings.Join(segs, "/")
	if strings.Index(newPath, "http") != 0 {
		newPath = a.config.BaseUrl + newPath
	}
	if q := url.Values(query).Encode(); q != "" {
		newPath = newPath + "?" + q
	}
	return newPath, nil
}

// newRequest 组装通用请求（header + 鉴权），body 直传 io.Reader 不经内存缓冲。
// header 装配顺序对齐 core.newHTTPRequest：option.Header -> config.Header -> UA -> Content-Type -> 鉴权。
func (a *FileService) newRequest(ctx context.Context, method, rawURL, contentType string,
	body io.Reader, bodyLen int64, options ...core.RequestOptionFunc) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, rawURL, body)
	if err != nil {
		return nil, err
	}
	if bodyLen >= 0 {
		req.ContentLength = bodyLen // 避免 chunked，明确告知长度
	}
	option := &core.RequestOption{Header: make(http.Header)}
	for _, f := range options {
		f(option)
	}
	for k, vs := range option.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	for k, vs := range a.config.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	req.Header.Set(headerUserAgent, userAgentValue)
	if contentType != "" {
		req.Header.Set(headerContentType, contentType)
	}
	if a.config.EnableTokenCache && req.Header.Get(core.HTTPHeaderAccessToken) == "" {
		token, err := core.GetAccessToken(ctx, a.config)
		if err != nil {
			return nil, err
		}
		req.Header.Set(core.HTTPHeaderAccessToken, token)
	}
	return req, nil
}

func (a *FileService) httpClient() core.HttpClient {
	if a.config.HttpClient != nil {
		return a.config.HttpClient
	}
	return http.DefaultClient
}

// doStreamRequest 上传方向：body 流式发送；响应体小（JSON），ReadAll 进 APIResp 复用既有解析。
func (a *FileService) doStreamRequest(ctx context.Context, method, apiPath string,
	pathParams core.PathParams, query core.QueryParams, contentType string,
	body io.Reader, bodyLen int64, options ...core.RequestOptionFunc) (*core.APIResp, error) {
	rawURL, err := a.buildURL(apiPath, pathParams, query)
	if err != nil {
		return nil, err
	}
	req, err := a.newRequest(ctx, method, rawURL, contentType, body, bodyLen, options...)
	if err != nil {
		return nil, err
	}
	resp, err := a.httpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &core.APIResp{StatusCode: resp.StatusCode, Header: resp.Header, RawBody: raw}, nil
}

// doStreamDownload 下载方向：不读 body，直接把 *http.Response 交回调用方（由其决定读/Close）。
func (a *FileService) doStreamDownload(ctx context.Context, method, apiPath string,
	pathParams core.PathParams, query core.QueryParams, options ...core.RequestOptionFunc) (*http.Response, error) {
	rawURL, err := a.buildURL(apiPath, pathParams, query)
	if err != nil {
		return nil, err
	}
	req, err := a.newRequest(ctx, method, rawURL, "", nil, -1, options...)
	if err != nil {
		return nil, err
	}
	return a.httpClient().Do(req)
}
