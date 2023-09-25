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

package workitem_conf

import (
	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem"
)

type ConfirmForm struct {
	Action int64 `json:"action"`

	StateKey string `json:"state_key"`
}

type Connection struct {
	SourceStateKey string `json:"source_state_key"`

	TargetStateKey string `json:"target_state_key"`
}

type NodeConf struct {
	StateKey string `json:"state_key"`

	NodeName string `json:"node_name"`

	NodeTags []string `json:"node_tags"`

	NodeType string `json:"node_type"`

	IsVisibility int32 `json:"is_visibility"`

	NeedSchedule bool `json:"need_schedule"`

	Owner *OwnerConf `json:"owner"`

	WbsStatusMap *WbsStatusMap `json:"wbs_status_map"`

	NodeSubProcess *SubProcessConf `json:"node_sub_process"`

	WbsNodeMap *WbsNodeMap `json:"wbs_node_map"`
}

type OwnerConf struct {
	OwnerUsageMode string `json:"owner_usage_mode"`

	OwnerRoles []string `json:"owner_roles"`

	UserKeys []string `json:"user_keys"`
}

type StateFlowConfInfo struct {
	StateKey string `json:"state_key"`

	Name string `json:"name"`

	StateType int64 `json:"state_type"`

	AuthorizedRoles []string `json:"authorized_roles"`

	ConfirmFormList []*ConfirmForm `json:"confirm_form_list"`

	Action int64 `json:"action"`
}

type StatusConf struct {
	StatusKey string `json:"status_key"`

	StatusName string `json:"status_name"`

	StatusOrderIndex int32 `json:"status_order_index"`
}

type SubProcessConf struct {
	TemplateKey string `json:"template_key"`

	TemplateName string `json:"template_name"`

	TemplateID string `json:"template_id"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	FlowMode string `json:"flow_mode"`
}

type TaskConf struct {
	Action int64 `json:"action"`

	Name string `json:"name"`

	ID string `json:"id"`

	DeliverableFieldID string `json:"deliverable_field_id"`

	PassMode int64 `json:"pass_mode"`

	NodePassRequiredMode int64 `json:"node_pass_required_mode"`
}

type TemplateConf struct {
	TemplateID string `json:"template_id"`

	TemplateName string `json:"template_name"`

	IsDisabled int32 `json:"is_disabled"`

	Version int64 `json:"version"`

	UniqueKey string `json:"unique_key"`

	TemplateKey string `json:"template_key"`
}

type TemplateDetail struct {
	WorkflowConfs []*WorkflowConfInfo `json:"workflow_confs"`

	StateFlowConfs []*StateFlowConfInfo `json:"state_flow_confs"`

	Connections []*workitem.Connection `json:"connections"`

	TemplateID int64 `json:"template_id"`

	TemplateName string `json:"template_name"`

	Version int64 `json:"version"`

	IsDisabled int64 `json:"is_disabled"`
}

type WbsNodeMap struct {
	StateKey string `json:"state_key"`

	StatusName string `json:"status_name"`
}

type WbsStatusMap struct {
	StatusKey string `json:"status_key"`

	StatusName string `json:"status_name"`
}

type WbsTemplate struct {
	TemplateKey string `json:"template_key"`

	TemplateName string `json:"template_name"`

	TemplateID string `json:"template_id"`

	IsDisabled bool `json:"is_disabled"`

	Version int64 `json:"version"`

	WorkflowConf *WorkflowConf `json:"workflow_conf"`
}

type WorkflowConf struct {
	StatusInfos []*StatusConf `json:"status_infos"`

	NodeInfos []*NodeConf `json:"node_infos"`

	Connections []*Connection `json:"connections"`
}

type WorkflowConfInfo struct {
	StateKey string `json:"state_key"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	OwnerUsageMode int64 `json:"owner_usage_mode"`

	OwnerRoles []string `json:"owner_roles"`

	Owners []string `json:"owners"`

	NeedSchedule bool `json:"need_schedule"`

	DifferentSchedule bool `json:"different_schedule"`

	VisibilityUsageMode int64 `json:"visibility_usage_mode"`

	Deletable bool `json:"deletable"`

	DeletableOperationRole []string `json:"deletable_operation_role"`

	PassMode int64 `json:"pass_mode"`

	IsLimitNode bool `json:"is_limit_node"`

	DoneOperationRole []string `json:"done_operation_role"`

	DoneSchedule bool `json:"done_schedule"`

	DoneAllocateOwner bool `json:"done_allocate_owner"`

	Action int64 `json:"action"`

	PreNodeStateKey []string `json:"pre_node_state_key"`

	CompletionTips string `json:"completion_tips"`

	TaskConfs []*TaskConf `json:"task_confs"`
}
