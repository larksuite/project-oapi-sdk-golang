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

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

// 添加附件接口参数构造器
type UploadAttachmentReqBuilder struct {
	apiReq *core.APIReq
	body   *core.FormData
}

// 添加附件接口参数
type UploadAttachmentReq struct {
	apiReq *core.APIReq
}

// 添加附件接口响应
type UploadAttachmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

func NewUploadAttachmentReqBuilder() *UploadAttachmentReqBuilder {
	builder := &UploadAttachmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	builder.body = &core.FormData{}
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
	builder.body.AddFile("unknown-file", file)
	return builder
}

func (builder *UploadAttachmentReqBuilder) FileWithFileName(fileName string, file io.Reader) *UploadAttachmentReqBuilder {
	builder.body.AddFile(fileName, file)
	return builder
}

func (builder *UploadAttachmentReqBuilder) FieldKey(fieldKey string) *UploadAttachmentReqBuilder {
	builder.body.AddField("field_key", fieldKey)
	return builder
}

func (builder *UploadAttachmentReqBuilder) FieldAlias(fieldAlias string) *UploadAttachmentReqBuilder {
	builder.body.AddField("field_alias", fieldAlias)
	return builder
}

func (builder *UploadAttachmentReqBuilder) Build() *UploadAttachmentReq {
	req := &UploadAttachmentReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type DownloadAttachmentReqBody struct {
	UUID string `json:"uuid"`
}

type DownloadAttachmentReqBuilder struct {
	apiReq *core.APIReq
	body   *DownloadAttachmentReqBody
}

type DownloadAttachmentReq struct {
	apiReq *core.APIReq
}

type DownloadAttachmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	File     io.Reader `json:"-"`
	FileName string    `json:"-"`
}

func NewDownloadAttachmentReqBuilder() *DownloadAttachmentReqBuilder {
	builder := &DownloadAttachmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	builder.body = &DownloadAttachmentReqBody{}
	return builder
}

func (builder *DownloadAttachmentReqBuilder) ProjectKey(projectKey string) *DownloadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *DownloadAttachmentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DownloadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", workItemTypeKey)
	return builder
}

func (builder *DownloadAttachmentReqBuilder) WorkItemID(workItemID int64) *DownloadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}

func (builder *DownloadAttachmentReqBuilder) UUID(uuid string) *DownloadAttachmentReqBuilder {
	builder.body.UUID = uuid
	return builder
}

func (builder *DownloadAttachmentReqBuilder) Build() *DownloadAttachmentReq {
	req := &DownloadAttachmentReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}

type SpecialUploadAttachmentReqBuilder struct {
	apiReq *core.APIReq
	body   *core.FormData
}

type SpecialUploadAttachmentReq struct {
	apiReq *core.APIReq
}

type SpecialUploadAttachmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []string `json:"data"`
}

func NewSpecialUploadAttachmentReqBuilder() *SpecialUploadAttachmentReqBuilder {
	builder := &SpecialUploadAttachmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams: core.PathParams{},
	}
	builder.body = &core.FormData{}
	return builder
}

func (builder *SpecialUploadAttachmentReqBuilder) ProjectKey(projectKey string) *SpecialUploadAttachmentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", projectKey)
	return builder
}

func (builder *SpecialUploadAttachmentReqBuilder) File(file io.Reader) *SpecialUploadAttachmentReqBuilder {
	builder.body.AddFile("unknown-file", file)
	return builder
}

func (builder *SpecialUploadAttachmentReqBuilder) FileWithFileName(fileName string, file io.Reader) *SpecialUploadAttachmentReqBuilder {
	builder.body.AddFile(fileName, file)
	return builder
}

func (builder *SpecialUploadAttachmentReqBuilder) Build() *SpecialUploadAttachmentReq {
	req := &SpecialUploadAttachmentReq{}
	req.apiReq = builder.apiReq
	req.apiReq.Body = builder.body
	return req
}
