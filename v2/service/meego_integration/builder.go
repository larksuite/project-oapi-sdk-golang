package meego_integration

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type OAPIInviteBotJoinChatReq struct {
	apiReq *core.APIReq
}
type OAPIInviteBotJoinChatReqBody struct {

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    AppIDs  []string `json:"app_ids,omitempty"`

}

type OAPIInviteBotJoinChatResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *BotJoinChatInfo         `json:"data"`
	
}

type OAPIInviteBotJoinChatReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIInviteBotJoinChatReqBuilder() *OAPIInviteBotJoinChatReqBuilder {
	builder := &OAPIInviteBotJoinChatReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIInviteBotJoinChatReqBody{},
	}
	return builder
}
func (builder *OAPIInviteBotJoinChatReqBuilder) ProjectKey(projectKey *string) *OAPIInviteBotJoinChatReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIInviteBotJoinChatReqBuilder) WorkItemID(workItemID *int64) *OAPIInviteBotJoinChatReqBuilder {
	builder.apiReq.PathParams.Set("work_item_id", fmt.Sprint(*workItemID))
	return builder
}
func (builder *OAPIInviteBotJoinChatReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIInviteBotJoinChatReqBuilder {
	builder.apiReq.Body.(*OAPIInviteBotJoinChatReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIInviteBotJoinChatReqBuilder) AppIDs(appIDs []string) *OAPIInviteBotJoinChatReqBuilder {
	builder.apiReq.Body.(*OAPIInviteBotJoinChatReqBody).AppIDs = appIDs
	return builder
}
func (builder *OAPIInviteBotJoinChatReqBuilder) Build() *OAPIInviteBotJoinChatReq {
	req := &OAPIInviteBotJoinChatReq{}
	req.apiReq = builder.apiReq
	return req
}

