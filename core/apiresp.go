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
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

type APIResp struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	RawBody    []byte      `json:"-"`
}

func (resp APIResp) JSONUnmarshalBody(val interface{}, config *Config) error {
	if !strings.Contains(resp.Header.Get(contentTypeHeader), contentTypeJson) {
		return fmt.Errorf("response content-type not json, response: %v", resp)
	}
	return json.Unmarshal(resp.RawBody, val)
}

func (resp APIResp) RequestId() string {
	logID := resp.Header.Get(HTTPHeaderKeyLogID)
	if logID != "" {
		return logID
	}
	return resp.Header.Get(HTTPHeaderKeyRequestID)
}

type CodeError struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Err     struct {
		Code int    `json:"code,omitempty"`
		Msg  string `json:"msg,omitempty"`
	} `json:"err"`
}

func (ce CodeError) Success() bool {
	return ce.ErrCode == 0
}

func (ce CodeError) Code() int {
	return ce.ErrCode
}

func (ce CodeError) Error() string {
	return ce.String()
}

func (ce CodeError) String() string {
	sb := strings.Builder{}
	sb.WriteString("msg:")
	sb.WriteString(ce.ErrMsg)
	sb.WriteString(",code:")
	sb.WriteString(strconv.Itoa(ce.ErrCode))
	sb.WriteString(",msg_detail:")
	sb.WriteString(ce.Err.Msg)
	sb.WriteString(",code_detail:")
	sb.WriteString(strconv.Itoa(ce.Err.Code))
	return sb.String()
}

type AccessTokenErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (ce AccessTokenErr) Error() string {
	return ce.String()
}

func (ce AccessTokenErr) String() string {
	sb := strings.Builder{}
	sb.WriteString("msg:")
	sb.WriteString(ce.Msg)
	sb.WriteString(",code:")
	sb.WriteString(strconv.Itoa(ce.Code))
	return sb.String()
}

func FileNameByHeader(header http.Header) string {
	filename := ""
	_, media, _ := mime.ParseMediaType(header.Get("Content-Disposition"))
	if len(media) > 0 {
		filename = media["filename"]
	}
	return filename
}
