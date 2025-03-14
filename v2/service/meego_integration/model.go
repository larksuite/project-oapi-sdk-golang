package meego_integration


type BotJoinChatInfo struct {

    ChatID  *string `json:"chat_id,omitempty"`

    SuccessMembers  []string `json:"success_members,omitempty"`

    FailedMembers  []string `json:"failed_members,omitempty"`

}
