package open_api

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type OAPIBatchWebHookReq struct {
	apiReq *core.APIReq
}
type OAPIBatchWebHookReqBody struct {

    RequestTime  *string `json:"request_time,omitempty"`

    Signature  *string `json:"signature,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    ViewID  *string `json:"view_id,omitempty"`

    Source  *string `json:"source,omitempty"`

    UserLanguage  *string `json:"user_language,omitempty"`

    PluginInfo  *PluginInfo `json:"plugin_info,omitempty"`

    UserInfo  *UserInfo `json:"user_info,omitempty"`

    WorkItemInfos  []WorkItemInfo `json:"work_item_infos,omitempty"`

}

type OAPIBatchWebHookResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	WorkItemDatas       []WorkItemData         `json:"work_item_datas"`
	
}

type OAPIBatchWebHookReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIBatchWebHookReqBuilder() *OAPIBatchWebHookReqBuilder {
	builder := &OAPIBatchWebHookReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIBatchWebHookReqBody{},
	}
	return builder
}

func (builder *OAPIBatchWebHookReqBuilder) RequestTime(requestTime string) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).RequestTime = &requestTime
	return builder
}


func (builder *OAPIBatchWebHookReqBuilder) Signature(signature string) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).Signature = &signature
	return builder
}


func (builder *OAPIBatchWebHookReqBuilder) ProjectKey(projectKey string) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *OAPIBatchWebHookReqBuilder) ViewID(viewID string) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).ViewID = &viewID
	return builder
}


func (builder *OAPIBatchWebHookReqBuilder) Source(source string) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).Source = &source
	return builder
}


func (builder *OAPIBatchWebHookReqBuilder) UserLanguage(userLanguage string) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).UserLanguage = &userLanguage
	return builder
}


func (builder *OAPIBatchWebHookReqBuilder) PluginInfo(pluginInfo *PluginInfo) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).PluginInfo = pluginInfo
	return builder
}

func (builder *OAPIBatchWebHookReqBuilder) UserInfo(userInfo *UserInfo) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).UserInfo = userInfo
	return builder
}

func (builder *OAPIBatchWebHookReqBuilder) WorkItemInfos(workItemInfos []WorkItemInfo) *OAPIBatchWebHookReqBuilder {
	builder.apiReq.Body.(*OAPIBatchWebHookReqBody).WorkItemInfos = workItemInfos
	return builder
}
func (builder *OAPIBatchWebHookReqBuilder) Build() *OAPIBatchWebHookReq {
	req := &OAPIBatchWebHookReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryTaskResultReq struct {
	apiReq *core.APIReq
}

type OAPIQueryTaskResultResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	TaskID       *string         `json:"task_id"`
	
	TaskStatus       *string         `json:"task_status"`
	
	Total       *int64         `json:"total"`
	
	SuccessTotal       *int64         `json:"success_total"`
	
	ErrorTotal       *int64         `json:"error_total"`
	
	SuccessSubTaskList       []string         `json:"success_sub_task_list"`
	
	FailSubTaskList       []string         `json:"fail_sub_task_list"`
	
	ErrorScenes       []string         `json:"error_scenes"`
	
}

type OAPIQueryTaskResultReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryTaskResultReqBuilder() *OAPIQueryTaskResultReqBuilder {
	builder := &OAPIQueryTaskResultReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *OAPIQueryTaskResultReqBuilder) TaskID(taskID string) *OAPIQueryTaskResultReqBuilder {
	builder.apiReq.QueryParams.Set("task_id", fmt.Sprint(taskID))
	return builder
}

func (builder *OAPIQueryTaskResultReqBuilder) Build() *OAPIQueryTaskResultReq {
	req := &OAPIQueryTaskResultReq{}
	req.apiReq = builder.apiReq
	return req
}

