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
	"context"
	"fmt"
	"net/http"
	"net/url"
)

var reqTranslator ReqTranslator

func NewHttpClient(config *Config) {
	if config.HttpClient == nil {
		if config.ReqTimeout == 0 {
			config.HttpClient = http.DefaultClient
		} else {
			config.HttpClient = &http.Client{Timeout: config.ReqTimeout}
		}
	}
}

func doSend(ctx context.Context, rawRequest *http.Request, httpClient HttpClient, logger Logger) (*APIResp, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(rawRequest)
	if err != nil {
		if er, ok := err.(*url.Error); ok {
			if er.Timeout() {
				return nil, &ClientTimeoutError{msg: er.Error()}
			}
		}
		return nil, err
	}

	if resp.StatusCode == http.StatusGatewayTimeout {
		logID := resp.Header.Get(HTTPHeaderKeyLogID)
		if logID == "" {
			logID = resp.Header.Get(HTTPHeaderKeyRequestID)
		}
		logger.Info(ctx, fmt.Sprintf("req path:%s, server time out,requestId:%s",
			rawRequest.URL.RequestURI(), logID))
		return nil, &ServerTimeoutError{msg: "server time out error"}
	}
	body, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	return &APIResp{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		RawBody:    body,
	}, nil
}

func Request(ctx context.Context, req *APIReq, config *Config, options ...RequestOptionFunc) (*APIResp, error) {
	option := &RequestOption{
		Header: make(http.Header),
	}
	for _, optionFunc := range options {
		optionFunc(option)
	}

	httpReq, err := reqTranslator.translate(ctx, req, config, option)
	if err != nil {
		return nil, err
	}

	if config.LogReqAtDebug {
		config.Logger.Debug(ctx, fmt.Sprintf("req:%v", httpReq))
	} else {
		config.Logger.Debug(ctx, fmt.Sprintf("req:%s,%s", req.HttpMethod, req.ApiPath))
	}
	rawResp, err := doSend(ctx, httpReq, config.HttpClient, config.Logger)
	if err != nil {
		return nil, err
	}
	if config.LogReqAtDebug {
		config.Logger.Debug(ctx, fmt.Sprintf("resp:%v", rawResp))
	}
	return rawResp, nil

}
