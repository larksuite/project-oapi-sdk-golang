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

package project_relation

type ProjectRelationRule struct {
	RemoteProjectKey string `json:"remote_project_key"`

	RemoteProjectName string `json:"remote_project_name"`

	Rules []*RelationRule `json:"rules"`
}

type RelationBindInstance struct {
	ProjectKey string `json:"project_key"`

	WorkItemTypeKey string `json:"work_item_type_key"`

	WorkItemID int64 `json:"work_item_id"`

	ChatGroupMerge int64 `json:"chat_group_merge"`
}

type RelationInstance struct {
	RelationWorkItemID int64 `json:"relation_work_item_id"`

	RelationWorkItemName string `json:"relation_work_item_name"`

	RelationWorkItemTypeName string `json:"relation_work_item_type_name"`

	RelationWorkItemTypeKey string `json:"relation_work_item_type_key"`

	ProjectRelationRuleID string `json:"project_relation_rule_id"`

	ProjectRelationRuleName string `json:"project_relation_rule_name"`

	RelationProjectKey string `json:"relation_project_key"`

	RelationProjectName string `json:"relation_project_name"`
}

type RelationRule struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Disabled int64 `json:"disabled"`

	WorkItemRelationID string `json:"work_item_relation_id"`

	WorkItemRelationName string `json:"work_item_relation_name"`

	CurrentWorkItemTypeKey string `json:"current_work_item_type_key"`

	CurrentWorkItemTypeName string `json:"current_work_item_type_name"`

	RemoteWorkItemTypeKey string `json:"remote_work_item_type_key"`

	RemoteWorkItemTypeName string `json:"remote_work_item_type_name"`

	ChatGroupMerge int64 `json:"chat_group_merge"`
}
