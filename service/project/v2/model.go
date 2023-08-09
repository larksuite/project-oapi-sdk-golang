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

package project

import (
	"code.byted.org/bits/project-oapi-sdk-golang/service/user/v2"
)

type Business struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Project string `json:"project"`

	Labels []string `json:"labels"`

	RoleOwners map[string]*user.RoleOwner `json:"role_owners"`

	Watchers []string `json:"watchers"`

	Order float64 `json:"order"`

	SuperMasters []string `json:"super_masters"`

	Parent string `json:"parent"`

	Disabled bool `json:"disabled"`

	LevelID int64 `json:"level_id"`

	Children []*Business `json:"children"`
}

type Project struct {
	ProjectKey string `json:"project_key"`

	Name string `json:"name"`

	SimpleName string `json:"simple_name"`

	Administrators []string `json:"administrators"`
}

type Team struct {
	TeamID int64 `json:"team_id"`

	TeamName string `json:"team_name"`

	UserKeys []string `json:"user_keys"`

	Administrators []string `json:"administrators"`
}
