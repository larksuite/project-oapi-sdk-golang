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

package workitem

import (
	"github.com/larksuite/project-oapi-sdk-golang/service/common"
	"github.com/larksuite/project-oapi-sdk-golang/service/field"
	"github.com/larksuite/project-oapi-sdk-golang/service/role_conf"
	"github.com/larksuite/project-oapi-sdk-golang/service/user"
)

type Checker struct {
	CheckedTime int64 `json:"checked_time"`

	Owner string `json:"owner"`

	Status int32 `json:"status"`
}

type CompInfo struct {
	ID string `json:"ID"`

	Name string `json:"name"`

	WorkItemTypeKey string `json:"WorkItemTypeKey"`

	ProjectKey string `json:"ProjectKey"`

	CreatedBy string `json:"CreatedBy"`

	CreatedAt int64 `json:"CreatedAt"`

	SearchHit []string `json:"SearchHit"`

	ViewScopeKey string `json:"ViewScopeKey"`

	ProjectKeys []string `json:"ProjectKeys"`
}

type Connection struct {
	SourceStateKey string `json:"source_state_key"`

	TargetStateKey string `json:"target_state_key"`

	TransitionID int64 `json:"transition_id"`
}

type CreateWorkItemRelationData struct {
	RelationID string `json:"relation_id"`
}

type CreateWorkingHourRecord struct {
	ResourceType string `json:"resource_type"`

	ResourceID string `json:"resource_id"`

	WorkTime string `json:"work_time"`

	WorkDescription string `json:"work_description"`
}

type DefaultValue struct {
	DefaultAppear int32 `json:"default_appear"`

	Value interface{} `json:"value"`
}

type Expand struct {
	NeedWorkflow bool `json:"need_workflow"`

	RelationFieldsDetail bool `json:"relation_fields_detail"`

	NeedMultiText bool `json:"need_multi_text"`

	NeedUserDetail bool `json:"need_user_detail"`

	NeedSubTaskParent bool `json:"need_sub_task_parent"`
}

type FieldConf struct {
	IsRequired int32 `json:"is_required"`

	IsVisibility int32 `json:"is_visibility"`

	RoleAssign []*user.RoleAssign `json:"role_assign"`

	FieldName string `json:"field_name"`

	FieldKey string `json:"field_key"`

	FieldAlias string `json:"field_alias"`

	FieldTypeKey string `json:"field_type_key"`

	DefaultValue *DefaultValue `json:"default_value"`

	Options []*OptionConf `json:"options"`

	CompoundFields []*FieldConf `json:"compound_fields"`

	IsValidity int32 `json:"is_validity"`

	Label string `json:"label"`
}

type FieldDetail struct {
	StoryID int64 `json:"story_id"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	ProjectKey string `json:"project_key"`
}

type ManHourRecord struct {
	ID int64 `json:"id"`

	RelatedWorkItemID int64 `json:"related_work_item_id"`

	RelatedWorkItemTypeKey string `json:"related_work_item_type_key"`

	RelatedWorkItemName string `json:"related_work_item_name"`

	ResourceType string `json:"resource_type"`

	ResourceID string `json:"resource_id"`

	WorkDescription string `json:"work_description"`

	WorkTime float64 `json:"work_time"`

	WorkUserKey string `json:"work_user_key"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`

	ResourceName string `json:"resource_name"`

	WorkDate int64 `json:"work_date"`
}

type MultiText struct {
	FieldKey string `json:"field_key"`

	FieldValue *MultiTextDetail `json:"field_value"`
}

type MultiTextDetail struct {
	Doc string `json:"doc"`

	DocText string `json:"doc_text"`

	IsEmpty bool `json:"is_empty"`

	NotifyUserList []string `json:"notify_user_list"`

	NotifyUserType string `json:"notify_user_type"`

	DocHTML string `json:"doc_html"`
}

type NodeBasicInfo struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Owners []string `json:"owners"`

	Milestone bool `json:"milestone"`
}

type NodeTask struct {
	ID string `json:"id"`

	StateKey string `json:"state_key"`

	SubTasks []*SubTask `json:"sub_tasks"`
}

type NodesConnections struct {
	WorkflowNodes []*WorkflowNode `json:"workflow_nodes"`

	Connections []*Connection `json:"connections"`

	StateFlowNodes []*StateFlowNode `json:"state_flow_nodes"`

	UserDetails []*common.UserDetail `json:"user_details"`
}

type OptionConf struct {
	Label string `json:"label"`

	Value string `json:"value"`

	IsDisabled int32 `json:"is_disabled"`

	IsVisibility int32 `json:"is_visibility"`

	Children []*OptionConf `json:"children"`
}

type RelationDetail struct {
	WorkItemTypeKey string `json:"work_item_type_key"`

	WorkItemTypeName string `json:"work_item_type_name"`

	ProjectKey string `json:"project_key"`

	ProjectName string `json:"project_name"`
}

type RelationFieldDetail struct {
	FieldKey string `json:"field_key"`

	Detail []*FieldDetail `json:"detail"`
}

type Schedule struct {
	Points float64 `json:"points"`

	EstimateStartDate int64 `json:"estimate_start_date"`

	EstimateEndDate int64 `json:"estimate_end_date"`

	Owners []string `json:"owners"`

	ActualWorkTime float64 `json:"actual_work_time"`
}

type SearchGroup struct {
	SearchParams []*SearchParam `json:"search_params"`

	Conjunction string `json:"conjunction"`

	SearchGroups []*SearchGroup `json:"search_groups"`
}

type SearchParam struct {
	ParamKey string `json:"param_key"`

	Value interface{} `json:"value"`

	Operator string `json:"operator"`
}

type SearchUser struct {
	UserKeys []string `json:"user_keys"`

	FieldKey string `json:"field_key"`

	Role string `json:"role"`
}

type StateFlowNode struct {
	ID string `json:"id"`

	Name string `json:"name"`

	RoleOwners []*user.RoleOwner `json:"role_owners"`

	Status int32 `json:"status"`

	ActualBeginTime string `json:"actual_begin_time"`

	ActualFinishTime string `json:"actual_finish_time"`

	Fields []*field.FieldValuePair `json:"fields"`
}

type StateTime struct {
	StateKey string `json:"state_key"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	Name string `json:"name"`
}

type SubTask struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Schedules []*Schedule `json:"schedules"`

	Order float64 `json:"order"`

	Details string `json:"details"`

	Passed bool `json:"passed"`

	Owners []string `json:"owners"`

	Note string `json:"note"`

	ActualBeginTime string `json:"actual_begin_time"`

	ActualFinishTime string `json:"actual_finish_time"`

	Assignee []string `json:"assignee"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`

	Deliverable []*field.FieldValuePair `json:"deliverable"`
}

type SubTaskParentInfo struct {
	WorkItemID int64 `json:"work_item_id"`

	WorkItemName string `json:"work_item_name"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	NodeID string `json:"node_id"`
}

type TimeInterval struct {
	Start int64 `json:"start"`

	End int64 `json:"end"`
}

type UpdateWorkingHourRecord struct {
	ID int64 `json:"id"`

	WorkTime string `json:"work_time"`

	WorkDescription string `json:"work_description"`
}

type WBSParentWorkItem struct {
	IsTop bool `json:"is_top"`

	WorkItemID int64 `json:"work_item_id"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	Name string `json:"name"`

	TemplateID string `json:"template_id"`

	TemplateKey string `json:"template_key"`

	TemplateName string `json:"template_name"`

	TemplateVersion int64 `json:"template_version"`

	RelationNodeID string `json:"relation_node_id"`

	RelationNodeName string `json:"relation_node_name"`

	RelationNodeTags []string `json:"relation_node_tags"`

	RelationNodeUUID string `json:"relation_node_uuid"`
}

type WBSWorkItem struct {
	NodeUUID string `json:"node_uuid"`

	WorkItemID int64 `json:"work_item_id"`

	Type string `json:"type"`

	WbsStatus string `json:"wbs_status"`

	WbsStatusMap map[string]string `json:"wbs_status_map"`

	SubWorkItem []*WBSWorkItem `json:"sub_work_item"`

	Name string `json:"name"`

	Deliverable []*field.FieldValuePair `json:"deliverable"`

	Schedule *Schedule `json:"schedule"`

	Schedules []*Schedule `json:"schedules"`

	Points float64 `json:"points"`

	RoleOwners []*user.RoleOwner `json:"role_owners"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	Milestone bool `json:"milestone"`
}

type WbsViewResponse struct {
	TemplateKey string `json:"template_key"`

	RelatedSubWorkItems []*WBSWorkItem `json:"related_sub_work_items"`

	TemplateVersion int64 `json:"template_version"`

	TemplateName string `json:"template_name"`

	TemplateID string `json:"template_id"`

	RelatedParentWorkItem *WBSParentWorkItem `json:"related_parent_work_item"`

	UserDetails []*common.UserDetail `json:"user_details"`
}

type WorkItemInfo struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	ProjectKey string `json:"project_key"`

	TemplateType string `json:"template_type"`

	Pattern string `json:"pattern"`

	SubStage string `json:"sub_stage"`

	CurrentNodes []*NodeBasicInfo `json:"current_nodes"`

	CreatedBy string `json:"created_by"`

	UpdatedBy string `json:"updated_by"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`

	Fields []*field.FieldValuePair `json:"fields"`

	WorkItemStatus *WorkItemStatus `json:"work_item_status"`

	DeletedBy string `json:"deleted_by"`

	DeletedAt int64 `json:"deleted_at"`

	SimpleName string `json:"simple_name"`

	TemplateID int64 `json:"template_id"`

	WorkflowInfos *NodesConnections `json:"workflow_infos"`

	StateTimes []*StateTime `json:"state_times"`

	MultiTexts []*MultiText `json:"multi_texts"`

	RelationFieldsDetail []*RelationFieldDetail `json:"relation_fields_detail"`

	UserDetails []*common.UserDetail `json:"user_details"`

	SubTaskParentInfo *SubTaskParentInfo `json:"sub_task_parent_info"`
}

type WorkItemKeyType struct {
	TypeKey string `json:"type_key"`

	Name string `json:"name"`

	IsDisable int32 `json:"is_disable"`

	APIName string `json:"api_name"`
}

type WorkItemRelation struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Disabled bool `json:"disabled"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	WorkItemTypeName string `json:"work_item_type_name"`

	RelationDetails []*RelationDetail `json:"relation_details"`
}

type WorkItemStatus struct {
	StateKey string `json:"state_key"`

	IsArchivedState bool `json:"is_archived_state"`

	IsInitState bool `json:"is_init_state"`

	UpdatedAt int64 `json:"updated_at"`

	UpdatedBy string `json:"updated_by"`

	History []*WorkItemStatus `json:"history"`
}

type WorkItemTypeInfo struct {
	TypeKey string `json:"type_key"`

	Name string `json:"name"`

	FlowMode string `json:"flow_mode"`

	APIName string `json:"api_name"`

	Description string `json:"description"`

	IsDisabled bool `json:"is_disabled"`

	IsPinned bool `json:"is_pinned"`

	EnableSchedule bool `json:"enable_schedule"`

	ScheduleFieldKey string `json:"schedule_field_key"`

	ScheduleFieldName string `json:"schedule_field_name"`

	EstimatePointFieldKey string `json:"estimate_point_field_key"`

	EstimatePointFieldName string `json:"estimate_point_field_name"`

	ActualWorkTimeFieldKey string `json:"actual_work_time_field_key"`

	ActualWorkTimeFieldName string `json:"actual_work_time_field_name"`

	BelongRoles []*role_conf.SimpleRoleConf `json:"belong_roles"`
}

type WorkflowNode struct {
	ID string `json:"id"`

	StateKey string `json:"state_key"`

	Name string `json:"name"`

	Status int32 `json:"status"`

	Fields []*field.FieldValuePair `json:"fields"`

	Owners []string `json:"owners"`

	NodeSchedule *Schedule `json:"node_schedule"`

	Schedules []*Schedule `json:"schedules"`

	SubTasks []*SubTask `json:"sub_tasks"`

	ActualBeginTime string `json:"actual_begin_time"`

	ActualFinishTime string `json:"actual_finish_time"`

	RoleAssignee []*user.RoleOwner `json:"role_assignee"`

	DifferentSchedule bool `json:"different_schedule"`

	SubStatus []*Checker `json:"sub_status"`

	Milestone bool `json:"milestone"`

	Participants []string `json:"participants"`
}
