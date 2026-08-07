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
func (builder *CreateCommentReqBuilder) Build() *CreateCommentReq {
	req := &CreateCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateCommentNewReq struct {
	apiReq *core.APIReq
}
type CreateCommentNewReqBody struct {
    Object  *BizObject `json:"object,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
    Content  *MutateCreateContent `json:"content,omitempty"`
}
type CreateCommentNewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *CreateCommentNewRespData        `json:"data,omitempty"`
}

type CreateCommentNewRespData struct {
	CommentID       *string         `json:"comment_id,omitempty"`
}

type CreateCommentNewReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateCommentNewReqBuilder() *CreateCommentNewReqBuilder {
	builder := &CreateCommentNewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateCommentNewReqBody{},
	}
	return builder
}

func (builder *CreateCommentNewReqBuilder) Object(object *BizObject) *CreateCommentNewReqBuilder {
	builder.apiReq.Body.(*CreateCommentNewReqBody).Object = object
	return builder
}

func (builder *CreateCommentNewReqBuilder) ProjectKey(projectKey string) *CreateCommentNewReqBuilder {
	builder.apiReq.Body.(*CreateCommentNewReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *CreateCommentNewReqBuilder) Content(content *MutateCreateContent) *CreateCommentNewReqBuilder {
	builder.apiReq.Body.(*CreateCommentNewReqBody).Content = content
	return builder
}
func (builder *CreateCommentNewReqBuilder) Build() *CreateCommentNewReq {
	req := &CreateCommentNewReq{}
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

type DeleteCommentNewReq struct {
	apiReq *core.APIReq
}
type DeleteCommentNewReqBody struct {
    CommentID  *string `json:"comment_id,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
}
type DeleteCommentNewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteCommentNewReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteCommentNewReqBuilder() *DeleteCommentNewReqBuilder {
	builder := &DeleteCommentNewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &DeleteCommentNewReqBody{},
	}
	return builder
}

func (builder *DeleteCommentNewReqBuilder) CommentID(commentID string) *DeleteCommentNewReqBuilder {
	builder.apiReq.Body.(*DeleteCommentNewReqBody).CommentID = &commentID
	return builder
}


func (builder *DeleteCommentNewReqBuilder) ProjectKey(projectKey string) *DeleteCommentNewReqBuilder {
	builder.apiReq.Body.(*DeleteCommentNewReqBody).ProjectKey = &projectKey
	return builder
}

func (builder *DeleteCommentNewReqBuilder) Build() *DeleteCommentNewReq {
	req := &DeleteCommentNewReq{}
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

type QueryCommentNewReq struct {
	apiReq *core.APIReq
}
type QueryCommentNewReqBody struct {
    Object  *BizObject `json:"object,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
    NeedRichTextMarkDown  *bool `json:"need_rich_text_mark_down,omitempty"`
    Paginator  *Paginator `json:"paginator,omitempty"`
}
type QueryCommentNewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *QueryCommentNewRespData        `json:"data,omitempty"`
}

type QueryCommentNewRespData struct {
	Comments       []Comment         `json:"comments,omitempty"`
	NextCursor       *string         `json:"next_cursor,omitempty"`
}

type QueryCommentNewReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryCommentNewReqBuilder() *QueryCommentNewReqBuilder {
	builder := &QueryCommentNewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryCommentNewReqBody{},
	}
	return builder
}

func (builder *QueryCommentNewReqBuilder) Object(object *BizObject) *QueryCommentNewReqBuilder {
	builder.apiReq.Body.(*QueryCommentNewReqBody).Object = object
	return builder
}

func (builder *QueryCommentNewReqBuilder) ProjectKey(projectKey string) *QueryCommentNewReqBuilder {
	builder.apiReq.Body.(*QueryCommentNewReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *QueryCommentNewReqBuilder) NeedRichTextMarkDown(needRichTextMarkDown bool) *QueryCommentNewReqBuilder {
	builder.apiReq.Body.(*QueryCommentNewReqBody).NeedRichTextMarkDown = &needRichTextMarkDown
	return builder
}


func (builder *QueryCommentNewReqBuilder) Paginator(paginator *Paginator) *QueryCommentNewReqBuilder {
	builder.apiReq.Body.(*QueryCommentNewReqBody).Paginator = paginator
	return builder
}
func (builder *QueryCommentNewReqBuilder) Build() *QueryCommentNewReq {
	req := &QueryCommentNewReq{}
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

type UpdateCommentNewReq struct {
	apiReq *core.APIReq
}
type UpdateCommentNewReqBody struct {
    CommentID  *string `json:"comment_id,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
    Content  *MutateUpdateContent `json:"content,omitempty"`
}
type UpdateCommentNewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateCommentNewReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateCommentNewReqBuilder() *UpdateCommentNewReqBuilder {
	builder := &UpdateCommentNewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateCommentNewReqBody{},
	}
	return builder
}

func (builder *UpdateCommentNewReqBuilder) CommentID(commentID string) *UpdateCommentNewReqBuilder {
	builder.apiReq.Body.(*UpdateCommentNewReqBody).CommentID = &commentID
	return builder
}


func (builder *UpdateCommentNewReqBuilder) ProjectKey(projectKey string) *UpdateCommentNewReqBuilder {
	builder.apiReq.Body.(*UpdateCommentNewReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *UpdateCommentNewReqBuilder) Content(content *MutateUpdateContent) *UpdateCommentNewReqBuilder {
	builder.apiReq.Body.(*UpdateCommentNewReqBody).Content = content
	return builder
}
func (builder *UpdateCommentNewReqBuilder) Build() *UpdateCommentNewReq {
	req := &UpdateCommentNewReq{}
	req.apiReq = builder.apiReq
	return req
}

