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

package view

type FixView struct {
	ViewID string `json:"view_id"`

	Name string `json:"name"`

	CreatedBy string `json:"created_by"`

	CreatedAt int64 `json:"created_at"`

	ModifiedBy string `json:"modified_by"`

	WorkItemIDList []int64 `json:"work_item_id_list"`

	Editable bool `json:"editable"`
}

type ViewConf struct {
	ViewID string `json:"view_id"`

	Name string `json:"name"`

	ViewUrl string `json:"view_url"`

	ViewType int32 `json:"view_type"`

	Auth int32 `json:"auth"`

	SystemView int32 `json:"system_view"`

	Collaborators []string `json:"collaborators"`

	CreatedAt int64 `json:"created_at"`

	CreatedBy string `json:"created_by"`
}
