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

package attachment

import (
	"fmt"
	"io"
	"net/http"

	"code.byted.org/bits/project-oapi-sdk-golang/core"
)

// 添加附件接口参数构造器
type UploadAttachmentReqBuilder struct {
	apiReq *core.ApiReq
}

// 添加附件接口参数
type UploadAttachmentReq struct {
	apiReq *core.ApiReq
}

// 添加附件接口响应
type UploadAttachmentResp struct {
	*core.ApiResp `json:"-"`
	core.CodeError
}

func NewUploadAttachmentReqBuilder() *UploadAttachmentReqBuilder {
	builder := &UploadAttachmentReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams: core.PathParams{},
		Body:       &core.Formdata{},
		Header:     make(http.Header),
	}
	return builder
}

// 可选，当选择使用插件身份凭证的时候，需要额外必选指定接口调用的用户user_key
func (builder *UploadAttachmentReqBuilder) AccessUser(userKey string) *UploadAttachmentReqBuilder {
	builder.apiReq.Header.Add("X-USER-KEY", fmt.Sprint(userKey))
	return builder
}

// 可选，写类型接口的幂等串，可以不设置，设置后会进行同一个X-PLUGIN-TOKEN下同一接口的幂等判断
func (builder *UploadAttachmentReqBuilder) UUID(uuid string) *UploadAttachmentReqBuilder {
	builder.apiReq.Header.Add("X-IDEM-UUID", fmt.Sprint(uuid))
	return builder
}

func (builder *UploadAttachmentReqBuilder) ProjectKey(projectKey string) *UploadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *UploadAttachmentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UploadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", workItemTypeKey)
	return builder
}

func (builder *UploadAttachmentReqBuilder) WorkItemID(workItemID int64) *UploadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}

func (builder *UploadAttachmentReqBuilder) File(file io.Reader) *UploadAttachmentReqBuilder {
	builder.apiReq.Body.(*core.Formdata).AddFile("file", file)
	return builder
}

func (builder *UploadAttachmentReqBuilder) FieldKey(fieldKey string) *UploadAttachmentReqBuilder {
	builder.apiReq.Body.(*core.Formdata).AddField("field_key", fieldKey)
	return builder
}

func (builder *UploadAttachmentReqBuilder) FieldAlias(fieldAlias string) *UploadAttachmentReqBuilder {
	builder.apiReq.Body.(*core.Formdata).AddField("field_alias", fieldAlias)
	return builder
}

func (builder *UploadAttachmentReqBuilder) Build() *UploadAttachmentReq {
	req := &UploadAttachmentReq{}
	req.apiReq = builder.apiReq
	return req
}
