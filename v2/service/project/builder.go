package project

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type ListProjectTeamReq struct {
	apiReq *core.APIReq
}

type ListProjectTeamResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []Team         `json:"data"`
	
}

type ListProjectTeamReqBuilder struct {
	apiReq *core.APIReq
}

func NewListProjectTeamReqBuilder() *ListProjectTeamReqBuilder {
	builder := &ListProjectTeamReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *ListProjectTeamReqBuilder) ProjectKey(projectKey string) *ListProjectTeamReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *ListProjectTeamReqBuilder) Build() *ListProjectTeamReq {
	req := &ListProjectTeamReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIBatchQueryProjectInfoReq struct {
	apiReq *core.APIReq
}
type OAPIBatchQueryProjectInfoReqBody struct {

    UserKey  *string `json:"user_key,omitempty"`

    SimpleNames  []string `json:"simple_names,omitempty"`

    TenantGroupID  *int64 `json:"tenant_group_id,omitempty"`

}

type OAPIBatchQueryProjectInfoResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       map[string]Project         `json:"data"`
	
}

type OAPIBatchQueryProjectInfoReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIBatchQueryProjectInfoReqBuilder() *OAPIBatchQueryProjectInfoReqBuilder {
	builder := &OAPIBatchQueryProjectInfoReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIBatchQueryProjectInfoReqBody{},
	}
	return builder
}

func (builder *OAPIBatchQueryProjectInfoReqBuilder) UserKey(userKey string) *OAPIBatchQueryProjectInfoReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryProjectInfoReqBody).UserKey = &userKey
	return builder
}


func (builder *OAPIBatchQueryProjectInfoReqBuilder) SimpleNames(simpleNames []string) *OAPIBatchQueryProjectInfoReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryProjectInfoReqBody).SimpleNames = simpleNames
	return builder
}

func (builder *OAPIBatchQueryProjectInfoReqBuilder) TenantGroupID(tenantGroupID *int64) *OAPIBatchQueryProjectInfoReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryProjectInfoReqBody).TenantGroupID = tenantGroupID
	return builder
}
func (builder *OAPIBatchQueryProjectInfoReqBuilder) Build() *OAPIBatchQueryProjectInfoReq {
	req := &OAPIBatchQueryProjectInfoReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateCommentReq struct {
	apiReq *core.APIReq
}
type OAPICreateCommentReqBody struct {

    Content  *string `json:"content,omitempty"`

    RichText  interface{} `json:"rich_text,omitempty"`

}

type OAPICreateCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type OAPICreateCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateCommentReqBuilder() *OAPICreateCommentReqBuilder {
	builder := &OAPICreateCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateCommentReqBody{},
	}
	return builder
}

func (builder *OAPICreateCommentReqBuilder) ProjectKey(projectKey string) *OAPICreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *OAPICreateCommentReqBuilder) WorkItemID(workItemID *int64) *OAPICreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}

func (builder *OAPICreateCommentReqBuilder) Content(content string) *OAPICreateCommentReqBuilder {
	builder.apiReq.Body.(*OAPICreateCommentReqBody).Content = &content
	return builder
}


func (builder *OAPICreateCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *OAPICreateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *OAPICreateCommentReqBuilder) RichText(richText interface{}) *OAPICreateCommentReqBuilder {
	builder.apiReq.Body.(*OAPICreateCommentReqBody).RichText = richText
	return builder
}
func (builder *OAPICreateCommentReqBuilder) Build() *OAPICreateCommentReq {
	req := &OAPICreateCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteCommentReq struct {
	apiReq *core.APIReq
}

type OAPIDeleteCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteCommentReqBuilder() *OAPIDeleteCommentReqBuilder {
	builder := &OAPIDeleteCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *OAPIDeleteCommentReqBuilder) ProjectKey(projectKey string) *OAPIDeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *OAPIDeleteCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *OAPIDeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *OAPIDeleteCommentReqBuilder) WorkItemID(workItemID *int64) *OAPIDeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}

func (builder *OAPIDeleteCommentReqBuilder) CommentID(commentID *int64) *OAPIDeleteCommentReqBuilder {
	builder.apiReq.PathParams.Set("comment_id", fmt.Sprint(*commentID))
	return builder
}
func (builder *OAPIDeleteCommentReqBuilder) Build() *OAPIDeleteCommentReq {
	req := &OAPIDeleteCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIListCommentsReq struct {
	apiReq *core.APIReq
}

type OAPIListCommentsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []CommentForOpenAPI         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type OAPIListCommentsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIListCommentsReqBuilder() *OAPIListCommentsReqBuilder {
	builder := &OAPIListCommentsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *OAPIListCommentsReqBuilder) ProjectKey(projectKey string) *OAPIListCommentsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *OAPIListCommentsReqBuilder) WorkItemID(workItemID *int64) *OAPIListCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}

func (builder *OAPIListCommentsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *OAPIListCommentsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *OAPIListCommentsReqBuilder) PageSize(pageSize *int64) *OAPIListCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(*pageSize))
	return builder
}

func (builder *OAPIListCommentsReqBuilder) PageNum(pageNum *int64) *OAPIListCommentsReqBuilder {
	builder.apiReq.QueryParams.Set("page_num", fmt.Sprint(*pageNum))
	return builder
}
func (builder *OAPIListCommentsReqBuilder) Build() *OAPIListCommentsReq {
	req := &OAPIListCommentsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIPageGetWorkItemOpRecordReq struct {
	apiReq *core.APIReq
}
type OAPIPageGetWorkItemOpRecordReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

    StartFrom  *string `json:"start_from,omitempty"`

    Operator  []string `json:"operator,omitempty"`

    OperatorType  []string `json:"operator_type,omitempty"`

    SourceType  []string `json:"source_type,omitempty"`

    Source  []string `json:"source,omitempty"`

    OperationType  []string `json:"operation_type,omitempty"`

    Start  *int64 `json:"start,omitempty"`

    End  *int64 `json:"end,omitempty"`

    OpRecordModule  []string `json:"op_record_module,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

}

type OAPIPageGetWorkItemOpRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	HasMore       *bool         `json:"has_more"`
	
	StartFrom       *string         `json:"start_from"`
	
	OpRecords       []OAPIOperationRecord         `json:"op_records"`
	
	Total       *int64         `json:"total"`
	
}

type OAPIPageGetWorkItemOpRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIPageGetWorkItemOpRecordReqBuilder() *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder := &OAPIPageGetWorkItemOpRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIPageGetWorkItemOpRecordReqBody{},
	}
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) ProjectKey(projectKey string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) WorkItemIDs(workItemIDs []int64) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).WorkItemIDs = workItemIDs
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) StartFrom(startFrom string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).StartFrom = &startFrom
	return builder
}


func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) Operator(operator []string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).Operator = operator
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) OperatorType(operatorType []string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).OperatorType = operatorType
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) SourceType(sourceType []string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).SourceType = sourceType
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) Source(source []string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).Source = source
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) OperationType(operationType []string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).OperationType = operationType
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) Start(start *int64) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).Start = start
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) End(end *int64) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).End = end
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) OpRecordModule(opRecordModule []string) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).OpRecordModule = opRecordModule
	return builder
}

func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) PageSize(pageSize *int64) *OAPIPageGetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*OAPIPageGetWorkItemOpRecordReqBody).PageSize = pageSize
	return builder
}
func (builder *OAPIPageGetWorkItemOpRecordReqBuilder) Build() *OAPIPageGetWorkItemOpRecordReq {
	req := &OAPIPageGetWorkItemOpRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryProjectsReq struct {
	apiReq *core.APIReq
}
type OAPIQueryProjectsReqBody struct {

    UserKey  *string `json:"user_key,omitempty"`

    TenantGroupID  *int64 `json:"tenant_group_id,omitempty"`

    AssetKey  *string `json:"asset_key,omitempty"`

    Order  []string `json:"order,omitempty"`

}

type OAPIQueryProjectsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []string         `json:"data"`
	
}

type OAPIQueryProjectsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryProjectsReqBuilder() *OAPIQueryProjectsReqBuilder {
	builder := &OAPIQueryProjectsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryProjectsReqBody{},
	}
	return builder
}

func (builder *OAPIQueryProjectsReqBuilder) UserKey(userKey string) *OAPIQueryProjectsReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectsReqBody).UserKey = &userKey
	return builder
}


func (builder *OAPIQueryProjectsReqBuilder) TenantGroupID(tenantGroupID *int64) *OAPIQueryProjectsReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectsReqBody).TenantGroupID = tenantGroupID
	return builder
}

func (builder *OAPIQueryProjectsReqBuilder) AssetKey(assetKey string) *OAPIQueryProjectsReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectsReqBody).AssetKey = &assetKey
	return builder
}


func (builder *OAPIQueryProjectsReqBuilder) Order(order []string) *OAPIQueryProjectsReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectsReqBody).Order = order
	return builder
}
func (builder *OAPIQueryProjectsReqBuilder) Build() *OAPIQueryProjectsReq {
	req := &OAPIQueryProjectsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateCommentReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateCommentReqBody struct {

    Content  *string `json:"content,omitempty"`

    RichText  interface{} `json:"rich_text,omitempty"`

}

type OAPIUpdateCommentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateCommentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateCommentReqBuilder() *OAPIUpdateCommentReqBuilder {
	builder := &OAPIUpdateCommentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateCommentReqBody{},
	}
	return builder
}

func (builder *OAPIUpdateCommentReqBuilder) ProjectKey(projectKey string) *OAPIUpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *OAPIUpdateCommentReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}

func (builder *OAPIUpdateCommentReqBuilder) CommentID(commentID *int64) *OAPIUpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("comment_id", fmt.Sprint(*commentID))
	return builder
}

func (builder *OAPIUpdateCommentReqBuilder) Content(content string) *OAPIUpdateCommentReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateCommentReqBody).Content = &content
	return builder
}


func (builder *OAPIUpdateCommentReqBuilder) WorkItemTypeKey(workItemTypeKey string) *OAPIUpdateCommentReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *OAPIUpdateCommentReqBuilder) RichText(richText interface{}) *OAPIUpdateCommentReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateCommentReqBody).RichText = richText
	return builder
}
func (builder *OAPIUpdateCommentReqBuilder) Build() *OAPIUpdateCommentReq {
	req := &OAPIUpdateCommentReq{}
	req.apiReq = builder.apiReq
	return req
}

