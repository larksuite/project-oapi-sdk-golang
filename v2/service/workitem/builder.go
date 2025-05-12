package workitem

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type AbortWorkItemReq struct {
	apiReq *core.APIReq
}
type AbortWorkItemReqBody struct {
    IsAborted  *bool `json:"is_aborted,omitempty"`
    Reason  *string `json:"reason,omitempty"`
    ReasonOption  *string `json:"reason_option,omitempty"`
}
type AbortWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *AbortWorkItemRespData        `json:"data,omitempty"`
}

type AbortWorkItemRespData struct {
}

type AbortWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewAbortWorkItemReqBuilder() *AbortWorkItemReqBuilder {
	builder := &AbortWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &AbortWorkItemReqBody{},
	}
	return builder
}

func (builder *AbortWorkItemReqBuilder) ProjectKey(projectKey string) *AbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *AbortWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *AbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *AbortWorkItemReqBuilder) WorkItemID(workItemID int64) *AbortWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *AbortWorkItemReqBuilder) IsAborted(isAborted bool) *AbortWorkItemReqBuilder {
	builder.apiReq.Body.(*AbortWorkItemReqBody).IsAborted = &isAborted
	return builder
}


func (builder *AbortWorkItemReqBuilder) Reason(reason string) *AbortWorkItemReqBuilder {
	builder.apiReq.Body.(*AbortWorkItemReqBody).Reason = &reason
	return builder
}


func (builder *AbortWorkItemReqBuilder) ReasonOption(reasonOption string) *AbortWorkItemReqBuilder {
	builder.apiReq.Body.(*AbortWorkItemReqBody).ReasonOption = &reasonOption
	return builder
}

func (builder *AbortWorkItemReqBuilder) Build() *AbortWorkItemReq {
	req := &AbortWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type ActualTimeUpdateReq struct {
	apiReq *core.APIReq
}
type ActualTimeUpdateReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    NodeID  *string `json:"node_id,omitempty"`
    ActualTimeInfo  *ActualTimeInfo `json:"actual_time_info,omitempty"`
}
type ActualTimeUpdateResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *ActualTimeUpdateRespData        `json:"data,omitempty"`
}

type ActualTimeUpdateRespData struct {
}

type ActualTimeUpdateReqBuilder struct {
	apiReq *core.APIReq
}

func NewActualTimeUpdateReqBuilder() *ActualTimeUpdateReqBuilder {
	builder := &ActualTimeUpdateReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ActualTimeUpdateReqBody{},
	}
	return builder
}

func (builder *ActualTimeUpdateReqBuilder) ProjectKey(projectKey string) *ActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*ActualTimeUpdateReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *ActualTimeUpdateReqBuilder) WorkItemID(workItemID int64) *ActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*ActualTimeUpdateReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *ActualTimeUpdateReqBuilder) NodeID(nodeID string) *ActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*ActualTimeUpdateReqBody).NodeID = &nodeID
	return builder
}


func (builder *ActualTimeUpdateReqBuilder) ActualTimeInfo(actualTimeInfo *ActualTimeInfo) *ActualTimeUpdateReqBuilder {
	builder.apiReq.Body.(*ActualTimeUpdateReqBody).ActualTimeInfo = actualTimeInfo
	return builder
}
func (builder *ActualTimeUpdateReqBuilder) Build() *ActualTimeUpdateReq {
	req := &ActualTimeUpdateReq{}
	req.apiReq = builder.apiReq
	return req
}

type BatchQueryConclusionOptionReq struct {
	apiReq *core.APIReq
}
type BatchQueryConclusionOptionReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    NodeIDs  []string `json:"node_ids,omitempty"`
}
type BatchQueryConclusionOptionResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []OAPIBatchQueryConclusionOptionItem         `json:"data"`
	
}

type BatchQueryConclusionOptionReqBuilder struct {
	apiReq *core.APIReq
}

func NewBatchQueryConclusionOptionReqBuilder() *BatchQueryConclusionOptionReqBuilder {
	builder := &BatchQueryConclusionOptionReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &BatchQueryConclusionOptionReqBody{},
	}
	return builder
}

func (builder *BatchQueryConclusionOptionReqBuilder) ProjectKey(projectKey string) *BatchQueryConclusionOptionReqBuilder {
	builder.apiReq.Body.(*BatchQueryConclusionOptionReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *BatchQueryConclusionOptionReqBuilder) WorkItemID(workItemID int64) *BatchQueryConclusionOptionReqBuilder {
	builder.apiReq.Body.(*BatchQueryConclusionOptionReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *BatchQueryConclusionOptionReqBuilder) NodeIDs(nodeIDs []string) *BatchQueryConclusionOptionReqBuilder {
	builder.apiReq.Body.(*BatchQueryConclusionOptionReqBody).NodeIDs = nodeIDs
	return builder
}
func (builder *BatchQueryConclusionOptionReqBuilder) Build() *BatchQueryConclusionOptionReq {
	req := &BatchQueryConclusionOptionReq{}
	req.apiReq = builder.apiReq
	return req
}

type BatchQueryDeliverableReq struct {
	apiReq *core.APIReq
}
type BatchQueryDeliverableReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`
}
type BatchQueryDeliverableResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []OAPIBatchQueryDeliverable         `json:"data"`
	
}

type BatchQueryDeliverableReqBuilder struct {
	apiReq *core.APIReq
}

func NewBatchQueryDeliverableReqBuilder() *BatchQueryDeliverableReqBuilder {
	builder := &BatchQueryDeliverableReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &BatchQueryDeliverableReqBody{},
	}
	return builder
}

func (builder *BatchQueryDeliverableReqBuilder) ProjectKey(projectKey string) *BatchQueryDeliverableReqBuilder {
	builder.apiReq.Body.(*BatchQueryDeliverableReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *BatchQueryDeliverableReqBuilder) WorkItemIDs(workItemIDs []int64) *BatchQueryDeliverableReqBuilder {
	builder.apiReq.Body.(*BatchQueryDeliverableReqBody).WorkItemIDs = workItemIDs
	return builder
}
func (builder *BatchQueryDeliverableReqBuilder) Build() *BatchQueryDeliverableReq {
	req := &BatchQueryDeliverableReq{}
	req.apiReq = builder.apiReq
	return req
}

type BatchQueryFinishedReq struct {
	apiReq *core.APIReq
}
type BatchQueryFinishedReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    NodeIDs  []string `json:"node_ids,omitempty"`
}
type BatchQueryFinishedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *BatchQueryFinishedRespData        `json:"data,omitempty"`
}

type BatchQueryFinishedRespData struct {
	ProjectKey       *string         `json:"project_key,omitempty"`
	WorkItemID       *int64         `json:"work_item_id,omitempty"`
	FinishedInfos       []OAPIFinishedInfoItem         `json:"finished_infos,omitempty"`
}

type BatchQueryFinishedReqBuilder struct {
	apiReq *core.APIReq
}

func NewBatchQueryFinishedReqBuilder() *BatchQueryFinishedReqBuilder {
	builder := &BatchQueryFinishedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &BatchQueryFinishedReqBody{},
	}
	return builder
}

func (builder *BatchQueryFinishedReqBuilder) ProjectKey(projectKey string) *BatchQueryFinishedReqBuilder {
	builder.apiReq.Body.(*BatchQueryFinishedReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *BatchQueryFinishedReqBuilder) WorkItemID(workItemID int64) *BatchQueryFinishedReqBuilder {
	builder.apiReq.Body.(*BatchQueryFinishedReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *BatchQueryFinishedReqBuilder) NodeIDs(nodeIDs []string) *BatchQueryFinishedReqBuilder {
	builder.apiReq.Body.(*BatchQueryFinishedReqBody).NodeIDs = nodeIDs
	return builder
}
func (builder *BatchQueryFinishedReqBuilder) Build() *BatchQueryFinishedReq {
	req := &BatchQueryFinishedReq{}
	req.apiReq = builder.apiReq
	return req
}

type BatchUpdateBasicWorkItemReq struct {
	apiReq *core.APIReq
}
type BatchUpdateBasicWorkItemReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`
    UpdateMode  *string `json:"update_mode,omitempty"`
    FieldKey  *string `json:"field_key,omitempty"`
    BeforeFieldValue  interface{} `json:"before_field_value,omitempty"`
    AfterFieldValue  interface{} `json:"after_field_value,omitempty"`
}
type BatchUpdateBasicWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *BatchUpdateBasicWorkItemRespData        `json:"data,omitempty"`
}

type BatchUpdateBasicWorkItemRespData struct {
	TaskID       *string         `json:"task_id,omitempty"`
}

type BatchUpdateBasicWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewBatchUpdateBasicWorkItemReqBuilder() *BatchUpdateBasicWorkItemReqBuilder {
	builder := &BatchUpdateBasicWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &BatchUpdateBasicWorkItemReqBody{},
	}
	return builder
}

func (builder *BatchUpdateBasicWorkItemReqBuilder) ProjectKey(projectKey string) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *BatchUpdateBasicWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *BatchUpdateBasicWorkItemReqBuilder) WorkItemIDs(workItemIDs []int64) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).WorkItemIDs = workItemIDs
	return builder
}

func (builder *BatchUpdateBasicWorkItemReqBuilder) UpdateMode(updateMode string) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).UpdateMode = &updateMode
	return builder
}


func (builder *BatchUpdateBasicWorkItemReqBuilder) FieldKey(fieldKey string) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).FieldKey = &fieldKey
	return builder
}


func (builder *BatchUpdateBasicWorkItemReqBuilder) BeforeFieldValue(beforeFieldValue interface{}) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).BeforeFieldValue = beforeFieldValue
	return builder
}

func (builder *BatchUpdateBasicWorkItemReqBuilder) AfterFieldValue(afterFieldValue interface{}) *BatchUpdateBasicWorkItemReqBuilder {
	builder.apiReq.Body.(*BatchUpdateBasicWorkItemReqBody).AfterFieldValue = afterFieldValue
	return builder
}
func (builder *BatchUpdateBasicWorkItemReqBuilder) Build() *BatchUpdateBasicWorkItemReq {
	req := &BatchUpdateBasicWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type CompositiveSearchReq struct {
	apiReq *core.APIReq
}
type CompositiveSearchReqBody struct {
    QueryType  *string `json:"query_type,omitempty"`
    Query  *string `json:"query,omitempty"`
    QuerySubType  []string `json:"query_sub_type,omitempty"`
    PageSize  *int64 `json:"page_size,omitempty"`
    PageNum  *int64 `json:"page_num,omitempty"`
    SimpleNames  []string `json:"simple_names,omitempty"`
}
type CompositiveSearchResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []CompInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type CompositiveSearchReqBuilder struct {
	apiReq *core.APIReq
}

func NewCompositiveSearchReqBuilder() *CompositiveSearchReqBuilder {
	builder := &CompositiveSearchReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CompositiveSearchReqBody{},
	}
	return builder
}

func (builder *CompositiveSearchReqBuilder) QueryType(queryType string) *CompositiveSearchReqBuilder {
	builder.apiReq.Body.(*CompositiveSearchReqBody).QueryType = &queryType
	return builder
}


func (builder *CompositiveSearchReqBuilder) Query(query string) *CompositiveSearchReqBuilder {
	builder.apiReq.Body.(*CompositiveSearchReqBody).Query = &query
	return builder
}


func (builder *CompositiveSearchReqBuilder) QuerySubType(querySubType []string) *CompositiveSearchReqBuilder {
	builder.apiReq.Body.(*CompositiveSearchReqBody).QuerySubType = querySubType
	return builder
}

func (builder *CompositiveSearchReqBuilder) PageSize(pageSize int64) *CompositiveSearchReqBuilder {
	builder.apiReq.Body.(*CompositiveSearchReqBody).PageSize = &pageSize
	return builder
}


func (builder *CompositiveSearchReqBuilder) PageNum(pageNum int64) *CompositiveSearchReqBuilder {
	builder.apiReq.Body.(*CompositiveSearchReqBody).PageNum = &pageNum
	return builder
}


func (builder *CompositiveSearchReqBuilder) SimpleNames(simpleNames []string) *CompositiveSearchReqBuilder {
	builder.apiReq.Body.(*CompositiveSearchReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *CompositiveSearchReqBuilder) Build() *CompositiveSearchReq {
	req := &CompositiveSearchReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateFieldReq struct {
	apiReq *core.APIReq
}
type CreateFieldReqBody struct {
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
    TeamOption  *TeamOption `json:"team_option,omitempty"`
    ParentFieldKey  *string `json:"parent_field_key,omitempty"`
}
type CreateFieldResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *string         `json:"data"`
	
}

type CreateFieldReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateFieldReqBuilder() *CreateFieldReqBuilder {
	builder := &CreateFieldReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateFieldReqBody{},
	}
	return builder
}

func (builder *CreateFieldReqBuilder) ProjectKey(projectKey string) *CreateFieldReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateFieldReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateFieldReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateFieldReqBuilder) FieldName(fieldName string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldName = &fieldName
	return builder
}


func (builder *CreateFieldReqBuilder) FieldTypeKey(fieldTypeKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldTypeKey = &fieldTypeKey
	return builder
}


func (builder *CreateFieldReqBuilder) ValueType(valueType int64) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ValueType = &valueType
	return builder
}


func (builder *CreateFieldReqBuilder) ReferenceWorkItemTypeKey(referenceWorkItemTypeKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ReferenceWorkItemTypeKey = &referenceWorkItemTypeKey
	return builder
}


func (builder *CreateFieldReqBuilder) ReferenceFieldKey(referenceFieldKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ReferenceFieldKey = &referenceFieldKey
	return builder
}


func (builder *CreateFieldReqBuilder) FieldValue(fieldValue interface{}) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldValue = fieldValue
	return builder
}

func (builder *CreateFieldReqBuilder) FreeAdd(freeAdd int64) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FreeAdd = &freeAdd
	return builder
}


func (builder *CreateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).WorkItemRelationUUID = &workItemRelationUUID
	return builder
}


func (builder *CreateFieldReqBuilder) DefaultValue(defaultValue interface{}) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).DefaultValue = defaultValue
	return builder
}

func (builder *CreateFieldReqBuilder) FieldAlias(fieldAlias string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).FieldAlias = &fieldAlias
	return builder
}


func (builder *CreateFieldReqBuilder) HelpDescription(helpDescription string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).HelpDescription = &helpDescription
	return builder
}


func (builder *CreateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).AuthorizedRoles = authorizedRoles
	return builder
}

func (builder *CreateFieldReqBuilder) IsMulti(isMulti bool) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).IsMulti = &isMulti
	return builder
}


func (builder *CreateFieldReqBuilder) Format(format bool) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).Format = &format
	return builder
}


func (builder *CreateFieldReqBuilder) RelatedFieldExtraDisplayInfos(relatedFieldExtraDisplayInfos []RelatedFieldExtraDisplayInfo) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).RelatedFieldExtraDisplayInfos = relatedFieldExtraDisplayInfos
	return builder
}

func (builder *CreateFieldReqBuilder) TeamOption(teamOption *TeamOption) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).TeamOption = teamOption
	return builder
}

func (builder *CreateFieldReqBuilder) ParentFieldKey(parentFieldKey string) *CreateFieldReqBuilder {
	builder.apiReq.Body.(*CreateFieldReqBody).ParentFieldKey = &parentFieldKey
	return builder
}

func (builder *CreateFieldReqBuilder) Build() *CreateFieldReq {
	req := &CreateFieldReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateProjectRelationInstancesReq struct {
	apiReq *core.APIReq
}
type CreateProjectRelationInstancesReqBody struct {
    RelationRuleID  *string `json:"relation_rule_id,omitempty"`
    Instances  []RelationBindInstance `json:"instances,omitempty"`
}
type CreateProjectRelationInstancesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *CreateProjectRelationInstancesRespData        `json:"data,omitempty"`
}

type CreateProjectRelationInstancesRespData struct {
}

type CreateProjectRelationInstancesReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateProjectRelationInstancesReqBuilder() *CreateProjectRelationInstancesReqBuilder {
	builder := &CreateProjectRelationInstancesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateProjectRelationInstancesReqBody{},
	}
	return builder
}

func (builder *CreateProjectRelationInstancesReqBuilder) ProjectKey(projectKey string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateProjectRelationInstancesReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateProjectRelationInstancesReqBuilder) WorkItemID(workItemID int64) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *CreateProjectRelationInstancesReqBuilder) RelationRuleID(relationRuleID string) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Body.(*CreateProjectRelationInstancesReqBody).RelationRuleID = &relationRuleID
	return builder
}


func (builder *CreateProjectRelationInstancesReqBuilder) Instances(instances []RelationBindInstance) *CreateProjectRelationInstancesReqBuilder {
	builder.apiReq.Body.(*CreateProjectRelationInstancesReqBody).Instances = instances
	return builder
}
func (builder *CreateProjectRelationInstancesReqBuilder) Build() *CreateProjectRelationInstancesReq {
	req := &CreateProjectRelationInstancesReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateStoryRelationsReq struct {
	apiReq *core.APIReq
}
type CreateStoryRelationsReqBody struct {
    StoryRelations  []StoryRelationEntity `json:"story_relations,omitempty"`
}
type CreateStoryRelationsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       map[string]int64         `json:"data"`
	
}

type CreateStoryRelationsReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateStoryRelationsReqBuilder() *CreateStoryRelationsReqBuilder {
	builder := &CreateStoryRelationsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateStoryRelationsReqBody{},
	}
	return builder
}

func (builder *CreateStoryRelationsReqBuilder) StoryRelations(storyRelations []StoryRelationEntity) *CreateStoryRelationsReqBuilder {
	builder.apiReq.Body.(*CreateStoryRelationsReqBody).StoryRelations = storyRelations
	return builder
}

func (builder *CreateStoryRelationsReqBuilder) ProjectKey(projectKey string) *CreateStoryRelationsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *CreateStoryRelationsReqBuilder) Build() *CreateStoryRelationsReq {
	req := &CreateStoryRelationsReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateTemplateDetailReq struct {
	apiReq *core.APIReq
}
type CreateTemplateDetailReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    TemplateName  *string `json:"template_name,omitempty"`
    CopyTemplateID  *int64 `json:"copy_template_id,omitempty"`
}
type CreateTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type CreateTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateTemplateDetailReqBuilder() *CreateTemplateDetailReqBuilder {
	builder := &CreateTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateTemplateDetailReqBody{},
	}
	return builder
}

func (builder *CreateTemplateDetailReqBuilder) ProjectKey(projectKey string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *CreateTemplateDetailReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *CreateTemplateDetailReqBuilder) TemplateName(templateName string) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).TemplateName = &templateName
	return builder
}


func (builder *CreateTemplateDetailReqBuilder) CopyTemplateID(copyTemplateID int64) *CreateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*CreateTemplateDetailReqBody).CopyTemplateID = &copyTemplateID
	return builder
}

func (builder *CreateTemplateDetailReqBuilder) Build() *CreateTemplateDetailReq {
	req := &CreateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateWorkItemReq struct {
	apiReq *core.APIReq
}
type CreateWorkItemReqBody struct {
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    FieldValuePairs  []FieldValuePair `json:"field_value_pairs,omitempty"`
    TemplateID  *int64 `json:"template_id,omitempty"`
    Name  *string `json:"name,omitempty"`
}
type CreateWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type CreateWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateWorkItemReqBuilder() *CreateWorkItemReqBuilder {
	builder := &CreateWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateWorkItemReqBody{},
	}
	return builder
}

func (builder *CreateWorkItemReqBuilder) ProjectKey(projectKey string) *CreateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateWorkItemReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *CreateWorkItemReqBuilder) FieldValuePairs(fieldValuePairs []FieldValuePair) *CreateWorkItemReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemReqBody).FieldValuePairs = fieldValuePairs
	return builder
}

func (builder *CreateWorkItemReqBuilder) TemplateID(templateID int64) *CreateWorkItemReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemReqBody).TemplateID = &templateID
	return builder
}


func (builder *CreateWorkItemReqBuilder) Name(name string) *CreateWorkItemReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemReqBody).Name = &name
	return builder
}

func (builder *CreateWorkItemReqBuilder) Build() *CreateWorkItemReq {
	req := &CreateWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type CreateWorkItemRelationReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    Name  *string `json:"name,omitempty"`
    RelationDetails  []RelationDetail `json:"relation_details,omitempty"`
}
type CreateWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *CreateWorkItemRelationData         `json:"data"`
	
}

type CreateWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateWorkItemRelationReqBuilder() *CreateWorkItemRelationReqBuilder {
	builder := &CreateWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateWorkItemRelationReqBody{},
	}
	return builder
}

func (builder *CreateWorkItemRelationReqBuilder) ProjectKey(projectKey string) *CreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemRelationReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *CreateWorkItemRelationReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemRelationReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *CreateWorkItemRelationReqBuilder) Name(name string) *CreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemRelationReqBody).Name = &name
	return builder
}


func (builder *CreateWorkItemRelationReqBuilder) RelationDetails(relationDetails []RelationDetail) *CreateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemRelationReqBody).RelationDetails = relationDetails
	return builder
}
func (builder *CreateWorkItemRelationReqBuilder) Build() *CreateWorkItemRelationReq {
	req := &CreateWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}
type CreateWorkItemSubTaskReqBody struct {
    NodeID  *string `json:"node_id,omitempty"`
    Name  *string `json:"name,omitempty"`
    AliasKey  *string `json:"alias_key,omitempty"`
    Assignee  []string `json:"assignee,omitempty"`
    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`
    Schedule  *Schedule `json:"schedule,omitempty"`
    Note  *string `json:"note,omitempty"`
    FieldValuePairs  []FieldValuePair `json:"field_value_pairs,omitempty"`
}
type CreateWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *int64         `json:"data"`
	
}

type CreateWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateWorkItemSubTaskReqBuilder() *CreateWorkItemSubTaskReqBuilder {
	builder := &CreateWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateWorkItemSubTaskReqBody{},
	}
	return builder
}

func (builder *CreateWorkItemSubTaskReqBuilder) ProjectKey(projectKey string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) WorkItemID(workItemID int64) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) NodeID(nodeID string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).NodeID = &nodeID
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) Name(name string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).Name = &name
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) AliasKey(aliasKey string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).AliasKey = &aliasKey
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) Assignee(assignee []string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).Assignee = assignee
	return builder
}

func (builder *CreateWorkItemSubTaskReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).RoleAssignee = roleAssignee
	return builder
}

func (builder *CreateWorkItemSubTaskReqBuilder) Schedule(schedule *Schedule) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).Schedule = schedule
	return builder
}

func (builder *CreateWorkItemSubTaskReqBuilder) Note(note string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).Note = &note
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateWorkItemSubTaskReqBuilder) FieldValuePairs(fieldValuePairs []FieldValuePair) *CreateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*CreateWorkItemSubTaskReqBody).FieldValuePairs = fieldValuePairs
	return builder
}
func (builder *CreateWorkItemSubTaskReqBuilder) Build() *CreateWorkItemSubTaskReq {
	req := &CreateWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateWorkingHourRecordReq struct {
	apiReq *core.APIReq
}
type CreateWorkingHourRecordReqBody struct {
    WorkBeginDate  *int64 `json:"work_begin_date,omitempty"`
    WorkEndDate  *int64 `json:"work_end_date,omitempty"`
    IncludeHolidays  *bool `json:"include_holidays,omitempty"`
    WorkingHourRecords  []CreateWorkingHourRecord `json:"working_hour_records,omitempty"`
}
type CreateWorkingHourRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ManHourRecord         `json:"data"`
	
}

type CreateWorkingHourRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateWorkingHourRecordReqBuilder() *CreateWorkingHourRecordReqBuilder {
	builder := &CreateWorkingHourRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateWorkingHourRecordReqBody{},
	}
	return builder
}

func (builder *CreateWorkingHourRecordReqBuilder) ProjectKey(projectKey string) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateWorkingHourRecordReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateWorkingHourRecordReqBuilder) WorkItemID(workItemID int64) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *CreateWorkingHourRecordReqBuilder) WorkBeginDate(workBeginDate int64) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*CreateWorkingHourRecordReqBody).WorkBeginDate = &workBeginDate
	return builder
}


func (builder *CreateWorkingHourRecordReqBuilder) WorkEndDate(workEndDate int64) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*CreateWorkingHourRecordReqBody).WorkEndDate = &workEndDate
	return builder
}


func (builder *CreateWorkingHourRecordReqBuilder) IncludeHolidays(includeHolidays bool) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*CreateWorkingHourRecordReqBody).IncludeHolidays = &includeHolidays
	return builder
}


func (builder *CreateWorkingHourRecordReqBuilder) WorkingHourRecords(workingHourRecords []CreateWorkingHourRecord) *CreateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*CreateWorkingHourRecordReqBody).WorkingHourRecords = workingHourRecords
	return builder
}
func (builder *CreateWorkingHourRecordReqBuilder) Build() *CreateWorkingHourRecordReq {
	req := &CreateWorkingHourRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteFileReq struct {
	apiReq *core.APIReq
}
type DeleteFileReqBody struct {
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
    FieldKey  *string `json:"field_key,omitempty"`
    FieldAlias  *string `json:"field_alias,omitempty"`
    Uuids  []string `json:"uuids,omitempty"`
}
type DeleteFileResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *DeleteFileRespData        `json:"data,omitempty"`
}

type DeleteFileRespData struct {
}

type DeleteFileReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteFileReqBuilder() *DeleteFileReqBuilder {
	builder := &DeleteFileReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &DeleteFileReqBody{},
	}
	return builder
}

func (builder *DeleteFileReqBuilder) WorkItemID(workItemID int64) *DeleteFileReqBuilder {
	builder.apiReq.Body.(*DeleteFileReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *DeleteFileReqBuilder) ProjectKey(projectKey string) *DeleteFileReqBuilder {
	builder.apiReq.Body.(*DeleteFileReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *DeleteFileReqBuilder) FieldKey(fieldKey string) *DeleteFileReqBuilder {
	builder.apiReq.Body.(*DeleteFileReqBody).FieldKey = &fieldKey
	return builder
}


func (builder *DeleteFileReqBuilder) FieldAlias(fieldAlias string) *DeleteFileReqBuilder {
	builder.apiReq.Body.(*DeleteFileReqBody).FieldAlias = &fieldAlias
	return builder
}


func (builder *DeleteFileReqBuilder) Uuids(uuids []string) *DeleteFileReqBuilder {
	builder.apiReq.Body.(*DeleteFileReqBody).Uuids = uuids
	return builder
}
func (builder *DeleteFileReqBuilder) Build() *DeleteFileReq {
	req := &DeleteFileReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteProjectRelationInstanceReq struct {
	apiReq *core.APIReq
}
type DeleteProjectRelationInstanceReqBody struct {
    RelationRuleID  *string `json:"relation_rule_id,omitempty"`
    RelationWorkItemID  *int64 `json:"relation_work_item_id,omitempty"`
}
type DeleteProjectRelationInstanceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *DeleteProjectRelationInstanceRespData        `json:"data,omitempty"`
}

type DeleteProjectRelationInstanceRespData struct {
}

type DeleteProjectRelationInstanceReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteProjectRelationInstanceReqBuilder() *DeleteProjectRelationInstanceReqBuilder {
	builder := &DeleteProjectRelationInstanceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &DeleteProjectRelationInstanceReqBody{},
	}
	return builder
}

func (builder *DeleteProjectRelationInstanceReqBuilder) ProjectKey(projectKey string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteProjectRelationInstanceReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *DeleteProjectRelationInstanceReqBuilder) WorkItemID(workItemID int64) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *DeleteProjectRelationInstanceReqBuilder) RelationRuleID(relationRuleID string) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*DeleteProjectRelationInstanceReqBody).RelationRuleID = &relationRuleID
	return builder
}


func (builder *DeleteProjectRelationInstanceReqBuilder) RelationWorkItemID(relationWorkItemID int64) *DeleteProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*DeleteProjectRelationInstanceReqBody).RelationWorkItemID = &relationWorkItemID
	return builder
}

func (builder *DeleteProjectRelationInstanceReqBuilder) Build() *DeleteProjectRelationInstanceReq {
	req := &DeleteProjectRelationInstanceReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteTemplateDetailReq struct {
	apiReq *core.APIReq
}
type DeleteTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *DeleteTemplateDetailRespData        `json:"data,omitempty"`
}

type DeleteTemplateDetailRespData struct {
}

type DeleteTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteTemplateDetailReqBuilder() *DeleteTemplateDetailReqBuilder {
	builder := &DeleteTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *DeleteTemplateDetailReqBuilder) ProjectKey(projectKey string) *DeleteTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteTemplateDetailReqBuilder) TemplateID(templateID int64) *DeleteTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("template_id", fmt.Sprint(templateID))
	return builder
}

func (builder *DeleteTemplateDetailReqBuilder) Build() *DeleteTemplateDetailReq {
	req := &DeleteTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteWorkItemReq struct {
	apiReq *core.APIReq
}
type DeleteWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *DeleteWorkItemRespData        `json:"data,omitempty"`
}

type DeleteWorkItemRespData struct {
}

type DeleteWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteWorkItemReqBuilder() *DeleteWorkItemReqBuilder {
	builder := &DeleteWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *DeleteWorkItemReqBuilder) ProjectKey(projectKey string) *DeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *DeleteWorkItemReqBuilder) WorkItemID(workItemID int64) *DeleteWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}

func (builder *DeleteWorkItemReqBuilder) Build() *DeleteWorkItemReq {
	req := &DeleteWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type DeleteWorkItemRelationReqBody struct {
    RelationID  *string `json:"relation_id,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
}
type DeleteWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *string         `json:"data"`
	
}

type DeleteWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteWorkItemRelationReqBuilder() *DeleteWorkItemRelationReqBuilder {
	builder := &DeleteWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &DeleteWorkItemRelationReqBody{},
	}
	return builder
}

func (builder *DeleteWorkItemRelationReqBuilder) RelationID(relationID string) *DeleteWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*DeleteWorkItemRelationReqBody).RelationID = &relationID
	return builder
}


func (builder *DeleteWorkItemRelationReqBuilder) ProjectKey(projectKey string) *DeleteWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*DeleteWorkItemRelationReqBody).ProjectKey = &projectKey
	return builder
}

func (builder *DeleteWorkItemRelationReqBuilder) Build() *DeleteWorkItemRelationReq {
	req := &DeleteWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}
type DeleteWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *DeleteWorkItemSubTaskRespData        `json:"data,omitempty"`
}

type DeleteWorkItemSubTaskRespData struct {
}

type DeleteWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteWorkItemSubTaskReqBuilder() *DeleteWorkItemSubTaskReqBuilder {
	builder := &DeleteWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *DeleteWorkItemSubTaskReqBuilder) ProjectKey(projectKey string) *DeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteWorkItemSubTaskReqBuilder) WorkItemID(workItemID int64) *DeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *DeleteWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *DeleteWorkItemSubTaskReqBuilder) TaskID(taskID int64) *DeleteWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("task_id", fmt.Sprint(taskID))
	return builder
}

func (builder *DeleteWorkItemSubTaskReqBuilder) Build() *DeleteWorkItemSubTaskReq {
	req := &DeleteWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteWorkingHourRecordReq struct {
	apiReq *core.APIReq
}
type DeleteWorkingHourRecordReqBody struct {
    WorkingHourRecordIDs  []int64 `json:"working_hour_record_ids,omitempty"`
}
type DeleteWorkingHourRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *DeleteWorkingHourRecordRespData        `json:"data,omitempty"`
}

type DeleteWorkingHourRecordRespData struct {
}

type DeleteWorkingHourRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteWorkingHourRecordReqBuilder() *DeleteWorkingHourRecordReqBuilder {
	builder := &DeleteWorkingHourRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &DeleteWorkingHourRecordReqBody{},
	}
	return builder
}

func (builder *DeleteWorkingHourRecordReqBuilder) ProjectKey(projectKey string) *DeleteWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteWorkingHourRecordReqBuilder) WorkItemTypeKey(workItemTypeKey string) *DeleteWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *DeleteWorkingHourRecordReqBuilder) WorkItemID(workItemID int64) *DeleteWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *DeleteWorkingHourRecordReqBuilder) WorkingHourRecordIDs(workingHourRecordIDs []int64) *DeleteWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*DeleteWorkingHourRecordReqBody).WorkingHourRecordIDs = workingHourRecordIDs
	return builder
}
func (builder *DeleteWorkingHourRecordReqBuilder) Build() *DeleteWorkingHourRecordReq {
	req := &DeleteWorkingHourRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type ElementTemplateCreateReq struct {
	apiReq *core.APIReq
}
type ElementTemplateCreateReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    ElementType  *string `json:"element_type,omitempty"`
    NodeElement  *NodeElement `json:"node_element,omitempty"`
    TaskElement  *TaskElement `json:"task_element,omitempty"`
}
type ElementTemplateCreateResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *ElementTemplateCreateRespData        `json:"data,omitempty"`
}

type ElementTemplateCreateRespData struct {
	ElementKey       *string         `json:"element_key,omitempty"`
}

type ElementTemplateCreateReqBuilder struct {
	apiReq *core.APIReq
}

func NewElementTemplateCreateReqBuilder() *ElementTemplateCreateReqBuilder {
	builder := &ElementTemplateCreateReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ElementTemplateCreateReqBody{},
	}
	return builder
}

func (builder *ElementTemplateCreateReqBuilder) ProjectKey(projectKey string) *ElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*ElementTemplateCreateReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *ElementTemplateCreateReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*ElementTemplateCreateReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *ElementTemplateCreateReqBuilder) ElementType(elementType string) *ElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*ElementTemplateCreateReqBody).ElementType = &elementType
	return builder
}


func (builder *ElementTemplateCreateReqBuilder) NodeElement(nodeElement *NodeElement) *ElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*ElementTemplateCreateReqBody).NodeElement = nodeElement
	return builder
}

func (builder *ElementTemplateCreateReqBuilder) TaskElement(taskElement *TaskElement) *ElementTemplateCreateReqBuilder {
	builder.apiReq.Body.(*ElementTemplateCreateReqBody).TaskElement = taskElement
	return builder
}
func (builder *ElementTemplateCreateReqBuilder) Build() *ElementTemplateCreateReq {
	req := &ElementTemplateCreateReq{}
	req.apiReq = builder.apiReq
	return req
}

type ElementTemplateQueryReq struct {
	apiReq *core.APIReq
}
type ElementTemplateQueryReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    ElementType  *string `json:"element_type,omitempty"`
    PageNum  *int64 `json:"page_num,omitempty"`
    PageSize  *int64 `json:"page_size,omitempty"`
}
type ElementTemplateQueryResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *ElementTemplateQueryRespData        `json:"data,omitempty"`
}

type ElementTemplateQueryRespData struct {
	NodeElements       []NodeElement         `json:"node_elements,omitempty"`
	TaskElements       []TaskElement         `json:"task_elements,omitempty"`
	Total       *int64         `json:"total,omitempty"`
}

type ElementTemplateQueryReqBuilder struct {
	apiReq *core.APIReq
}

func NewElementTemplateQueryReqBuilder() *ElementTemplateQueryReqBuilder {
	builder := &ElementTemplateQueryReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ElementTemplateQueryReqBody{},
	}
	return builder
}

func (builder *ElementTemplateQueryReqBuilder) ProjectKey(projectKey string) *ElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*ElementTemplateQueryReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *ElementTemplateQueryReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*ElementTemplateQueryReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *ElementTemplateQueryReqBuilder) ElementType(elementType string) *ElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*ElementTemplateQueryReqBody).ElementType = &elementType
	return builder
}


func (builder *ElementTemplateQueryReqBuilder) PageNum(pageNum int64) *ElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*ElementTemplateQueryReqBody).PageNum = &pageNum
	return builder
}


func (builder *ElementTemplateQueryReqBuilder) PageSize(pageSize int64) *ElementTemplateQueryReqBuilder {
	builder.apiReq.Body.(*ElementTemplateQueryReqBody).PageSize = &pageSize
	return builder
}

func (builder *ElementTemplateQueryReqBuilder) Build() *ElementTemplateQueryReq {
	req := &ElementTemplateQueryReq{}
	req.apiReq = builder.apiReq
	return req
}

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

func (builder *FilterReqBuilder) ProjectKey(projectKey string) *FilterReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *FilterReqBuilder) WorkItemName(workItemName string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).WorkItemName = &workItemName
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

func (builder *FilterReqBuilder) PageNum(pageNum int64) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).PageNum = &pageNum
	return builder
}


func (builder *FilterReqBuilder) PageSize(pageSize int64) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).PageSize = &pageSize
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

func (builder *FilterReqBuilder) SearchID(searchID string) *FilterReqBuilder {
	builder.apiReq.Body.(*FilterReqBody).SearchID = &searchID
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

func (builder *FilterAcrossProjectReqBuilder) WorkItemTypeKey(workItemTypeKey string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).WorkItemTypeKey = &workItemTypeKey
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

func (builder *FilterAcrossProjectReqBuilder) WorkItemName(workItemName string) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).WorkItemName = &workItemName
	return builder
}


func (builder *FilterAcrossProjectReqBuilder) PageNum(pageNum int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).PageNum = &pageNum
	return builder
}


func (builder *FilterAcrossProjectReqBuilder) PageSize(pageSize int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).PageSize = &pageSize
	return builder
}


func (builder *FilterAcrossProjectReqBuilder) TenantGroupID(tenantGroupID int64) *FilterAcrossProjectReqBuilder {
	builder.apiReq.Body.(*FilterAcrossProjectReqBody).TenantGroupID = &tenantGroupID
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

type FreezeWorkItemReq struct {
	apiReq *core.APIReq
}
type FreezeWorkItemReqBody struct {
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    IsFrozen  *bool `json:"is_frozen,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
}
type FreezeWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *FreezeWorkItemRespData        `json:"data,omitempty"`
}

type FreezeWorkItemRespData struct {
}

type FreezeWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewFreezeWorkItemReqBuilder() *FreezeWorkItemReqBuilder {
	builder := &FreezeWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &FreezeWorkItemReqBody{},
	}
	return builder
}

func (builder *FreezeWorkItemReqBuilder) WorkItemID(workItemID int64) *FreezeWorkItemReqBuilder {
	builder.apiReq.Body.(*FreezeWorkItemReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *FreezeWorkItemReqBuilder) IsFrozen(isFrozen bool) *FreezeWorkItemReqBuilder {
	builder.apiReq.Body.(*FreezeWorkItemReqBody).IsFrozen = &isFrozen
	return builder
}


func (builder *FreezeWorkItemReqBuilder) ProjectKey(projectKey string) *FreezeWorkItemReqBuilder {
	builder.apiReq.Body.(*FreezeWorkItemReqBody).ProjectKey = &projectKey
	return builder
}

func (builder *FreezeWorkItemReqBuilder) Build() *FreezeWorkItemReq {
	req := &FreezeWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWBSInfoReq struct {
	apiReq *core.APIReq
}
type GetWBSInfoReqBody struct {
    Expand  *Expand `json:"expand,omitempty"`
}
type GetWBSInfoResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *WBSInfo         `json:"data"`
	
}

type GetWBSInfoReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWBSInfoReqBuilder() *GetWBSInfoReqBuilder {
	builder := &GetWBSInfoReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &GetWBSInfoReqBody{},
	}
	return builder
}

func (builder *GetWBSInfoReqBuilder) ProjectKey(projectKey string) *GetWBSInfoReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *GetWBSInfoReqBuilder) WorkItemID(workItemID int64) *GetWBSInfoReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *GetWBSInfoReqBuilder) WorkItemTypeKey(workItemTypeKey string) *GetWBSInfoReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *GetWBSInfoReqBuilder) Expand(expand *Expand) *GetWBSInfoReqBuilder {
	builder.apiReq.Body.(*GetWBSInfoReqBody).Expand = expand
	return builder
}

func (builder *GetWBSInfoReqBuilder) NeedUnionDeliverable(needUnionDeliverable bool) *GetWBSInfoReqBuilder {
	builder.apiReq.QueryParams.Set("need_union_deliverable", fmt.Sprint(needUnionDeliverable))
	return builder
}

func (builder *GetWBSInfoReqBuilder) Build() *GetWBSInfoReq {
	req := &GetWBSInfoReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWBSViewSubWorkItemConfReq struct {
	apiReq *core.APIReq
}
type GetWBSViewSubWorkItemConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *GetWBSViewSubWorkItemConfRespData        `json:"data,omitempty"`
}

type GetWBSViewSubWorkItemConfRespData struct {
	DraftViewSubWorkItemConfs       []DraftViewSubWorkItemConf         `json:"draft_view_sub_work_item_confs,omitempty"`
}

type GetWBSViewSubWorkItemConfReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWBSViewSubWorkItemConfReqBuilder() *GetWBSViewSubWorkItemConfReqBuilder {
	builder := &GetWBSViewSubWorkItemConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *GetWBSViewSubWorkItemConfReqBuilder) DraftID(draftID string) *GetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("draft_id", fmt.Sprint(draftID))
	return builder
}


func (builder *GetWBSViewSubWorkItemConfReqBuilder) WorkItemID(workItemID int64) *GetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *GetWBSViewSubWorkItemConfReqBuilder) NodeUUID(nodeUUID string) *GetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("node_uuid", fmt.Sprint(nodeUUID))
	return builder
}


func (builder *GetWBSViewSubWorkItemConfReqBuilder) ProjectKey(projectKey string) *GetWBSViewSubWorkItemConfReqBuilder {
	builder.apiReq.QueryParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *GetWBSViewSubWorkItemConfReqBuilder) Build() *GetWBSViewSubWorkItemConfReq {
	req := &GetWBSViewSubWorkItemConfReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWorkFlowReq struct {
	apiReq *core.APIReq
}
type GetWorkFlowReqBody struct {
    Fields  []string `json:"fields,omitempty"`
    FlowType  *int64 `json:"flow_type,omitempty"`
    Expand  *Expand `json:"expand,omitempty"`
}
type GetWorkFlowResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *NodesConnections         `json:"data"`
	
}

type GetWorkFlowReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWorkFlowReqBuilder() *GetWorkFlowReqBuilder {
	builder := &GetWorkFlowReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &GetWorkFlowReqBody{},
	}
	return builder
}

func (builder *GetWorkFlowReqBuilder) ProjectKey(projectKey string) *GetWorkFlowReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *GetWorkFlowReqBuilder) WorkItemID(workItemID int64) *GetWorkFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *GetWorkFlowReqBuilder) Fields(fields []string) *GetWorkFlowReqBuilder {
	builder.apiReq.Body.(*GetWorkFlowReqBody).Fields = fields
	return builder
}

func (builder *GetWorkFlowReqBuilder) FlowType(flowType int64) *GetWorkFlowReqBuilder {
	builder.apiReq.Body.(*GetWorkFlowReqBody).FlowType = &flowType
	return builder
}


func (builder *GetWorkFlowReqBuilder) WorkItemTypeKey(workItemTypeKey string) *GetWorkFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *GetWorkFlowReqBuilder) Expand(expand *Expand) *GetWorkFlowReqBuilder {
	builder.apiReq.Body.(*GetWorkFlowReqBody).Expand = expand
	return builder
}
func (builder *GetWorkFlowReqBuilder) Build() *GetWorkFlowReq {
	req := &GetWorkFlowReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWorkItemManHourRecordsReq struct {
	apiReq *core.APIReq
}
type GetWorkItemManHourRecordsReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    PageNum  *int64 `json:"page_num,omitempty"`
    PageSize  *int64 `json:"page_size,omitempty"`
}
type GetWorkItemManHourRecordsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ManHourRecord         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type GetWorkItemManHourRecordsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWorkItemManHourRecordsReqBuilder() *GetWorkItemManHourRecordsReqBuilder {
	builder := &GetWorkItemManHourRecordsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &GetWorkItemManHourRecordsReqBody{},
	}
	return builder
}

func (builder *GetWorkItemManHourRecordsReqBuilder) ProjectKey(projectKey string) *GetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemManHourRecordsReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *GetWorkItemManHourRecordsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *GetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemManHourRecordsReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *GetWorkItemManHourRecordsReqBuilder) WorkItemID(workItemID int64) *GetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemManHourRecordsReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *GetWorkItemManHourRecordsReqBuilder) PageNum(pageNum int64) *GetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemManHourRecordsReqBody).PageNum = &pageNum
	return builder
}


func (builder *GetWorkItemManHourRecordsReqBuilder) PageSize(pageSize int64) *GetWorkItemManHourRecordsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemManHourRecordsReqBody).PageSize = &pageSize
	return builder
}

func (builder *GetWorkItemManHourRecordsReqBuilder) Build() *GetWorkItemManHourRecordsReq {
	req := &GetWorkItemManHourRecordsReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWorkItemOpRecordReq struct {
	apiReq *core.APIReq
}
type GetWorkItemOpRecordReqBody struct {
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
type GetWorkItemOpRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *GetWorkItemOpRecordRespData        `json:"data,omitempty"`
}

type GetWorkItemOpRecordRespData struct {
	HasMore       *bool         `json:"has_more,omitempty"`
	StartFrom       *string         `json:"start_from,omitempty"`
	OpRecords       []OAPIOperationRecord         `json:"op_records,omitempty"`
	Total       *int64         `json:"total,omitempty"`
}

type GetWorkItemOpRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWorkItemOpRecordReqBuilder() *GetWorkItemOpRecordReqBuilder {
	builder := &GetWorkItemOpRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &GetWorkItemOpRecordReqBody{},
	}
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) ProjectKey(projectKey string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *GetWorkItemOpRecordReqBuilder) WorkItemIDs(workItemIDs []int64) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).WorkItemIDs = workItemIDs
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) StartFrom(startFrom string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).StartFrom = &startFrom
	return builder
}


func (builder *GetWorkItemOpRecordReqBuilder) Operator(operator []string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).Operator = operator
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) OperatorType(operatorType []string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).OperatorType = operatorType
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) SourceType(sourceType []string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).SourceType = sourceType
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) Source(source []string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).Source = source
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) OperationType(operationType []string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).OperationType = operationType
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) Start(start int64) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).Start = &start
	return builder
}


func (builder *GetWorkItemOpRecordReqBuilder) End(end int64) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).End = &end
	return builder
}


func (builder *GetWorkItemOpRecordReqBuilder) OpRecordModule(opRecordModule []string) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).OpRecordModule = opRecordModule
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) PageSize(pageSize int64) *GetWorkItemOpRecordReqBuilder {
	builder.apiReq.Body.(*GetWorkItemOpRecordReqBody).PageSize = &pageSize
	return builder
}

func (builder *GetWorkItemOpRecordReqBuilder) Build() *GetWorkItemOpRecordReq {
	req := &GetWorkItemOpRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWorkItemTypeInfoByKeyReq struct {
	apiReq *core.APIReq
}
type GetWorkItemTypeInfoByKeyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *WorkItemTypeInfo         `json:"data"`
	
}

type GetWorkItemTypeInfoByKeyReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWorkItemTypeInfoByKeyReqBuilder() *GetWorkItemTypeInfoByKeyReqBuilder {
	builder := &GetWorkItemTypeInfoByKeyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *GetWorkItemTypeInfoByKeyReqBuilder) ProjectKey(projectKey string) *GetWorkItemTypeInfoByKeyReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *GetWorkItemTypeInfoByKeyReqBuilder) WorkItemTypeKey(workItemTypeKey string) *GetWorkItemTypeInfoByKeyReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}

func (builder *GetWorkItemTypeInfoByKeyReqBuilder) Build() *GetWorkItemTypeInfoByKeyReq {
	req := &GetWorkItemTypeInfoByKeyReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetWorkItemsByIdsReq struct {
	apiReq *core.APIReq
}
type GetWorkItemsByIdsReqBody struct {
    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`
    Fields  []string `json:"fields,omitempty"`
    Expand  *Expand `json:"expand,omitempty"`
}
type GetWorkItemsByIdsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
}

type GetWorkItemsByIdsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetWorkItemsByIdsReqBuilder() *GetWorkItemsByIdsReqBuilder {
	builder := &GetWorkItemsByIdsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &GetWorkItemsByIdsReqBody{},
	}
	return builder
}

func (builder *GetWorkItemsByIdsReqBuilder) ProjectKey(projectKey string) *GetWorkItemsByIdsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *GetWorkItemsByIdsReqBuilder) WorkItemIDs(workItemIDs []int64) *GetWorkItemsByIdsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemsByIdsReqBody).WorkItemIDs = workItemIDs
	return builder
}

func (builder *GetWorkItemsByIdsReqBuilder) Fields(fields []string) *GetWorkItemsByIdsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemsByIdsReqBody).Fields = fields
	return builder
}

func (builder *GetWorkItemsByIdsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *GetWorkItemsByIdsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *GetWorkItemsByIdsReqBuilder) Expand(expand *Expand) *GetWorkItemsByIdsReqBuilder {
	builder.apiReq.Body.(*GetWorkItemsByIdsReqBody).Expand = expand
	return builder
}
func (builder *GetWorkItemsByIdsReqBuilder) Build() *GetWorkItemsByIdsReq {
	req := &GetWorkItemsByIdsReq{}
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

func (builder *IntegrateSearchReqBuilder) SearchID(searchID string) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).SearchID = &searchID
	return builder
}


func (builder *IntegrateSearchReqBuilder) ViewID(viewID string) *IntegrateSearchReqBuilder {
	builder.apiReq.Body.(*IntegrateSearchReqBody).ViewID = &viewID
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

type InviteBotJoinChatReq struct {
	apiReq *core.APIReq
}
type InviteBotJoinChatReqBody struct {
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    AppIDs  []string `json:"app_ids,omitempty"`
}
type InviteBotJoinChatResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *BotJoinChatInfo         `json:"data"`
	
}

type InviteBotJoinChatReqBuilder struct {
	apiReq *core.APIReq
}

func NewInviteBotJoinChatReqBuilder() *InviteBotJoinChatReqBuilder {
	builder := &InviteBotJoinChatReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &InviteBotJoinChatReqBody{},
	}
	return builder
}

func (builder *InviteBotJoinChatReqBuilder) ProjectKey(projectKey string) *InviteBotJoinChatReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *InviteBotJoinChatReqBuilder) WorkItemID(workItemID int64) *InviteBotJoinChatReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *InviteBotJoinChatReqBuilder) WorkItemTypeKey(workItemTypeKey string) *InviteBotJoinChatReqBuilder {
	builder.apiReq.Body.(*InviteBotJoinChatReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *InviteBotJoinChatReqBuilder) AppIDs(appIDs []string) *InviteBotJoinChatReqBuilder {
	builder.apiReq.Body.(*InviteBotJoinChatReqBody).AppIDs = appIDs
	return builder
}
func (builder *InviteBotJoinChatReqBuilder) Build() *InviteBotJoinChatReq {
	req := &InviteBotJoinChatReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListTemplateConfReq struct {
	apiReq *core.APIReq
}
type ListTemplateConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []TemplateConf         `json:"data"`
	
}

type ListTemplateConfReqBuilder struct {
	apiReq *core.APIReq
}

func NewListTemplateConfReqBuilder() *ListTemplateConfReqBuilder {
	builder := &ListTemplateConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *ListTemplateConfReqBuilder) ProjectKey(projectKey string) *ListTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *ListTemplateConfReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ListTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}

func (builder *ListTemplateConfReqBuilder) Build() *ListTemplateConfReq {
	req := &ListTemplateConfReq{}
	req.apiReq = builder.apiReq
	return req
}

type PatchWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type PatchWBSViewDraftReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    DraftID  *string `json:"draft_id,omitempty"`
    Operations  []Operation `json:"operations,omitempty"`
}
type PatchWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *PatchWBSViewDraftRespData        `json:"data,omitempty"`
}

type PatchWBSViewDraftRespData struct {
	RelationValues       []Operation         `json:"relation_values,omitempty"`
}

type PatchWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewPatchWBSViewDraftReqBuilder() *PatchWBSViewDraftReqBuilder {
	builder := &PatchWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &PatchWBSViewDraftReqBody{},
	}
	return builder
}

func (builder *PatchWBSViewDraftReqBuilder) ProjectKey(projectKey string) *PatchWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*PatchWBSViewDraftReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *PatchWBSViewDraftReqBuilder) DraftID(draftID string) *PatchWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*PatchWBSViewDraftReqBody).DraftID = &draftID
	return builder
}


func (builder *PatchWBSViewDraftReqBuilder) Operations(operations []Operation) *PatchWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*PatchWBSViewDraftReqBody).Operations = operations
	return builder
}
func (builder *PatchWBSViewDraftReqBuilder) Build() *PatchWBSViewDraftReq {
	req := &PatchWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type PublishWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type PublishWBSViewDraftReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    DraftID  *string `json:"draft_id,omitempty"`
}
type PublishWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *PublishWBSViewDraftRespData        `json:"data,omitempty"`
}

type PublishWBSViewDraftRespData struct {
}

type PublishWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewPublishWBSViewDraftReqBuilder() *PublishWBSViewDraftReqBuilder {
	builder := &PublishWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &PublishWBSViewDraftReqBody{},
	}
	return builder
}

func (builder *PublishWBSViewDraftReqBuilder) ProjectKey(projectKey string) *PublishWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*PublishWBSViewDraftReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *PublishWBSViewDraftReqBuilder) DraftID(draftID string) *PublishWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*PublishWBSViewDraftReqBody).DraftID = &draftID
	return builder
}

func (builder *PublishWBSViewDraftReqBuilder) Build() *PublishWBSViewDraftReq {
	req := &PublishWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryAWorkItemTypesReq struct {
	apiReq *core.APIReq
}
type QueryAWorkItemTypesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemKeyType         `json:"data"`
	
}

type QueryAWorkItemTypesReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryAWorkItemTypesReqBuilder() *QueryAWorkItemTypesReqBuilder {
	builder := &QueryAWorkItemTypesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryAWorkItemTypesReqBuilder) ProjectKey(projectKey string) *QueryAWorkItemTypesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *QueryAWorkItemTypesReqBuilder) Build() *QueryAWorkItemTypesReq {
	req := &QueryAWorkItemTypesReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryBusinessesReq struct {
	apiReq *core.APIReq
}
type QueryBusinessesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []Business         `json:"data"`
	
}

type QueryBusinessesReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryBusinessesReqBuilder() *QueryBusinessesReqBuilder {
	builder := &QueryBusinessesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryBusinessesReqBuilder) ProjectKey(projectKey string) *QueryBusinessesReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *QueryBusinessesReqBuilder) Build() *QueryBusinessesReq {
	req := &QueryBusinessesReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectFieldsReq struct {
	apiReq *core.APIReq
}
type QueryProjectFieldsReqBody struct {
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
}
type QueryProjectFieldsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []SimpleField         `json:"data"`
	
}

type QueryProjectFieldsReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryProjectFieldsReqBuilder() *QueryProjectFieldsReqBuilder {
	builder := &QueryProjectFieldsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryProjectFieldsReqBody{},
	}
	return builder
}

func (builder *QueryProjectFieldsReqBuilder) ProjectKey(projectKey string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryProjectFieldsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryProjectFieldsReqBuilder {
	builder.apiReq.Body.(*QueryProjectFieldsReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}

func (builder *QueryProjectFieldsReqBuilder) Build() *QueryProjectFieldsReq {
	req := &QueryProjectFieldsReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectRelationReq struct {
	apiReq *core.APIReq
}
type QueryProjectRelationReqBody struct {
    RemoteProjects  []string `json:"remote_projects,omitempty"`
}
type QueryProjectRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ProjectRelationRule         `json:"data"`
	
}

type QueryProjectRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryProjectRelationReqBuilder() *QueryProjectRelationReqBuilder {
	builder := &QueryProjectRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryProjectRelationReqBody{},
	}
	return builder
}

func (builder *QueryProjectRelationReqBuilder) ProjectKey(projectKey string) *QueryProjectRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryProjectRelationReqBuilder) RemoteProjects(remoteProjects []string) *QueryProjectRelationReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationReqBody).RemoteProjects = remoteProjects
	return builder
}
func (builder *QueryProjectRelationReqBuilder) Build() *QueryProjectRelationReq {
	req := &QueryProjectRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectRelationInstanceReq struct {
	apiReq *core.APIReq
}
type QueryProjectRelationInstanceReqBody struct {
    RelationRuleID  *string `json:"relation_rule_id,omitempty"`
    RelationWorkItemID  *int64 `json:"relation_work_item_id,omitempty"`
    RelationWorkItemTypeKey  *string `json:"relation_work_item_type_key,omitempty"`
    RelationProjectKey  *string `json:"relation_project_key,omitempty"`
}
type QueryProjectRelationInstanceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []RelationInstance         `json:"data"`
	
}

type QueryProjectRelationInstanceReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryProjectRelationInstanceReqBuilder() *QueryProjectRelationInstanceReqBuilder {
	builder := &QueryProjectRelationInstanceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryProjectRelationInstanceReqBody{},
	}
	return builder
}

func (builder *QueryProjectRelationInstanceReqBuilder) ProjectKey(projectKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryProjectRelationInstanceReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *QueryProjectRelationInstanceReqBuilder) WorkItemID(workItemID int64) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *QueryProjectRelationInstanceReqBuilder) RelationRuleID(relationRuleID string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationRuleID = &relationRuleID
	return builder
}


func (builder *QueryProjectRelationInstanceReqBuilder) RelationWorkItemID(relationWorkItemID int64) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationWorkItemID = &relationWorkItemID
	return builder
}


func (builder *QueryProjectRelationInstanceReqBuilder) RelationWorkItemTypeKey(relationWorkItemTypeKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationWorkItemTypeKey = &relationWorkItemTypeKey
	return builder
}


func (builder *QueryProjectRelationInstanceReqBuilder) RelationProjectKey(relationProjectKey string) *QueryProjectRelationInstanceReqBuilder {
	builder.apiReq.Body.(*QueryProjectRelationInstanceReqBody).RelationProjectKey = &relationProjectKey
	return builder
}

func (builder *QueryProjectRelationInstanceReqBuilder) Build() *QueryProjectRelationInstanceReq {
	req := &QueryProjectRelationInstanceReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryRoleConfDetailsReq struct {
	apiReq *core.APIReq
}
type QueryRoleConfDetailsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []RoleConfDetail         `json:"data"`
	
}

type QueryRoleConfDetailsReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryRoleConfDetailsReqBuilder() *QueryRoleConfDetailsReqBuilder {
	builder := &QueryRoleConfDetailsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryRoleConfDetailsReqBuilder) ProjectKey(projectKey string) *QueryRoleConfDetailsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryRoleConfDetailsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryRoleConfDetailsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}

func (builder *QueryRoleConfDetailsReqBuilder) Build() *QueryRoleConfDetailsReq {
	req := &QueryRoleConfDetailsReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryStoryRelationsReq struct {
	apiReq *core.APIReq
}
type QueryStoryRelationsReqBody struct {
    WorkItemIDs  []int64 `json:"work_item_ids,omitempty"`
}
type QueryStoryRelationsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *QueryStoryRelationData         `json:"data"`
	
}

type QueryStoryRelationsReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryStoryRelationsReqBuilder() *QueryStoryRelationsReqBuilder {
	builder := &QueryStoryRelationsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryStoryRelationsReqBody{},
	}
	return builder
}

func (builder *QueryStoryRelationsReqBuilder) WorkItemIDs(workItemIDs []int64) *QueryStoryRelationsReqBuilder {
	builder.apiReq.Body.(*QueryStoryRelationsReqBody).WorkItemIDs = workItemIDs
	return builder
}

func (builder *QueryStoryRelationsReqBuilder) ProjectKey(projectKey string) *QueryStoryRelationsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *QueryStoryRelationsReqBuilder) Build() *QueryStoryRelationsReq {
	req := &QueryStoryRelationsReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryTaskResultReq struct {
	apiReq *core.APIReq
}
type QueryTaskResultResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *QueryTaskResultRespData        `json:"data,omitempty"`
}

type QueryTaskResultRespData struct {
	TaskID       *string         `json:"task_id,omitempty"`
	TaskStatus       *string         `json:"task_status,omitempty"`
	Total       *int64         `json:"total,omitempty"`
	SuccessTotal       *int64         `json:"success_total,omitempty"`
	ErrorTotal       *int64         `json:"error_total,omitempty"`
	SuccessSubTaskList       []string         `json:"success_sub_task_list,omitempty"`
	FailSubTaskList       []string         `json:"fail_sub_task_list,omitempty"`
	ErrorScenes       []string         `json:"error_scenes,omitempty"`
}

type QueryTaskResultReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryTaskResultReqBuilder() *QueryTaskResultReqBuilder {
	builder := &QueryTaskResultReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryTaskResultReqBuilder) TaskID(taskID string) *QueryTaskResultReqBuilder {
	builder.apiReq.QueryParams.Set("task_id", fmt.Sprint(taskID))
	return builder
}

func (builder *QueryTaskResultReqBuilder) Build() *QueryTaskResultReq {
	req := &QueryTaskResultReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryTemplateDetailReq struct {
	apiReq *core.APIReq
}
type QueryTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *TemplateDetail         `json:"data"`
	
}

type QueryTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryTemplateDetailReqBuilder() *QueryTemplateDetailReqBuilder {
	builder := &QueryTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryTemplateDetailReqBuilder) ProjectKey(projectKey string) *QueryTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryTemplateDetailReqBuilder) TemplateID(templateID int64) *QueryTemplateDetailReqBuilder {
	builder.apiReq.PathParams.Set("template_id", fmt.Sprint(templateID))
	return builder
}

func (builder *QueryTemplateDetailReqBuilder) Build() *QueryTemplateDetailReq {
	req := &QueryTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type QueryWBSViewDraftReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    NeedInit  *bool `json:"need_init,omitempty"`
}
type QueryWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *QueryWBSViewDraftRespData        `json:"data,omitempty"`
}

type QueryWBSViewDraftRespData struct {
	WbsDraft       *WbsDraft         `json:"wbs_draft,omitempty"`
}

type QueryWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWBSViewDraftReqBuilder() *QueryWBSViewDraftReqBuilder {
	builder := &QueryWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryWBSViewDraftReqBody{},
	}
	return builder
}

func (builder *QueryWBSViewDraftReqBuilder) ProjectKey(projectKey string) *QueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*QueryWBSViewDraftReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *QueryWBSViewDraftReqBuilder) WorkItemID(workItemID int64) *QueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*QueryWBSViewDraftReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *QueryWBSViewDraftReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*QueryWBSViewDraftReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *QueryWBSViewDraftReqBuilder) NeedInit(needInit bool) *QueryWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*QueryWBSViewDraftReqBody).NeedInit = &needInit
	return builder
}

func (builder *QueryWBSViewDraftReqBuilder) Build() *QueryWBSViewDraftReq {
	req := &QueryWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWbsTemplateConfReq struct {
	apiReq *core.APIReq
}
type QueryWbsTemplateConfReqBody struct {
    TemplateKey  *string `json:"template_key,omitempty"`
}
type QueryWbsTemplateConfResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *WbsTemplate         `json:"data"`
	
}

type QueryWbsTemplateConfReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWbsTemplateConfReqBuilder() *QueryWbsTemplateConfReqBuilder {
	builder := &QueryWbsTemplateConfReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryWbsTemplateConfReqBody{},
	}
	return builder
}

func (builder *QueryWbsTemplateConfReqBuilder) ProjectKey(projectKey string) *QueryWbsTemplateConfReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryWbsTemplateConfReqBuilder) TemplateKey(templateKey string) *QueryWbsTemplateConfReqBuilder {
	builder.apiReq.Body.(*QueryWbsTemplateConfReqBody).TemplateKey = &templateKey
	return builder
}

func (builder *QueryWbsTemplateConfReqBuilder) Build() *QueryWbsTemplateConfReq {
	req := &QueryWbsTemplateConfReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkItemMetaDataReq struct {
	apiReq *core.APIReq
}
type QueryWorkItemMetaDataResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []FieldConf         `json:"data"`
	
}

type QueryWorkItemMetaDataReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWorkItemMetaDataReqBuilder() *QueryWorkItemMetaDataReqBuilder {
	builder := &QueryWorkItemMetaDataReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryWorkItemMetaDataReqBuilder) ProjectKey(projectKey string) *QueryWorkItemMetaDataReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryWorkItemMetaDataReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryWorkItemMetaDataReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}

func (builder *QueryWorkItemMetaDataReqBuilder) Build() *QueryWorkItemMetaDataReq {
	req := &QueryWorkItemMetaDataReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type QueryWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemRelation         `json:"data"`
	
}

type QueryWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWorkItemRelationReqBuilder() *QueryWorkItemRelationReqBuilder {
	builder := &QueryWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryWorkItemRelationReqBuilder) ProjectKey(projectKey string) *QueryWorkItemRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *QueryWorkItemRelationReqBuilder) Build() *QueryWorkItemRelationReq {
	req := &QueryWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}
type QueryWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []NodeTask         `json:"data"`
	
}

type QueryWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWorkItemSubTaskReqBuilder() *QueryWorkItemSubTaskReqBuilder {
	builder := &QueryWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *QueryWorkItemSubTaskReqBuilder) ProjectKey(projectKey string) *QueryWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryWorkItemSubTaskReqBuilder) WorkItemID(workItemID int64) *QueryWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *QueryWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *QueryWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *QueryWorkItemSubTaskReqBuilder) NodeID(nodeID string) *QueryWorkItemSubTaskReqBuilder {
	builder.apiReq.QueryParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}

func (builder *QueryWorkItemSubTaskReqBuilder) Build() *QueryWorkItemSubTaskReq {
	req := &QueryWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type ResetWBSViewDraftReq struct {
	apiReq *core.APIReq
}
type ResetWBSViewDraftReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
}
type ResetWBSViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *ResetWBSViewDraftRespData        `json:"data,omitempty"`
}

type ResetWBSViewDraftRespData struct {
	WbsDraft       *WbsDraft         `json:"wbs_draft,omitempty"`
}

type ResetWBSViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewResetWBSViewDraftReqBuilder() *ResetWBSViewDraftReqBuilder {
	builder := &ResetWBSViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ResetWBSViewDraftReqBody{},
	}
	return builder
}

func (builder *ResetWBSViewDraftReqBuilder) ProjectKey(projectKey string) *ResetWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*ResetWBSViewDraftReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *ResetWBSViewDraftReqBuilder) WorkItemID(workItemID int64) *ResetWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*ResetWBSViewDraftReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *ResetWBSViewDraftReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ResetWBSViewDraftReqBuilder {
	builder.apiReq.Body.(*ResetWBSViewDraftReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}

func (builder *ResetWBSViewDraftReqBuilder) Build() *ResetWBSViewDraftReq {
	req := &ResetWBSViewDraftReq{}
	req.apiReq = builder.apiReq
	return req
}

type ResourceCreateWorkItemReq struct {
	apiReq *core.APIReq
}
type ResourceCreateWorkItemReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    TemplateID  *int64 `json:"template_id,omitempty"`
    FieldValuePairs  []FieldValuePair `json:"field_value_pairs,omitempty"`
}
type ResourceCreateWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *OAPICreateWorkItemInfo         `json:"data"`
	
}

type ResourceCreateWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewResourceCreateWorkItemReqBuilder() *ResourceCreateWorkItemReqBuilder {
	builder := &ResourceCreateWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ResourceCreateWorkItemReqBody{},
	}
	return builder
}

func (builder *ResourceCreateWorkItemReqBuilder) ProjectKey(projectKey string) *ResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*ResourceCreateWorkItemReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *ResourceCreateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*ResourceCreateWorkItemReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *ResourceCreateWorkItemReqBuilder) TemplateID(templateID int64) *ResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*ResourceCreateWorkItemReqBody).TemplateID = &templateID
	return builder
}


func (builder *ResourceCreateWorkItemReqBuilder) FieldValuePairs(fieldValuePairs []FieldValuePair) *ResourceCreateWorkItemReqBuilder {
	builder.apiReq.Body.(*ResourceCreateWorkItemReqBody).FieldValuePairs = fieldValuePairs
	return builder
}
func (builder *ResourceCreateWorkItemReqBuilder) Build() *ResourceCreateWorkItemReq {
	req := &ResourceCreateWorkItemReq{}
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

func (builder *SearchByParamsReqBuilder) ProjectKey(projectKey string) *SearchByParamsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *SearchByParamsReqBuilder) SearchGroup(searchGroup *SearchGroup) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).SearchGroup = searchGroup
	return builder
}

func (builder *SearchByParamsReqBuilder) WorkItemTypeKey(workItemTypeKey string) *SearchByParamsReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *SearchByParamsReqBuilder) PageNum(pageNum int64) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).PageNum = &pageNum
	return builder
}


func (builder *SearchByParamsReqBuilder) PageSize(pageSize int64) *SearchByParamsReqBuilder {
	builder.apiReq.Body.(*SearchByParamsReqBody).PageSize = &pageSize
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

type SearchSubtaskReq struct {
	apiReq *core.APIReq
}
type SearchSubtaskReqBody struct {
    PageSize  *int64 `json:"page_size,omitempty"`
    PageNum  *int64 `json:"page_num,omitempty"`
    Name  *string `json:"name,omitempty"`
    UserKeys  []string `json:"user_keys,omitempty"`
    Status  *int32 `json:"status,omitempty"`
    CreatedAt  *TimeInterval `json:"created_at,omitempty"`
    SimpleNames  []string `json:"simple_names,omitempty"`
}
type SearchSubtaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Pagination       *Pagination         `json:"pagination"`
	
	Data       []SubDetail         `json:"data"`
	
}

type SearchSubtaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewSearchSubtaskReqBuilder() *SearchSubtaskReqBuilder {
	builder := &SearchSubtaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SearchSubtaskReqBody{},
	}
	return builder
}

func (builder *SearchSubtaskReqBuilder) PageSize(pageSize int64) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).PageSize = &pageSize
	return builder
}


func (builder *SearchSubtaskReqBuilder) PageNum(pageNum int64) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).PageNum = &pageNum
	return builder
}


func (builder *SearchSubtaskReqBuilder) Name(name string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).Name = &name
	return builder
}


func (builder *SearchSubtaskReqBuilder) UserKeys(userKeys []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).UserKeys = userKeys
	return builder
}

func (builder *SearchSubtaskReqBuilder) Status(status int32) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).Status = &status
	return builder
}


func (builder *SearchSubtaskReqBuilder) CreatedAt(createdAt *TimeInterval) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).CreatedAt = createdAt
	return builder
}

func (builder *SearchSubtaskReqBuilder) SimpleNames(simpleNames []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *SearchSubtaskReqBuilder) Build() *SearchSubtaskReq {
	req := &SearchSubtaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type SearchWorkItemsRelationReq struct {
	apiReq *core.APIReq
}
type SearchWorkItemsRelationReqBody struct {
    RelationWorkItemTypeKey  *string `json:"relation_work_item_type_key,omitempty"`
    RelationKey  *string `json:"relation_key,omitempty"`
    PageNum  *int64 `json:"page_num,omitempty"`
    PageSize  *int64 `json:"page_size,omitempty"`
    RelationType  *int32 `json:"relation_type,omitempty"`
    Expand  *Expand `json:"expand,omitempty"`
}
type SearchWorkItemsRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type SearchWorkItemsRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewSearchWorkItemsRelationReqBuilder() *SearchWorkItemsRelationReqBuilder {
	builder := &SearchWorkItemsRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SearchWorkItemsRelationReqBody{},
	}
	return builder
}

func (builder *SearchWorkItemsRelationReqBuilder) ProjectKey(projectKey string) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) WorkItemTypeKey(workItemTypeKey string) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) WorkItemID(workItemID int64) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) RelationWorkItemTypeKey(relationWorkItemTypeKey string) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*SearchWorkItemsRelationReqBody).RelationWorkItemTypeKey = &relationWorkItemTypeKey
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) RelationKey(relationKey string) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*SearchWorkItemsRelationReqBody).RelationKey = &relationKey
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) PageNum(pageNum int64) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*SearchWorkItemsRelationReqBody).PageNum = &pageNum
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) PageSize(pageSize int64) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*SearchWorkItemsRelationReqBody).PageSize = &pageSize
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) RelationType(relationType int32) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*SearchWorkItemsRelationReqBody).RelationType = &relationType
	return builder
}


func (builder *SearchWorkItemsRelationReqBuilder) Expand(expand *Expand) *SearchWorkItemsRelationReqBuilder {
	builder.apiReq.Body.(*SearchWorkItemsRelationReqBody).Expand = expand
	return builder
}
func (builder *SearchWorkItemsRelationReqBuilder) Build() *SearchWorkItemsRelationReq {
	req := &SearchWorkItemsRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type SubTaskModifyReq struct {
	apiReq *core.APIReq
}
type SubTaskModifyReqBody struct {
    NodeID  *string `json:"node_id,omitempty"`
    TaskID  *int64 `json:"task_id,omitempty"`
    Action  *string `json:"action,omitempty"`
    Assignee  []string `json:"assignee,omitempty"`
    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`
    Schedules  []Schedule `json:"schedules,omitempty"`
    Deliverable  []FieldValuePair `json:"deliverable,omitempty"`
    Note  *string `json:"note,omitempty"`
}
type SubTaskModifyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *SubTaskModifyRespData        `json:"data,omitempty"`
}

type SubTaskModifyRespData struct {
}

type SubTaskModifyReqBuilder struct {
	apiReq *core.APIReq
}

func NewSubTaskModifyReqBuilder() *SubTaskModifyReqBuilder {
	builder := &SubTaskModifyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SubTaskModifyReqBody{},
	}
	return builder
}

func (builder *SubTaskModifyReqBuilder) NodeID(nodeID string) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).NodeID = &nodeID
	return builder
}


func (builder *SubTaskModifyReqBuilder) TaskID(taskID int64) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).TaskID = &taskID
	return builder
}


func (builder *SubTaskModifyReqBuilder) Action(action string) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).Action = &action
	return builder
}


func (builder *SubTaskModifyReqBuilder) Assignee(assignee []string) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).Assignee = assignee
	return builder
}

func (builder *SubTaskModifyReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).RoleAssignee = roleAssignee
	return builder
}

func (builder *SubTaskModifyReqBuilder) Schedules(schedules []Schedule) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).Schedules = schedules
	return builder
}

func (builder *SubTaskModifyReqBuilder) Deliverable(deliverable []FieldValuePair) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).Deliverable = deliverable
	return builder
}

func (builder *SubTaskModifyReqBuilder) Note(note string) *SubTaskModifyReqBuilder {
	builder.apiReq.Body.(*SubTaskModifyReqBody).Note = &note
	return builder
}


func (builder *SubTaskModifyReqBuilder) ProjectKey(projectKey string) *SubTaskModifyReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *SubTaskModifyReqBuilder) WorkItemTypeKey(workItemTypeKey string) *SubTaskModifyReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *SubTaskModifyReqBuilder) WorkItemID(workItemID int64) *SubTaskModifyReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}

func (builder *SubTaskModifyReqBuilder) Build() *SubTaskModifyReq {
	req := &SubTaskModifyReq{}
	req.apiReq = builder.apiReq
	return req
}

type SwitchBackToWbsViewDraftReq struct {
	apiReq *core.APIReq
}
type SwitchBackToWbsViewDraftReqBody struct {
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    DraftID  *string `json:"draft_id,omitempty"`
    CommitID  *string `json:"commit_id,omitempty"`
}
type SwitchBackToWbsViewDraftResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *SwitchBackToWbsViewDraftRespData        `json:"data,omitempty"`
}

type SwitchBackToWbsViewDraftRespData struct {
	Success       *bool         `json:"success,omitempty"`
}

type SwitchBackToWbsViewDraftReqBuilder struct {
	apiReq *core.APIReq
}

func NewSwitchBackToWbsViewDraftReqBuilder() *SwitchBackToWbsViewDraftReqBuilder {
	builder := &SwitchBackToWbsViewDraftReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SwitchBackToWbsViewDraftReqBody{},
	}
	return builder
}

func (builder *SwitchBackToWbsViewDraftReqBuilder) ProjectKey(projectKey string) *SwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *SwitchBackToWbsViewDraftReqBuilder) WorkItemID(workItemID int64) *SwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.Body.(*SwitchBackToWbsViewDraftReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *SwitchBackToWbsViewDraftReqBuilder) DraftID(draftID string) *SwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.Body.(*SwitchBackToWbsViewDraftReqBody).DraftID = &draftID
	return builder
}


func (builder *SwitchBackToWbsViewDraftReqBuilder) CommitID(commitID string) *SwitchBackToWbsViewDraftReqBuilder {
	builder.apiReq.Body.(*SwitchBackToWbsViewDraftReqBody).CommitID = &commitID
	return builder
}

func (builder *SwitchBackToWbsViewDraftReqBuilder) Build() *SwitchBackToWbsViewDraftReq {
	req := &SwitchBackToWbsViewDraftReq{}
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

func (builder *UniversalSearchReqBuilder) UserKey(userKey string) *UniversalSearchReqBuilder {
	builder.apiReq.Body.(*UniversalSearchReqBody).UserKey = &userKey
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

type UpdateFieldReq struct {
	apiReq *core.APIReq
}
type UpdateFieldReqBody struct {
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
    TeamOption  *TeamOption `json:"team_option,omitempty"`
}
type UpdateFieldResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateFieldRespData        `json:"data,omitempty"`
}

type UpdateFieldRespData struct {
}

type UpdateFieldReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateFieldReqBuilder() *UpdateFieldReqBuilder {
	builder := &UpdateFieldReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateFieldReqBody{},
	}
	return builder
}

func (builder *UpdateFieldReqBuilder) ProjectKey(projectKey string) *UpdateFieldReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateFieldReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateFieldReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateFieldReqBuilder) FieldName(fieldName string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldName = &fieldName
	return builder
}


func (builder *UpdateFieldReqBuilder) FieldKey(fieldKey string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldKey = &fieldKey
	return builder
}


func (builder *UpdateFieldReqBuilder) FieldValue(fieldValue interface{}) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldValue = fieldValue
	return builder
}

func (builder *UpdateFieldReqBuilder) FreeAdd(freeAdd int64) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FreeAdd = &freeAdd
	return builder
}


func (builder *UpdateFieldReqBuilder) WorkItemRelationUUID(workItemRelationUUID string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).WorkItemRelationUUID = &workItemRelationUUID
	return builder
}


func (builder *UpdateFieldReqBuilder) DefaultValue(defaultValue interface{}) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).DefaultValue = defaultValue
	return builder
}

func (builder *UpdateFieldReqBuilder) FieldAlias(fieldAlias string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).FieldAlias = &fieldAlias
	return builder
}


func (builder *UpdateFieldReqBuilder) HelpDescription(helpDescription string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).HelpDescription = &helpDescription
	return builder
}


func (builder *UpdateFieldReqBuilder) AuthorizedRoles(authorizedRoles []string) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).AuthorizedRoles = authorizedRoles
	return builder
}

func (builder *UpdateFieldReqBuilder) RelatedFieldExtraDisplayInfos(relatedFieldExtraDisplayInfos []RelatedFieldExtraDisplayInfo) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).RelatedFieldExtraDisplayInfos = relatedFieldExtraDisplayInfos
	return builder
}

func (builder *UpdateFieldReqBuilder) TeamOption(teamOption *TeamOption) *UpdateFieldReqBuilder {
	builder.apiReq.Body.(*UpdateFieldReqBody).TeamOption = teamOption
	return builder
}
func (builder *UpdateFieldReqBuilder) Build() *UpdateFieldReq {
	req := &UpdateFieldReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateFinishedReq struct {
	apiReq *core.APIReq
}
type UpdateFinishedReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    NodeID  *string `json:"node_id,omitempty"`
    Opinion  *string `json:"opinion,omitempty"`
    FinishedConclusionOptionKey  *string `json:"finished_conclusion_option_key,omitempty"`
    OperationType  *string `json:"operation_type,omitempty"`
    Reset  *bool `json:"reset,omitempty"`
}
type UpdateFinishedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateFinishedRespData        `json:"data,omitempty"`
}

type UpdateFinishedRespData struct {
}

type UpdateFinishedReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateFinishedReqBuilder() *UpdateFinishedReqBuilder {
	builder := &UpdateFinishedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateFinishedReqBody{},
	}
	return builder
}

func (builder *UpdateFinishedReqBuilder) ProjectKey(projectKey string) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *UpdateFinishedReqBuilder) WorkItemID(workItemID int64) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *UpdateFinishedReqBuilder) NodeID(nodeID string) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).NodeID = &nodeID
	return builder
}


func (builder *UpdateFinishedReqBuilder) Opinion(opinion string) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).Opinion = &opinion
	return builder
}


func (builder *UpdateFinishedReqBuilder) FinishedConclusionOptionKey(finishedConclusionOptionKey string) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).FinishedConclusionOptionKey = &finishedConclusionOptionKey
	return builder
}


func (builder *UpdateFinishedReqBuilder) OperationType(operationType string) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).OperationType = &operationType
	return builder
}


func (builder *UpdateFinishedReqBuilder) Reset(reset bool) *UpdateFinishedReqBuilder {
	builder.apiReq.Body.(*UpdateFinishedReqBody).Reset = &reset
	return builder
}

func (builder *UpdateFinishedReqBuilder) Build() *UpdateFinishedReq {
	req := &UpdateFinishedReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateMultiSignalReq struct {
	apiReq *core.APIReq
}
type UpdateMultiSignalReqBody struct {
    FieldKey  *string `json:"field_key,omitempty"`
    FieldAlias  *string `json:"field_alias,omitempty"`
    Details  []MultiSignalDetail `json:"details,omitempty"`
    UpdateType  *string `json:"update_type,omitempty"`
}
type UpdateMultiSignalResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *MultiSignal         `json:"data"`
	
}

type UpdateMultiSignalReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateMultiSignalReqBuilder() *UpdateMultiSignalReqBuilder {
	builder := &UpdateMultiSignalReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateMultiSignalReqBody{},
	}
	return builder
}

func (builder *UpdateMultiSignalReqBuilder) ProjectKey(projectKey string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateMultiSignalReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateMultiSignalReqBuilder) WorkItemID(workItemID int64) *UpdateMultiSignalReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateMultiSignalReqBuilder) FieldKey(fieldKey string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*UpdateMultiSignalReqBody).FieldKey = &fieldKey
	return builder
}


func (builder *UpdateMultiSignalReqBuilder) FieldAlias(fieldAlias string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*UpdateMultiSignalReqBody).FieldAlias = &fieldAlias
	return builder
}


func (builder *UpdateMultiSignalReqBuilder) Details(details []MultiSignalDetail) *UpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*UpdateMultiSignalReqBody).Details = details
	return builder
}

func (builder *UpdateMultiSignalReqBuilder) UpdateType(updateType string) *UpdateMultiSignalReqBuilder {
	builder.apiReq.Body.(*UpdateMultiSignalReqBody).UpdateType = &updateType
	return builder
}

func (builder *UpdateMultiSignalReqBuilder) Build() *UpdateMultiSignalReq {
	req := &UpdateMultiSignalReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateNodeStateReq struct {
	apiReq *core.APIReq
}
type UpdateNodeStateReqBody struct {
    Action  *string `json:"action,omitempty"`
    RollbackReason  *string `json:"rollback_reason,omitempty"`
    NodeOwners  []string `json:"node_owners,omitempty"`
    NodeSchedule  *Schedule `json:"node_schedule,omitempty"`
    Schedules  []Schedule `json:"schedules,omitempty"`
    Fields  []FieldValuePair `json:"fields,omitempty"`
    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`
}
type UpdateNodeStateResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateNodeStateRespData        `json:"data,omitempty"`
}

type UpdateNodeStateRespData struct {
}

type UpdateNodeStateReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateNodeStateReqBuilder() *UpdateNodeStateReqBuilder {
	builder := &UpdateNodeStateReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateNodeStateReqBody{},
	}
	return builder
}

func (builder *UpdateNodeStateReqBuilder) ProjectKey(projectKey string) *UpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateNodeStateReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateNodeStateReqBuilder) WorkItemID(workItemID int64) *UpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateNodeStateReqBuilder) NodeID(nodeID string) *UpdateNodeStateReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}


func (builder *UpdateNodeStateReqBuilder) Action(action string) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).Action = &action
	return builder
}


func (builder *UpdateNodeStateReqBuilder) RollbackReason(rollbackReason string) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).RollbackReason = &rollbackReason
	return builder
}


func (builder *UpdateNodeStateReqBuilder) NodeOwners(nodeOwners []string) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).NodeOwners = nodeOwners
	return builder
}

func (builder *UpdateNodeStateReqBuilder) NodeSchedule(nodeSchedule *Schedule) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).NodeSchedule = nodeSchedule
	return builder
}

func (builder *UpdateNodeStateReqBuilder) Schedules(schedules []Schedule) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).Schedules = schedules
	return builder
}

func (builder *UpdateNodeStateReqBuilder) Fields(fields []FieldValuePair) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).Fields = fields
	return builder
}

func (builder *UpdateNodeStateReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *UpdateNodeStateReqBuilder {
	builder.apiReq.Body.(*UpdateNodeStateReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *UpdateNodeStateReqBuilder) Build() *UpdateNodeStateReq {
	req := &UpdateNodeStateReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateStateFlowReq struct {
	apiReq *core.APIReq
}
type UpdateStateFlowReqBody struct {
    TransitionID  *int64 `json:"transition_id,omitempty"`
    Fields  []FieldValuePair `json:"fields,omitempty"`
    RoleOwners  []RoleOwner `json:"role_owners,omitempty"`
}
type UpdateStateFlowResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateStateFlowRespData        `json:"data,omitempty"`
}

type UpdateStateFlowRespData struct {
}

type UpdateStateFlowReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateStateFlowReqBuilder() *UpdateStateFlowReqBuilder {
	builder := &UpdateStateFlowReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateStateFlowReqBody{},
	}
	return builder
}

func (builder *UpdateStateFlowReqBuilder) ProjectKey(projectKey string) *UpdateStateFlowReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateStateFlowReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateStateFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateStateFlowReqBuilder) WorkItemID(workItemID int64) *UpdateStateFlowReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateStateFlowReqBuilder) TransitionID(transitionID int64) *UpdateStateFlowReqBuilder {
	builder.apiReq.Body.(*UpdateStateFlowReqBody).TransitionID = &transitionID
	return builder
}


func (builder *UpdateStateFlowReqBuilder) Fields(fields []FieldValuePair) *UpdateStateFlowReqBuilder {
	builder.apiReq.Body.(*UpdateStateFlowReqBody).Fields = fields
	return builder
}

func (builder *UpdateStateFlowReqBuilder) RoleOwners(roleOwners []RoleOwner) *UpdateStateFlowReqBuilder {
	builder.apiReq.Body.(*UpdateStateFlowReqBody).RoleOwners = roleOwners
	return builder
}
func (builder *UpdateStateFlowReqBuilder) Build() *UpdateStateFlowReq {
	req := &UpdateStateFlowReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateTemplateDetailReq struct {
	apiReq *core.APIReq
}
type UpdateTemplateDetailReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    TemplateID  *int64 `json:"template_id,omitempty"`
    WorkflowConfs  []WorkflowConfInfo `json:"workflow_confs,omitempty"`
    StateFlowConfs  []StateFlowConfInfo `json:"state_flow_confs,omitempty"`
}
type UpdateTemplateDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateTemplateDetailRespData        `json:"data,omitempty"`
}

type UpdateTemplateDetailRespData struct {
}

type UpdateTemplateDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateTemplateDetailReqBuilder() *UpdateTemplateDetailReqBuilder {
	builder := &UpdateTemplateDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateTemplateDetailReqBody{},
	}
	return builder
}

func (builder *UpdateTemplateDetailReqBuilder) ProjectKey(projectKey string) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *UpdateTemplateDetailReqBuilder) TemplateID(templateID int64) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).TemplateID = &templateID
	return builder
}


func (builder *UpdateTemplateDetailReqBuilder) WorkflowConfs(workflowConfs []WorkflowConfInfo) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).WorkflowConfs = workflowConfs
	return builder
}

func (builder *UpdateTemplateDetailReqBuilder) StateFlowConfs(stateFlowConfs []StateFlowConfInfo) *UpdateTemplateDetailReqBuilder {
	builder.apiReq.Body.(*UpdateTemplateDetailReqBody).StateFlowConfs = stateFlowConfs
	return builder
}
func (builder *UpdateTemplateDetailReqBuilder) Build() *UpdateTemplateDetailReq {
	req := &UpdateTemplateDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateWorkItemReq struct {
	apiReq *core.APIReq
}
type UpdateWorkItemReqBody struct {
    UpdateFields  []FieldValuePair `json:"update_fields,omitempty"`
}
type UpdateWorkItemResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateWorkItemRespData        `json:"data,omitempty"`
}

type UpdateWorkItemRespData struct {
}

type UpdateWorkItemReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateWorkItemReqBuilder() *UpdateWorkItemReqBuilder {
	builder := &UpdateWorkItemReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateWorkItemReqBody{},
	}
	return builder
}

func (builder *UpdateWorkItemReqBuilder) ProjectKey(projectKey string) *UpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateWorkItemReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateWorkItemReqBuilder) UpdateFields(updateFields []FieldValuePair) *UpdateWorkItemReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemReqBody).UpdateFields = updateFields
	return builder
}

func (builder *UpdateWorkItemReqBuilder) WorkItemID(workItemID int64) *UpdateWorkItemReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}

func (builder *UpdateWorkItemReqBuilder) Build() *UpdateWorkItemReq {
	req := &UpdateWorkItemReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateWorkItemRelationReq struct {
	apiReq *core.APIReq
}
type UpdateWorkItemRelationReqBody struct {
    RelationID  *string `json:"relation_id,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`
    Name  *string `json:"name,omitempty"`
    RelationDetails  []RelationDetail `json:"relation_details,omitempty"`
}
type UpdateWorkItemRelationResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateWorkItemRelationRespData        `json:"data,omitempty"`
}

type UpdateWorkItemRelationRespData struct {
}

type UpdateWorkItemRelationReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateWorkItemRelationReqBuilder() *UpdateWorkItemRelationReqBuilder {
	builder := &UpdateWorkItemRelationReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateWorkItemRelationReqBody{},
	}
	return builder
}

func (builder *UpdateWorkItemRelationReqBuilder) RelationID(relationID string) *UpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemRelationReqBody).RelationID = &relationID
	return builder
}


func (builder *UpdateWorkItemRelationReqBuilder) ProjectKey(projectKey string) *UpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemRelationReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *UpdateWorkItemRelationReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemRelationReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *UpdateWorkItemRelationReqBuilder) Name(name string) *UpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemRelationReqBody).Name = &name
	return builder
}


func (builder *UpdateWorkItemRelationReqBuilder) RelationDetails(relationDetails []RelationDetail) *UpdateWorkItemRelationReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemRelationReqBody).RelationDetails = relationDetails
	return builder
}
func (builder *UpdateWorkItemRelationReqBuilder) Build() *UpdateWorkItemRelationReq {
	req := &UpdateWorkItemRelationReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateWorkItemSubTaskReq struct {
	apiReq *core.APIReq
}
type UpdateWorkItemSubTaskReqBody struct {
    Name  *string `json:"name,omitempty"`
    Assignee  []string `json:"assignee,omitempty"`
    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`
    Schedule  *Schedule `json:"schedule,omitempty"`
    Note  *string `json:"note,omitempty"`
    Deliverable  []FieldValuePair `json:"deliverable,omitempty"`
    UpdateFields  []FieldValuePair `json:"update_fields,omitempty"`
}
type UpdateWorkItemSubTaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateWorkItemSubTaskRespData        `json:"data,omitempty"`
}

type UpdateWorkItemSubTaskRespData struct {
}

type UpdateWorkItemSubTaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateWorkItemSubTaskReqBuilder() *UpdateWorkItemSubTaskReqBuilder {
	builder := &UpdateWorkItemSubTaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateWorkItemSubTaskReqBody{},
	}
	return builder
}

func (builder *UpdateWorkItemSubTaskReqBuilder) ProjectKey(projectKey string) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) WorkItemID(workItemID int64) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) NodeID(nodeID string) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) TaskID(taskID int64) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.PathParams.Set("task_id", fmt.Sprint(taskID))
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) Name(name string) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).Name = &name
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) Assignee(assignee []string) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).Assignee = assignee
	return builder
}

func (builder *UpdateWorkItemSubTaskReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).RoleAssignee = roleAssignee
	return builder
}

func (builder *UpdateWorkItemSubTaskReqBuilder) Schedule(schedule *Schedule) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).Schedule = schedule
	return builder
}

func (builder *UpdateWorkItemSubTaskReqBuilder) Note(note string) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).Note = &note
	return builder
}


func (builder *UpdateWorkItemSubTaskReqBuilder) Deliverable(deliverable []FieldValuePair) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).Deliverable = deliverable
	return builder
}

func (builder *UpdateWorkItemSubTaskReqBuilder) UpdateFields(updateFields []FieldValuePair) *UpdateWorkItemSubTaskReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemSubTaskReqBody).UpdateFields = updateFields
	return builder
}
func (builder *UpdateWorkItemSubTaskReqBuilder) Build() *UpdateWorkItemSubTaskReq {
	req := &UpdateWorkItemSubTaskReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateWorkItemTypeInfoReq struct {
	apiReq *core.APIReq
}
type UpdateWorkItemTypeInfoReqBody struct {
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
type UpdateWorkItemTypeInfoResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateWorkItemTypeInfoRespData        `json:"data,omitempty"`
}

type UpdateWorkItemTypeInfoRespData struct {
}

type UpdateWorkItemTypeInfoReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateWorkItemTypeInfoReqBuilder() *UpdateWorkItemTypeInfoReqBuilder {
	builder := &UpdateWorkItemTypeInfoReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateWorkItemTypeInfoReqBody{},
	}
	return builder
}

func (builder *UpdateWorkItemTypeInfoReqBuilder) ProjectKey(projectKey string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) Description(description string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).Description = &description
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) IsDisabled(isDisabled bool) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).IsDisabled = &isDisabled
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) IsPinned(isPinned bool) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).IsPinned = &isPinned
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) EnableSchedule(enableSchedule bool) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).EnableSchedule = &enableSchedule
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) ScheduleFieldKey(scheduleFieldKey string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).ScheduleFieldKey = &scheduleFieldKey
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) EstimatePointFieldKey(estimatePointFieldKey string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).EstimatePointFieldKey = &estimatePointFieldKey
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) ActualWorkTimeFieldKey(actualWorkTimeFieldKey string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).ActualWorkTimeFieldKey = &actualWorkTimeFieldKey
	return builder
}


func (builder *UpdateWorkItemTypeInfoReqBuilder) BelongRoleKeys(belongRoleKeys []string) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).BelongRoleKeys = belongRoleKeys
	return builder
}

func (builder *UpdateWorkItemTypeInfoReqBuilder) ActualWorkTimeSwitch(actualWorkTimeSwitch bool) *UpdateWorkItemTypeInfoReqBuilder {
	builder.apiReq.Body.(*UpdateWorkItemTypeInfoReqBody).ActualWorkTimeSwitch = &actualWorkTimeSwitch
	return builder
}

func (builder *UpdateWorkItemTypeInfoReqBuilder) Build() *UpdateWorkItemTypeInfoReq {
	req := &UpdateWorkItemTypeInfoReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateWorkflowNodeReq struct {
	apiReq *core.APIReq
}
type UpdateWorkflowNodeReqBody struct {
    NodeOwners  []string `json:"node_owners,omitempty"`
    NodeSchedule  *Schedule `json:"node_schedule,omitempty"`
    Schedules  []Schedule `json:"schedules,omitempty"`
    Fields  []FieldValuePair `json:"fields,omitempty"`
    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`
}
type UpdateWorkflowNodeResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateWorkflowNodeRespData        `json:"data,omitempty"`
}

type UpdateWorkflowNodeRespData struct {
}

type UpdateWorkflowNodeReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateWorkflowNodeReqBuilder() *UpdateWorkflowNodeReqBuilder {
	builder := &UpdateWorkflowNodeReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateWorkflowNodeReqBody{},
	}
	return builder
}

func (builder *UpdateWorkflowNodeReqBuilder) ProjectKey(projectKey string) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateWorkflowNodeReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateWorkflowNodeReqBuilder) WorkItemID(workItemID int64) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateWorkflowNodeReqBuilder) NodeID(nodeID string) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.PathParams.Set("node_id", fmt.Sprint(nodeID))
	return builder
}


func (builder *UpdateWorkflowNodeReqBuilder) NodeOwners(nodeOwners []string) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*UpdateWorkflowNodeReqBody).NodeOwners = nodeOwners
	return builder
}

func (builder *UpdateWorkflowNodeReqBuilder) NodeSchedule(nodeSchedule *Schedule) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*UpdateWorkflowNodeReqBody).NodeSchedule = nodeSchedule
	return builder
}

func (builder *UpdateWorkflowNodeReqBuilder) Schedules(schedules []Schedule) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*UpdateWorkflowNodeReqBody).Schedules = schedules
	return builder
}

func (builder *UpdateWorkflowNodeReqBuilder) Fields(fields []FieldValuePair) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*UpdateWorkflowNodeReqBody).Fields = fields
	return builder
}

func (builder *UpdateWorkflowNodeReqBuilder) RoleAssignee(roleAssignee []RoleOwner) *UpdateWorkflowNodeReqBuilder {
	builder.apiReq.Body.(*UpdateWorkflowNodeReqBody).RoleAssignee = roleAssignee
	return builder
}
func (builder *UpdateWorkflowNodeReqBuilder) Build() *UpdateWorkflowNodeReq {
	req := &UpdateWorkflowNodeReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateWorkingHourRecordReq struct {
	apiReq *core.APIReq
}
type UpdateWorkingHourRecordReqBody struct {
    WorkingHourRecords  []UpdateWorkingHourRecord `json:"working_hour_records,omitempty"`
}
type UpdateWorkingHourRecordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UpdateWorkingHourRecordRespData        `json:"data,omitempty"`
}

type UpdateWorkingHourRecordRespData struct {
}

type UpdateWorkingHourRecordReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateWorkingHourRecordReqBuilder() *UpdateWorkingHourRecordReqBuilder {
	builder := &UpdateWorkingHourRecordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateWorkingHourRecordReqBody{},
	}
	return builder
}

func (builder *UpdateWorkingHourRecordReqBuilder) ProjectKey(projectKey string) *UpdateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateWorkingHourRecordReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateWorkingHourRecordReqBuilder) WorkItemID(workItemID int64) *UpdateWorkingHourRecordReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(workItemID))
	return builder
}


func (builder *UpdateWorkingHourRecordReqBuilder) WorkingHourRecords(workingHourRecords []UpdateWorkingHourRecord) *UpdateWorkingHourRecordReqBuilder {
	builder.apiReq.Body.(*UpdateWorkingHourRecordReqBody).WorkingHourRecords = workingHourRecords
	return builder
}
func (builder *UpdateWorkingHourRecordReqBuilder) Build() *UpdateWorkingHourRecordReq {
	req := &UpdateWorkingHourRecordReq{}
	req.apiReq = builder.apiReq
	return req
}

type WBSUpdateDraftFrozenRowsReq struct {
	apiReq *core.APIReq
}
type WBSUpdateDraftFrozenRowsReqBody struct {
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    DraftID  *string `json:"draft_id,omitempty"`
    UpdateType  *int32 `json:"update_type,omitempty"`
    CommitID  *string `json:"commit_id,omitempty"`
}
type WBSUpdateDraftFrozenRowsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *WBSUpdateDraftFrozenRowsRespData        `json:"data,omitempty"`
}

type WBSUpdateDraftFrozenRowsRespData struct {
}

type WBSUpdateDraftFrozenRowsReqBuilder struct {
	apiReq *core.APIReq
}

func NewWBSUpdateDraftFrozenRowsReqBuilder() *WBSUpdateDraftFrozenRowsReqBuilder {
	builder := &WBSUpdateDraftFrozenRowsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &WBSUpdateDraftFrozenRowsReqBody{},
	}
	return builder
}

func (builder *WBSUpdateDraftFrozenRowsReqBuilder) ProjectKey(projectKey string) *WBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *WBSUpdateDraftFrozenRowsReqBuilder) WorkItemID(workItemID int64) *WBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*WBSUpdateDraftFrozenRowsReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *WBSUpdateDraftFrozenRowsReqBuilder) DraftID(draftID string) *WBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*WBSUpdateDraftFrozenRowsReqBody).DraftID = &draftID
	return builder
}


func (builder *WBSUpdateDraftFrozenRowsReqBuilder) UpdateType(updateType int32) *WBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*WBSUpdateDraftFrozenRowsReqBody).UpdateType = &updateType
	return builder
}


func (builder *WBSUpdateDraftFrozenRowsReqBuilder) CommitID(commitID string) *WBSUpdateDraftFrozenRowsReqBuilder {
	builder.apiReq.Body.(*WBSUpdateDraftFrozenRowsReqBody).CommitID = &commitID
	return builder
}

func (builder *WBSUpdateDraftFrozenRowsReqBuilder) Build() *WBSUpdateDraftFrozenRowsReq {
	req := &WBSUpdateDraftFrozenRowsReq{}
	req.apiReq = builder.apiReq
	return req
}

type WbsCollaborationPublishReq struct {
	apiReq *core.APIReq
}
type WbsCollaborationPublishReqBody struct {
    WorkItemID  *int64 `json:"work_item_id,omitempty"`
    DraftID  *string `json:"draft_id,omitempty"`
    CommitID  *string `json:"commit_id,omitempty"`
}
type WbsCollaborationPublishResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *WbsCollaborationPublishRespData        `json:"data,omitempty"`
}

type WbsCollaborationPublishRespData struct {
	Success       *bool         `json:"success,omitempty"`
}

type WbsCollaborationPublishReqBuilder struct {
	apiReq *core.APIReq
}

func NewWbsCollaborationPublishReqBuilder() *WbsCollaborationPublishReqBuilder {
	builder := &WbsCollaborationPublishReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &WbsCollaborationPublishReqBody{},
	}
	return builder
}

func (builder *WbsCollaborationPublishReqBuilder) ProjectKey(projectKey string) *WbsCollaborationPublishReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *WbsCollaborationPublishReqBuilder) WorkItemID(workItemID int64) *WbsCollaborationPublishReqBuilder {
	builder.apiReq.Body.(*WbsCollaborationPublishReqBody).WorkItemID = &workItemID
	return builder
}


func (builder *WbsCollaborationPublishReqBuilder) DraftID(draftID string) *WbsCollaborationPublishReqBuilder {
	builder.apiReq.Body.(*WbsCollaborationPublishReqBody).DraftID = &draftID
	return builder
}


func (builder *WbsCollaborationPublishReqBuilder) CommitID(commitID string) *WbsCollaborationPublishReqBuilder {
	builder.apiReq.Body.(*WbsCollaborationPublishReqBody).CommitID = &commitID
	return builder
}

func (builder *WbsCollaborationPublishReqBuilder) Build() *WbsCollaborationPublishReq {
	req := &WbsCollaborationPublishReq{}
	req.apiReq = builder.apiReq
	return req
}

