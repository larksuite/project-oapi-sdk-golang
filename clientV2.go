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

package sdk

import (
	"github.com/larksuite/project-oapi-sdk-golang/core"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/comment"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/measure"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/project"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/user"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/view"
	"github.com/larksuite/project-oapi-sdk-golang/v2/service/workitem"
)

type ClientV2 struct {
	config   *core.Config            // sdk配置
	Measure  *measure.MeasureService // 度量
	Project  *project.ProjectService
	User     *user.UserService
	WorkItem *workitem.WorkItemService
	View     *view.ViewService
	Comment  *comment.CommentService
}

func NewClientV2(appId, appSecret string, options ...ClientOptionFunc) *ClientV2 {
	// 构建配置
	config := &core.Config{
		AppID:            appId,
		AppSecret:        appSecret,
		EnableTokenCache: true,
		AccessTokenType:  core.AccessTokenTypePlugin,
		BaseUrl:          "https://project.feishu.cn",
	}
	for _, option := range options {
		option(config)
	}

	// 构建日志器
	core.NewLogger(config)

	// 构建缓存
	core.NewCache(config)

	// 创建httpclient
	core.NewHttpClient(config)

	// 创建sdk-client，并初始化服务
	client := &ClientV2{config: config}
	initServiceV2(client, config)

	return client
}

func initServiceV2(client *ClientV2, config *core.Config) {
	client.Project = project.NewService(config)
	client.User = user.NewService(config)
	client.WorkItem = workitem.NewService(config)
	client.View = view.NewService(config)
	client.Measure = measure.NewService(config)
	client.Comment = comment.NewService(config)
}
