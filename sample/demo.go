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

package sample

import (
	"code.byted.org/bits/project-oapi-sdk-golang/service/field"
	"code.byted.org/bits/project-oapi-sdk-golang/service/project"
	sub_task "code.byted.org/bits/project-oapi-sdk-golang/service/task"
	"code.byted.org/bits/project-oapi-sdk-golang/service/workitem"
	"context"
	"fmt"

	sdk "code.byted.org/bits/project-oapi-sdk-golang"
	sdkcore "code.byted.org/bits/project-oapi-sdk-golang/core"
)

// 获取空间下工作项类型
func listProjectWorkItemType(client *sdk.Client) {
	resp, err := client.Project.ListProjectWorkItemType(context.Background(), project.NewListProjectWorkItemTypeReqBuilder().
		AccessUser("user_key").
		ProjectKey("project_key").
		Build(),
	)

	//处理错误
	if err != nil {
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	fmt.Println(sdkcore.Prettify(resp.Data))
}

// 创建工作项类型
func createWorkItem(client *sdk.Client) {
	resp, err := client.WorkItem.CreateWorkItem(context.Background(), workitem.workitem.NewCreateWorkItemReqBuilder().
		AccessUser("user_key").
		ProjectKey("project_key").
		WorkItemTypeKey("work_item_type_key").
		Name("name").
		TemplateID(1024).
		FieldValuePairs([]*field.FieldValuePair{
			{
				FieldKey:   "priority",
				FieldValue: "这是描述",
			},
		}).
		Build(),
	)

	//处理错误
	if err != nil {
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	fmt.Println(sdkcore.Prettify(resp.Data))
}

// 创建子任务
func CreateSubTask(client *sdk.Client) {
	resp, err := client.Task.CreateSubTask(context.Background(), sub_task.NewCreateSubTaskReqBuilder().
		AccessUser("user_key").
		ProjectKey("project_key").
		WorkItemTypeKey("work_item_type_key").
		WorkItemID(0).
		NodeID("node_id").
		Name("name").
		Build(),
	)

	//处理错误
	if err != nil {
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	fmt.Println(sdkcore.Prettify(resp.Data))
}
