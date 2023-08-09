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
	"reflect"
	"strings"
)

type ReqTranslator struct {
}

func (translator *ReqTranslator) translate(ctx context.Context, req *ApiReq, config *Config, option *RequestOption) (*http.Request, error) {
	body := req.Body
	skipAuth := req.SkipAuth
	contentType, rawBody, err := translator.payload(body)
	if err != nil {
		return nil, err
	}

	// path
	var pathSegs []string
	for _, p := range strings.Split(req.ApiPath, "/") {
		if strings.Index(p, ":") == 0 {
			varName := p[1:]
			v, ok := req.PathParams[varName]
			if !ok {
				return nil, fmt.Errorf("http path:%s, name: %s, not found value", req.ApiPath, varName)
			}
			val := fmt.Sprint(v)
			if val == "" {
				return nil, fmt.Errorf("http path:%s, name: %s, value is empty", req.ApiPath, varName)
			}
			val = url.PathEscape(val)
			pathSegs = append(pathSegs, val)
			continue
		}
		pathSegs = append(pathSegs, p)
	}
	newPath := strings.Join(pathSegs, "/")
	if strings.Index(newPath, "http") != 0 {
		newPath = fmt.Sprintf("%s%s", config.BaseUrl, newPath)
	}

	queryPath := req.QueryParams.Encode()
	if queryPath != "" {
		newPath = fmt.Sprintf("%s?%s", newPath, queryPath)
	}

	req1, err := translator.newHTTPRequest(ctx, req.HttpMethod, newPath, contentType, rawBody, skipAuth, option, config)
	if err != nil {
		return nil, err
	}
	return req1, nil
}

func authorizationToHeader(req *http.Request, token string) {
	req.Header.Set("X-PLUGIN-TOKEN", token)
}

func (translator *ReqTranslator) newHTTPRequest(ctx context.Context,
	httpMethod, url, contentType string, body []byte,
	skipAuth bool, option *RequestOption, config *Config) (*http.Request, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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
		switch config.AccessTokenType {
		case AccessTokenTypePlugin:
			accessToken := option.PluginAccessToken
			if config.EnableTokenCache && accessToken == "" {
				accessToken, err = tokenManager.getPluginAccessToken(ctx, config)
				if err != nil {
					return nil, err
				}
			}
			authorizationToHeader(httpRequest, accessToken)

		case AccessTokenTypeVirtualPlugin:
			accessToken := option.VirtualPluginAccessToken
			if config.EnableTokenCache && accessToken == "" {
				accessToken, err = tokenManager.getVirtualPluginAccessToken(ctx, config)
				if err != nil {
					return nil, err
				}
			}
			authorizationToHeader(httpRequest, accessToken)

		case AccessTokenTypeUserPlugin:
			authorizationToHeader(httpRequest, option.UserPluginAccessToken)
		}
		if err != nil {
			return nil, err
		}
	}
	return httpRequest, nil
}

func (translator *ReqTranslator) getFullReqUrl(domain string, httpPath string, pathVars, queries map[string]interface{}) (string, error) {
	// path
	var pathSegs []string
	for _, p := range strings.Split(httpPath, "/") {
		if strings.Index(p, ":") == 0 {
			varName := p[1:]
			v, ok := pathVars[varName]
			if !ok {
				return "", fmt.Errorf("http path:%s, name: %s, not found value", httpPath, varName)
			}
			val := fmt.Sprint(v)
			if val == "" {
				return "", fmt.Errorf("http path:%s, name: %s, value is empty", httpPath, varName)
			}
			val = url.PathEscape(val)
			pathSegs = append(pathSegs, val)
			continue
		}
		pathSegs = append(pathSegs, p)
	}
	newPath := strings.Join(pathSegs, "/")
	if strings.Index(newPath, "http") != 0 {
		newPath = fmt.Sprintf("%s%s", domain, newPath)
	}
	// query
	query := make(url.Values)
	for k, v := range queries {
		sv := reflect.ValueOf(v)
		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			for i := 0; i < sv.Len(); i++ {
				query.Add(k, fmt.Sprint(sv.Index(i)))
			}
		} else {
			query.Set(k, fmt.Sprint(v))
		}
	}
	if len(query) > 0 {
		newPath = fmt.Sprintf("%s?%s", newPath, query.Encode())
	}
	return newPath, nil
}

func (translator *ReqTranslator) payload(body interface{}) (string, []byte, error) {
	if fd, ok := body.(*Formdata); ok {
		return fd.content()
	}
	contentType := defaultContentType
	if body == nil {
		return contentType, nil, nil
	}
	bs, err := json.Marshal(body)
	return contentType, bs, err
}

func NewFormdata() *Formdata {
	return &Formdata{}
}

func (fd *Formdata) AddField(field string, val interface{}) *Formdata {
	if fd.fields == nil {
		fd.fields = map[string]interface{}{}
	}
	fd.fields[field] = val
	return fd
}

func (fd *Formdata) AddFile(field string, r io.Reader) *Formdata {
	return fd.AddField(field, r)
}

func (fd *Formdata) content() (string, []byte, error) {
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

type Formdata struct {
	fields map[string]interface{}
	data   *struct {
		content     []byte
		contentType string
	}
}

func (translator *ReqTranslator) parseInput(input interface{}, option *RequestOption) (map[string]interface{}, map[string]interface{}, interface{}) {
	if input == nil {
		return nil, nil, nil
	}
	if _, ok := input.(*Formdata); ok {
		return nil, nil, input
	}
	var hasHTTPTag bool
	paths, queries := map[string]interface{}{}, map[string]interface{}{}
	vv := reflect.ValueOf(input)
	vt := reflect.TypeOf(input)
	if vt.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}
	if vt.Kind() != reflect.Struct {
		return nil, nil, input
	}
	var body interface{}
	for i := 0; i < vt.NumField(); i++ {
		fieldValue := vv.Field(i)
		fieldType := vt.Field(i)
		if path, ok := fieldType.Tag.Lookup("path"); ok {
			hasHTTPTag = true
			if path != "" && !isEmptyValue(fieldValue) {
				paths[path] = reflect.Indirect(fieldValue).Interface()
			}
			continue
		}
		if query, ok := fieldType.Tag.Lookup("query"); ok {
			hasHTTPTag = true
			if query != "" && !isEmptyValue(fieldValue) {
				queries[query] = reflect.Indirect(fieldValue).Interface()
			}
			continue
		}
		if _, ok := fieldType.Tag.Lookup("body"); ok {
			hasHTTPTag = true
			body = fieldValue.Interface()
		}
	}
	if !hasHTTPTag {
		body = input
		return nil, nil, body
	}
	return paths, queries, body
}

func toFormdata(body interface{}) *Formdata {
	formdata := &Formdata{}
	v := reflect.ValueOf(body)
	t := reflect.TypeOf(body)
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		if isEmptyValue(fieldValue) {
			continue
		}
		if fieldName := fieldType.Tag.Get("json"); fieldName != "" {
			fieldName = strings.TrimSuffix(fieldName, ",omitempty")
			formdata.AddField(fieldName, reflect.Indirect(fieldValue).Interface())
		}
	}
	return formdata
}
