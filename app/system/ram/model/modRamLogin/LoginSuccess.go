package modRamLogin

type LoginSuccessInfo struct {
	Account  string `json:"account"`
	RealName string `json:"realName"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type LoginSuccess struct {
	Token       string           `json:"token"`
	AccessToken string           `json:"accessToken"`
	AuthCode    []string         `json:"authCode"`
	Info        LoginSuccessInfo `json:"info"`
}
