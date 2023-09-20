/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type ReqTranslator struct {
}

func (translator *ReqTranslator) translate(ctx context.Context, req *APIReq, config *Config, option *RequestOption) (*http.Request, error) {
	body := req.Body
	skipAuth := req.SkipAuth
	contentType, rawBody, err := translator.payload(body)
	if err != nil {
		return nil, err
	}

	newPath, err := translator.pathRebuild(req)
	if err != nil {
		return nil, err
	}
	if strings.Index(newPath, "http") != 0 {
		newPath = fmt.Sprintf("%s%s", config.BaseUrl, newPath)
	}

	queryPath := req.QueryParams.Encode()
	if queryPath != "" {
		newPath = fmt.Sprintf("%s?%s", newPath, queryPath)
	}

	req1, err := translator.newHTTPRequest(ctx, req, newPath, contentType, rawBody, skipAuth, option, config)
	if err != nil {
		return nil, err
	}
	return req1, nil
}

func (translator *ReqTranslator) pathRebuild(req *APIReq) (string, error) {
	var pathSegs []string
	for _, p := range strings.Split(req.ApiPath, "/") {
		if strings.Index(p, ":") == 0 {
			varName := p[1:]
			v, ok := req.PathParams[varName]
			if !ok {
				return "", fmt.Errorf("http path:%s, name: %s, not found value", req.ApiPath, varName)
			}
			val := fmt.Sprint(v)
			if val == "" {
				return "", fmt.Errorf("http path:%s, name: %s, value is empty", req.ApiPath, varName)
			}
			val = url.PathEscape(val)
			pathSegs = append(pathSegs, val)
			continue
		}
		pathSegs = append(pathSegs, p)
	}
	return strings.Join(pathSegs, "/"), nil
}

func (translator *ReqTranslator) newHTTPRequest(ctx context.Context,
	req *APIReq, url, contentType string, body []byte,
	skipAuth bool, option *RequestOption, config *Config) (*http.Request, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, req.HttpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if option.RequestId != "" {
		httpRequest.Header.Add(customRequestId, option.RequestId)
	}
	for k, vs := range option.Header {
		for _, v := range vs {
			httpRequest.Header.Add(k, v)
		}
	}
	for k, vs := range config.Header {
		for _, v := range vs {
			httpRequest.Header.Add(k, v)
		}
	}
	httpRequest.Header.Set(userAgentHeader, userAgent())
	if contentType != "" {
		httpRequest.Header.Set(contentTypeHeader, contentType)
	}
	if !skipAuth {
		httpHeaderAccessToken := httpRequest.Header.Get(HTTPHeaderAccessToken)
		if config.EnableTokenCache && len(httpHeaderAccessToken) == 0 {
			accessToken, err := tokenManager.getAccessToken(ctx, config)
			if err != nil {
				return nil, err
			}
			httpRequest.Header.Set(HTTPHeaderAccessToken, accessToken)
		}
	}
	return httpRequest, nil
}

func (translator *ReqTranslator) payload(body interface{}) (string, []byte, error) {
	if fd, ok := body.(*FormData); ok {
		return fd.content()
	}
	contentType := defaultContentType
	if body == nil {
		return contentType, nil, nil
	}
	bs, err := json.Marshal(body)
	return contentType, bs, err
}

func NewFormdata() *FormData {
	return &FormData{}
}

func (fd *FormData) AddField(field string, val interface{}) *FormData {
	if fd.fields == nil {
		fd.fields = map[string]interface{}{}
	}
	fd.fields[field] = val
	return fd
}

func (fd *FormData) AddFile(field string, r io.Reader) *FormData {
	return fd.AddField(field, r)
}

func (fd *FormData) content() (string, []byte, error) {
	if fd.data != nil {
		return fd.data.contentType, fd.data.content, nil
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, val := range fd.fields {
		if r, ok := val.(io.Reader); ok {
			part, err := writer.CreateFormFile(key, "unknown-file")
			if err != nil {
				return "", nil, err
			}
			_, err = io.Copy(part, r)
			if err != nil {
				return "", nil, err
			}
			continue
		}
		err := writer.WriteField(key, fmt.Sprint(val))
		if err != nil {
			return "", nil, err
		}
	}
	contentType := writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		return "", nil, err
	}
	fd.data = &struct {
		content     []byte
		contentType string
	}{content: buf.Bytes(), contentType: contentType}
	return fd.data.contentType, fd.data.content, nil
}

type FormData struct {
	fields map[string]interface{}
	data   *struct {
		content     []byte
		contentType string
	}
}
