package comment

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type CreateCommentReq struct {
	apiReq *core.APIReq
}
type CreateCommentReqBody struct {
    Content  *string `json:"content,omitempty"`
    RichText  interface{} `json:"rich_text,omitempty"`
    DocRichText  *DocTypeRichText `json:"doc_rich_text,omitempty"`
}
type CreateCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type CreateCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateCommentReqBuilder() *CreateCommentReqBuilder {
	builder := &CreateCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateCommentReqBody{},
	}
	return builder
}

func (builder *CreateCommentReqBuilder) ProjectKey(projectKey string) *CreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateCommentReqBuilder) WorkItemID(workItemID int64) *CreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *CreateCommentReqBuilder) Content(content string) *CreateCommentReqBuilder {
	builder.apiReq.Body.(*CreateCommentReqBody).Content = &content
	return builder
}


func (builder *CreateCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateCommentReqBuilder) RichText(richText interface{}) *CreateCommentReqBuilder {
	builder.apiReq.Body.(*CreateCommentReqBody).RichText = richText
	return builder
}

func (builder *CreateCommentReqBuilder) DocRichText(docRichText *DocTypeRichText) *CreateCommentReqBuilder {
	builder.apiReq.Body.(*CreateCommentReqBody).DocRichText = docRichText
	return builder
}
func (builder *CreateCommentReqBuilder) Build() *CreateCommentReq {
	req := &CreateCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateFileCommentReq struct {
	apiReq *core.APIReq
}
type CreateFileCommentReqBody struct {
    FileToken  *string `json:"file_token,omitempty"`
}
type CreateFileCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type CreateFileCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateFileCommentReqBuilder() *CreateFileCommentReqBuilder {
	builder := &CreateFileCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateFileCommentReqBody{},
	}
	return builder
}

func (builder *CreateFileCommentReqBuilder) ProjectKey(projectKey string) *CreateFileCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateFileCommentReqBuilder) WorkItemID(workItemID int64) *CreateFileCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *CreateFileCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateFileCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateFileCommentReqBuilder) FileToken(fileToken string) *CreateFileCommentReqBuilder {
	builder.apiReq.Body.(*CreateFileCommentReqBody).FileToken = &fileToken
	return builder
}

func (builder *CreateFileCommentReqBuilder) Build() *CreateFileCommentReq {
	req := &CreateFileCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteCommentReq struct {
	apiReq *core.APIReq
}
type DeleteCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteCommentReqBuilder() *DeleteCommentReqBuilder {
	builder := &DeleteCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *DeleteCommentReqBuilder) ProjectKey(projectKey string) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *DeleteCommentReqBuilder) WorkItemID(workItemID int64) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *DeleteCommentReqBuilder) CommentID(commentID int64) *DeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("comment_id", fmt.Sprint(commentID))
	return builder
}

func (builder *DeleteCommentReqBuilder) Build() *DeleteCommentReq {
	req := &DeleteCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListCommentsReq struct {
	apiReq *core.APIReq
}
type ListCommentsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []CommentForOpenAPI         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type ListCommentsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListCommentsReqBuilder() *ListCommentsReqBuilder {
	builder := &ListCommentsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *ListCommentsReqBuilder) ProjectKey(projectKey string) *ListCommentsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *ListCommentsReqBuilder) WorkItemID(workItemID int64) *ListCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *ListCommentsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ListCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *ListCommentsReqBuilder) PageSize(pageSize int64) *ListCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}


func (builder *ListCommentsReqBuilder) PageNum(pageNum int64) *ListCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_num", fmt.Sprint(pageNum))
	return builder
}

func (builder *ListCommentsReqBuilder) Build() *ListCommentsReq {
	req := &ListCommentsReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListFileCommentsReq struct {
	apiReq *core.APIReq
}
type ListFileCommentsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []FileCommentForOpenAPI         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type ListFileCommentsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListFileCommentsReqBuilder() *ListFileCommentsReqBuilder {
	builder := &ListFileCommentsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *ListFileCommentsReqBuilder) ProjectKey(projectKey string) *ListFileCommentsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *ListFileCommentsReqBuilder) WorkItemID(workItemID int64) *ListFileCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *ListFileCommentsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ListFileCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *ListFileCommentsReqBuilder) PageSize(pageSize int64) *ListFileCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}


func (builder *ListFileCommentsReqBuilder) PageNum(pageNum int64) *ListFileCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_num", fmt.Sprint(pageNum))
	return builder
}

func (builder *ListFileCommentsReqBuilder) Build() *ListFileCommentsReq {
	req := &ListFileCommentsReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateCommentReq struct {
	apiReq *core.APIReq
}
type UpdateCommentReqBody struct {
    Content  *string `json:"content,omitempty"`
    RichText  interface{} `json:"rich_text,omitempty"`
}
type UpdateCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateCommentReqBuilder() *UpdateCommentReqBuilder {
	builder := &UpdateCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateCommentReqBody{},
	}
	return builder
}

func (builder *UpdateCommentReqBuilder) ProjectKey(projectKey string) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateCommentReqBuilder) WorkItemID(workItemID int64) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateCommentReqBuilder) CommentID(commentID int64) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("comment_id", fmt.Sprint(commentID))
	return builder
}


func (builder *UpdateCommentReqBuilder) Content(content string) *UpdateCommentReqBuilder {
	builder.apiReq.Body.(*UpdateCommentReqBody).Content = &content
	return builder
}


func (builder *UpdateCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateCommentReqBuilder) RichText(richText interface{}) *UpdateCommentReqBuilder {
	builder.apiReq.Body.(*UpdateCommentReqBody).RichText = richText
	return builder
}
func (builder *UpdateCommentReqBuilder) Build() *UpdateCommentReq {
	req := &UpdateCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

