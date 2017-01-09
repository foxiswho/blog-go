package csdn

import (
	"net/url"
	"strings"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
)

type AuthorizeWeb struct {
	RedirectUri string `json:"redirect_uri"` //登录成功后浏览器回跳的URL。
	Code        string `json:"code"`         //Authorization Code
}
//设置回调URl
func (t *AuthorizeWeb)SetRedirectUri(str string) {
	t.RedirectUri = strings.TrimSpace(str)
}
//登录URL
func (t *AuthorizeWeb)GetAuthorizeUrl() string {
	if ACCESS_KEY == "" {
		return "ACCESS_KEY 必须赋值"
	}
	if SECRET_KEY == "" {
		return "SECRET_KEY 必须赋值"
	}
	//client_id：在开发者中心注册应用时获得的API Key。
	//redirect_uri：登录成功后浏览器回跳的URL。
	//response_type：服务端流程，此值固定为“code”。
	str := WEB_URL
	str += "?client_id=" + url.QueryEscape(ACCESS_KEY)
	str += "&redirect_uri=" + url.QueryEscape(t.RedirectUri)
	str += "&response_type=code"
	return str
}
//获取Token
func (t *AuthorizeWeb)getAccessToken(token string) string {
	token = strings.TrimSpace(token)
	//client_id：在开发者中心注册应用时获得的API Key。
	//client_secret：在开发者中心注册应用时获得的API Secret。
	//grant_type：此值为“authorization_code”。
	//redirect_uri：流程结束后要跳转回得URL。
	//code：用户登录成功后获得的 Authorization Code。
	//
	//这里获取ACCESS_TOKEN 的URL 和APP_URL一样，直接使用
	str := APP_URL
	str += "?client_id=" + url.QueryEscape(ACCESS_KEY)
	str += "&client_secret=" + url.QueryEscape(SECRET_KEY)
	str += "&grant_type=authorization_code"
	str += "&redirect_uri=" + url.QueryEscape(t.RedirectUri)
	str += "&code=" + url.QueryEscape(token)
	return str
}
//获取内容
func (t *AuthorizeWeb)Get(token string) (string, error) {
	req := httplib.Get(t.getAccessToken(token))
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}
//获取AccessToken
func (t *AuthorizeWeb)GetAccessToken(token string) (*AccessToken, error) {
	s, err := t.Get(token)
	if err != nil {
		return nil, err
	}
	var access AccessToken
	err = json.Unmarshal(s, &access)
	if err != nil {
		return nil, err
	}
	return access, nil
}