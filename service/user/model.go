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

package user

type Channel struct {
	TenantName string `json:"tenant_name"`

	TenantGroupID int64 `json:"tenant_group_id"`
}

type RoleAssign struct {
	Role string `json:"role"`

	Name string `json:"name"`

	DefaultAppear int32 `json:"default_appear"`

	Deletable int32 `json:"deletable"`

	MemberAssign int32 `json:"member_assign"`

	Members []string `json:"members"`
}

type RoleOwner struct {
	Role string `json:"role"`

	Name string `json:"name"`

	Owners []string `json:"owners"`
}

type UserBasicInfo struct {
	UserID int64 `json:"user_id"`

	UserKey string `json:"user_key"`

	Username string `json:"username"`

	Email string `json:"email"`

	AvatarUrl string `json:"avatar_url"`

	NameCn string `json:"name_cn"`

	NameEn string `json:"name_en"`

	OutID string `json:"out_id"`

	Channels []*Channel `json:"channels"`

	Status string `json:"status"`
}
