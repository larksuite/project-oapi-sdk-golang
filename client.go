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
	"net/http"
	"time"

	"github.com/larksuite/project-oapi-sdk-golang/core"
	"github.com/larksuite/project-oapi-sdk-golang/service/attachment"
	"github.com/larksuite/project-oapi-sdk-golang/service/chat"
	"github.com/larksuite/project-oapi-sdk-golang/service/comment"
	"github.com/larksuite/project-oapi-sdk-golang/service/field"
	"github.com/larksuite/project-oapi-sdk-golang/service/measure"
	"github.com/larksuite/project-oapi-sdk-golang/service/plugin"
	"github.com/larksuite/project-oapi-sdk-golang/service/project"
	"github.com/larksuite/project-oapi-sdk-golang/service/project_relation"
	"github.com/larksuite/project-oapi-sdk-golang/service/role_conf"
	"github.com/larksuite/project-oapi-sdk-golang/service/task"
	"github.com/larksuite/project-oapi-sdk-golang/service/user"
	"github.com/larksuite/project-oapi-sdk-golang/service/view"
	"github.com/larksuite/project-oapi-sdk-golang/service/wbs_view_draft"
	"github.com/larksuite/project-oapi-sdk-golang/service/workitem"
	"github.com/larksuite/project-oapi-sdk-golang/service/workitem_conf"
)

type Client struct {
	config          *core.Config                             // sdk配置
	View            *view.ViewService                        // 视图
	Attachment      *attachment.AttachmentService            // 附件
	Field           *field.FieldService                      // 字段
	WorkItemConf    *workitem_conf.WorkItemConfService       // 工作项配置
	Chat            *chat.ChatService                        // 群组
	Comment         *comment.CommentService                  // 评论
	WorkItem        *workitem.WorkItemService                // 工作项
	User            *user.UserService                        // 用户
	Project         *project.ProjectService                  // 空间
	Task            *task.TaskService                        // 子任务
	RoleConf        *role_conf.RoleConfService               // 流程角色
	ProjectRelation *project_relation.ProjectRelationService // 空间关联
	Plugin          *plugin.PluginService                    // token
	Measure         *measure.MeasureService                  // 度量
	WbsViewDraft    *wbs_view_draft.WbsViewDraftService      // 计划表草稿
}

type ClientOptionFunc func(config *core.Config)

func WithAccessTokenType(accessTokenType core.AccessTokenType) ClientOptionFunc {
	return func(config *core.Config) {
		config.AccessTokenType = accessTokenType
	}
}

func WithEnableTokenCache(enableTokenCache bool) ClientOptionFunc {
	return func(config *core.Config) {
		config.EnableTokenCache = enableTokenCache
	}
}

func WithLogger(logger core.Logger) ClientOptionFunc {
	return func(config *core.Config) {
		config.Logger = logger
	}
}

func WithOpenBaseUrl(baseUrl string) ClientOptionFunc {
	return func(config *core.Config) {
		config.BaseUrl = baseUrl
	}
}

func WithLogLevel(logLevel core.LogLevel) ClientOptionFunc {
	return func(config *core.Config) {
		config.LogLevel = logLevel
	}
}

func WithTokenCache(cache core.Cache) ClientOptionFunc {
	return func(config *core.Config) {
		config.TokenCache = cache
	}
}

func WithLogReqAtDebug(printReqRespLog bool) ClientOptionFunc {
	return func(config *core.Config) {
		config.LogReqAtDebug = printReqRespLog
	}
}

func WithHttpClient(httpClient core.HttpClient) ClientOptionFunc {
	return func(config *core.Config) {
		config.HttpClient = httpClient
	}
}

func WithReqTimeout(reqTimeout time.Duration) ClientOptionFunc {
	return func(config *core.Config) {
		config.ReqTimeout = reqTimeout
	}
}

// 设置每次请求都会携带的 header
func WithHeaders(header http.Header) ClientOptionFunc {
	return func(config *core.Config) {
		config.Header = header
	}
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) *Client {
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
	client := &Client{config: config}
	initService(client, config)

	return client
}

func initService(client *Client, config *core.Config) {
	client.Project = project.NewService(config)
	client.User = user.NewService(config)
	client.WorkItem = workitem.NewService(config)
	client.Task = task.NewService(config)
	client.Field = field.NewService(config)
	client.View = view.NewService(config)
	client.Attachment = attachment.NewService(config)
	client.Comment = comment.NewService(config)
	client.Chat = chat.NewService(config)
	client.WorkItemConf = workitem_conf.NewService(config)
	client.RoleConf = role_conf.NewService(config)
	client.ProjectRelation = project_relation.NewService(config)
	client.Plugin = plugin.NewService(config)
	client.Measure = measure.NewService(config)
	client.WbsViewDraft = wbs_view_draft.NewService(config)
}
