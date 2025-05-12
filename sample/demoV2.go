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
	"context"
	"fmt"
	sdk "github.com/larksuite/project-oapi-sdk-golang"
	sdkcore "github.com/larksuite/project-oapi-sdk-golang/core"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/workitem"
)

// 创建工作项类型
func createWorkItemV2(client *sdk.ClientV2) {
	project_key := "project_key"
	work_item_type_key := "work_item_type_key"
	name := "name"
	resp, err := client.WorkItem.CreateWorkItem(context.Background(), workitem.NewCreateWorkItemReqBuilder().
		ProjectKey(project_key).
		WorkItemTypeKey(work_item_type_key).
		Name(name).
		Build(),
		sdkcore.WithUserKey("user_key"),
	)

	//处理错误
	if err != nil {
		// 处理err
		fmt.Printf("error creating work item: %v\n", err)
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

// 更新工作项类型
func updateWorkItemV2(client *sdk.ClientV2) {
	project_key := "project_key"
	work_item_type_key := "work_item_type_key"
	var work_item_id int64 = 0
	fieldKey := "name"
	fieldValue := "fieldValue"
	fieldValuePairs := []workitem.FieldValuePair{
		workitem.FieldValuePair{
			FieldKey:   &fieldKey,
			FieldValue: &fieldValue,
		},
	}
	resp, err := client.WorkItem.UpdateWorkItem(context.Background(), workitem.NewUpdateWorkItemReqBuilder().
		ProjectKey(project_key).
		WorkItemTypeKey(work_item_type_key).
		WorkItemID(work_item_id).
		UpdateFields(fieldValuePairs).
		Build(),
		sdkcore.WithUserKey("user_key"),
	)

	//处理错误
	if err != nil {
		// 处理err
		fmt.Printf("error UpdateWorkItem work item: %v\n", err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	fmt.Println(sdkcore.Prettify(resp.String()))
}

// 获取工作项操作记录
func GetWorkItemOpRecordV2(client *sdk.ClientV2) {
	project_key := "project_key"
	resp, err := client.WorkItem.GetWorkItemOpRecord(context.Background(), workitem.NewGetWorkItemOpRecordReqBuilder().
		ProjectKey(project_key).
		WorkItemIDs([]int64{0}).
		Build(),
		sdkcore.WithUserKey("user_key"),
	)

	//处理错误
	if err != nil {
		// 处理err
		fmt.Printf("error GetWorkItemOpRecordV2 work item: %v\n", err)
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
