package plugin

type TokenErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserPluginToken struct {
	ExpireTime             int    `json:"expire_time"`
	Token                  string `json:"token"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpireTime int    `json:"refresh_token_expire_time"`
	UserKey                string `json:"user_key"`
	TenantKey              string `json:"tenant_key"`
}

type RefreshToken struct {
	ExpireTime             int    `json:"expire_time"`
	Token                  string `json:"token"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpireTime int    `json:"refresh_token_expire_time"`
}
