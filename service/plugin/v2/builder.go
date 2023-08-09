package plugin

import "code.byted.org/bits/project-oapi-sdk-golang/core"

type GetUserPluginTokenReq struct {
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

type GetUserPluginTokenResp struct {
	*core.ApiResp `json:"-"`
	Error         *TokenErr        `json:"error"`
	Data          *UserPluginToken `json:"data"`
}

type RefreshTokenReq struct {
	apiReq *core.ApiReq
}

type RefreshTokenReqBuilder struct {
	apiReq *core.ApiReq
}

func NewRefreshTokenReqBuilder() *RefreshTokenReqBuilder {
	builder := &RefreshTokenReqBuilder{}
	builder.apiReq = &core.ApiReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &RefreshTokenReqBody{},
	}
	return builder
}

type RefreshTokenReqBody struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    int    `json:"type,omitempty"`
}

// 刷新token
func (builder *RefreshTokenReqBuilder) RefreshToken(refreshToken string) *RefreshTokenReqBuilder {
	builder.apiReq.Body.(*RefreshTokenReqBody).RefreshToken = refreshToken
	return builder
}

// 要刷新的token类型，目前固定填1
func (builder *RefreshTokenReqBuilder) TokenType(tokenType int) *RefreshTokenReqBuilder {
	builder.apiReq.Body.(*RefreshTokenReqBody).TokenType = tokenType
	return builder
}

func (builder *RefreshTokenReqBuilder) Build() *RefreshTokenReq {
	req := &RefreshTokenReq{}
	req.apiReq = builder.apiReq
	return req
}

type RefreshTokenResp struct {
	*core.ApiResp `json:"-"`
	Error         *TokenErr     `json:"error"`
	Data          *RefreshToken `json:"data"`
}
