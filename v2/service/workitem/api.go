package workitem

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_Filter = "/open_api/:project_key/work_item/filter"

const APIPath_FilterAcrossProject = "/open_api/work_items/filter_across_project"

const APIPath_IntegrateSearch = "/open_api/view_search/integrate_search"

const APIPath_OAPIAbortWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/abort"

const APIPath_OAPIActualTimeUpdate = "/open_api/work_item/actual_time/update"

const APIPath_OAPIBatchQueryConclusionOption = "/open_api/work_item/finished/query_conclusion_option"

const APIPath_OAPIBatchQueryDeliverable = "/open_api/work_item/deliverable/batch_query"

const APIPath_OAPIBatchQueryFinished = "/open_api/work_item/finished/batch_query"

const APIPath_OAPIBatchUpdateBasicWorkItem = "/open_api/work_item/batch_update"

const APIPath_OAPICompositiveSearch = "/open_api/compositive_search"

const APIPath_OAPICreateField = "/open_api/:project_key/field/:work_item_type_key/create"

const APIPath_OAPICreateProjectRelationInstances = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/batch_bind"

const APIPath_OAPICreateStoryRelations = "/open_api/:project_key/story_relations/create"

const APIPath_OAPICreateTemplateDetail = "/open_api/template/v2/create_template"

const APIPath_OAPICreateWorkItem = "/open_api/:project_key/work_item/create"

const APIPath_OAPICreateWorkItemRelation = "/open_api/work_item/relation/create"

const APIPath_OAPICreateWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/task"

const APIPath_OAPICreateWorkingHourRecord = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/work_hour_record"

const APIPath_OAPIDeleteFile = "/open_api/file/delete"

const APIPath_OAPIDeleteProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id"

const APIPath_OAPIDeleteTemplateDetail = "/open_api/template/v2/delete_template/:project_key/:template_id"

const APIPath_OAPIDeleteWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id"

const APIPath_OAPIDeleteWorkItemRelation = "/open_api/work_item/relation/delete"

const APIPath_OAPIDeleteWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/task/:task_id"

const APIPath_OAPIDeleteWorkingHourRecord = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/work_hour_record"

const APIPath_OAPIElementTemplateCreate = "/open_api/work_item/element_template/create"

const APIPath_OAPIElementTemplateQuery = "/open_api/work_item/element_template/query"

const APIPath_OAPIFreezeWorkItem = "/open_api/work_item/freeze"

const APIPath_OAPIGetWBSInfo = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/wbs_view"

const APIPath_OAPIGetWBSViewSubWorkItemConf = "/open_api/work_item/wbs_view_draft/sub_work_item_conf"

const APIPath_OAPIGetWorkFlow = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/query"

const APIPath_OAPIGetWorkItemManHourRecords = "/open_api/work_item/man_hour/records"

const APIPath_OAPIGetWorkItemTypeInfoByKey = "/open_api/:project_key/work_item/type/:work_item_type_key"

const APIPath_OAPIGetWorkItemsByIds = "/open_api/:project_key/work_item/:work_item_type_key/query"

const APIPath_OAPIListTemplateConf = "/open_api/:project_key/template_list/:work_item_type_key"

const APIPath_OAPIPatchWBSViewDraft = "/open_api/work_item/wbs_view_draft/patch"

const APIPath_OAPIPublishWBSViewDraft = "/open_api/work_item/wbs_view_draft/publish"

const APIPath_OAPIQueryAWorkItemTypes = "/open_api/:project_key/work_item/all-types"

const APIPath_OAPIQueryBusinesses = "/open_api/:project_key/business/all"

const APIPath_OAPIQueryProjectFields = "/open_api/:project_key/field/all"

const APIPath_OAPIQueryProjectRelation = "/open_api/:project_key/relation/rules"

const APIPath_OAPIQueryProjectRelationInstance = "/open_api/:project_key/relation/:work_item_type_key/:work_item_id/work_item_list"

const APIPath_OAPIQueryRoleConfDetails = "/open_api/:project_key/flow_roles/:work_item_type_key"

const APIPath_OAPIQueryStoryRelations = "/open_api/:project_key/story_relations/query"

const APIPath_OAPIQueryTemplateDetail = "/open_api/:project_key/template_detail/:template_id"

const APIPath_OAPIQueryWBSViewDraft = "/open_api/work_item/wbs_view_draft/query"

const APIPath_OAPIQueryWbsTemplateConf = "/open_api/:project_key/wbs_template"

const APIPath_OAPIQueryWorkItemMetaData = "/open_api/:project_key/work_item/:work_item_type_key/meta"

const APIPath_OAPIQueryWorkItemRelation = "/open_api/:project_key/work_item/relation"

const APIPath_OAPIQueryWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/task"

const APIPath_OAPIResetWBSViewDraft = "/open_api/work_item/wbs_view_draft/reset"

const APIPath_OAPIResourceCreateWorkItem = "/open_api/work_item/resource/create_work_item"

const APIPath_OAPISearchWorkItemsRelation = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/search_by_relation"

const APIPath_OAPISubTaskModify = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/subtask/modify"

const APIPath_OAPISwitchBackToWbsViewDraft = "/open_api/:project_key/wbs_view_draft/switch"

const APIPath_OAPIUpdateField = "/open_api/:project_key/field/:work_item_type_key"

const APIPath_OAPIUpdateFinished = "/open_api/work_item/finished/update"

const APIPath_OAPIUpdateMultiSignal = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/update/multi_signal"

const APIPath_OAPIUpdateNodeState = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/:node_id/operate"

const APIPath_OAPIUpdateStateFlow = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/state_change"

const APIPath_OAPIUpdateTemplateDetail = "/open_api/template/v2/update_template"

const APIPath_OAPIUpdateWorkItem = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id"

const APIPath_OAPIUpdateWorkItemRelation = "/open_api/work_item/relation/update"

const APIPath_OAPIUpdateWorkItemSubTask = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/workflow/:node_id/task/:task_id"

const APIPath_OAPIUpdateWorkItemTypeInfo = "/open_api/:project_key/work_item/type/:work_item_type_key"

const APIPath_OAPIUpdateWorkflowNode = "/open_api/:project_key/workflow/:work_item_type_key/:work_item_id/node/:node_id"

const APIPath_OAPIUpdateWorkingHourRecord = "/open_api/:project_key/work_item/:work_item_type_key/:work_item_id/work_hour_record"

const APIPath_OAPIWBSUpdateDraftFrozenRows = "/open_api/:project_key/wbs_view_draft/update/frozen_rows"

const APIPath_OAPIWbsCollaborationPublish = "/open_api/:project_key/wbs_view_draft/publish"

const APIPath_SearchByParams = "/open_api/:project_key/work_item/:work_item_type_key/search/params"

const APIPath_UniversalSearch = "/open_api/view_search/universal_search"


func NewService(config *core.Config) *WorkItemService {
	a := &WorkItemService{config: config}
	return a
}

type WorkItemService struct {
	config *core.Config
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
 * @name: OAPIAbortWorkItem
 * @desc: 终止/恢复工作项
 */
func (a *WorkItemService) OAPIAbortWorkItem(ctx context.Context, req *OAPIAbortWorkItemReq, options ...core.RequestOptionFunc) (*OAPIAbortWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIAbortWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIAbortWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIAbortWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIAbortWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIActualTimeUpdate
 * @desc: 更新节点实际开始时间和结束时间
 */
func (a *WorkItemService) OAPIActualTimeUpdate(ctx context.Context, req *OAPIActualTimeUpdateReq, options ...core.RequestOptionFunc) (*OAPIActualTimeUpdateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIActualTimeUpdate
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIActualTimeUpdate] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIActualTimeUpdateResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIActualTimeUpdate] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIBatchQueryConclusionLabel
 * @desc: 批量查询工作项交付物详情信息
 */
func (a *WorkItemService) OAPIBatchQueryConclusionOption(ctx context.Context, req *OAPIBatchQueryConclusionOptionReq, options ...core.RequestOptionFunc) (*OAPIBatchQueryConclusionOptionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIBatchQueryConclusionOption
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryConclusionOption] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIBatchQueryConclusionOptionResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryConclusionOption] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * 
 */
func (a *WorkItemService) OAPIBatchQueryDeliverable(ctx context.Context, req *OAPIBatchQueryDeliverableReq, options ...core.RequestOptionFunc) (*OAPIBatchQueryDeliverableResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIBatchQueryDeliverable
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryDeliverable] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIBatchQueryDeliverableResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryDeliverable] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIBatchQueryFinished
 * @desc: 批量查询工作项交付物详情信息
 */
func (a *WorkItemService) OAPIBatchQueryFinished(ctx context.Context, req *OAPIBatchQueryFinishedReq, options ...core.RequestOptionFunc) (*OAPIBatchQueryFinishedResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIBatchQueryFinished
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryFinished] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIBatchQueryFinishedResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchQueryFinished] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIBatchUpdateBasicWorkItem
 * @desc: 批量更新工作项字段-异步
 */
func (a *WorkItemService) OAPIBatchUpdateBasicWorkItem(ctx context.Context, req *OAPIBatchUpdateBasicWorkItemReq, options ...core.RequestOptionFunc) (*OAPIBatchUpdateBasicWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIBatchUpdateBasicWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchUpdateBasicWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIBatchUpdateBasicWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIBatchUpdateBasicWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: openapi获取指定的工作项列表（全文搜索）
 * @desc: openapi获取指定的工作项列表（全文搜索）
 */
func (a *WorkItemService) OAPICompositiveSearch(ctx context.Context, req *OAPICompositiveSearchReq, options ...core.RequestOptionFunc) (*OAPICompositiveSearchResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICompositiveSearch
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICompositiveSearch] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICompositiveSearchResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICompositiveSearch] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateField
 * @desc: 创建自定义字段
 */
func (a *WorkItemService) OAPICreateField(ctx context.Context, req *OAPICreateFieldReq, options ...core.RequestOptionFunc) (*OAPICreateFieldResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateField
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateField] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateFieldResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateField] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateProjectRelationInstances
 * @desc: 通过空间关联绑定关联工作项
 */
func (a *WorkItemService) OAPICreateProjectRelationInstances(ctx context.Context, req *OAPICreateProjectRelationInstancesReq, options ...core.RequestOptionFunc) (*OAPICreateProjectRelationInstancesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateProjectRelationInstances
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateProjectRelationInstances] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateProjectRelationInstancesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateProjectRelationInstances] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateStoryRelations
 * @desc: 创建需求关联关系
 */
func (a *WorkItemService) OAPICreateStoryRelations(ctx context.Context, req *OAPICreateStoryRelationsReq, options ...core.RequestOptionFunc) (*OAPICreateStoryRelationsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateStoryRelations
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateStoryRelations] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateStoryRelationsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateStoryRelations] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateTemplateDetail
 * @desc: 新增流程类型配置
 */
func (a *WorkItemService) OAPICreateTemplateDetail(ctx context.Context, req *OAPICreateTemplateDetailReq, options ...core.RequestOptionFunc) (*OAPICreateTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateTemplateDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkItem
 * @desc: 创建工作项
 */
func (a *WorkItemService) OAPICreateWorkItem(ctx context.Context, req *OAPICreateWorkItemReq, options ...core.RequestOptionFunc) (*OAPICreateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkItemRelation
 * @desc: 创建工作项关系
 */
func (a *WorkItemService) OAPICreateWorkItemRelation(ctx context.Context, req *OAPICreateWorkItemRelationReq, options ...core.RequestOptionFunc) (*OAPICreateWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateWorkItemRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkItemSubTask
 * @desc: 创建子任务
 */
func (a *WorkItemService) OAPICreateWorkItemSubTask(ctx context.Context, req *OAPICreateWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*OAPICreateWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateWorkItemSubTask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateWorkingHourRecord
 * @desc: 创建实际工时
 */
func (a *WorkItemService) OAPICreateWorkingHourRecord(ctx context.Context, req *OAPICreateWorkingHourRecordReq, options ...core.RequestOptionFunc) (*OAPICreateWorkingHourRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateWorkingHourRecord
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkingHourRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateWorkingHourRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateWorkingHourRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteFile
 * @desc: 删除实例对应的文件
 */
func (a *WorkItemService) OAPIDeleteFile(ctx context.Context, req *OAPIDeleteFileReq, options ...core.RequestOptionFunc) (*OAPIDeleteFileResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteFile
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteFile] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteFileResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteFile] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteProjectRelationInstance
 * @desc: 通过空间关联解绑关联工作项
 */
func (a *WorkItemService) OAPIDeleteProjectRelationInstance(ctx context.Context, req *OAPIDeleteProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*OAPIDeleteProjectRelationInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteProjectRelationInstance
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteProjectRelationInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteTemplateDetail
 * @desc: 删除流程类型配置
 */
func (a *WorkItemService) OAPIDeleteTemplateDetail(ctx context.Context, req *OAPIDeleteTemplateDetailReq, options ...core.RequestOptionFunc) (*OAPIDeleteTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteTemplateDetail
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkItem
 * @desc: 删除工作项
 */
func (a *WorkItemService) OAPIDeleteWorkItem(ctx context.Context, req *OAPIDeleteWorkItemReq, options ...core.RequestOptionFunc) (*OAPIDeleteWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteWorkItem
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkItemRelation
 * @desc: 删除工作项关系
 */
func (a *WorkItemService) OAPIDeleteWorkItemRelation(ctx context.Context, req *OAPIDeleteWorkItemRelationReq, options ...core.RequestOptionFunc) (*OAPIDeleteWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteWorkItemRelation
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkItemSubTask
 * @desc: 删除子任务
 */
func (a *WorkItemService) OAPIDeleteWorkItemSubTask(ctx context.Context, req *OAPIDeleteWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*OAPIDeleteWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteWorkItemSubTask
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteWorkingHourRecord
 * @desc: 删除实际工时
 */
func (a *WorkItemService) OAPIDeleteWorkingHourRecord(ctx context.Context, req *OAPIDeleteWorkingHourRecordReq, options ...core.RequestOptionFunc) (*OAPIDeleteWorkingHourRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteWorkingHourRecord
	apiReq.HttpMethod = "DELETE"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkingHourRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteWorkingHourRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteWorkingHourRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemTemplate
 * @desc: 创建资源库资源(节点/任务)
 */
func (a *WorkItemService) OAPIElementTemplateCreate(ctx context.Context, req *OAPIElementTemplateCreateReq, options ...core.RequestOptionFunc) (*OAPIElementTemplateCreateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIElementTemplateCreate
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIElementTemplateCreate] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIElementTemplateCreateResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIElementTemplateCreate] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIElementTemplateQueryAll
 * @desc: 查询资源库资源(节点/任务)
 */
func (a *WorkItemService) OAPIElementTemplateQuery(ctx context.Context, req *OAPIElementTemplateQueryReq, options ...core.RequestOptionFunc) (*OAPIElementTemplateQueryResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIElementTemplateQuery
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIElementTemplateQuery] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIElementTemplateQueryResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIElementTemplateQuery] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIFreezeWorkItem
 * @desc: 冻结/终止工作项
 */
func (a *WorkItemService) OAPIFreezeWorkItem(ctx context.Context, req *OAPIFreezeWorkItemReq, options ...core.RequestOptionFunc) (*OAPIFreezeWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIFreezeWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIFreezeWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIFreezeWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIFreezeWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWBSInfo
 * @desc: 获取WBS
 */
func (a *WorkItemService) OAPIGetWBSInfo(ctx context.Context, req *OAPIGetWBSInfoReq, options ...core.RequestOptionFunc) (*OAPIGetWBSInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIGetWBSInfo
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWBSInfo] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIGetWBSInfoResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWBSInfo] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIActualTimeUpdate
 * @desc: 更新节点实际开始时间和结束时间
 */
func (a *WorkItemService) OAPIGetWBSViewSubWorkItemConf(ctx context.Context, req *OAPIGetWBSViewSubWorkItemConfReq, options ...core.RequestOptionFunc) (*OAPIGetWBSViewSubWorkItemConfResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIGetWBSViewSubWorkItemConf
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWBSViewSubWorkItemConf] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIGetWBSViewSubWorkItemConfResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWBSViewSubWorkItemConf] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkFlow
 * @desc: 获取工作流
 */
func (a *WorkItemService) OAPIGetWorkFlow(ctx context.Context, req *OAPIGetWorkFlowReq, options ...core.RequestOptionFunc) (*OAPIGetWorkFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIGetWorkFlow
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkFlow] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIGetWorkFlowResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkFlow] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemManHourRecords
 * @desc: 获取工作项的工时记录列表
 */
func (a *WorkItemService) OAPIGetWorkItemManHourRecords(ctx context.Context, req *OAPIGetWorkItemManHourRecordsReq, options ...core.RequestOptionFunc) (*OAPIGetWorkItemManHourRecordsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIGetWorkItemManHourRecords
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkItemManHourRecords] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIGetWorkItemManHourRecordsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkItemManHourRecords] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemTypeInfoByKey
 * @desc: 查询工作项类型
 */
func (a *WorkItemService) OAPIGetWorkItemTypeInfoByKey(ctx context.Context, req *OAPIGetWorkItemTypeInfoByKeyReq, options ...core.RequestOptionFunc) (*OAPIGetWorkItemTypeInfoByKeyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIGetWorkItemTypeInfoByKey
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkItemTypeInfoByKey] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIGetWorkItemTypeInfoByKeyResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkItemTypeInfoByKey] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIGetWorkItemsByIds
 * @desc: 批量查询工作项
 */
func (a *WorkItemService) OAPIGetWorkItemsByIds(ctx context.Context, req *OAPIGetWorkItemsByIdsReq, options ...core.RequestOptionFunc) (*OAPIGetWorkItemsByIdsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIGetWorkItemsByIds
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkItemsByIds] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIGetWorkItemsByIdsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIGetWorkItemsByIds] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIListTemplateConf
 * @desc: 获取工作项下的流程类型
 */
func (a *WorkItemService) OAPIListTemplateConf(ctx context.Context, req *OAPIListTemplateConfReq, options ...core.RequestOptionFunc) (*OAPIListTemplateConfResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIListTemplateConf
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIListTemplateConf] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIListTemplateConfResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIListTemplateConf] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: wbs草稿态编辑
 * @desc: wbs草稿态编辑
 */
func (a *WorkItemService) OAPIPatchWBSViewDraft(ctx context.Context, req *OAPIPatchWBSViewDraftReq, options ...core.RequestOptionFunc) (*OAPIPatchWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIPatchWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPatchWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIPatchWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPatchWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIWbsCollaborationPublish
 * @desc: 计划表基于草稿发布,包含插件逻辑
 */
func (a *WorkItemService) OAPIPublishWBSViewDraft(ctx context.Context, req *OAPIPublishWBSViewDraftReq, options ...core.RequestOptionFunc) (*OAPIPublishWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIPublishWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPublishWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIPublishWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPublishWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryAWorkItemTypes
 * @desc: 获取空间下工作项类型
 */
func (a *WorkItemService) OAPIQueryAWorkItemTypes(ctx context.Context, req *OAPIQueryAWorkItemTypesReq, options ...core.RequestOptionFunc) (*OAPIQueryAWorkItemTypesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryAWorkItemTypes
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryAWorkItemTypes] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryAWorkItemTypesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryAWorkItemTypes] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryBusinesses
 * @desc: 获取空间下业务线
 */
func (a *WorkItemService) OAPIQueryBusinesses(ctx context.Context, req *OAPIQueryBusinessesReq, options ...core.RequestOptionFunc) (*OAPIQueryBusinessesResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryBusinesses
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryBusinesses] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryBusinessesResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryBusinesses] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryProjectFields
 * @desc: 获取空间字段
 */
func (a *WorkItemService) OAPIQueryProjectFields(ctx context.Context, req *OAPIQueryProjectFieldsReq, options ...core.RequestOptionFunc) (*OAPIQueryProjectFieldsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryProjectFields
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjectFields] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryProjectFieldsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjectFields] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryProjectRelation
 * @desc: 查询空间关系
 */
func (a *WorkItemService) OAPIQueryProjectRelation(ctx context.Context, req *OAPIQueryProjectRelationReq, options ...core.RequestOptionFunc) (*OAPIQueryProjectRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryProjectRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjectRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryProjectRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjectRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryProjectRelationInstance
 * @desc: 获取空间关联下的关联工作项列表
 */
func (a *WorkItemService) OAPIQueryProjectRelationInstance(ctx context.Context, req *OAPIQueryProjectRelationInstanceReq, options ...core.RequestOptionFunc) (*OAPIQueryProjectRelationInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryProjectRelationInstance
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjectRelationInstance] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryProjectRelationInstanceResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryProjectRelationInstance] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryRoleConfDetails
 * @desc: 获取流程角色配置详情
 */
func (a *WorkItemService) OAPIQueryRoleConfDetails(ctx context.Context, req *OAPIQueryRoleConfDetailsReq, options ...core.RequestOptionFunc) (*OAPIQueryRoleConfDetailsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryRoleConfDetails
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryRoleConfDetails] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryRoleConfDetailsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryRoleConfDetails] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryStoryRelations
 * @desc: 查询需求关联关系
 */
func (a *WorkItemService) OAPIQueryStoryRelations(ctx context.Context, req *OAPIQueryStoryRelationsReq, options ...core.RequestOptionFunc) (*OAPIQueryStoryRelationsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryStoryRelations
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryStoryRelations] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryStoryRelationsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryStoryRelations] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryTemplateDetail
 * @desc: 获取流程类型配置详情
 */
func (a *WorkItemService) OAPIQueryTemplateDetail(ctx context.Context, req *OAPIQueryTemplateDetailReq, options ...core.RequestOptionFunc) (*OAPIQueryTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryTemplateDetail
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIInitWBSViewDraft
 * @desc: 查询计划表草稿
 */
func (a *WorkItemService) OAPIQueryWBSViewDraft(ctx context.Context, req *OAPIQueryWBSViewDraftReq, options ...core.RequestOptionFunc) (*OAPIQueryWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWbsTemplateConf
 * @desc: 获取流程类型配置（wbs）
 */
func (a *WorkItemService) OAPIQueryWbsTemplateConf(ctx context.Context, req *OAPIQueryWbsTemplateConfReq, options ...core.RequestOptionFunc) (*OAPIQueryWbsTemplateConfResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryWbsTemplateConf
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWbsTemplateConf] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryWbsTemplateConfResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWbsTemplateConf] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemMetaData
 * @desc: 获取创建工作项元信息
 */
func (a *WorkItemService) OAPIQueryWorkItemMetaData(ctx context.Context, req *OAPIQueryWorkItemMetaDataReq, options ...core.RequestOptionFunc) (*OAPIQueryWorkItemMetaDataResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryWorkItemMetaData
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWorkItemMetaData] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryWorkItemMetaDataResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWorkItemMetaData] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemRelation
 * @desc: 查询工作项关系
 */
func (a *WorkItemService) OAPIQueryWorkItemRelation(ctx context.Context, req *OAPIQueryWorkItemRelationReq, options ...core.RequestOptionFunc) (*OAPIQueryWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryWorkItemRelation
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryWorkItemSubTask
 * @desc: 获取子任务详情
 */
func (a *WorkItemService) OAPIQueryWorkItemSubTask(ctx context.Context, req *OAPIQueryWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*OAPIQueryWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryWorkItemSubTask
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIInitWBSViewDraft
 * @desc: 重置计划表草稿
 */
func (a *WorkItemService) OAPIResetWBSViewDraft(ctx context.Context, req *OAPIResetWBSViewDraftReq, options ...core.RequestOptionFunc) (*OAPIResetWBSViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIResetWBSViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIResetWBSViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIResetWBSViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIResetWBSViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIResourceCreateWorkItem
 * @desc: 创建工作项资源库
 */
func (a *WorkItemService) OAPIResourceCreateWorkItem(ctx context.Context, req *OAPIResourceCreateWorkItemReq, options ...core.RequestOptionFunc) (*OAPIResourceCreateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIResourceCreateWorkItem
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIResourceCreateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIResourceCreateWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIResourceCreateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISearchWorkItemsRelation
 * @desc: 获取指定的关联工作项列表（单空间）
 */
func (a *WorkItemService) OAPISearchWorkItemsRelation(ctx context.Context, req *OAPISearchWorkItemsRelationReq, options ...core.RequestOptionFunc) (*OAPISearchWorkItemsRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPISearchWorkItemsRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISearchWorkItemsRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPISearchWorkItemsRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISearchWorkItemsRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISubTaskModify
 * @desc: 子任务完成/回滚
 */
func (a *WorkItemService) OAPISubTaskModify(ctx context.Context, req *OAPISubTaskModifyReq, options ...core.RequestOptionFunc) (*OAPISubTaskModifyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPISubTaskModify
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISubTaskModify] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPISubTaskModifyResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISubTaskModify] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISwitchBackToWbsViewDraft
 * @desc: 审批拒绝切回草稿
 */
func (a *WorkItemService) OAPISwitchBackToWbsViewDraft(ctx context.Context, req *OAPISwitchBackToWbsViewDraftReq, options ...core.RequestOptionFunc) (*OAPISwitchBackToWbsViewDraftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPISwitchBackToWbsViewDraft
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISwitchBackToWbsViewDraft] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPISwitchBackToWbsViewDraftResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISwitchBackToWbsViewDraft] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateField
 * @desc: 更新自定义字段
 */
func (a *WorkItemService) OAPIUpdateField(ctx context.Context, req *OAPIUpdateFieldReq, options ...core.RequestOptionFunc) (*OAPIUpdateFieldResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateField
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateField] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateFieldResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateField] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateFinished
 * @desc: 批量查询工作项交付物详情信息
 */
func (a *WorkItemService) OAPIUpdateFinished(ctx context.Context, req *OAPIUpdateFinishedReq, options ...core.RequestOptionFunc) (*OAPIUpdateFinishedResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateFinished
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateFinished] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateFinishedResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateFinished] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateMultiSignal
 * @desc: 更新多值系统外信号
 */
func (a *WorkItemService) OAPIUpdateMultiSignal(ctx context.Context, req *OAPIUpdateMultiSignalReq, options ...core.RequestOptionFunc) (*OAPIUpdateMultiSignalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateMultiSignal
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateMultiSignal] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateMultiSignalResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateMultiSignal] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateNodeState
 * @desc: 节点完成/回滚
 */
func (a *WorkItemService) OAPIUpdateNodeState(ctx context.Context, req *OAPIUpdateNodeStateReq, options ...core.RequestOptionFunc) (*OAPIUpdateNodeStateResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateNodeState
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateNodeState] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateNodeStateResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateNodeState] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateStateFlow
 * @desc: 状态流转
 */
func (a *WorkItemService) OAPIUpdateStateFlow(ctx context.Context, req *OAPIUpdateStateFlowReq, options ...core.RequestOptionFunc) (*OAPIUpdateStateFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateStateFlow
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateStateFlow] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateStateFlowResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateStateFlow] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateTemplateDetail
 * @desc: 更新流程类型配置.
 */
func (a *WorkItemService) OAPIUpdateTemplateDetail(ctx context.Context, req *OAPIUpdateTemplateDetailReq, options ...core.RequestOptionFunc) (*OAPIUpdateTemplateDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateTemplateDetail
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateTemplateDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateTemplateDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateTemplateDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItem
 * @desc: 更新工作项
 */
func (a *WorkItemService) OAPIUpdateWorkItem(ctx context.Context, req *OAPIUpdateWorkItemReq, options ...core.RequestOptionFunc) (*OAPIUpdateWorkItemResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateWorkItem
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItem] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateWorkItemResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItem] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItemRelation
 * @desc: 更新工作项关系
 */
func (a *WorkItemService) OAPIUpdateWorkItemRelation(ctx context.Context, req *OAPIUpdateWorkItemRelationReq, options ...core.RequestOptionFunc) (*OAPIUpdateWorkItemRelationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateWorkItemRelation
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItemRelation] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateWorkItemRelationResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItemRelation] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItemSubTask
 * @desc: 更新子任务
 */
func (a *WorkItemService) OAPIUpdateWorkItemSubTask(ctx context.Context, req *OAPIUpdateWorkItemSubTaskReq, options ...core.RequestOptionFunc) (*OAPIUpdateWorkItemSubTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateWorkItemSubTask
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItemSubTask] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateWorkItemSubTaskResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItemSubTask] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkItemTypeInfo
 * @desc: 更新工作项类型
 */
func (a *WorkItemService) OAPIUpdateWorkItemTypeInfo(ctx context.Context, req *OAPIUpdateWorkItemTypeInfoReq, options ...core.RequestOptionFunc) (*OAPIUpdateWorkItemTypeInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateWorkItemTypeInfo
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItemTypeInfo] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateWorkItemTypeInfoResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkItemTypeInfo] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkflowNode
 * @desc: 更新节点
 */
func (a *WorkItemService) OAPIUpdateWorkflowNode(ctx context.Context, req *OAPIUpdateWorkflowNodeReq, options ...core.RequestOptionFunc) (*OAPIUpdateWorkflowNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateWorkflowNode
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkflowNode] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateWorkflowNodeResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkflowNode] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateWorkingHourRecord
 * @desc: 更新实际工时
 */
func (a *WorkItemService) OAPIUpdateWorkingHourRecord(ctx context.Context, req *OAPIUpdateWorkingHourRecordReq, options ...core.RequestOptionFunc) (*OAPIUpdateWorkingHourRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateWorkingHourRecord
	apiReq.HttpMethod = "PUT"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkingHourRecord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateWorkingHourRecordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateWorkingHourRecord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIWBSUpdateDraftFrozenRows
 * @desc: 更新草稿的冻结行
 */
func (a *WorkItemService) OAPIWBSUpdateDraftFrozenRows(ctx context.Context, req *OAPIWBSUpdateDraftFrozenRowsReq, options ...core.RequestOptionFunc) (*OAPIWBSUpdateDraftFrozenRowsResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIWBSUpdateDraftFrozenRows
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIWBSUpdateDraftFrozenRows] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIWBSUpdateDraftFrozenRowsResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIWBSUpdateDraftFrozenRows] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIWbsCollaborationPublish
 * @desc: 计划表基于草稿发布
 */
func (a *WorkItemService) OAPIWbsCollaborationPublish(ctx context.Context, req *OAPIWbsCollaborationPublishReq, options ...core.RequestOptionFunc) (*OAPIWbsCollaborationPublishResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIWbsCollaborationPublish
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIWbsCollaborationPublish] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIWbsCollaborationPublishResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIWbsCollaborationPublish] fail to unmarshal response body, error: %v", err.Error()))
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

