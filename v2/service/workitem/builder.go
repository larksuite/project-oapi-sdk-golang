package workitem

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type FilterReq struct {
	apiReq *core.APIReq
}
type FilterReqBody struct {

    WorkItemName  *string `json:"work_item_name,omitempty"`

    UserKeys  []string `json:"user_keys,omitempty"`

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

    WorkItemTypeKeys  []string `json:"work_item_type_keys,omitempty"`

    CreatedAt  *TimeInterval `json:"created_at,omitempty"`

    UpdatedAt  *TimeInterval `json:"updated_at,omitempty"`

    SubStages  []string `json:"sub_stages,omitempty"`

    Businesses  []string `json:"businesses,omitempty"`

    Priorities  []string `json:"priorities,omitempty"`

    Tags  []string `json:"tags,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    WorkItemStatus  []WorkItemStatus `json:"work_item_status,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

    SearchID  *string `json:"search_id,omitempty"`

}

type FilterResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type FilterReqBuilder struct {
	apiReq *core.APIReq
}

func NewFilterReqBuilder() *FilterReqBuilder {
	builder := &FilterReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &FilterReqBody{},
	}
	return builder
}
func (builder *FilterReqBuilder) ProjectKey(projectKey *string) *FilterReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *FilterReqBuilder) WorkItemName(workItemName *string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).WorkItemName = workItemName
	return builder
}
func (builder *FilterReqBuilder) UserKeys(userKeys []string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).UserKeys = userKeys
	return builder
}
func (builder *FilterReqBuilder) WorkItemIDs(workItemIDs []int64) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *FilterReqBuilder) WorkItemTypeKeys(workItemTypeKeys []string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).WorkItemTypeKeys = workItemTypeKeys
	return builder
}
func (builder *FilterReqBuilder) CreatedAt(createdAt *TimeInterval) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).CreatedAt = createdAt
	return builder
}
func (builder *FilterReqBuilder) UpdatedAt(updatedAt *TimeInterval) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).UpdatedAt = updatedAt
	return builder
}
func (builder *FilterReqBuilder) SubStages(subStages []string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).SubStages = subStages
	return builder
}
func (builder *FilterReqBuilder) Businesses(businesses []string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).Businesses = businesses
	return builder
}
func (builder *FilterReqBuilder) Priorities(priorities []string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).Priorities = priorities
	return builder
}
func (builder *FilterReqBuilder) Tags(tags []string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).Tags = tags
	return builder
}
func (builder *FilterReqBuilder) PageNum(pageNum *int64) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).PageNum = pageNum
	return builder
}
func (builder *FilterReqBuilder) PageSize(pageSize *int64) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).PageSize = pageSize
	return builder
}
func (builder *FilterReqBuilder) WorkItemStatus(workItemStatus []WorkItemStatus) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).WorkItemStatus = workItemStatus
	return builder
}
func (builder *FilterReqBuilder) Expand(expand *Expand) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).Expand = expand
	return builder
}
func (builder *FilterReqBuilder) SearchID(searchID *string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).SearchID = searchID
	return builder
}
func (builder *FilterReqBuilder) Build() *FilterReq {
	req := &FilterReq{}
	req.apiReq = builder.apiReq
	return req
}

type FilterAcrossProjectReq struct {
	apiReq *core.APIReq
}
type FilterAcrossProjectReqBody struct {

    SearchUser  *SearchUser `json:"search_user,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    CreatedAt  *TimeInterval `json:"created_at,omitempty"`

    UpdatedAt  *TimeInterval `json:"updated_at,omitempty"`

    WorkItemStatus  []WorkItemStatus `json:"work_item_status,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    TenantGroupID  *int64 `json:"tenant_group_id,omitempty"`

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

    Businesses  []string `json:"businesses,omitempty"`

    Priorities  []string `json:"priorities,omitempty"`

    Tags  []string `json:"tags,omitempty"`

    SimpleNames  []string `json:"simple_names,omitempty"`

    TemplateIDs  []int64 `json:"template_ids,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

}

type FilterAcrossProjectResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type FilterAcrossProjectReqBuilder struct {
	apiReq *core.APIReq
}

func NewFilterAcrossProjectReqBuilder() *FilterAcrossProjectReqBuilder {
	builder := &FilterAcrossProjectReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &FilterAcrossProjectReqBody{},
	}
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) SearchUser(searchUser *SearchUser) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).SearchUser = searchUser
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) CreatedAt(createdAt *TimeInterval) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).CreatedAt = createdAt
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) UpdatedAt(updatedAt *TimeInterval) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).UpdatedAt = updatedAt
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemStatus(workItemStatus []WorkItemStatus) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).WorkItemStatus = workItemStatus
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemName(workItemName *string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).WorkItemName = workItemName
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) PageNum(pageNum *int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).PageNum = pageNum
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) PageSize(pageSize *int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).PageSize = pageSize
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) TenantGroupID(tenantGroupID *int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).TenantGroupID = tenantGroupID
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) WorkItemIDs(workItemIDs []int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Businesses(businesses []string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).Businesses = businesses
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Priorities(priorities []string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).Priorities = priorities
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Tags(tags []string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).Tags = tags
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) SimpleNames(simpleNames []string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) TemplateIDs(templateIDs []int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).TemplateIDs = templateIDs
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Expand(expand *Expand) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).Expand = expand
	return builder
}
func (builder *FilterAcrossProjectReqBuilder) Build() *FilterAcrossProjectReq {
	req := &FilterAcrossProjectReq{}
	req.apiReq = builder.apiReq
	return req
}

type IntegrateSearchReq struct {
	apiReq *core.APIReq
}
type IntegrateSearchReqBody struct {

    Query  *Query `json:"query,omitempty"`

    SearchID  *string `json:"search_id,omitempty"`

    ViewID  *string `json:"view_id,omitempty"`

    FieldSelected  []string `json:"field_selected,omitempty"`

    Features  map[string]string `json:"features,omitempty"`

    DataSources  []DataSource `json:"data_sources,omitempty"`

}

type IntegrateSearchResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *string         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
	ExtraInfo       map[string]string         `json:"extra_info"`
	
}

type IntegrateSearchReqBuilder struct {
	apiReq *core.APIReq
}

func NewIntegrateSearchReqBuilder() *IntegrateSearchReqBuilder {
	builder := &IntegrateSearchReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &IntegrateSearchReqBody{},
	}
	return builder
}
func (builder *IntegrateSearchReqBuilder) Query(query *Query) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).Query = query
	return builder
}
func (builder *IntegrateSearchReqBuilder) SearchID(searchID *string) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).SearchID = searchID
	return builder
}
func (builder *IntegrateSearchReqBuilder) ViewID(viewID *string) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).ViewID = viewID
	return builder
}
func (builder *IntegrateSearchReqBuilder) FieldSelected(fieldSelected []string) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).FieldSelected = fieldSelected
	return builder
}
func (builder *IntegrateSearchReqBuilder) Features(features map[string]string) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).Features = features
	return builder
}
func (builder *IntegrateSearchReqBuilder) DataSources(dataSources []DataSource) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).DataSources = dataSources
	return builder
}
func (builder *IntegrateSearchReqBuilder) Build() *IntegrateSearchReq {
	req := &IntegrateSearchReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIAbortWorkItemReq struct {
	apiReq *core.APIReq
}
type OAPIAbortWorkItemReqBody struct {

    IsAborted  *bool `json:"is_aborted,omitempty"`

    Reason  *string `json:"reason,omitempty"`

    ReasonOption  *string `json:"reason_option,omitempty"`

}

type OAPIAbortWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIAbortWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIAbortWorkItemReqBuilder() *OAPIAbortWorkItemReqBuilder {
	builder := &OAPIAbortWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIAbortWorkItemReqBody{},
	}
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPIAbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIAbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) WorkItemID(workItemID *int64) *OAPIAbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) IsAborted(isAborted *bool) *OAPIAbortWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIAbortWorkItemReqBody).IsAborted = isAborted
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) Reason(reason *string) *OAPIAbortWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIAbortWorkItemReqBody).Reason = reason
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) ReasonOption(reasonOption *string) *OAPIAbortWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIAbortWorkItemReqBody).ReasonOption = reasonOption
	return builder
}
func (builder *OAPIAbortWorkItemReqBuilder) Build() *OAPIAbortWorkItemReq {
	req := &OAPIAbortWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIActualTimeUpdateReq struct {
	apiReq *core.APIReq
}
type OAPIActualTimeUpdateReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    NodeID  *string `json:"node_id,omitempty"`

    ActualTimeInfo  *ActualTimeInfo `json:"actual_time_info,omitempty"`

}

type OAPIActualTimeUpdateResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIActualTimeUpdateReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIActualTimeUpdateReqBuilder() *OAPIActualTimeUpdateReqBuilder {
	builder := &OAPIActualTimeUpdateReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIActualTimeUpdateReqBody{},
	}
	return builder
}
func (builder *OAPIActualTimeUpdateReqBuilder) ProjectKey(projectKey *string) *OAPIActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*OAPIActualTimeUpdateReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIActualTimeUpdateReqBuilder) WorkItemID(workItemID *int64) *OAPIActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*OAPIActualTimeUpdateReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIActualTimeUpdateReqBuilder) NodeID(nodeID *string) *OAPIActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*OAPIActualTimeUpdateReqBody).NodeID = nodeID
	return builder
}
func (builder *OAPIActualTimeUpdateReqBuilder) ActualTimeInfo(actualTimeInfo *ActualTimeInfo) *OAPIActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*OAPIActualTimeUpdateReqBody).ActualTimeInfo = actualTimeInfo
	return builder
}
func (builder *OAPIActualTimeUpdateReqBuilder) Build() *OAPIActualTimeUpdateReq {
	req := &OAPIActualTimeUpdateReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIBatchQueryConclusionOptionReq struct {
	apiReq *core.APIReq
}
type OAPIBatchQueryConclusionOptionReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    NodeIDs  []string `json:"node_ids,omitempty"`

}

type OAPIBatchQueryConclusionOptionResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []OAPIBatchQueryConclusionOptionItem         `json:"data"`
	
}

type OAPIBatchQueryConclusionOptionReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIBatchQueryConclusionOptionReqBuilder() *OAPIBatchQueryConclusionOptionReqBuilder {
	builder := &OAPIBatchQueryConclusionOptionReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIBatchQueryConclusionOptionReqBody{},
	}
	return builder
}
func (builder *OAPIBatchQueryConclusionOptionReqBuilder) ProjectKey(projectKey *string) *OAPIBatchQueryConclusionOptionReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryConclusionOptionReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIBatchQueryConclusionOptionReqBuilder) WorkItemID(workItemID *int64) *OAPIBatchQueryConclusionOptionReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryConclusionOptionReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIBatchQueryConclusionOptionReqBuilder) NodeIDs(nodeIDs []string) *OAPIBatchQueryConclusionOptionReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryConclusionOptionReqBody).NodeIDs = nodeIDs
	return builder
}
func (builder *OAPIBatchQueryConclusionOptionReqBuilder) Build() *OAPIBatchQueryConclusionOptionReq {
	req := &OAPIBatchQueryConclusionOptionReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIBatchQueryDeliverableReq struct {
	apiReq *core.APIReq
}
type OAPIBatchQueryDeliverableReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

}

type OAPIBatchQueryDeliverableResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []OAPIBatchQueryDeliverable         `json:"data"`
	
}

type OAPIBatchQueryDeliverableReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIBatchQueryDeliverableReqBuilder() *OAPIBatchQueryDeliverableReqBuilder {
	builder := &OAPIBatchQueryDeliverableReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIBatchQueryDeliverableReqBody{},
	}
	return builder
}
func (builder *OAPIBatchQueryDeliverableReqBuilder) ProjectKey(projectKey *string) *OAPIBatchQueryDeliverableReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryDeliverableReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIBatchQueryDeliverableReqBuilder) WorkItemIDs(workItemIDs []int64) *OAPIBatchQueryDeliverableReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryDeliverableReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *OAPIBatchQueryDeliverableReqBuilder) Build() *OAPIBatchQueryDeliverableReq {
	req := &OAPIBatchQueryDeliverableReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIBatchQueryFinishedReq struct {
	apiReq *core.APIReq
}
type OAPIBatchQueryFinishedReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    NodeIDs  []string `json:"node_ids,omitempty"`

}

type OAPIBatchQueryFinishedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	ProjectKey       *string         `json:"project_key"`
	
	WorkItemID       *int64         `json:"work_item_id"`
	
	FinishedInfos       []OAPIFinishedInfoItem         `json:"finished_infos"`
	
}

type OAPIBatchQueryFinishedReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIBatchQueryFinishedReqBuilder() *OAPIBatchQueryFinishedReqBuilder {
	builder := &OAPIBatchQueryFinishedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIBatchQueryFinishedReqBody{},
	}
	return builder
}
func (builder *OAPIBatchQueryFinishedReqBuilder) ProjectKey(projectKey *string) *OAPIBatchQueryFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryFinishedReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIBatchQueryFinishedReqBuilder) WorkItemID(workItemID *int64) *OAPIBatchQueryFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryFinishedReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIBatchQueryFinishedReqBuilder) NodeIDs(nodeIDs []string) *OAPIBatchQueryFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIBatchQueryFinishedReqBody).NodeIDs = nodeIDs
	return builder
}
func (builder *OAPIBatchQueryFinishedReqBuilder) Build() *OAPIBatchQueryFinishedReq {
	req := &OAPIBatchQueryFinishedReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIBatchUpdateBasicWorkItemReq struct {
	apiReq *core.APIReq
}
type OAPIBatchUpdateBasicWorkItemReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

    UpdateMode  *string `json:"update_mode,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

    BeforeFieldValue  interface{} `json:"before_field_value,omitempty"`

    AfterFieldValue  interface{} `json:"after_field_value,omitempty"`

}

type OAPIBatchUpdateBasicWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	TaskID       *string         `json:"task_id"`
	
}

type OAPIBatchUpdateBasicWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIBatchUpdateBasicWorkItemReqBuilder() *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder := &OAPIBatchUpdateBasicWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIBatchUpdateBasicWorkItemReqBody{},
	}
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) WorkItemIDs(workItemIDs []int64) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) UpdateMode(updateMode *string) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).UpdateMode = updateMode
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) FieldKey(fieldKey *string) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).FieldKey = fieldKey
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) BeforeFieldValue(beforeFieldValue interface{}) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).BeforeFieldValue = beforeFieldValue
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) AfterFieldValue(afterFieldValue interface{}) *OAPIBatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIBatchUpdateBasicWorkItemReqBody).AfterFieldValue = afterFieldValue
	return builder
}
func (builder *OAPIBatchUpdateBasicWorkItemReqBuilder) Build() *OAPIBatchUpdateBasicWorkItemReq {
	req := &OAPIBatchUpdateBasicWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICompositiveSearchReq struct {
	apiReq *core.APIReq
}
type OAPICompositiveSearchReqBody struct {

    QueryType  *string `json:"query_type,omitempty"`

    Query  *string `json:"query,omitempty"`

    QuerySubType  []string `json:"query_sub_type,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    SimpleNames  []string `json:"simple_names,omitempty"`

}

type OAPICompositiveSearchResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []CompInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type OAPICompositiveSearchReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICompositiveSearchReqBuilder() *OAPICompositiveSearchReqBuilder {
	builder := &OAPICompositiveSearchReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICompositiveSearchReqBody{},
	}
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) QueryType(queryType *string) *OAPICompositiveSearchReqBuilder {
	builder.apiReq.Body.(*OAPICompositiveSearchReqBody).QueryType = queryType
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) Query(query *string) *OAPICompositiveSearchReqBuilder {
	builder.apiReq.Body.(*OAPICompositiveSearchReqBody).Query = query
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) QuerySubType(querySubType []string) *OAPICompositiveSearchReqBuilder {
	builder.apiReq.Body.(*OAPICompositiveSearchReqBody).QuerySubType = querySubType
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) PageSize(pageSize *int64) *OAPICompositiveSearchReqBuilder {
	builder.apiReq.Body.(*OAPICompositiveSearchReqBody).PageSize = pageSize
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) PageNum(pageNum *int64) *OAPICompositiveSearchReqBuilder {
	builder.apiReq.Body.(*OAPICompositiveSearchReqBody).PageNum = pageNum
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) SimpleNames(simpleNames []string) *OAPICompositiveSearchReqBuilder {
	builder.apiReq.Body.(*OAPICompositiveSearchReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *OAPICompositiveSearchReqBuilder) Build() *OAPICompositiveSearchReq {
	req := &OAPICompositiveSearchReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateFieldReq struct {
	apiReq *core.APIReq
}
type OAPICreateFieldReqBody struct {

    FieldName  *string `json:"field_name,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    ValueType  *int64 `json:"value_type,omitempty"`

    ReferenceWorkItemTypeKey  *string `json:"reference_work_item_type_key,omitempty"`

    ReferenceFieldKey  *string `json:"reference_field_key,omitempty"`

    FieldValue  interface{} `json:"field_value,omitempty"`

    FreeAdd  *int64 `json:"free_add,omitempty"`

    WorkItemRelationUUID  *string `json:"work_item_relation_uuid,omitempty"`

    DefaultValue  interface{} `json:"default_value,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    HelpDescription  *string `json:"help_description,omitempty"`

    AuthorizedRoles  []string `json:"authorized_roles,omitempty"`

    IsMulti  *bool `json:"is_multi,omitempty"`

    Format  *bool `json:"format,omitempty"`

    RelatedFieldExtraDisplayInfos  []RelatedFieldExtraDisplayInfo `json:"related_field_extra_display_infos,omitempty"`

}

type OAPICreateFieldResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *string         `json:"data"`
	
}

type OAPICreateFieldReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateFieldReqBuilder() *OAPICreateFieldReqBuilder {
	builder := &OAPICreateFieldReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateFieldReqBody{},
	}
	return builder
}
func (builder *OAPICreateFieldReqBuilder) ProjectKey(projectKey *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateFieldReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPICreateFieldReqBuilder) FieldName(fieldName *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).FieldName = fieldName
	return builder
}
func (builder *OAPICreateFieldReqBuilder) FieldTypeKey(fieldTypeKey *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).FieldTypeKey = fieldTypeKey
	return builder
}
func (builder *OAPICreateFieldReqBuilder) ValueType(valueType *int64) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).ValueType = valueType
	return builder
}
func (builder *OAPICreateFieldReqBuilder) ReferenceWorkItemTypeKey(referenceWorkItemTypeKey *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).ReferenceWorkItemTypeKey = referenceWorkItemTypeKey
	return builder
}
func (builder *OAPICreateFieldReqBuilder) ReferenceFieldKey(referenceFieldKey *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).ReferenceFieldKey = referenceFieldKey
	return builder
}
func (builder *OAPICreateFieldReqBuilder) FieldValue(fieldValue interface{}) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).FieldValue = fieldValue
	return builder
}
func (builder *OAPICreateFieldReqBuilder) FreeAdd(freeAdd *int64) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).FreeAdd = freeAdd
	return builder
}
func (builder *OAPICreateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).WorkItemRelationUUID = workItemRelationUUID
	return builder
}
func (builder *OAPICreateFieldReqBuilder) DefaultValue(defaultValue interface{}) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).DefaultValue = defaultValue
	return builder
}
func (builder *OAPICreateFieldReqBuilder) FieldAlias(fieldAlias *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).FieldAlias = fieldAlias
	return builder
}
func (builder *OAPICreateFieldReqBuilder) HelpDescription(helpDescription *string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).HelpDescription = helpDescription
	return builder
}
func (builder *OAPICreateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).AuthorizedRoles = authorizedRoles
	return builder
}
func (builder *OAPICreateFieldReqBuilder) IsMulti(isMulti *bool) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).IsMulti = isMulti
	return builder
}
func (builder *OAPICreateFieldReqBuilder) Format(format *bool) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).Format = format
	return builder
}
func (builder *OAPICreateFieldReqBuilder) RelatedFieldExtraDisplayInfos(relatedFieldExtraDisplayInfos []RelatedFieldExtraDisplayInfo) *OAPICreateFieldReqBuilder {
	builder.apiReq.Body.(*OAPICreateFieldReqBody).RelatedFieldExtraDisplayInfos = relatedFieldExtraDisplayInfos
	return builder
}
func (builder *OAPICreateFieldReqBuilder) Build() *OAPICreateFieldReq {
	req := &OAPICreateFieldReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateProjectRelationInstancesReq struct {
	apiReq *core.APIReq
}
type OAPICreateProjectRelationInstancesReqBody struct {

    RelationRuleID  *string `json:"relation_rule_id,omitempty"`

    Instances  []RelationBindInstance `json:"instances,omitempty"`

}

type OAPICreateProjectRelationInstancesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPICreateProjectRelationInstancesReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateProjectRelationInstancesReqBuilder() *OAPICreateProjectRelationInstancesReqBuilder {
	builder := &OAPICreateProjectRelationInstancesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateProjectRelationInstancesReqBody{},
	}
	return builder
}
func (builder *OAPICreateProjectRelationInstancesReqBuilder) ProjectKey(projectKey *string) *OAPICreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateProjectRelationInstancesReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPICreateProjectRelationInstancesReqBuilder) WorkItemID(workItemID *int64) *OAPICreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPICreateProjectRelationInstancesReqBuilder) RelationRuleID(relationRuleID *string) *OAPICreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Body.(*OAPICreateProjectRelationInstancesReqBody).RelationRuleID = relationRuleID
	return builder
}
func (builder *OAPICreateProjectRelationInstancesReqBuilder) Instances(instances []RelationBindInstance) *OAPICreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Body.(*OAPICreateProjectRelationInstancesReqBody).Instances = instances
	return builder
}
func (builder *OAPICreateProjectRelationInstancesReqBuilder) Build() *OAPICreateProjectRelationInstancesReq {
	req := &OAPICreateProjectRelationInstancesReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateStoryRelationsReq struct {
	apiReq *core.APIReq
}
type OAPICreateStoryRelationsReqBody struct {

    StoryRelations  []StoryRelationEntity `json:"story_relations,omitempty"`

}

type OAPICreateStoryRelationsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       map[string]int64         `json:"data"`
	
}

type OAPICreateStoryRelationsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateStoryRelationsReqBuilder() *OAPICreateStoryRelationsReqBuilder {
	builder := &OAPICreateStoryRelationsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateStoryRelationsReqBody{},
	}
	return builder
}
func (builder *OAPICreateStoryRelationsReqBuilder) StoryRelations(storyRelations []StoryRelationEntity) *OAPICreateStoryRelationsReqBuilder {
	builder.apiReq.Body.(*OAPICreateStoryRelationsReqBody).StoryRelations = storyRelations
	return builder
}
func (builder *OAPICreateStoryRelationsReqBuilder) ProjectKey(projectKey *string) *OAPICreateStoryRelationsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateStoryRelationsReqBuilder) Build() *OAPICreateStoryRelationsReq {
	req := &OAPICreateStoryRelationsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateTemplateDetailReq struct {
	apiReq *core.APIReq
}
type OAPICreateTemplateDetailReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    CopyTemplateID  *int64 `json:"copy_template_id,omitempty"`

}

type OAPICreateTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type OAPICreateTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateTemplateDetailReqBuilder() *OAPICreateTemplateDetailReqBuilder {
	builder := &OAPICreateTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateTemplateDetailReqBody{},
	}
	return builder
}
func (builder *OAPICreateTemplateDetailReqBuilder) ProjectKey(projectKey *string) *OAPICreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPICreateTemplateDetailReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPICreateTemplateDetailReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPICreateTemplateDetailReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPICreateTemplateDetailReqBuilder) TemplateName(templateName *string) *OAPICreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPICreateTemplateDetailReqBody).TemplateName = templateName
	return builder
}
func (builder *OAPICreateTemplateDetailReqBuilder) CopyTemplateID(copyTemplateID *int64) *OAPICreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPICreateTemplateDetailReqBody).CopyTemplateID = copyTemplateID
	return builder
}
func (builder *OAPICreateTemplateDetailReqBuilder) Build() *OAPICreateTemplateDetailReq {
	req := &OAPICreateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateWorkItemReq struct {
	apiReq *core.APIReq
}
type OAPICreateWorkItemReqBody struct {

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    FieldValuePairs  []FieldValuePair `json:"field_value_pairs,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    Name  *string `json:"name,omitempty"`

}

type OAPICreateWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type OAPICreateWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateWorkItemReqBuilder() *OAPICreateWorkItemReqBuilder {
	builder := &OAPICreateWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateWorkItemReqBody{},
	}
	return builder
}
func (builder *OAPICreateWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPICreateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPICreateWorkItemReqBuilder) FieldValuePairs(fieldValuePairs []FieldValuePair) *OAPICreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemReqBody).FieldValuePairs = fieldValuePairs
	return builder
}
func (builder *OAPICreateWorkItemReqBuilder) TemplateID(templateID *int64) *OAPICreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemReqBody).TemplateID = templateID
	return builder
}
func (builder *OAPICreateWorkItemReqBuilder) Name(name *string) *OAPICreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemReqBody).Name = name
	return builder
}
func (builder *OAPICreateWorkItemReqBuilder) Build() *OAPICreateWorkItemReq {
	req := &OAPICreateWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type OAPICreateWorkItemRelationReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    RelationDetails  []RelationDetail `json:"relation_details,omitempty"`

}

type OAPICreateWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *CreateWorkItemRelationData         `json:"data"`
	
}

type OAPICreateWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateWorkItemRelationReqBuilder() *OAPICreateWorkItemRelationReqBuilder {
	builder := &OAPICreateWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateWorkItemRelationReqBody{},
	}
	return builder
}
func (builder *OAPICreateWorkItemRelationReqBuilder) ProjectKey(projectKey *string) *OAPICreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemRelationReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPICreateWorkItemRelationReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemRelationReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPICreateWorkItemRelationReqBuilder) Name(name *string) *OAPICreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemRelationReqBody).Name = name
	return builder
}
func (builder *OAPICreateWorkItemRelationReqBuilder) RelationDetails(relationDetails []RelationDetail) *OAPICreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemRelationReqBody).RelationDetails = relationDetails
	return builder
}
func (builder *OAPICreateWorkItemRelationReqBuilder) Build() *OAPICreateWorkItemRelationReq {
	req := &OAPICreateWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}
type OAPICreateWorkItemSubTaskReqBody struct {

    NodeID  *string `json:"node_id,omitempty"`

    Name  *string `json:"name,omitempty"`

    AliasKey  *string `json:"alias_key,omitempty"`

    Assignee  []string `json:"assignee,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

    Schedule  *Schedule `json:"schedule,omitempty"`

    Note  *string `json:"note,omitempty"`

    FieldValuePairs  []FieldValuePair `json:"field_value_pairs,omitempty"`

}

type OAPICreateWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type OAPICreateWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateWorkItemSubTaskReqBuilder() *OAPICreateWorkItemSubTaskReqBuilder {
	builder := &OAPICreateWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateWorkItemSubTaskReqBody{},
	}
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) ProjectKey(projectKey *string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) WorkItemID(workItemID *int64) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) NodeID(nodeID *string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).NodeID = nodeID
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) Name(name *string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).Name = name
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) AliasKey(aliasKey *string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).AliasKey = aliasKey
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) Assignee(assignee []string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).Assignee = assignee
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) Schedule(schedule *Schedule) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).Schedule = schedule
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) Note(note *string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).Note = note
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) FieldValuePairs(fieldValuePairs []FieldValuePair) *OAPICreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkItemSubTaskReqBody).FieldValuePairs = fieldValuePairs
	return builder
}
func (builder *OAPICreateWorkItemSubTaskReqBuilder) Build() *OAPICreateWorkItemSubTaskReq {
	req := &OAPICreateWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateWorkingHourRecordReq struct {
	apiReq *core.APIReq
}
type OAPICreateWorkingHourRecordReqBody struct {

    WorkBeginDate  *int64 `json:"work_begin_date,omitempty"`

    WorkEndDate  *int64 `json:"work_end_date,omitempty"`

    IncludeHolidays  *bool `json:"include_holidays,omitempty"`

    WorkingHourRecords  []CreateWorkingHourRecord `json:"working_hour_records,omitempty"`

}

type OAPICreateWorkingHourRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ManHourRecord         `json:"data"`
	
}

type OAPICreateWorkingHourRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateWorkingHourRecordReqBuilder() *OAPICreateWorkingHourRecordReqBuilder {
	builder := &OAPICreateWorkingHourRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateWorkingHourRecordReqBody{},
	}
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) ProjectKey(projectKey *string) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) WorkItemID(workItemID *int64) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) WorkBeginDate(workBeginDate *int64) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkingHourRecordReqBody).WorkBeginDate = workBeginDate
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) WorkEndDate(workEndDate *int64) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkingHourRecordReqBody).WorkEndDate = workEndDate
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) IncludeHolidays(includeHolidays *bool) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkingHourRecordReqBody).IncludeHolidays = includeHolidays
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) WorkingHourRecords(workingHourRecords []CreateWorkingHourRecord) *OAPICreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*OAPICreateWorkingHourRecordReqBody).WorkingHourRecords = workingHourRecords
	return builder
}
func (builder *OAPICreateWorkingHourRecordReqBuilder) Build() *OAPICreateWorkingHourRecordReq {
	req := &OAPICreateWorkingHourRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteFileReq struct {
	apiReq *core.APIReq
}
type OAPIDeleteFileReqBody struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    Uuids  []string `json:"uuids,omitempty"`

}

type OAPIDeleteFileResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteFileReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteFileReqBuilder() *OAPIDeleteFileReqBuilder {
	builder := &OAPIDeleteFileReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIDeleteFileReqBody{},
	}
	return builder
}
func (builder *OAPIDeleteFileReqBuilder) WorkItemID(workItemID *int64) *OAPIDeleteFileReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteFileReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIDeleteFileReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteFileReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteFileReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIDeleteFileReqBuilder) FieldKey(fieldKey *string) *OAPIDeleteFileReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteFileReqBody).FieldKey = fieldKey
	return builder
}
func (builder *OAPIDeleteFileReqBuilder) FieldAlias(fieldAlias *string) *OAPIDeleteFileReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteFileReqBody).FieldAlias = fieldAlias
	return builder
}
func (builder *OAPIDeleteFileReqBuilder) Uuids(uuids []string) *OAPIDeleteFileReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteFileReqBody).Uuids = uuids
	return builder
}
func (builder *OAPIDeleteFileReqBuilder) Build() *OAPIDeleteFileReq {
	req := &OAPIDeleteFileReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteProjectRelationInstanceReq struct {
	apiReq *core.APIReq
}
type OAPIDeleteProjectRelationInstanceReqBody struct {

    RelationRuleID  *string `json:"relation_rule_id,omitempty"`

    RelationWorkItemID  *int64 `json:"relation_work_item_id,omitempty"`

}

type OAPIDeleteProjectRelationInstanceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteProjectRelationInstanceReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteProjectRelationInstanceReqBuilder() *OAPIDeleteProjectRelationInstanceReqBuilder {
	builder := &OAPIDeleteProjectRelationInstanceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIDeleteProjectRelationInstanceReqBody{},
	}
	return builder
}
func (builder *OAPIDeleteProjectRelationInstanceReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIDeleteProjectRelationInstanceReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIDeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIDeleteProjectRelationInstanceReqBuilder) WorkItemID(workItemID *int64) *OAPIDeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIDeleteProjectRelationInstanceReqBuilder) RelationRuleID(relationRuleID *string) *OAPIDeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteProjectRelationInstanceReqBody).RelationRuleID = relationRuleID
	return builder
}
func (builder *OAPIDeleteProjectRelationInstanceReqBuilder) RelationWorkItemID(relationWorkItemID *int64) *OAPIDeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteProjectRelationInstanceReqBody).RelationWorkItemID = relationWorkItemID
	return builder
}
func (builder *OAPIDeleteProjectRelationInstanceReqBuilder) Build() *OAPIDeleteProjectRelationInstanceReq {
	req := &OAPIDeleteProjectRelationInstanceReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteTemplateDetailReq struct {
	apiReq *core.APIReq
}

type OAPIDeleteTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteTemplateDetailReqBuilder() *OAPIDeleteTemplateDetailReqBuilder {
	builder := &OAPIDeleteTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIDeleteTemplateDetailReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIDeleteTemplateDetailReqBuilder) TemplateID(templateID *int64) *OAPIDeleteTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("template_id", fmt.Sprint(*templateID))
	return builder
}
func (builder *OAPIDeleteTemplateDetailReqBuilder) Build() *OAPIDeleteTemplateDetailReq {
	req := &OAPIDeleteTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteWorkItemReq struct {
	apiReq *core.APIReq
}

type OAPIDeleteWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteWorkItemReqBuilder() *OAPIDeleteWorkItemReqBuilder {
	builder := &OAPIDeleteWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIDeleteWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIDeleteWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIDeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIDeleteWorkItemReqBuilder) WorkItemID(workItemID *int64) *OAPIDeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIDeleteWorkItemReqBuilder) Build() *OAPIDeleteWorkItemReq {
	req := &OAPIDeleteWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type OAPIDeleteWorkItemRelationReqBody struct {

    RelationID  *string `json:"relation_id,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

}

type OAPIDeleteWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *string         `json:"data"`
	
}

type OAPIDeleteWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteWorkItemRelationReqBuilder() *OAPIDeleteWorkItemRelationReqBuilder {
	builder := &OAPIDeleteWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIDeleteWorkItemRelationReqBody{},
	}
	return builder
}
func (builder *OAPIDeleteWorkItemRelationReqBuilder) RelationID(relationID *string) *OAPIDeleteWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteWorkItemRelationReqBody).RelationID = relationID
	return builder
}
func (builder *OAPIDeleteWorkItemRelationReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteWorkItemRelationReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIDeleteWorkItemRelationReqBuilder) Build() *OAPIDeleteWorkItemRelationReq {
	req := &OAPIDeleteWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}

type OAPIDeleteWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteWorkItemSubTaskReqBuilder() *OAPIDeleteWorkItemSubTaskReqBuilder {
	builder := &OAPIDeleteWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIDeleteWorkItemSubTaskReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIDeleteWorkItemSubTaskReqBuilder) WorkItemID(workItemID *int64) *OAPIDeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIDeleteWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIDeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIDeleteWorkItemSubTaskReqBuilder) TaskID(taskID *int64) *OAPIDeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("task_id", fmt.Sprint(*taskID))
	return builder
}
func (builder *OAPIDeleteWorkItemSubTaskReqBuilder) Build() *OAPIDeleteWorkItemSubTaskReq {
	req := &OAPIDeleteWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteWorkingHourRecordReq struct {
	apiReq *core.APIReq
}
type OAPIDeleteWorkingHourRecordReqBody struct {

    WorkingHourRecordIDs  []int64 `json:"working_hour_record_ids,omitempty"`

}

type OAPIDeleteWorkingHourRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteWorkingHourRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteWorkingHourRecordReqBuilder() *OAPIDeleteWorkingHourRecordReqBuilder {
	builder := &OAPIDeleteWorkingHourRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIDeleteWorkingHourRecordReqBody{},
	}
	return builder
}
func (builder *OAPIDeleteWorkingHourRecordReqBuilder) ProjectKey(projectKey *string) *OAPIDeleteWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIDeleteWorkingHourRecordReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIDeleteWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIDeleteWorkingHourRecordReqBuilder) WorkItemID(workItemID *int64) *OAPIDeleteWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIDeleteWorkingHourRecordReqBuilder) WorkingHourRecordIDs(workingHourRecordIDs []int64) *OAPIDeleteWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteWorkingHourRecordReqBody).WorkingHourRecordIDs = workingHourRecordIDs
	return builder
}
func (builder *OAPIDeleteWorkingHourRecordReqBuilder) Build() *OAPIDeleteWorkingHourRecordReq {
	req := &OAPIDeleteWorkingHourRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIElementTemplateCreateReq struct {
	apiReq *core.APIReq
}
type OAPIElementTemplateCreateReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ElementType  *string `json:"element_type,omitempty"`

    NodeElement  *NodeElement `json:"node_element,omitempty"`

    TaskElement  *TaskElement `json:"task_element,omitempty"`

}

type OAPIElementTemplateCreateResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	ElementKey       *string         `json:"element_key"`
	
}

type OAPIElementTemplateCreateReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIElementTemplateCreateReqBuilder() *OAPIElementTemplateCreateReqBuilder {
	builder := &OAPIElementTemplateCreateReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIElementTemplateCreateReqBody{},
	}
	return builder
}
func (builder *OAPIElementTemplateCreateReqBuilder) ProjectKey(projectKey *string) *OAPIElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateCreateReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIElementTemplateCreateReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateCreateReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIElementTemplateCreateReqBuilder) ElementType(elementType *string) *OAPIElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateCreateReqBody).ElementType = elementType
	return builder
}
func (builder *OAPIElementTemplateCreateReqBuilder) NodeElement(nodeElement *NodeElement) *OAPIElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateCreateReqBody).NodeElement = nodeElement
	return builder
}
func (builder *OAPIElementTemplateCreateReqBuilder) TaskElement(taskElement *TaskElement) *OAPIElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateCreateReqBody).TaskElement = taskElement
	return builder
}
func (builder *OAPIElementTemplateCreateReqBuilder) Build() *OAPIElementTemplateCreateReq {
	req := &OAPIElementTemplateCreateReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIElementTemplateQueryReq struct {
	apiReq *core.APIReq
}
type OAPIElementTemplateQueryReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ElementType  *string `json:"element_type,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

}

type OAPIElementTemplateQueryResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	NodeElements       []NodeElement         `json:"node_elements"`
	
	TaskElements       []TaskElement         `json:"task_elements"`
	
	Total       *int64         `json:"total"`
	
}

type OAPIElementTemplateQueryReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIElementTemplateQueryReqBuilder() *OAPIElementTemplateQueryReqBuilder {
	builder := &OAPIElementTemplateQueryReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIElementTemplateQueryReqBody{},
	}
	return builder
}
func (builder *OAPIElementTemplateQueryReqBuilder) ProjectKey(projectKey *string) *OAPIElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateQueryReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIElementTemplateQueryReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateQueryReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIElementTemplateQueryReqBuilder) ElementType(elementType *string) *OAPIElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateQueryReqBody).ElementType = elementType
	return builder
}
func (builder *OAPIElementTemplateQueryReqBuilder) PageNum(pageNum *int64) *OAPIElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateQueryReqBody).PageNum = pageNum
	return builder
}
func (builder *OAPIElementTemplateQueryReqBuilder) PageSize(pageSize *int64) *OAPIElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*OAPIElementTemplateQueryReqBody).PageSize = pageSize
	return builder
}
func (builder *OAPIElementTemplateQueryReqBuilder) Build() *OAPIElementTemplateQueryReq {
	req := &OAPIElementTemplateQueryReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIFreezeWorkItemReq struct {
	apiReq *core.APIReq
}
type OAPIFreezeWorkItemReqBody struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    IsFrozen  *bool `json:"is_frozen,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

}

type OAPIFreezeWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIFreezeWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIFreezeWorkItemReqBuilder() *OAPIFreezeWorkItemReqBuilder {
	builder := &OAPIFreezeWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIFreezeWorkItemReqBody{},
	}
	return builder
}
func (builder *OAPIFreezeWorkItemReqBuilder) WorkItemID(workItemID *int64) *OAPIFreezeWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIFreezeWorkItemReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIFreezeWorkItemReqBuilder) IsFrozen(isFrozen *bool) *OAPIFreezeWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIFreezeWorkItemReqBody).IsFrozen = isFrozen
	return builder
}
func (builder *OAPIFreezeWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPIFreezeWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIFreezeWorkItemReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIFreezeWorkItemReqBuilder) Build() *OAPIFreezeWorkItemReq {
	req := &OAPIFreezeWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIGetWBSInfoReq struct {
	apiReq *core.APIReq
}
type OAPIGetWBSInfoReqBody struct {

    Expand  *Expand `json:"expand,omitempty"`

}

type OAPIGetWBSInfoResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *WBSInfo         `json:"data"`
	
}

type OAPIGetWBSInfoReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetWBSInfoReqBuilder() *OAPIGetWBSInfoReqBuilder {
	builder := &OAPIGetWBSInfoReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIGetWBSInfoReqBody{},
	}
	return builder
}
func (builder *OAPIGetWBSInfoReqBuilder) ProjectKey(projectKey *string) *OAPIGetWBSInfoReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIGetWBSInfoReqBuilder) WorkItemID(workItemID *int64) *OAPIGetWBSInfoReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIGetWBSInfoReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIGetWBSInfoReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIGetWBSInfoReqBuilder) Expand(expand *Expand) *OAPIGetWBSInfoReqBuilder {
	builder.apiReq.Body.(*OAPIGetWBSInfoReqBody).Expand = expand
	return builder
}
func (builder *OAPIGetWBSInfoReqBuilder) NeedUnionDeliverable(needUnionDeliverable *bool) *OAPIGetWBSInfoReqBuilder {
	builder.apiReq.QueryParams.Set("need_union_deliverable", fmt.Sprint(needUnionDeliverable))
	return builder
}
func (builder *OAPIGetWBSInfoReqBuilder) Build() *OAPIGetWBSInfoReq {
	req := &OAPIGetWBSInfoReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIGetWBSViewSubWorkItemConfReq struct {
	apiReq *core.APIReq
}

type OAPIGetWBSViewSubWorkItemConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	DraftViewSubWorkItemConfs       []DraftViewSubWorkItemConf         `json:"draft_view_sub_work_item_confs"`
	
}

type OAPIGetWBSViewSubWorkItemConfReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetWBSViewSubWorkItemConfReqBuilder() *OAPIGetWBSViewSubWorkItemConfReqBuilder {
	builder := &OAPIGetWBSViewSubWorkItemConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIGetWBSViewSubWorkItemConfReqBuilder) DraftID(draftID *string) *OAPIGetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("draft_id", fmt.Sprint(draftID))
	return builder
}
func (builder *OAPIGetWBSViewSubWorkItemConfReqBuilder) WorkItemID(workItemID *int64) *OAPIGetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}
func (builder *OAPIGetWBSViewSubWorkItemConfReqBuilder) NodeUUID(nodeUUID *string) *OAPIGetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("node_uuid", fmt.Sprint(nodeUUID))
	return builder
}
func (builder *OAPIGetWBSViewSubWorkItemConfReqBuilder) ProjectKey(projectKey *string) *OAPIGetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *OAPIGetWBSViewSubWorkItemConfReqBuilder) Build() *OAPIGetWBSViewSubWorkItemConfReq {
	req := &OAPIGetWBSViewSubWorkItemConfReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIGetWorkFlowReq struct {
	apiReq *core.APIReq
}
type OAPIGetWorkFlowReqBody struct {

    Fields  []string `json:"fields,omitempty"`

    FlowType  *int64 `json:"flow_type,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

}

type OAPIGetWorkFlowResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *NodesConnections         `json:"data"`
	
}

type OAPIGetWorkFlowReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetWorkFlowReqBuilder() *OAPIGetWorkFlowReqBuilder {
	builder := &OAPIGetWorkFlowReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIGetWorkFlowReqBody{},
	}
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) ProjectKey(projectKey *string) *OAPIGetWorkFlowReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) WorkItemID(workItemID *int64) *OAPIGetWorkFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) Fields(fields []string) *OAPIGetWorkFlowReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkFlowReqBody).Fields = fields
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) FlowType(flowType *int64) *OAPIGetWorkFlowReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkFlowReqBody).FlowType = flowType
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIGetWorkFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) Expand(expand *Expand) *OAPIGetWorkFlowReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkFlowReqBody).Expand = expand
	return builder
}
func (builder *OAPIGetWorkFlowReqBuilder) Build() *OAPIGetWorkFlowReq {
	req := &OAPIGetWorkFlowReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIGetWorkItemManHourRecordsReq struct {
	apiReq *core.APIReq
}
type OAPIGetWorkItemManHourRecordsReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

}

type OAPIGetWorkItemManHourRecordsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ManHourRecord         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type OAPIGetWorkItemManHourRecordsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetWorkItemManHourRecordsReqBuilder() *OAPIGetWorkItemManHourRecordsReqBuilder {
	builder := &OAPIGetWorkItemManHourRecordsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIGetWorkItemManHourRecordsReqBody{},
	}
	return builder
}
func (builder *OAPIGetWorkItemManHourRecordsReqBuilder) ProjectKey(projectKey *string) *OAPIGetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemManHourRecordsReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIGetWorkItemManHourRecordsReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIGetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemManHourRecordsReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIGetWorkItemManHourRecordsReqBuilder) WorkItemID(workItemID *int64) *OAPIGetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemManHourRecordsReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIGetWorkItemManHourRecordsReqBuilder) PageNum(pageNum *int64) *OAPIGetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemManHourRecordsReqBody).PageNum = pageNum
	return builder
}
func (builder *OAPIGetWorkItemManHourRecordsReqBuilder) PageSize(pageSize *int64) *OAPIGetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemManHourRecordsReqBody).PageSize = pageSize
	return builder
}
func (builder *OAPIGetWorkItemManHourRecordsReqBuilder) Build() *OAPIGetWorkItemManHourRecordsReq {
	req := &OAPIGetWorkItemManHourRecordsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIGetWorkItemTypeInfoByKeyReq struct {
	apiReq *core.APIReq
}

type OAPIGetWorkItemTypeInfoByKeyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *WorkItemTypeInfo         `json:"data"`
	
}

type OAPIGetWorkItemTypeInfoByKeyReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetWorkItemTypeInfoByKeyReqBuilder() *OAPIGetWorkItemTypeInfoByKeyReqBuilder {
	builder := &OAPIGetWorkItemTypeInfoByKeyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIGetWorkItemTypeInfoByKeyReqBuilder) ProjectKey(projectKey *string) *OAPIGetWorkItemTypeInfoByKeyReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIGetWorkItemTypeInfoByKeyReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIGetWorkItemTypeInfoByKeyReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIGetWorkItemTypeInfoByKeyReqBuilder) Build() *OAPIGetWorkItemTypeInfoByKeyReq {
	req := &OAPIGetWorkItemTypeInfoByKeyReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIGetWorkItemsByIdsReq struct {
	apiReq *core.APIReq
}
type OAPIGetWorkItemsByIdsReqBody struct {

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

    Fields  []string `json:"fields,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

}

type OAPIGetWorkItemsByIdsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
}

type OAPIGetWorkItemsByIdsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetWorkItemsByIdsReqBuilder() *OAPIGetWorkItemsByIdsReqBuilder {
	builder := &OAPIGetWorkItemsByIdsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIGetWorkItemsByIdsReqBody{},
	}
	return builder
}
func (builder *OAPIGetWorkItemsByIdsReqBuilder) ProjectKey(projectKey *string) *OAPIGetWorkItemsByIdsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIGetWorkItemsByIdsReqBuilder) WorkItemIDs(workItemIDs []int64) *OAPIGetWorkItemsByIdsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemsByIdsReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *OAPIGetWorkItemsByIdsReqBuilder) Fields(fields []string) *OAPIGetWorkItemsByIdsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemsByIdsReqBody).Fields = fields
	return builder
}
func (builder *OAPIGetWorkItemsByIdsReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIGetWorkItemsByIdsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIGetWorkItemsByIdsReqBuilder) Expand(expand *Expand) *OAPIGetWorkItemsByIdsReqBuilder {
	builder.apiReq.Body.(*OAPIGetWorkItemsByIdsReqBody).Expand = expand
	return builder
}
func (builder *OAPIGetWorkItemsByIdsReqBuilder) Build() *OAPIGetWorkItemsByIdsReq {
	req := &OAPIGetWorkItemsByIdsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIListTemplateConfReq struct {
	apiReq *core.APIReq
}

type OAPIListTemplateConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []TemplateConf         `json:"data"`
	
}

type OAPIListTemplateConfReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIListTemplateConfReqBuilder() *OAPIListTemplateConfReqBuilder {
	builder := &OAPIListTemplateConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIListTemplateConfReqBuilder) ProjectKey(projectKey *string) *OAPIListTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIListTemplateConfReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIListTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIListTemplateConfReqBuilder) Build() *OAPIListTemplateConfReq {
	req := &OAPIListTemplateConfReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIPatchWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type OAPIPatchWBSViewDraftReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    DraftID  *string `json:"draft_id,omitempty"`

    Operations  []Operation `json:"operations,omitempty"`

}

type OAPIPatchWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	RelationValues       []Operation         `json:"relation_values"`
	
}

type OAPIPatchWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIPatchWBSViewDraftReqBuilder() *OAPIPatchWBSViewDraftReqBuilder {
	builder := &OAPIPatchWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIPatchWBSViewDraftReqBody{},
	}
	return builder
}
func (builder *OAPIPatchWBSViewDraftReqBuilder) ProjectKey(projectKey *string) *OAPIPatchWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIPatchWBSViewDraftReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIPatchWBSViewDraftReqBuilder) DraftID(draftID *string) *OAPIPatchWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIPatchWBSViewDraftReqBody).DraftID = draftID
	return builder
}
func (builder *OAPIPatchWBSViewDraftReqBuilder) Operations(operations []Operation) *OAPIPatchWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIPatchWBSViewDraftReqBody).Operations = operations
	return builder
}
func (builder *OAPIPatchWBSViewDraftReqBuilder) Build() *OAPIPatchWBSViewDraftReq {
	req := &OAPIPatchWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIPublishWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type OAPIPublishWBSViewDraftReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    DraftID  *string `json:"draft_id,omitempty"`

}

type OAPIPublishWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIPublishWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIPublishWBSViewDraftReqBuilder() *OAPIPublishWBSViewDraftReqBuilder {
	builder := &OAPIPublishWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIPublishWBSViewDraftReqBody{},
	}
	return builder
}
func (builder *OAPIPublishWBSViewDraftReqBuilder) ProjectKey(projectKey *string) *OAPIPublishWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIPublishWBSViewDraftReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIPublishWBSViewDraftReqBuilder) DraftID(draftID *string) *OAPIPublishWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIPublishWBSViewDraftReqBody).DraftID = draftID
	return builder
}
func (builder *OAPIPublishWBSViewDraftReqBuilder) Build() *OAPIPublishWBSViewDraftReq {
	req := &OAPIPublishWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryAWorkItemTypesReq struct {
	apiReq *core.APIReq
}

type OAPIQueryAWorkItemTypesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemKeyType         `json:"data"`
	
}

type OAPIQueryAWorkItemTypesReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryAWorkItemTypesReqBuilder() *OAPIQueryAWorkItemTypesReqBuilder {
	builder := &OAPIQueryAWorkItemTypesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryAWorkItemTypesReqBuilder) ProjectKey(projectKey *string) *OAPIQueryAWorkItemTypesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryAWorkItemTypesReqBuilder) Build() *OAPIQueryAWorkItemTypesReq {
	req := &OAPIQueryAWorkItemTypesReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryBusinessesReq struct {
	apiReq *core.APIReq
}

type OAPIQueryBusinessesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []Business         `json:"data"`
	
}

type OAPIQueryBusinessesReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryBusinessesReqBuilder() *OAPIQueryBusinessesReqBuilder {
	builder := &OAPIQueryBusinessesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryBusinessesReqBuilder) ProjectKey(projectKey *string) *OAPIQueryBusinessesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryBusinessesReqBuilder) Build() *OAPIQueryBusinessesReq {
	req := &OAPIQueryBusinessesReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryProjectFieldsReq struct {
	apiReq *core.APIReq
}
type OAPIQueryProjectFieldsReqBody struct {

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

}

type OAPIQueryProjectFieldsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []SimpleField         `json:"data"`
	
}

type OAPIQueryProjectFieldsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryProjectFieldsReqBuilder() *OAPIQueryProjectFieldsReqBuilder {
	builder := &OAPIQueryProjectFieldsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryProjectFieldsReqBody{},
	}
	return builder
}
func (builder *OAPIQueryProjectFieldsReqBuilder) ProjectKey(projectKey *string) *OAPIQueryProjectFieldsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryProjectFieldsReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIQueryProjectFieldsReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectFieldsReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIQueryProjectFieldsReqBuilder) Build() *OAPIQueryProjectFieldsReq {
	req := &OAPIQueryProjectFieldsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryProjectRelationReq struct {
	apiReq *core.APIReq
}
type OAPIQueryProjectRelationReqBody struct {

    RemoteProjects  []string `json:"remote_projects,omitempty"`

}

type OAPIQueryProjectRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ProjectRelationRule         `json:"data"`
	
}

type OAPIQueryProjectRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryProjectRelationReqBuilder() *OAPIQueryProjectRelationReqBuilder {
	builder := &OAPIQueryProjectRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryProjectRelationReqBody{},
	}
	return builder
}
func (builder *OAPIQueryProjectRelationReqBuilder) ProjectKey(projectKey *string) *OAPIQueryProjectRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryProjectRelationReqBuilder) RemoteProjects(remoteProjects []string) *OAPIQueryProjectRelationReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectRelationReqBody).RemoteProjects = remoteProjects
	return builder
}
func (builder *OAPIQueryProjectRelationReqBuilder) Build() *OAPIQueryProjectRelationReq {
	req := &OAPIQueryProjectRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryProjectRelationInstanceReq struct {
	apiReq *core.APIReq
}
type OAPIQueryProjectRelationInstanceReqBody struct {

    RelationRuleID  *string `json:"relation_rule_id,omitempty"`

    RelationWorkItemID  *int64 `json:"relation_work_item_id,omitempty"`

    RelationWorkItemTypeKey  *string `json:"relation_work_item_type_key,omitempty"`

    RelationProjectKey  *string `json:"relation_project_key,omitempty"`

}

type OAPIQueryProjectRelationInstanceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []RelationInstance         `json:"data"`
	
}

type OAPIQueryProjectRelationInstanceReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryProjectRelationInstanceReqBuilder() *OAPIQueryProjectRelationInstanceReqBuilder {
	builder := &OAPIQueryProjectRelationInstanceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryProjectRelationInstanceReqBody{},
	}
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) ProjectKey(projectKey *string) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) WorkItemID(workItemID *int64) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) RelationRuleID(relationRuleID *string) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectRelationInstanceReqBody).RelationRuleID = relationRuleID
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) RelationWorkItemID(relationWorkItemID *int64) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectRelationInstanceReqBody).RelationWorkItemID = relationWorkItemID
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) RelationWorkItemTypeKey(relationWorkItemTypeKey *string) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectRelationInstanceReqBody).RelationWorkItemTypeKey = relationWorkItemTypeKey
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) RelationProjectKey(relationProjectKey *string) *OAPIQueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*OAPIQueryProjectRelationInstanceReqBody).RelationProjectKey = relationProjectKey
	return builder
}
func (builder *OAPIQueryProjectRelationInstanceReqBuilder) Build() *OAPIQueryProjectRelationInstanceReq {
	req := &OAPIQueryProjectRelationInstanceReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryRoleConfDetailsReq struct {
	apiReq *core.APIReq
}

type OAPIQueryRoleConfDetailsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []RoleConfDetail         `json:"data"`
	
}

type OAPIQueryRoleConfDetailsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryRoleConfDetailsReqBuilder() *OAPIQueryRoleConfDetailsReqBuilder {
	builder := &OAPIQueryRoleConfDetailsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryRoleConfDetailsReqBuilder) ProjectKey(projectKey *string) *OAPIQueryRoleConfDetailsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryRoleConfDetailsReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIQueryRoleConfDetailsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIQueryRoleConfDetailsReqBuilder) Build() *OAPIQueryRoleConfDetailsReq {
	req := &OAPIQueryRoleConfDetailsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryStoryRelationsReq struct {
	apiReq *core.APIReq
}
type OAPIQueryStoryRelationsReqBody struct {

    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`

}

type OAPIQueryStoryRelationsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *QueryStoryRelationData         `json:"data"`
	
}

type OAPIQueryStoryRelationsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryStoryRelationsReqBuilder() *OAPIQueryStoryRelationsReqBuilder {
	builder := &OAPIQueryStoryRelationsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryStoryRelationsReqBody{},
	}
	return builder
}
func (builder *OAPIQueryStoryRelationsReqBuilder) WorkItemIDs(workItemIDs []int64) *OAPIQueryStoryRelationsReqBuilder {
	builder.apiReq.Body.(*OAPIQueryStoryRelationsReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *OAPIQueryStoryRelationsReqBuilder) ProjectKey(projectKey *string) *OAPIQueryStoryRelationsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryStoryRelationsReqBuilder) Build() *OAPIQueryStoryRelationsReq {
	req := &OAPIQueryStoryRelationsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryTemplateDetailReq struct {
	apiReq *core.APIReq
}

type OAPIQueryTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *TemplateDetail         `json:"data"`
	
}

type OAPIQueryTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryTemplateDetailReqBuilder() *OAPIQueryTemplateDetailReqBuilder {
	builder := &OAPIQueryTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryTemplateDetailReqBuilder) ProjectKey(projectKey *string) *OAPIQueryTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryTemplateDetailReqBuilder) TemplateID(templateID *int64) *OAPIQueryTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("template_id", fmt.Sprint(*templateID))
	return builder
}
func (builder *OAPIQueryTemplateDetailReqBuilder) Build() *OAPIQueryTemplateDetailReq {
	req := &OAPIQueryTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type OAPIQueryWBSViewDraftReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    NeedInit  *bool `json:"need_init,omitempty"`

}

type OAPIQueryWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	WbsDraft       *WbsDraft         `json:"wbs_draft"`
	
}

type OAPIQueryWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryWBSViewDraftReqBuilder() *OAPIQueryWBSViewDraftReqBuilder {
	builder := &OAPIQueryWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryWBSViewDraftReqBody{},
	}
	return builder
}
func (builder *OAPIQueryWBSViewDraftReqBuilder) ProjectKey(projectKey *string) *OAPIQueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIQueryWBSViewDraftReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIQueryWBSViewDraftReqBuilder) WorkItemID(workItemID *int64) *OAPIQueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIQueryWBSViewDraftReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIQueryWBSViewDraftReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIQueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIQueryWBSViewDraftReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIQueryWBSViewDraftReqBuilder) NeedInit(needInit *bool) *OAPIQueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIQueryWBSViewDraftReqBody).NeedInit = needInit
	return builder
}
func (builder *OAPIQueryWBSViewDraftReqBuilder) Build() *OAPIQueryWBSViewDraftReq {
	req := &OAPIQueryWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryWbsTemplateConfReq struct {
	apiReq *core.APIReq
}
type OAPIQueryWbsTemplateConfReqBody struct {

    TemplateKey  *string `json:"template_key,omitempty"`

}

type OAPIQueryWbsTemplateConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *WbsTemplate         `json:"data"`
	
}

type OAPIQueryWbsTemplateConfReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryWbsTemplateConfReqBuilder() *OAPIQueryWbsTemplateConfReqBuilder {
	builder := &OAPIQueryWbsTemplateConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryWbsTemplateConfReqBody{},
	}
	return builder
}
func (builder *OAPIQueryWbsTemplateConfReqBuilder) ProjectKey(projectKey *string) *OAPIQueryWbsTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryWbsTemplateConfReqBuilder) TemplateKey(templateKey *string) *OAPIQueryWbsTemplateConfReqBuilder {
	builder.apiReq.Body.(*OAPIQueryWbsTemplateConfReqBody).TemplateKey = templateKey
	return builder
}
func (builder *OAPIQueryWbsTemplateConfReqBuilder) Build() *OAPIQueryWbsTemplateConfReq {
	req := &OAPIQueryWbsTemplateConfReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryWorkItemMetaDataReq struct {
	apiReq *core.APIReq
}

type OAPIQueryWorkItemMetaDataResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []FieldConf         `json:"data"`
	
}

type OAPIQueryWorkItemMetaDataReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryWorkItemMetaDataReqBuilder() *OAPIQueryWorkItemMetaDataReqBuilder {
	builder := &OAPIQueryWorkItemMetaDataReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryWorkItemMetaDataReqBuilder) ProjectKey(projectKey *string) *OAPIQueryWorkItemMetaDataReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryWorkItemMetaDataReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIQueryWorkItemMetaDataReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIQueryWorkItemMetaDataReqBuilder) Build() *OAPIQueryWorkItemMetaDataReq {
	req := &OAPIQueryWorkItemMetaDataReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryWorkItemRelationReq struct {
	apiReq *core.APIReq
}

type OAPIQueryWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemRelation         `json:"data"`
	
}

type OAPIQueryWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryWorkItemRelationReqBuilder() *OAPIQueryWorkItemRelationReqBuilder {
	builder := &OAPIQueryWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryWorkItemRelationReqBuilder) ProjectKey(projectKey *string) *OAPIQueryWorkItemRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryWorkItemRelationReqBuilder) Build() *OAPIQueryWorkItemRelationReq {
	req := &OAPIQueryWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}

type OAPIQueryWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []NodeTask         `json:"data"`
	
}

type OAPIQueryWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryWorkItemSubTaskReqBuilder() *OAPIQueryWorkItemSubTaskReqBuilder {
	builder := &OAPIQueryWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *OAPIQueryWorkItemSubTaskReqBuilder) ProjectKey(projectKey *string) *OAPIQueryWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIQueryWorkItemSubTaskReqBuilder) WorkItemID(workItemID *int64) *OAPIQueryWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIQueryWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIQueryWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIQueryWorkItemSubTaskReqBuilder) NodeID(nodeID *string) *OAPIQueryWorkItemSubTaskReqBuilder {
	builder.apiReq.QueryParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}
func (builder *OAPIQueryWorkItemSubTaskReqBuilder) Build() *OAPIQueryWorkItemSubTaskReq {
	req := &OAPIQueryWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIResetWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type OAPIResetWBSViewDraftReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

}

type OAPIResetWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	WbsDraft       *WbsDraft         `json:"wbs_draft"`
	
}

type OAPIResetWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIResetWBSViewDraftReqBuilder() *OAPIResetWBSViewDraftReqBuilder {
	builder := &OAPIResetWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIResetWBSViewDraftReqBody{},
	}
	return builder
}
func (builder *OAPIResetWBSViewDraftReqBuilder) ProjectKey(projectKey *string) *OAPIResetWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIResetWBSViewDraftReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIResetWBSViewDraftReqBuilder) WorkItemID(workItemID *int64) *OAPIResetWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIResetWBSViewDraftReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIResetWBSViewDraftReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIResetWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPIResetWBSViewDraftReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIResetWBSViewDraftReqBuilder) Build() *OAPIResetWBSViewDraftReq {
	req := &OAPIResetWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIResourceCreateWorkItemReq struct {
	apiReq *core.APIReq
}
type OAPIResourceCreateWorkItemReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    FieldValuePairs  []FieldValuePair `json:"field_value_pairs,omitempty"`

}

type OAPIResourceCreateWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *OAPICreateWorkItemInfo         `json:"data"`
	
}

type OAPIResourceCreateWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIResourceCreateWorkItemReqBuilder() *OAPIResourceCreateWorkItemReqBuilder {
	builder := &OAPIResourceCreateWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIResourceCreateWorkItemReqBody{},
	}
	return builder
}
func (builder *OAPIResourceCreateWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPIResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIResourceCreateWorkItemReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIResourceCreateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIResourceCreateWorkItemReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIResourceCreateWorkItemReqBuilder) TemplateID(templateID *int64) *OAPIResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIResourceCreateWorkItemReqBody).TemplateID = templateID
	return builder
}
func (builder *OAPIResourceCreateWorkItemReqBuilder) FieldValuePairs(fieldValuePairs []FieldValuePair) *OAPIResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIResourceCreateWorkItemReqBody).FieldValuePairs = fieldValuePairs
	return builder
}
func (builder *OAPIResourceCreateWorkItemReqBuilder) Build() *OAPIResourceCreateWorkItemReq {
	req := &OAPIResourceCreateWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPISearchWorkItemsRelationReq struct {
	apiReq *core.APIReq
}
type OAPISearchWorkItemsRelationReqBody struct {

    RelationWorkItemTypeKey  *string `json:"relation_work_item_type_key,omitempty"`

    RelationKey  *string `json:"relation_key,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    RelationType  *int32 `json:"relation_type,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

}

type OAPISearchWorkItemsRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type OAPISearchWorkItemsRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPISearchWorkItemsRelationReqBuilder() *OAPISearchWorkItemsRelationReqBuilder {
	builder := &OAPISearchWorkItemsRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPISearchWorkItemsRelationReqBody{},
	}
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) ProjectKey(projectKey *string) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) WorkItemID(workItemID *int64) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) RelationWorkItemTypeKey(relationWorkItemTypeKey *string) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*OAPISearchWorkItemsRelationReqBody).RelationWorkItemTypeKey = relationWorkItemTypeKey
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) RelationKey(relationKey *string) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*OAPISearchWorkItemsRelationReqBody).RelationKey = relationKey
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) PageNum(pageNum *int64) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*OAPISearchWorkItemsRelationReqBody).PageNum = pageNum
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) PageSize(pageSize *int64) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*OAPISearchWorkItemsRelationReqBody).PageSize = pageSize
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) RelationType(relationType *int32) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*OAPISearchWorkItemsRelationReqBody).RelationType = relationType
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) Expand(expand *Expand) *OAPISearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*OAPISearchWorkItemsRelationReqBody).Expand = expand
	return builder
}
func (builder *OAPISearchWorkItemsRelationReqBuilder) Build() *OAPISearchWorkItemsRelationReq {
	req := &OAPISearchWorkItemsRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPISubTaskModifyReq struct {
	apiReq *core.APIReq
}
type OAPISubTaskModifyReqBody struct {

    NodeID  *string `json:"node_id,omitempty"`

    TaskID  *int64 `json:"task_id,omitempty"`

    Action  *string `json:"action,omitempty"`

    Assignee  []string `json:"assignee,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

    Schedules  []Schedule `json:"schedules,omitempty"`

    Deliverable  []FieldValuePair `json:"deliverable,omitempty"`

    Note  *string `json:"note,omitempty"`

}

type OAPISubTaskModifyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPISubTaskModifyReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPISubTaskModifyReqBuilder() *OAPISubTaskModifyReqBuilder {
	builder := &OAPISubTaskModifyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPISubTaskModifyReqBody{},
	}
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) NodeID(nodeID *string) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).NodeID = nodeID
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) TaskID(taskID *int64) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).TaskID = taskID
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) Action(action *string) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).Action = action
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) Assignee(assignee []string) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).Assignee = assignee
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) Schedules(schedules []Schedule) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).Schedules = schedules
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) Deliverable(deliverable []FieldValuePair) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).Deliverable = deliverable
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) Note(note *string) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.Body.(*OAPISubTaskModifyReqBody).Note = note
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) ProjectKey(projectKey *string) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) WorkItemID(workItemID *int64) *OAPISubTaskModifyReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPISubTaskModifyReqBuilder) Build() *OAPISubTaskModifyReq {
	req := &OAPISubTaskModifyReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPISwitchBackToWbsViewDraftReq struct {
	apiReq *core.APIReq
}
type OAPISwitchBackToWbsViewDraftReqBody struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    DraftID  *string `json:"draft_id,omitempty"`

    CommitID  *string `json:"commit_id,omitempty"`

}

type OAPISwitchBackToWbsViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Success       *bool         `json:"success"`
	
}

type OAPISwitchBackToWbsViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPISwitchBackToWbsViewDraftReqBuilder() *OAPISwitchBackToWbsViewDraftReqBuilder {
	builder := &OAPISwitchBackToWbsViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPISwitchBackToWbsViewDraftReqBody{},
	}
	return builder
}
func (builder *OAPISwitchBackToWbsViewDraftReqBuilder) ProjectKey(projectKey *string) *OAPISwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPISwitchBackToWbsViewDraftReqBuilder) WorkItemID(workItemID *int64) *OAPISwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPISwitchBackToWbsViewDraftReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPISwitchBackToWbsViewDraftReqBuilder) DraftID(draftID *string) *OAPISwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPISwitchBackToWbsViewDraftReqBody).DraftID = draftID
	return builder
}
func (builder *OAPISwitchBackToWbsViewDraftReqBuilder) CommitID(commitID *string) *OAPISwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.Body.(*OAPISwitchBackToWbsViewDraftReqBody).CommitID = commitID
	return builder
}
func (builder *OAPISwitchBackToWbsViewDraftReqBuilder) Build() *OAPISwitchBackToWbsViewDraftReq {
	req := &OAPISwitchBackToWbsViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateFieldReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateFieldReqBody struct {

    FieldName  *string `json:"field_name,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  interface{} `json:"field_value,omitempty"`

    FreeAdd  *int64 `json:"free_add,omitempty"`

    WorkItemRelationUUID  *string `json:"work_item_relation_uuid,omitempty"`

    DefaultValue  interface{} `json:"default_value,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    HelpDescription  *string `json:"help_description,omitempty"`

    AuthorizedRoles  []string `json:"authorized_roles,omitempty"`

    RelatedFieldExtraDisplayInfos  []RelatedFieldExtraDisplayInfo `json:"related_field_extra_display_infos,omitempty"`

}

type OAPIUpdateFieldResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateFieldReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateFieldReqBuilder() *OAPIUpdateFieldReqBuilder {
	builder := &OAPIUpdateFieldReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateFieldReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) FieldName(fieldName *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).FieldName = fieldName
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) FieldKey(fieldKey *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).FieldKey = fieldKey
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) FieldValue(fieldValue interface{}) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).FieldValue = fieldValue
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) FreeAdd(freeAdd *int64) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).FreeAdd = freeAdd
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).WorkItemRelationUUID = workItemRelationUUID
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) DefaultValue(defaultValue interface{}) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).DefaultValue = defaultValue
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) FieldAlias(fieldAlias *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).FieldAlias = fieldAlias
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) HelpDescription(helpDescription *string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).HelpDescription = helpDescription
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).AuthorizedRoles = authorizedRoles
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) RelatedFieldExtraDisplayInfos(relatedFieldExtraDisplayInfos []RelatedFieldExtraDisplayInfo) *OAPIUpdateFieldReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFieldReqBody).RelatedFieldExtraDisplayInfos = relatedFieldExtraDisplayInfos
	return builder
}
func (builder *OAPIUpdateFieldReqBuilder) Build() *OAPIUpdateFieldReq {
	req := &OAPIUpdateFieldReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateFinishedReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateFinishedReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    NodeID  *string `json:"node_id,omitempty"`

    Opinion  *string `json:"opinion,omitempty"`

    FinishedConclusionOptionKey  *string `json:"finished_conclusion_option_key,omitempty"`

    OperationType  *string `json:"operation_type,omitempty"`

    Reset  *bool `json:"reset,omitempty"`

}

type OAPIUpdateFinishedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateFinishedReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateFinishedReqBuilder() *OAPIUpdateFinishedReqBuilder {
	builder := &OAPIUpdateFinishedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateFinishedReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) NodeID(nodeID *string) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).NodeID = nodeID
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) Opinion(opinion *string) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).Opinion = opinion
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) FinishedConclusionOptionKey(finishedConclusionOptionKey *string) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).FinishedConclusionOptionKey = finishedConclusionOptionKey
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) OperationType(operationType *string) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).OperationType = operationType
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) Reset(reset *bool) *OAPIUpdateFinishedReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateFinishedReqBody).Reset = reset
	return builder
}
func (builder *OAPIUpdateFinishedReqBuilder) Build() *OAPIUpdateFinishedReq {
	req := &OAPIUpdateFinishedReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateMultiSignalReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateMultiSignalReqBody struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    Details  []MultiSignalDetail `json:"details,omitempty"`

    UpdateType  *string `json:"update_type,omitempty"`

}

type OAPIUpdateMultiSignalResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *MultiSignal         `json:"data"`
	
}

type OAPIUpdateMultiSignalReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateMultiSignalReqBuilder() *OAPIUpdateMultiSignalReqBuilder {
	builder := &OAPIUpdateMultiSignalReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateMultiSignalReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) FieldKey(fieldKey *string) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateMultiSignalReqBody).FieldKey = fieldKey
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) FieldAlias(fieldAlias *string) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateMultiSignalReqBody).FieldAlias = fieldAlias
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) Details(details []MultiSignalDetail) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateMultiSignalReqBody).Details = details
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) UpdateType(updateType *string) *OAPIUpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateMultiSignalReqBody).UpdateType = updateType
	return builder
}
func (builder *OAPIUpdateMultiSignalReqBuilder) Build() *OAPIUpdateMultiSignalReq {
	req := &OAPIUpdateMultiSignalReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateNodeStateReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateNodeStateReqBody struct {

    Action  *string `json:"action,omitempty"`

    RollbackReason  *string `json:"rollback_reason,omitempty"`

    NodeOwners  []string `json:"node_owners,omitempty"`

    NodeSchedule  *Schedule `json:"node_schedule,omitempty"`

    Schedules  []Schedule `json:"schedules,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

}

type OAPIUpdateNodeStateResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateNodeStateReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateNodeStateReqBuilder() *OAPIUpdateNodeStateReqBuilder {
	builder := &OAPIUpdateNodeStateReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateNodeStateReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) NodeID(nodeID *string) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(*nodeID))
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) Action(action *string) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).Action = action
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) RollbackReason(rollbackReason *string) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).RollbackReason = rollbackReason
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) NodeOwners(nodeOwners []string) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).NodeOwners = nodeOwners
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) NodeSchedule(nodeSchedule *Schedule) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).NodeSchedule = nodeSchedule
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) Schedules(schedules []Schedule) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).Schedules = schedules
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) Fields(fields []FieldValuePair) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).Fields = fields
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *OAPIUpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateNodeStateReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *OAPIUpdateNodeStateReqBuilder) Build() *OAPIUpdateNodeStateReq {
	req := &OAPIUpdateNodeStateReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateStateFlowReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateStateFlowReqBody struct {

    TransitionID  *int64 `json:"transition_id,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

    RoleOwners  []RoleOwner `json:"role_owners,omitempty"`

}

type OAPIUpdateStateFlowResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateStateFlowReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateStateFlowReqBuilder() *OAPIUpdateStateFlowReqBuilder {
	builder := &OAPIUpdateStateFlowReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateStateFlowReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateStateFlowReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateStateFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateStateFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) TransitionID(transitionID *int64) *OAPIUpdateStateFlowReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateStateFlowReqBody).TransitionID = transitionID
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) Fields(fields []FieldValuePair) *OAPIUpdateStateFlowReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateStateFlowReqBody).Fields = fields
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) RoleOwners(roleOwners []RoleOwner) *OAPIUpdateStateFlowReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateStateFlowReqBody).RoleOwners = roleOwners
	return builder
}
func (builder *OAPIUpdateStateFlowReqBuilder) Build() *OAPIUpdateStateFlowReq {
	req := &OAPIUpdateStateFlowReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateTemplateDetailReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateTemplateDetailReqBody struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    WorkflowConfs  []WorkflowConfInfo `json:"workflow_confs,omitempty"`

    StateFlowConfs  []StateFlowConfInfo `json:"state_flow_confs,omitempty"`

}

type OAPIUpdateTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateTemplateDetailReqBuilder() *OAPIUpdateTemplateDetailReqBuilder {
	builder := &OAPIUpdateTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateTemplateDetailReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateTemplateDetailReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateTemplateDetailReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIUpdateTemplateDetailReqBuilder) TemplateID(templateID *int64) *OAPIUpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateTemplateDetailReqBody).TemplateID = templateID
	return builder
}
func (builder *OAPIUpdateTemplateDetailReqBuilder) WorkflowConfs(workflowConfs []WorkflowConfInfo) *OAPIUpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateTemplateDetailReqBody).WorkflowConfs = workflowConfs
	return builder
}
func (builder *OAPIUpdateTemplateDetailReqBuilder) StateFlowConfs(stateFlowConfs []StateFlowConfInfo) *OAPIUpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateTemplateDetailReqBody).StateFlowConfs = stateFlowConfs
	return builder
}
func (builder *OAPIUpdateTemplateDetailReqBuilder) Build() *OAPIUpdateTemplateDetailReq {
	req := &OAPIUpdateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateWorkItemReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateWorkItemReqBody struct {

    UpdateFields  []FieldValuePair `json:"update_fields,omitempty"`

}

type OAPIUpdateWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateWorkItemReqBuilder() *OAPIUpdateWorkItemReqBuilder {
	builder := &OAPIUpdateWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateWorkItemReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateWorkItemReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateWorkItemReqBuilder) UpdateFields(updateFields []FieldValuePair) *OAPIUpdateWorkItemReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemReqBody).UpdateFields = updateFields
	return builder
}
func (builder *OAPIUpdateWorkItemReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateWorkItemReqBuilder) Build() *OAPIUpdateWorkItemReq {
	req := &OAPIUpdateWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateWorkItemRelationReqBody struct {

    RelationID  *string `json:"relation_id,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    RelationDetails  []RelationDetail `json:"relation_details,omitempty"`

}

type OAPIUpdateWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateWorkItemRelationReqBuilder() *OAPIUpdateWorkItemRelationReqBuilder {
	builder := &OAPIUpdateWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateWorkItemRelationReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateWorkItemRelationReqBuilder) RelationID(relationID *string) *OAPIUpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemRelationReqBody).RelationID = relationID
	return builder
}
func (builder *OAPIUpdateWorkItemRelationReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemRelationReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIUpdateWorkItemRelationReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemRelationReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIUpdateWorkItemRelationReqBuilder) Name(name *string) *OAPIUpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemRelationReqBody).Name = name
	return builder
}
func (builder *OAPIUpdateWorkItemRelationReqBuilder) RelationDetails(relationDetails []RelationDetail) *OAPIUpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemRelationReqBody).RelationDetails = relationDetails
	return builder
}
func (builder *OAPIUpdateWorkItemRelationReqBuilder) Build() *OAPIUpdateWorkItemRelationReq {
	req := &OAPIUpdateWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateWorkItemSubTaskReqBody struct {

    Name  *string `json:"name,omitempty"`

    Assignee  []string `json:"assignee,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

    Schedule  *Schedule `json:"schedule,omitempty"`

    Note  *string `json:"note,omitempty"`

    Deliverable  []FieldValuePair `json:"deliverable,omitempty"`

    UpdateFields  []FieldValuePair `json:"update_fields,omitempty"`

}

type OAPIUpdateWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateWorkItemSubTaskReqBuilder() *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder := &OAPIUpdateWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateWorkItemSubTaskReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) NodeID(nodeID *string) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(*nodeID))
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) TaskID(taskID *int64) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("task_id", fmt.Sprint(*taskID))
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) Name(name *string) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).Name = name
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) Assignee(assignee []string) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).Assignee = assignee
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) Schedule(schedule *Schedule) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).Schedule = schedule
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) Note(note *string) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).Note = note
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) Deliverable(deliverable []FieldValuePair) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).Deliverable = deliverable
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) UpdateFields(updateFields []FieldValuePair) *OAPIUpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemSubTaskReqBody).UpdateFields = updateFields
	return builder
}
func (builder *OAPIUpdateWorkItemSubTaskReqBuilder) Build() *OAPIUpdateWorkItemSubTaskReq {
	req := &OAPIUpdateWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateWorkItemTypeInfoReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateWorkItemTypeInfoReqBody struct {

    Description  *string `json:"description,omitempty"`

    IsDisabled  *bool `json:"is_disabled,omitempty"`

    IsPinned  *bool `json:"is_pinned,omitempty"`

    EnableSchedule  *bool `json:"enable_schedule,omitempty"`

    ScheduleFieldKey  *string `json:"schedule_field_key,omitempty"`

    EstimatePointFieldKey  *string `json:"estimate_point_field_key,omitempty"`

    ActualWorkTimeFieldKey  *string `json:"actual_work_time_field_key,omitempty"`

    BelongRoleKeys  []string `json:"belong_role_keys,omitempty"`

    ActualWorkTimeSwitch  *bool `json:"actual_work_time_switch,omitempty"`

}

type OAPIUpdateWorkItemTypeInfoResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateWorkItemTypeInfoReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateWorkItemTypeInfoReqBuilder() *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder := &OAPIUpdateWorkItemTypeInfoReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateWorkItemTypeInfoReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) Description(description *string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).Description = description
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) IsDisabled(isDisabled *bool) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).IsDisabled = isDisabled
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) IsPinned(isPinned *bool) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).IsPinned = isPinned
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) EnableSchedule(enableSchedule *bool) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).EnableSchedule = enableSchedule
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) ScheduleFieldKey(scheduleFieldKey *string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).ScheduleFieldKey = scheduleFieldKey
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) EstimatePointFieldKey(estimatePointFieldKey *string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).EstimatePointFieldKey = estimatePointFieldKey
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) ActualWorkTimeFieldKey(actualWorkTimeFieldKey *string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).ActualWorkTimeFieldKey = actualWorkTimeFieldKey
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) BelongRoleKeys(belongRoleKeys []string) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).BelongRoleKeys = belongRoleKeys
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) ActualWorkTimeSwitch(actualWorkTimeSwitch *bool) *OAPIUpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkItemTypeInfoReqBody).ActualWorkTimeSwitch = actualWorkTimeSwitch
	return builder
}
func (builder *OAPIUpdateWorkItemTypeInfoReqBuilder) Build() *OAPIUpdateWorkItemTypeInfoReq {
	req := &OAPIUpdateWorkItemTypeInfoReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateWorkflowNodeReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateWorkflowNodeReqBody struct {

    NodeOwners  []string `json:"node_owners,omitempty"`

    NodeSchedule  *Schedule `json:"node_schedule,omitempty"`

    Schedules  []Schedule `json:"schedules,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

}

type OAPIUpdateWorkflowNodeResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateWorkflowNodeReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateWorkflowNodeReqBuilder() *OAPIUpdateWorkflowNodeReqBuilder {
	builder := &OAPIUpdateWorkflowNodeReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateWorkflowNodeReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) NodeID(nodeID *string) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(*nodeID))
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) NodeOwners(nodeOwners []string) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkflowNodeReqBody).NodeOwners = nodeOwners
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) NodeSchedule(nodeSchedule *Schedule) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkflowNodeReqBody).NodeSchedule = nodeSchedule
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) Schedules(schedules []Schedule) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkflowNodeReqBody).Schedules = schedules
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) Fields(fields []FieldValuePair) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkflowNodeReqBody).Fields = fields
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *OAPIUpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkflowNodeReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *OAPIUpdateWorkflowNodeReqBuilder) Build() *OAPIUpdateWorkflowNodeReq {
	req := &OAPIUpdateWorkflowNodeReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateWorkingHourRecordReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateWorkingHourRecordReqBody struct {

    WorkingHourRecords  []UpdateWorkingHourRecord `json:"working_hour_records,omitempty"`

}

type OAPIUpdateWorkingHourRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateWorkingHourRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateWorkingHourRecordReqBuilder() *OAPIUpdateWorkingHourRecordReqBuilder {
	builder := &OAPIUpdateWorkingHourRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateWorkingHourRecordReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateWorkingHourRecordReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIUpdateWorkingHourRecordReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *OAPIUpdateWorkingHourRecordReqBuilder) WorkItemID(workItemID *int64) *OAPIUpdateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIUpdateWorkingHourRecordReqBuilder) WorkingHourRecords(workingHourRecords []UpdateWorkingHourRecord) *OAPIUpdateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateWorkingHourRecordReqBody).WorkingHourRecords = workingHourRecords
	return builder
}
func (builder *OAPIUpdateWorkingHourRecordReqBuilder) Build() *OAPIUpdateWorkingHourRecordReq {
	req := &OAPIUpdateWorkingHourRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIWBSUpdateDraftFrozenRowsReq struct {
	apiReq *core.APIReq
}
type OAPIWBSUpdateDraftFrozenRowsReqBody struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    DraftID  *string `json:"draft_id,omitempty"`

    UpdateType  *int32 `json:"update_type,omitempty"`

    CommitID  *string `json:"commit_id,omitempty"`

}

type OAPIWBSUpdateDraftFrozenRowsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIWBSUpdateDraftFrozenRowsReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIWBSUpdateDraftFrozenRowsReqBuilder() *OAPIWBSUpdateDraftFrozenRowsReqBuilder {
	builder := &OAPIWBSUpdateDraftFrozenRowsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIWBSUpdateDraftFrozenRowsReqBody{},
	}
	return builder
}
func (builder *OAPIWBSUpdateDraftFrozenRowsReqBuilder) ProjectKey(projectKey *string) *OAPIWBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIWBSUpdateDraftFrozenRowsReqBuilder) WorkItemID(workItemID *int64) *OAPIWBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*OAPIWBSUpdateDraftFrozenRowsReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIWBSUpdateDraftFrozenRowsReqBuilder) DraftID(draftID *string) *OAPIWBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*OAPIWBSUpdateDraftFrozenRowsReqBody).DraftID = draftID
	return builder
}
func (builder *OAPIWBSUpdateDraftFrozenRowsReqBuilder) UpdateType(updateType *int32) *OAPIWBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*OAPIWBSUpdateDraftFrozenRowsReqBody).UpdateType = updateType
	return builder
}
func (builder *OAPIWBSUpdateDraftFrozenRowsReqBuilder) CommitID(commitID *string) *OAPIWBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*OAPIWBSUpdateDraftFrozenRowsReqBody).CommitID = commitID
	return builder
}
func (builder *OAPIWBSUpdateDraftFrozenRowsReqBuilder) Build() *OAPIWBSUpdateDraftFrozenRowsReq {
	req := &OAPIWBSUpdateDraftFrozenRowsReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIWbsCollaborationPublishReq struct {
	apiReq *core.APIReq
}
type OAPIWbsCollaborationPublishReqBody struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    DraftID  *string `json:"draft_id,omitempty"`

    CommitID  *string `json:"commit_id,omitempty"`

}

type OAPIWbsCollaborationPublishResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Success       *bool         `json:"success"`
	
}

type OAPIWbsCollaborationPublishReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIWbsCollaborationPublishReqBuilder() *OAPIWbsCollaborationPublishReqBuilder {
	builder := &OAPIWbsCollaborationPublishReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIWbsCollaborationPublishReqBody{},
	}
	return builder
}
func (builder *OAPIWbsCollaborationPublishReqBuilder) ProjectKey(projectKey *string) *OAPIWbsCollaborationPublishReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIWbsCollaborationPublishReqBuilder) WorkItemID(workItemID *int64) *OAPIWbsCollaborationPublishReqBuilder {
	builder.apiReq.Body.(*OAPIWbsCollaborationPublishReqBody).WorkItemID = workItemID
	return builder
}
func (builder *OAPIWbsCollaborationPublishReqBuilder) DraftID(draftID *string) *OAPIWbsCollaborationPublishReqBuilder {
	builder.apiReq.Body.(*OAPIWbsCollaborationPublishReqBody).DraftID = draftID
	return builder
}
func (builder *OAPIWbsCollaborationPublishReqBuilder) CommitID(commitID *string) *OAPIWbsCollaborationPublishReqBuilder {
	builder.apiReq.Body.(*OAPIWbsCollaborationPublishReqBody).CommitID = commitID
	return builder
}
func (builder *OAPIWbsCollaborationPublishReqBuilder) Build() *OAPIWbsCollaborationPublishReq {
	req := &OAPIWbsCollaborationPublishReq{}
	req.apiReq = builder.apiReq
	return req
}

type SearchByParamsReq struct {
	apiReq *core.APIReq
}
type SearchByParamsReqBody struct {

    SearchGroup  *SearchGroup `json:"search_group,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Fields  []string `json:"fields,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

}

type SearchByParamsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type SearchByParamsReqBuilder struct {
	apiReq *core.APIReq
}

func NewSearchByParamsReqBuilder() *SearchByParamsReqBuilder {
	builder := &SearchByParamsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SearchByParamsReqBody{},
	}
	return builder
}
func (builder *SearchByParamsReqBuilder) ProjectKey(projectKey *string) *SearchByParamsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *SearchByParamsReqBuilder) SearchGroup(searchGroup *SearchGroup) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).SearchGroup = searchGroup
	return builder
}
func (builder *SearchByParamsReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *SearchByParamsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(*workItemTypeKey))
	return builder
}
func (builder *SearchByParamsReqBuilder) PageNum(pageNum *int64) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).PageNum = pageNum
	return builder
}
func (builder *SearchByParamsReqBuilder) PageSize(pageSize *int64) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).PageSize = pageSize
	return builder
}
func (builder *SearchByParamsReqBuilder) Fields(fields []string) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).Fields = fields
	return builder
}
func (builder *SearchByParamsReqBuilder) Expand(expand *Expand) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).Expand = expand
	return builder
}
func (builder *SearchByParamsReqBuilder) Build() *SearchByParamsReq {
	req := &SearchByParamsReq{}
	req.apiReq = builder.apiReq
	return req
}

type UniversalSearchReq struct {
	apiReq *core.APIReq
}
type UniversalSearchReqBody struct {

    DataSources  []DataSource `json:"data_sources,omitempty"`

    UserKey  *string `json:"user_key,omitempty"`

    SearchGroup  *SearchGroup `json:"search_group,omitempty"`

    Sort  *Sort `json:"sort,omitempty"`

    Pagination  *Pagination `json:"pagination,omitempty"`

    FieldSelected  []string `json:"field_selected,omitempty"`

    Features  map[string]string `json:"features,omitempty"`

}

type UniversalSearchResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *string         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
	ExtraInfo       map[string]string         `json:"extra_info"`
	
}

type UniversalSearchReqBuilder struct {
	apiReq *core.APIReq
}

func NewUniversalSearchReqBuilder() *UniversalSearchReqBuilder {
	builder := &UniversalSearchReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UniversalSearchReqBody{},
	}
	return builder
}
func (builder *UniversalSearchReqBuilder) DataSources(dataSources []DataSource) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).DataSources = dataSources
	return builder
}
func (builder *UniversalSearchReqBuilder) UserKey(userKey *string) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).UserKey = userKey
	return builder
}
func (builder *UniversalSearchReqBuilder) SearchGroup(searchGroup *SearchGroup) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).SearchGroup = searchGroup
	return builder
}
func (builder *UniversalSearchReqBuilder) Sort(sort *Sort) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).Sort = sort
	return builder
}
func (builder *UniversalSearchReqBuilder) Pagination(pagination *Pagination) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).Pagination = pagination
	return builder
}
func (builder *UniversalSearchReqBuilder) FieldSelected(fieldSelected []string) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).FieldSelected = fieldSelected
	return builder
}
func (builder *UniversalSearchReqBuilder) Features(features map[string]string) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).Features = features
	return builder
}
func (builder *UniversalSearchReqBuilder) Build() *UniversalSearchReq {
	req := &UniversalSearchReq{}
	req.apiReq = builder.apiReq
	return req
}

