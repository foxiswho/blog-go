package modRamLogin

// IdpLoginCt IDP OAuth 登录请求参数
type IdpLoginCt struct {
	SourceNo     string `json:"sourceNo" form:"sourceNo" label:"认证源编号"`           // 认证源编号
	Code         string `json:"code" form:"code" label:"授权码"`                     // OAuth 授权码
	State        string `json:"state" form:"state" label:"状态"`                    // OAuth state 参数
	RedirectUri  string `json:"redirectUri" form:"redirectUri" label:"回调地址"`      // OAuth 回调地址
	Method       string `json:"method" form:"method" label:"操作类型"`                // signup/signin/link
	CodeVerifier string `json:"codeVerifier" form:"codeVerifier" label:"PKCE验证码"` // PKCE code verifier
}

// IdpLoginSuccess IDP OAuth 登录成功返回
type IdpLoginSuccess struct {
	AccessToken  string           `json:"accessToken"`
	RefreshToken string           `json:"refreshToken"`
	Info         LoginSuccessInfo `json:"info"`
	IsSignup     bool             `json:"isSignup"`           // 是否为新注册用户
	MfaRequired  bool             `json:"mfaRequired"`        // 是否需要 MFA 验证
	MfaToken     string           `json:"mfaToken,omitempty"` // MFA 临时令牌
	MfaType      string           `json:"mfaType,omitempty"`  // 首选 MFA 类型
}
