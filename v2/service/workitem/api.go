package workitem

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

const APIPath_AbortWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/abort"

const APIPath_ActualTimeUpdate = "/open_api/work_item/actual_time/update"

const APIPath_BatchQueryConclusionOption = "/open_api/work_item/finished/query_conclusion_option"

const APIPath_BatchQueryDeliverable = "/open_api/work_item/deliverable/batch_query"

const APIPath_BatchQueryFinished = "/open_api/work_item/finished/batch_query"

const APIPath_BatchUpdateBasicWorkItem = "/open_api/work_item/batch_update"

const APIPath_CompleteCreateAuditDraft = "/open_api/wbs_view_draft/complete-create-audit"

const APIPath_CompositiveSearch = "/open_api/compositive_search"

const APIPath_CreateField = "/open_api/:project_key/field/:work_item_type_key/create"

const APIPath_CreateFlowRole = "/open_api/:project_key/flow_roles/:work_item_type_key/create_role"

const APIPath_CreateProjectRelationInstances = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/batch_bind"

const APIPath_CreateStoryRelations = "/open_api/:project_key/story_relations/create"

const APIPath_CreateTemplateDetail = "/open_api/template/v2/create_template"

const APIPath_CreateWorkItem = "/open_api/:project_key/work_item/create"

const APIPath_CreateWorkItemRelation = "/open_api/work_item/relation/create"

const APIPath_CreateWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/task"

const APIPath_CreateWorkingHourRecord = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/work_hour_record"

const APIPath_DeleteFile = "/open_api/file/delete"

const APIPath_DeleteFlowRole = "/open_api/:project_key/flow_roles/:work_item_type_key/delete_role"

const APIPath_DeleteProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id"

const APIPath_DeleteTemplateDetail = "/open_api/template/v2/delete_template/:project_key/:template_id"

const APIPath_DeleteWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id"

const APIPath_DeleteWorkItemRelation = "/open_api/work_item/relation/delete"

const APIPath_DeleteWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/task/:task_id"

const APIPath_DeleteWorkingHourRecord = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/work_hour_record"

const APIPath_ElementTemplateCreate = "/open_api/work_item/element_template/create"

const APIPath_ElementTemplateQuery = "/open_api/work_item/element_template/query"

const APIPath_Filter = "/open_api/:project_key/work_item/filter"

const APIPath_FilterAcrossProject = "/open_api/work_items/filter_across_project"

const APIPath_FreezeWorkItem = "/open_api/work_item/freeze"

const APIPath_GetResourceWorkItemsByIds = "/open_api/work_item/resource/query"

const APIPath_GetWBSInfo = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/wbs_view"

const APIPath_GetWBSViewSubWorkItemConf = "/open_api/work_item/wbs_view_draft/sub_work_item_conf"

const APIPath_GetWorkFlow = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/query"

const APIPath_GetWorkItemManHourRecords = "/open_api/work_item/man_hour/records"

const APIPath_GetWorkItemOpRecord = "/open_api/op_record/work_item/list"

const APIPath_GetWorkItemTransRequiredItem = "/open_api/work_item/transition_required_info/get"

const APIPath_GetWorkItemTypeInfoByKey = "/open_api/:project_key/work_item/type/:work_item_type_key"

const APIPath_GetWorkItemsByIds = "/open_api/:project_key/work_item/:work_item_type_key/query"

const APIPath_IntegrateSearch = "/open_api/view_search/integrate_search"

const APIPath_InviteBotJoinChat = "/open_api/:project_key/work_item/:work_item_id/bot_join_chat"

const APIPath_ListTemplateConf = "/open_api/:project_key/template_list/:work_item_type_key"

const APIPath_PatchWBSViewDraft = "/open_api/work_item/wbs_view_draft/patch"

const APIPath_PublishWBSViewDraft = "/open_api/work_item/wbs_view_draft/publish"

const APIPath_QueryAWorkItemTypes = "/open_api/:project_key/work_item/all-types"

const APIPath_QueryBusinesses = "/open_api/:project_key/business/all"

const APIPath_QueryProjectFields = "/open_api/:project_key/field/all"

const APIPath_QueryProjectRelation = "/open_api/:project_key/relation/rules"

const APIPath_QueryProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/work_item_list"

const APIPath_QueryRoleConfDetails = "/open_api/:project_key/flow_roles/:work_item_type_key"

const APIPath_QueryStoryRelations = "/open_api/:project_key/story_relations/query"

const APIPath_QueryTaskResult = "/open_api/task_result"

const APIPath_QueryTemplateDetail = "/open_api/:project_key/template_detail/:template_id"

const APIPath_QueryWBSViewDraft = "/open_api/work_item/wbs_view_draft/query"

const APIPath_QueryWbsTemplateConf = "/open_api/:project_key/wbs_template"

const APIPath_QueryWorkItemMetaData = "/open_api/:project_key/work_item/:work_item_type_key/meta"

const APIPath_QueryWorkItemRelation = "/open_api/:project_key/work_item/relation"

const APIPath_QueryWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/task"

const APIPath_ResetWBSViewDraft = "/open_api/work_item/wbs_view_draft/reset"

const APIPath_ResourceCreateInstance = "/open_api/work_item/resource/:project_key/:work_item_id/create_instance"

const APIPath_ResourceCreateWorkItem = "/open_api/work_item/resource/create_work_item"

const APIPath_ResourceSearchByParams = "/open_api/work_item/resource/search/params"

const APIPath_SearchByParams = "/open_api/:project_key/work_item/:work_item_type_key/search/params"

const APIPath_SearchSubtask = "/open_api/work_item/subtask/search"

const APIPath_SearchWorkItemsRelation = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/search_by_relation"

const APIPath_SubTaskModify = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/subtask/modify"

const APIPath_SwitchBackToWbsViewDraft = "/open_api/:project_key/wbs_view_draft/switch"

const APIPath_UniversalSearch = "/open_api/view_search/universal_search"

const APIPath_UpdateCompoundFieldValue = "/open_api/work_item/field_value/update_compound_field"

const APIPath_UpdateField = "/open_api/:project_key/field/:work_item_type_key"

const APIPath_UpdateFinished = "/open_api/work_item/finished/update"

const APIPath_UpdateFlowRole = "/open_api/:project_key/flow_roles/:work_item_type_key/update_role"

const APIPath_UpdateMultiSignal = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/update/multi_signal"

const APIPath_UpdateNodeState = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/:node_id/operate"

const APIPath_UpdateResourceWorkItem = "/open_api/work_item/resource/update"

const APIPath_UpdateStateFlow = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/state_change"

const APIPath_UpdateTemplateDetail = "/open_api/template/v2/update_template"

const APIPath_UpdateWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id"

const APIPath_UpdateWorkItemRelation = "/open_api/work_item/relation/update"

const APIPath_UpdateWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/:node_id/task/:task_id"

const APIPath_UpdateWorkItemTypeInfo = "/open_api/:project_key/work_item/type/:work_item_type_key"

const APIPath_UpdateWorkflowNode = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/:node_id"

const APIPath_UpdateWorkingHourRecord = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/work_hour_record"

const APIPath_WBSUpdateDraftFrozenRows = "/open_api/:project_key/wbs_view_draft/update/frozen_rows"

const APIPath_WbsCollaborationPublish = "/open_api/:project_key/wbs_view_draft/publish"

func NewService(config *core.Config) *WorkItemService {
	a := &WorkItemService{config: config}
	return a
}

type WorkItemService struct {
	config *core.Config
}

/*
 * @name: OAPIAbortWorkItem
 * @desc: 终止/恢复工作项
 */
func (a *WorkItemService) AbortWorkItem(ctx context.Context, req *AbortWorkItemReq, options ...core.RequestOptionFunc) (*AbortWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_AbortWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[AbortWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &AbortWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[AbortWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIActualTimeUpdate
 * @desc: 更新节点实际开始时间和结束时间
 */
func (a *WorkItemService) ActualTimeUpdate(ctx context.Context, req *ActualTimeUpdateReq, options ...core.RequestOptionFunc) (*ActualTimeUpdateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ActualTimeUpdate
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ActualTimeUpdate] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ActualTimeUpdateResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ActualTimeUpdate] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIBatchQueryConclusionLabel
 * @desc: 批量查询工作项交付物详情信息
 */
func (a *WorkItemService) BatchQueryConclusionOption(ctx context.Context, req *BatchQueryConclusionOptionReq, options ...core.RequestOptionFunc) (*BatchQueryConclusionOptionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_BatchQueryConclusionOption
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryConclusionOption] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchQueryConclusionOptionResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryConclusionOption] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 *
 */
func (a *WorkItemService) BatchQueryDeliverable(ctx context.Context, req *BatchQueryDeliverableReq, options ...core.RequestOptionFunc) (*BatchQueryDeliverableResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_BatchQueryDeliverable
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryDeliverable] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchQueryDeliverableResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryDeliverable] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIBatchQueryFinished
 * @desc: 批量查询工作项交付物详情信息
 */
func (a *WorkItemService) BatchQueryFinished(ctx context.Context, req *BatchQueryFinishedReq, options ...core.RequestOptionFunc) (*BatchQueryFinishedResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_BatchQueryFinished
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryFinished] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchQueryFinishedResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchQueryFinished] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIBatchUpdateBasicWorkItem
 * @desc: 批量更新工作项字段-异步
 */
func (a *WorkItemService) BatchUpdateBasicWorkItem(ctx context.Context, req *BatchUpdateBasicWorkItemReq, options ...core.RequestOptionFunc) (*BatchUpdateBasicWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_BatchUpdateBasicWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchUpdateBasicWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchUpdateBasicWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[BatchUpdateBasicWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICompleteCreateAuditDraft
 * @desc: 完成创建审批草稿
 */
func (a *WorkItemService) CompleteCreateAuditDraft(ctx context.Context, req *CompleteCreateAuditDraftReq, options ...core.RequestOptionFunc) (*CompleteCreateAuditDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CompleteCreateAuditDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CompleteCreateAuditDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CompleteCreateAuditDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CompleteCreateAuditDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（全文搜索）
 * @desc: openapi获取指定的工作项列表（全文搜索）
 */
func (a *WorkItemService) CompositiveSearch(ctx context.Context, req *CompositiveSearchReq, options ...core.RequestOptionFunc) (*CompositiveSearchResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CompositiveSearch
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CompositiveSearch] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CompositiveSearchResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CompositiveSearch] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateField
 * @desc: 创建自定义字段
 */
func (a *WorkItemService) CreateField(ctx context.Context, req *CreateFieldReq, options ...core.RequestOptionFunc) (*CreateFieldResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateField
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateField] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFieldResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateField] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateFlowRole
 * @desc: 创建流程角色
 */
func (a *WorkItemService) CreateFlowRole(ctx context.Context, req *CreateFlowRoleReq, options ...core.RequestOptionFunc) (*CreateFlowRoleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateFlowRole
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateFlowRole] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFlowRoleResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateFlowRole] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateProjectRelationInstances
 * @desc: 通过空间关联绑定关联工作项
 */
func (a *WorkItemService) CreateProjectRelationInstances(ctx context.Context, req *CreateProjectRelationInstancesReq, options ...core.RequestOptionFunc) (*CreateProjectRelationInstancesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateProjectRelationInstances
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateProjectRelationInstances] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateProjectRelationInstancesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateProjectRelationInstances] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateStoryRelations
 * @desc: 创建需求关联关系
 */
func (a *WorkItemService) CreateStoryRelations(ctx context.Context, req *CreateStoryRelationsReq, options ...core.RequestOptionFunc) (*CreateStoryRelationsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateStoryRelations
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateStoryRelations] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateStoryRelationsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateStoryRelations] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateTemplateDetail
 * @desc: 新增流程类型配置
 */
func (a *WorkItemService) CreateTemplateDetail(ctx context.Context, req *CreateTemplateDetailReq, options ...core.RequestOptionFunc) (*CreateTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateTemplateDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkItem
 * @desc: 创建工作项
 */
func (a *WorkItemService) CreateWorkItem(ctx context.Context, req *CreateWorkItemReq, options ...core.RequestOptionFunc) (*CreateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkItemRelation
 * @desc: 创建工作项关系
 */
func (a *WorkItemService) CreateWorkItemRelation(ctx context.Context, req *CreateWorkItemRelationReq, options ...core.RequestOptionFunc) (*CreateWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateWorkItemRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkItemSubTask
 * @desc: 创建子任务
 */
func (a *WorkItemService) CreateWorkItemSubTask(ctx context.Context, req *CreateWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*CreateWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateWorkItemSubTask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkingHourRecord
 * @desc: 创建实际工时
 */
func (a *WorkItemService) CreateWorkingHourRecord(ctx context.Context, req *CreateWorkingHourRecordReq, options ...core.RequestOptionFunc) (*CreateWorkingHourRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateWorkingHourRecord
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkingHourRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateWorkingHourRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateWorkingHourRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteFile
 * @desc: 删除实例对应的文件
 */
func (a *WorkItemService) DeleteFile(ctx context.Context, req *DeleteFileReq, options ...core.RequestOptionFunc) (*DeleteFileResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteFile
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFile] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFileResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFile] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteFlowRole
 * @desc: 删除流程角色
 */
func (a *WorkItemService) DeleteFlowRole(ctx context.Context, req *DeleteFlowRoleReq, options ...core.RequestOptionFunc) (*DeleteFlowRoleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteFlowRole
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFlowRole] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteFlowRoleResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteFlowRole] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteProjectRelationInstance
 * @desc: 通过空间关联解绑关联工作项
 */
func (a *WorkItemService) DeleteProjectRelationInstance(ctx context.Context, req *DeleteProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*DeleteProjectRelationInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteProjectRelationInstance
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteProjectRelationInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteTemplateDetail
 * @desc: 删除流程类型配置
 */
func (a *WorkItemService) DeleteTemplateDetail(ctx context.Context, req *DeleteTemplateDetailReq, options ...core.RequestOptionFunc) (*DeleteTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteTemplateDetail
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkItem
 * @desc: 删除工作项
 */
func (a *WorkItemService) DeleteWorkItem(ctx context.Context, req *DeleteWorkItemReq, options ...core.RequestOptionFunc) (*DeleteWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteWorkItem
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkItemRelation
 * @desc: 删除工作项关系
 */
func (a *WorkItemService) DeleteWorkItemRelation(ctx context.Context, req *DeleteWorkItemRelationReq, options ...core.RequestOptionFunc) (*DeleteWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteWorkItemRelation
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkItemSubTask
 * @desc: 删除子任务
 */
func (a *WorkItemService) DeleteWorkItemSubTask(ctx context.Context, req *DeleteWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*DeleteWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteWorkItemSubTask
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkingHourRecord
 * @desc: 删除实际工时
 */
func (a *WorkItemService) DeleteWorkingHourRecord(ctx context.Context, req *DeleteWorkingHourRecordReq, options ...core.RequestOptionFunc) (*DeleteWorkingHourRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteWorkingHourRecord
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkingHourRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteWorkingHourRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteWorkingHourRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemTemplate
 * @desc: 创建资源库资源(节点/任务)
 */
func (a *WorkItemService) ElementTemplateCreate(ctx context.Context, req *ElementTemplateCreateReq, options ...core.RequestOptionFunc) (*ElementTemplateCreateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ElementTemplateCreate
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ElementTemplateCreate] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ElementTemplateCreateResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ElementTemplateCreate] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIElementTemplateQueryAll
 * @desc: 查询资源库资源(节点/任务)
 */
func (a *WorkItemService) ElementTemplateQuery(ctx context.Context, req *ElementTemplateQueryReq, options ...core.RequestOptionFunc) (*ElementTemplateQueryResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ElementTemplateQuery
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ElementTemplateQuery] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ElementTemplateQueryResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ElementTemplateQuery] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（非跨空间）
 * @desc: openapi获取指定的工作项列表（非跨空间）
 */
func (a *WorkItemService) Filter(ctx context.Context, req *FilterReq, options ...core.RequestOptionFunc) (*FilterResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_Filter
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[Filter] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &FilterResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[Filter] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（跨空间）
 * @desc: openapi获取指定的工作项列表（跨空间）
 */
func (a *WorkItemService) FilterAcrossProject(ctx context.Context, req *FilterAcrossProjectReq, options ...core.RequestOptionFunc) (*FilterAcrossProjectResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_FilterAcrossProject
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[FilterAcrossProject] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &FilterAcrossProjectResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[FilterAcrossProject] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIFreezeWorkItem
 * @desc: 冻结/终止工作项
 */
func (a *WorkItemService) FreezeWorkItem(ctx context.Context, req *FreezeWorkItemReq, options ...core.RequestOptionFunc) (*FreezeWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_FreezeWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[FreezeWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &FreezeWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[FreezeWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemsByIds
 * @desc: 批量查询工作项
 */
func (a *WorkItemService) GetResourceWorkItemsByIds(ctx context.Context, req *GetResourceWorkItemsByIdsReq, options ...core.RequestOptionFunc) (*GetResourceWorkItemsByIdsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetResourceWorkItemsByIds
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetResourceWorkItemsByIds] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetResourceWorkItemsByIdsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetResourceWorkItemsByIds] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWBSInfo
 * @desc: 获取WBS
 */
func (a *WorkItemService) GetWBSInfo(ctx context.Context, req *GetWBSInfoReq, options ...core.RequestOptionFunc) (*GetWBSInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWBSInfo
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWBSInfo] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWBSInfoResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWBSInfo] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIActualTimeUpdate
 * @desc: 更新节点实际开始时间和结束时间
 */
func (a *WorkItemService) GetWBSViewSubWorkItemConf(ctx context.Context, req *GetWBSViewSubWorkItemConfReq, options ...core.RequestOptionFunc) (*GetWBSViewSubWorkItemConfResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWBSViewSubWorkItemConf
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWBSViewSubWorkItemConf] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWBSViewSubWorkItemConfResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWBSViewSubWorkItemConf] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkFlow
 * @desc: 获取工作流
 */
func (a *WorkItemService) GetWorkFlow(ctx context.Context, req *GetWorkFlowReq, options ...core.RequestOptionFunc) (*GetWorkFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWorkFlow
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkFlow] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWorkFlowResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkFlow] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemManHourRecords
 * @desc: 获取工作项的工时记录列表
 */
func (a *WorkItemService) GetWorkItemManHourRecords(ctx context.Context, req *GetWorkItemManHourRecordsReq, options ...core.RequestOptionFunc) (*GetWorkItemManHourRecordsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWorkItemManHourRecords
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemManHourRecords] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWorkItemManHourRecordsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemManHourRecords] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name:OAPIPageGetWorkItemOpRecord
 * @desc:OpenAPI，查询操作记录详情
 */
func (a *WorkItemService) GetWorkItemOpRecord(ctx context.Context, req *GetWorkItemOpRecordReq, options ...core.RequestOptionFunc) (*GetWorkItemOpRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWorkItemOpRecord
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemOpRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWorkItemOpRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemOpRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemTransRequiredItem
 * @desc: 获取节点/状态流转所需的必填信息
 */
func (a *WorkItemService) GetWorkItemTransRequiredItem(ctx context.Context, req *GetWorkItemTransRequiredItemReq, options ...core.RequestOptionFunc) (*GetWorkItemTransRequiredItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWorkItemTransRequiredItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemTransRequiredItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWorkItemTransRequiredItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemTransRequiredItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemTypeInfoByKey
 * @desc: 查询工作项类型
 */
func (a *WorkItemService) GetWorkItemTypeInfoByKey(ctx context.Context, req *GetWorkItemTypeInfoByKeyReq, options ...core.RequestOptionFunc) (*GetWorkItemTypeInfoByKeyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWorkItemTypeInfoByKey
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemTypeInfoByKey] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWorkItemTypeInfoByKeyResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemTypeInfoByKey] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemsByIds
 * @desc: 批量查询工作项
 */
func (a *WorkItemService) GetWorkItemsByIds(ctx context.Context, req *GetWorkItemsByIdsReq, options ...core.RequestOptionFunc) (*GetWorkItemsByIdsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetWorkItemsByIds
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemsByIds] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetWorkItemsByIdsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetWorkItemsByIds] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（极简链路检索能力，三合一能力）
 * @desc: openapi获取指定的工作项列表（极简链路检索能力，三合一能力）
 */
func (a *WorkItemService) IntegrateSearch(ctx context.Context, req *IntegrateSearchReq, options ...core.RequestOptionFunc) (*IntegrateSearchResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_IntegrateSearch
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[IntegrateSearch] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &IntegrateSearchResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[IntegrateSearch] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIInviteBotJoinChatV2
 * @desc: 邀请机器人入群
 */
func (a *WorkItemService) InviteBotJoinChat(ctx context.Context, req *InviteBotJoinChatReq, options ...core.RequestOptionFunc) (*InviteBotJoinChatResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_InviteBotJoinChat
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[InviteBotJoinChat] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &InviteBotJoinChatResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[InviteBotJoinChat] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIListTemplateConf
 * @desc: 获取工作项下的流程类型
 */
func (a *WorkItemService) ListTemplateConf(ctx context.Context, req *ListTemplateConfReq, options ...core.RequestOptionFunc) (*ListTemplateConfResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListTemplateConf
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListTemplateConf] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListTemplateConfResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListTemplateConf] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: wbs草稿态编辑
 * @desc: wbs草稿态编辑
 */
func (a *WorkItemService) PatchWBSViewDraft(ctx context.Context, req *PatchWBSViewDraftReq, options ...core.RequestOptionFunc) (*PatchWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_PatchWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[PatchWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[PatchWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIWbsCollaborationPublish
 * @desc: 计划表基于草稿发布,包含插件逻辑
 */
func (a *WorkItemService) PublishWBSViewDraft(ctx context.Context, req *PublishWBSViewDraftReq, options ...core.RequestOptionFunc) (*PublishWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_PublishWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[PublishWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &PublishWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[PublishWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryAWorkItemTypes
 * @desc: 获取空间下工作项类型
 */
func (a *WorkItemService) QueryAWorkItemTypes(ctx context.Context, req *QueryAWorkItemTypesReq, options ...core.RequestOptionFunc) (*QueryAWorkItemTypesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryAWorkItemTypes
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryAWorkItemTypes] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryAWorkItemTypesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryAWorkItemTypes] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryBusinesses
 * @desc: 获取空间下业务线
 */
func (a *WorkItemService) QueryBusinesses(ctx context.Context, req *QueryBusinessesReq, options ...core.RequestOptionFunc) (*QueryBusinessesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryBusinesses
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryBusinesses] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryBusinessesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryBusinesses] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryProjectFields
 * @desc: 获取空间字段
 */
func (a *WorkItemService) QueryProjectFields(ctx context.Context, req *QueryProjectFieldsReq, options ...core.RequestOptionFunc) (*QueryProjectFieldsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryProjectFields
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectFields] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryProjectFieldsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectFields] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryProjectRelation
 * @desc: 查询空间关系
 */
func (a *WorkItemService) QueryProjectRelation(ctx context.Context, req *QueryProjectRelationReq, options ...core.RequestOptionFunc) (*QueryProjectRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryProjectRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryProjectRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryProjectRelationInstance
 * @desc: 获取空间关联下的关联工作项列表
 */
func (a *WorkItemService) QueryProjectRelationInstance(ctx context.Context, req *QueryProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*QueryProjectRelationInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryProjectRelationInstance
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryProjectRelationInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryRoleConfDetails
 * @desc: 获取流程角色配置详情
 */
func (a *WorkItemService) QueryRoleConfDetails(ctx context.Context, req *QueryRoleConfDetailsReq, options ...core.RequestOptionFunc) (*QueryRoleConfDetailsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryRoleConfDetails
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryRoleConfDetails] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryRoleConfDetailsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryRoleConfDetails] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryStoryRelations
 * @desc: 查询需求关联关系
 */
func (a *WorkItemService) QueryStoryRelations(ctx context.Context, req *QueryStoryRelationsReq, options ...core.RequestOptionFunc) (*QueryStoryRelationsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryStoryRelations
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryStoryRelations] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryStoryRelationsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryStoryRelations] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryTaskResult
 * @desc: 查询任务执行情况——主要针对批量任务处理
 */
func (a *WorkItemService) QueryTaskResult(ctx context.Context, req *QueryTaskResultReq, options ...core.RequestOptionFunc) (*QueryTaskResultResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryTaskResult
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTaskResult] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTaskResultResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTaskResult] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryTemplateDetail
 * @desc: 获取流程类型配置详情
 */
func (a *WorkItemService) QueryTemplateDetail(ctx context.Context, req *QueryTemplateDetailReq, options ...core.RequestOptionFunc) (*QueryTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryTemplateDetail
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIInitWBSViewDraft
 * @desc: 查询计划表草稿
 */
func (a *WorkItemService) QueryWBSViewDraft(ctx context.Context, req *QueryWBSViewDraftReq, options ...core.RequestOptionFunc) (*QueryWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWbsTemplateConf
 * @desc: 获取流程类型配置（wbs）
 */
func (a *WorkItemService) QueryWbsTemplateConf(ctx context.Context, req *QueryWbsTemplateConfReq, options ...core.RequestOptionFunc) (*QueryWbsTemplateConfResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWbsTemplateConf
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWbsTemplateConf] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWbsTemplateConfResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWbsTemplateConf] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemMetaData
 * @desc: 获取创建工作项元信息
 */
func (a *WorkItemService) QueryWorkItemMetaData(ctx context.Context, req *QueryWorkItemMetaDataReq, options ...core.RequestOptionFunc) (*QueryWorkItemMetaDataResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemMetaData
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemMetaData] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemMetaDataResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemMetaData] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemRelation
 * @desc: 查询工作项关系
 */
func (a *WorkItemService) QueryWorkItemRelation(ctx context.Context, req *QueryWorkItemRelationReq, options ...core.RequestOptionFunc) (*QueryWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemRelation
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemSubTask
 * @desc: 获取子任务详情
 */
func (a *WorkItemService) QueryWorkItemSubTask(ctx context.Context, req *QueryWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*QueryWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryWorkItemSubTask
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIInitWBSViewDraft
 * @desc: 重置计划表草稿
 */
func (a *WorkItemService) ResetWBSViewDraft(ctx context.Context, req *ResetWBSViewDraftReq, options ...core.RequestOptionFunc) (*ResetWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ResetWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResetWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ResetWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResetWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: 0APIResourceCreateInstance
 * @desc: 通过资源创建实例
 */
func (a *WorkItemService) ResourceCreateInstance(ctx context.Context, req *ResourceCreateInstanceReq, options ...core.RequestOptionFunc) (*ResourceCreateInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ResourceCreateInstance
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResourceCreateInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ResourceCreateInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResourceCreateInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIResourceCreateWorkItem
 * @desc: 创建工作项资源库
 */
func (a *WorkItemService) ResourceCreateWorkItem(ctx context.Context, req *ResourceCreateWorkItemReq, options ...core.RequestOptionFunc) (*ResourceCreateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ResourceCreateWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResourceCreateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ResourceCreateWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResourceCreateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的资源工作项列表（单空间 | 复杂传参）
 * @desc: openapi获取指定的资源工作项列表（单空间 | 复杂传参）
 */
func (a *WorkItemService) ResourceSearchByParams(ctx context.Context, req *ResourceSearchByParamsReq, options ...core.RequestOptionFunc) (*ResourceSearchByParamsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ResourceSearchByParams
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResourceSearchByParams] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ResourceSearchByParamsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ResourceSearchByParams] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（单空间 | 复杂传参）
 * @desc: openapi获取指定的工作项列表（单空间 | 复杂传参）
 */
func (a *WorkItemService) SearchByParams(ctx context.Context, req *SearchByParamsReq, options ...core.RequestOptionFunc) (*SearchByParamsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchByParams
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchByParams] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchByParamsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchByParams] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的子任务列表（跨空间）
 * @desc: openapi获取指定的子任务列表（跨空间）
 */
func (a *WorkItemService) SearchSubtask(ctx context.Context, req *SearchSubtaskReq, options ...core.RequestOptionFunc) (*SearchSubtaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchSubtask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchSubtask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchSubtaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchSubtask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISearchWorkItemsRelation
 * @desc: 获取指定的关联工作项列表（单空间）
 */
func (a *WorkItemService) SearchWorkItemsRelation(ctx context.Context, req *SearchWorkItemsRelationReq, options ...core.RequestOptionFunc) (*SearchWorkItemsRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchWorkItemsRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchWorkItemsRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchWorkItemsRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchWorkItemsRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISubTaskModify
 * @desc: 子任务完成/回滚
 */
func (a *WorkItemService) SubTaskModify(ctx context.Context, req *SubTaskModifyReq, options ...core.RequestOptionFunc) (*SubTaskModifyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SubTaskModify
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SubTaskModify] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SubTaskModifyResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SubTaskModify] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISwitchBackToWbsViewDraft
 * @desc: 审批拒绝切回草稿
 */
func (a *WorkItemService) SwitchBackToWbsViewDraft(ctx context.Context, req *SwitchBackToWbsViewDraftReq, options ...core.RequestOptionFunc) (*SwitchBackToWbsViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SwitchBackToWbsViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SwitchBackToWbsViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SwitchBackToWbsViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SwitchBackToWbsViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（极简链路，替代原有的GeneralSearch系列）
 * @desc: openapi获取指定的工作项列表（极简链路，替代原有的GeneralSearch系列）
 */
func (a *WorkItemService) UniversalSearch(ctx context.Context, req *UniversalSearchReq, options ...core.RequestOptionFunc) (*UniversalSearchResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UniversalSearch
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UniversalSearch] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UniversalSearchResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UniversalSearch] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateCompoundFieldValue
 * @desc: 增量更新复合字段
 */
func (a *WorkItemService) UpdateCompoundFieldValue(ctx context.Context, req *UpdateCompoundFieldValueReq, options ...core.RequestOptionFunc) (*UpdateCompoundFieldValueResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateCompoundFieldValue
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateCompoundFieldValue] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateCompoundFieldValueResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateCompoundFieldValue] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateField
 * @desc: 更新自定义字段
 */
func (a *WorkItemService) UpdateField(ctx context.Context, req *UpdateFieldReq, options ...core.RequestOptionFunc) (*UpdateFieldResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateField
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateField] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateFieldResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateField] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateFinished
 * @desc: 批量查询工作项交付物详情信息
 */
func (a *WorkItemService) UpdateFinished(ctx context.Context, req *UpdateFinishedReq, options ...core.RequestOptionFunc) (*UpdateFinishedResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateFinished
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFinished] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateFinishedResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFinished] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateFlowRole
 * @desc: 更新流程角色
 */
func (a *WorkItemService) UpdateFlowRole(ctx context.Context, req *UpdateFlowRoleReq, options ...core.RequestOptionFunc) (*UpdateFlowRoleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateFlowRole
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFlowRole] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateFlowRoleResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateFlowRole] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateMultiSignal
 * @desc: 更新多值系统外信号
 */
func (a *WorkItemService) UpdateMultiSignal(ctx context.Context, req *UpdateMultiSignalReq, options ...core.RequestOptionFunc) (*UpdateMultiSignalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateMultiSignal
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateMultiSignal] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateMultiSignalResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateMultiSignal] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateNodeState
 * @desc: 节点完成/回滚
 */
func (a *WorkItemService) UpdateNodeState(ctx context.Context, req *UpdateNodeStateReq, options ...core.RequestOptionFunc) (*UpdateNodeStateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateNodeState
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateNodeState] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateNodeStateResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateNodeState] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateResourceWorkItem
 * @desc: 更新资源实例工作项
 */
func (a *WorkItemService) UpdateResourceWorkItem(ctx context.Context, req *UpdateResourceWorkItemReq, options ...core.RequestOptionFunc) (*UpdateResourceWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateResourceWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateResourceWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateResourceWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateResourceWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateStateFlow
 * @desc: 状态流转
 */
func (a *WorkItemService) UpdateStateFlow(ctx context.Context, req *UpdateStateFlowReq, options ...core.RequestOptionFunc) (*UpdateStateFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateStateFlow
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateStateFlow] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateStateFlowResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateStateFlow] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateTemplateDetail
 * @desc: 更新流程类型配置.
 */
func (a *WorkItemService) UpdateTemplateDetail(ctx context.Context, req *UpdateTemplateDetailReq, options ...core.RequestOptionFunc) (*UpdateTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateTemplateDetail
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItem
 * @desc: 更新工作项
 */
func (a *WorkItemService) UpdateWorkItem(ctx context.Context, req *UpdateWorkItemReq, options ...core.RequestOptionFunc) (*UpdateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItemRelation
 * @desc: 更新工作项关系
 */
func (a *WorkItemService) UpdateWorkItemRelation(ctx context.Context, req *UpdateWorkItemRelationReq, options ...core.RequestOptionFunc) (*UpdateWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkItemRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItemSubTask
 * @desc: 更新子任务
 */
func (a *WorkItemService) UpdateWorkItemSubTask(ctx context.Context, req *UpdateWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*UpdateWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkItemSubTask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItemTypeInfo
 * @desc: 更新工作项类型
 */
func (a *WorkItemService) UpdateWorkItemTypeInfo(ctx context.Context, req *UpdateWorkItemTypeInfoReq, options ...core.RequestOptionFunc) (*UpdateWorkItemTypeInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkItemTypeInfo
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemTypeInfo] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkItemTypeInfoResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkItemTypeInfo] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkflowNode
 * @desc: 更新节点
 */
func (a *WorkItemService) UpdateWorkflowNode(ctx context.Context, req *UpdateWorkflowNodeReq, options ...core.RequestOptionFunc) (*UpdateWorkflowNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkflowNode
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkflowNode] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkflowNodeResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkflowNode] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkingHourRecord
 * @desc: 更新实际工时
 */
func (a *WorkItemService) UpdateWorkingHourRecord(ctx context.Context, req *UpdateWorkingHourRecordReq, options ...core.RequestOptionFunc) (*UpdateWorkingHourRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateWorkingHourRecord
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkingHourRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateWorkingHourRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateWorkingHourRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIWBSUpdateDraftFrozenRows
 * @desc: 更新草稿的冻结行
 */
func (a *WorkItemService) WBSUpdateDraftFrozenRows(ctx context.Context, req *WBSUpdateDraftFrozenRowsReq, options ...core.RequestOptionFunc) (*WBSUpdateDraftFrozenRowsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_WBSUpdateDraftFrozenRows
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WBSUpdateDraftFrozenRows] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &WBSUpdateDraftFrozenRowsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WBSUpdateDraftFrozenRows] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIWbsCollaborationPublish
 * @desc: 计划表基于草稿发布
 */
func (a *WorkItemService) WbsCollaborationPublish(ctx context.Context, req *WbsCollaborationPublishReq, options ...core.RequestOptionFunc) (*WbsCollaborationPublishResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_WbsCollaborationPublish
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WbsCollaborationPublish] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &WbsCollaborationPublishResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[WbsCollaborationPublish] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
