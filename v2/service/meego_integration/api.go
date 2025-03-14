package meego_integration

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_OAPIInviteBotJoinChat = "/open_api/:project_key/work_item/:work_item_id/bot_join_chat"


func NewService(config *core.Config) *MeegoIntegrationService {
	a := &MeegoIntegrationService{config: config}
	return a
}

type MeegoIntegrationService struct {
	config *core.Config
}


/*
 * @name: OAPIInviteBotJoinChatV2
 * @desc: 邀请机器人入群
 */
func (a *MeegoIntegrationService) OAPIInviteBotJoinChat(ctx context.Context, req *OAPIInviteBotJoinChatReq, options ...core.RequestOptionFunc) (*OAPIInviteBotJoinChatResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIInviteBotJoinChat
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIInviteBotJoinChat] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIInviteBotJoinChatResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIInviteBotJoinChat] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

