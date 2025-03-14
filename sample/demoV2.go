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
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/workitem"
	"net/http"
	"time"

	sdk "github.com/larksuite/project-oapi-sdk-golang"
	sdkcore "github.com/larksuite/project-oapi-sdk-golang/core"
)

// 创建工作项类型
func createWorkItemV2(client *sdk.ClientV2) {
	project_key := "662e151cd7fd734ae1874213"
	work_item_type_key := "story"
	name := "generator1"
	resp, err := client.WorkItem.OAPICreateWorkItem(context.Background(), workitem.NewOAPICreateWorkItemReqBuilder().
		ProjectKey(&project_key).
		WorkItemTypeKey(&work_item_type_key).
		Name(&name).
		Build(),
		sdkcore.WithUserKey("7356795280408297476"),
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

// 创建工作项类型
func updateWorkItemV2(client *sdk.ClientV2) {
	project_key := "662e151cd7fd734ae1874213"
	work_item_type_key := "story"
	var work_item_id int64 = 6092143382
	fieldKey := "name"
	fieldValue := "generator-update"
	fieldValuePairs := []workitem.FieldValuePair{
		workitem.FieldValuePair{
			FieldKey:   &fieldKey,
			FieldValue: &fieldValue,
		},
	}
	resp, err := client.WorkItem.OAPIUpdateWorkItem(context.Background(), workitem.NewOAPIUpdateWorkItemReqBuilder().
		ProjectKey(&project_key).
		WorkItemTypeKey(&work_item_type_key).
		WorkItemID(&work_item_id).
		UpdateFields(fieldValuePairs).
		Build(),
		sdkcore.WithUserKey("7356795280408297476"),
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
	fmt.Println(sdkcore.Prettify(resp.String()))
}

func main() {
	var client = sdk.NewClientV2("MII_661E3D92B64A4004", "40F8C24AD263B9B12ECF55AF49C5AF45",
		sdk.WithOpenBaseUrl("https://meego.larkoffice.com/"),
		sdk.WithReqTimeout(3*time.Second),
		sdk.WithEnableTokenCache(true),
		sdk.WithHttpClient(http.DefaultClient))
	//createWorkItemV2(client)
	updateWorkItemV2(client)
}
