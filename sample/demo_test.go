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
	"github.com/larksuite/project-oapi-sdk-golang/service/plugin"
	"github.com/larksuite/project-oapi-sdk-golang/service/user"
	"net/http"
	"time"

	lark "code.byted.org/gopkg/lark/v2"
	sdk "github.com/larksuite/project-oapi-sdk-golang"
	constants "github.com/larksuite/project-oapi-sdk-golang/core"
	sdkcore "github.com/larksuite/project-oapi-sdk-golang/core"
	"github.com/larksuite/project-oapi-sdk-golang/service/field"
	"github.com/larksuite/project-oapi-sdk-golang/service/project"
	sub_task "github.com/larksuite/project-oapi-sdk-golang/service/task"
	"github.com/larksuite/project-oapi-sdk-golang/service/workitem"
	"testing"
)

// 获取空间下工作项类型
// func listProjectWorkItemType(client *sdk.Client) {
var meegoClient = sdk.NewClient(constants.MeegoPluginIdv2, constants.MeegoPluginSecretv2,
	sdk.WithLogLevel(constants.LogLevelDebug),
	sdk.WithReqTimeout(3*time.Second),
	sdk.WithEnableTokenCache(true),
	sdk.WithHttpClient(http.DefaultClient))

var larkClient = lark.New(
	lark.WithAppCredential(constants.AppId, constants.AppSecret),
	lark.WithOpenBaseURL("https://fsopen.bytedance.net"),
)

// 获取 pluginToken (1 虚拟 plugin_token)
func GetPluginToken(client *sdk.Client) string {
	resp, err := client.Plugin.GetPluginToken(context.Background(), 1)

	if err != nil {
		return resp.Data.Token
	}
	return ""
}

// 获取 VisualPluginToken (0 正式 plugin_token)
func GetVisualPluginToken(client *sdk.Client) string {
	resp, err := client.Plugin.GetPluginToken(context.Background(), 0)

	if err != nil {
		return resp.Data.Token
	}
	return ""
}

// 刷新/获取 userPluginToken
func RefreshUserToken(client *sdk.Client) string {
	resp, err := client.Plugin.RefreshToken(context.Background(), plugin.NewRefreshTokenReqBuilder().
		RefreshToken(constants.MeegoRefreshUserPluginToken).
		TokenType(1).
		Build(),
	)
	if err != nil {
		return resp.Data.Token
	}
	return ""
}

// 获取 userPluginToken
func TestGetUserToken(t *testing.T) {
	resp, _ := meegoClient.Plugin.GetUserPluginToken(context.Background(), "e356bd0fa3a642d8be107fddb31b7654",
		sdkcore.WithAccessToken(GetPluginToken(meegoClient)),
	)
	if resp != nil {
	}
}

func TestRefreshUserToken(t *testing.T) {
	resp, _ := meegoClient.Plugin.RefreshToken(context.Background(), plugin.NewRefreshTokenReqBuilder().
		RefreshToken(constants.MeegoRefreshUserPluginToken).
		TokenType(1).
		Build(),
	)
	if resp != nil {
	}
	return
}

// 获取 pluginToken (0 正式 plugin_token)
func TestGetVisualPluginToken(t *testing.T) {
	resp, _ := meegoClient.Plugin.GetPluginToken(context.Background(), 0)

	if resp != nil {
		return
	}
}

// 获取 pluginToken (1 虚拟 plugin_token)
func TestGetPluginToken(t *testing.T) {
	resp, _ := meegoClient.Plugin.GetPluginToken(context.Background(), 1)

	if resp != nil {
		return
	}
}

func TestListProjectWorkItemType(t *testing.T) {

	resp, err := meegoClient.Project.ListProjectWorkItemType(context.Background(), project.NewListProjectWorkItemTypeReqBuilder().
		ProjectKey(constants.MeegoPorjectKey).
		Build(),
		sdkcore.WithAccessToken(constants.MeegoPluginToken),
		//sdkcore.WithAccessToken(GetPluginToken(meegoClient)),
		sdkcore.WithUserKey(constants.MeegoUserKey),
	)

	//处理错误
	if err != nil {
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		//fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	//fmt.Println(sdkcore.Prettify(resp.Data))
}

// 创建工作项类型
func createWorkItem(client *sdk.Client) {
	resp, err := client.WorkItem.CreateWorkItem(context.Background(), workitem.NewCreateWorkItemReqBuilder().
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
		sdkcore.WithUserKey("user_key"),
	)

	//处理错误
	if err != nil {
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		//fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	//fmt.Println(sdkcore.Prettify(resp.Data))
}

// 创建子任务
func CreateSubTask(client *sdk.Client) {
	resp, err := client.Task.CreateSubTask(context.Background(), sub_task.NewCreateSubTaskReqBuilder().
		ProjectKey("project_key").
		WorkItemTypeKey("work_item_type_key").
		WorkItemID(0).
		NodeID("node_id").
		Name("name").
		Build(),
		sdkcore.WithUserKey("user_key"),
	)

	//处理错误
	if err != nil {
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		//fmt.Println(resp.ErrCode, resp.ErrMsg, resp.RequestId())
		return
	}

	// 业务数据处理
	//fmt.Println(sdkcore.Prettify(resp.Data))
}

// 获取工作项详情
func TestGetQueryWorkItemDetail(t *testing.T) {
	workitemIds := make([]int64, 0)
	//workitemIds = append(workitemIds, 5260430053)
	workitemIds = append(workitemIds, 4874734810)
	resp, err := meegoClient.WorkItem.QueryWorkItemDetail(context.Background(), workitem.NewQueryWorkItemDetailReqBuilder().
		ProjectKey(constants.MeegoPorjectKey).
		WorkItemTypeKey("story").
		WorkItemIDs(workitemIds).
		Build(),
		//sdkcore.WithAccessToken(constants.MeegoUserPluginToken),
		sdkcore.WithAccessToken(constants.MeegoPluginToken),
		sdkcore.WithUserKey(constants.MeegoUserKey),
	)
	print(resp)
	if err != nil {
		return
	}
}

// 获取用户信息
func TestQueryUserDetail(t *testing.T) {
	resp, _ := meegoClient.User.QueryUserDetail(context.Background(), user.NewQueryUserDetailReqBuilder().
		//UserKeys([]string{constants.MeegoUserKey}).
		UserKeys([]string{"7226723139252781057"}).
		Build(),
		sdkcore.WithAccessToken(constants.MeegoPluginToken),
		//sdkcore.WithUserKey(constants.MeegoUserKey),
	)
	if resp != nil {
		return
	}
}
