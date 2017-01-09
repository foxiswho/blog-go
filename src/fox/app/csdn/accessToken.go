package csdn
/**
令牌返回
 */
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int `json:"expires_in"`
	Username    string        `json:"username"`
}