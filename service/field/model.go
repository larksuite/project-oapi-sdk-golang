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

package field

type FieldValuePair struct {
	FieldKey string `json:"field_key"`

	FieldValue interface{} `json:"field_value"`

	TargetState *TargetState `json:"target_state"`

	FieldTypeKey string `json:"field_type_key"`

	FieldAlias string `json:"field_alias"`
}

type MultiSignal struct {
	Status string `json:"status"`

	Detail []*MultiSignalDetail `json:"detail"`
}

type MultiSignalDetail struct {
	ID string `json:"id"`

	Title string `json:"title"`

	Status string `json:"status"`

	ViewLink string `json:"view_link"`

	QueryLink *QueryLink `json:"query_link"`
}

type Option struct {
	Label string `json:"label"`

	Value string `json:"value"`

	Children []*Option `json:"children"`

	WorkItemTypeKey string `json:"work_item_type_key"`
}

type QueryLink struct {
	Url string `json:"url"`

	Method string `json:"method"`

	Headers interface{} `json:"headers"`

	Body interface{} `json:"body"`

	Params interface{} `json:"params"`
}

type SimpleField struct {
	FieldKey string `json:"field_key"`

	FieldAlias string `json:"field_alias"`

	FieldTypeKey string `json:"field_type_key"`

	FieldName string `json:"field_name"`

	IsCustomField bool `json:"is_custom_field"`

	Options []*Option `json:"options"`

	CompoundFields []*SimpleField `json:"compound_fields"`

	IsObsoleted bool `json:"is_obsoleted"`

	WorkItemScopes []string `json:"work_item_scopes"`

	ValueGenerateMode string `json:"value_generate_mode"`

	RelationID string `json:"relation_id"`
}

type TargetState struct {
	StateKey string `json:"state_key"`

	TransitionID int64 `json:"transition_id"`
}
