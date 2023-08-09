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
	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem/v2"
)

type ConfirmForm struct {
	Action int64 `json:"action"`

	StateKey string `json:"state_key"`
}

type StateFlowConfInfo struct {
	StateKey string `json:"state_key"`

	Name string `json:"name"`

	StateType int64 `json:"state_type"`

	AuthorizedRoles []string `json:"authorized_roles"`

	ConfirmFormList []*ConfirmForm `json:"confirm_form_list"`

	Action int64 `json:"action"`
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
