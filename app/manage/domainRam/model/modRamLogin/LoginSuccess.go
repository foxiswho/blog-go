package modRamLogin

type LoginSuccessInfo struct {
	Account  string   `json:"account"`
	RealName string   `json:"realName"`
	Name     string   `json:"name"`
	Avatar   string   `json:"avatar"`
	Roles    []string `json:"roles"`
}

type LoginSuccess struct {
	Token        string           `json:"token"`
	AccessToken  string           `json:"accessToken"`
	RefreshToken string           `json:"refreshToken"`
	AuthCode     []string         `json:"authCode"`
	Info         LoginSuccessInfo `json:"info"`
}
