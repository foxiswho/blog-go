package csdn

import (
	"github.com/foxiswho/blog-go/blog/app/csdn/conf"
	"github.com/foxiswho/blog-go/blog/app/csdn/entity"
	"github.com/foxiswho/blog-go/blog/fox"
	"github.com/foxiswho/blog-go/blog/fox/cache"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/foxiswho/blog-go/blog/fox/config"
	"net/url"
	"strings"
	"time"
)
//网页登陆
type AuthorizeWeb struct {
	RedirectUri string `json:"redirect_uri"` //登录成功后浏览器回跳的URL。
	Code        string `json:"code"`         //Authorization Code
	Config      map[string]string    `json:"-"`
}
//初始化
func NewAuthorizeWeb() *AuthorizeWeb {
	return new(AuthorizeWeb)
}
//获取配置
func (t *AuthorizeWeb)loadConfig() (bool, error) {
	maps, err := config.GetSection("csdn")
	if err != nil {
		return false, err
	}
	t.Config = maps
	return true, nil
}
//配置读取
func (t *AuthorizeWeb)SetConfig() (bool, error) {
	//获取配置
	ok, err := t.loadConfig()
	if err != nil {
		fmt.Println("setConfig err:", err)
		return false, err
	}
	fmt.Println("setConfig", ok)
	//配置文件是否读取
	if len(t.Config) < 1 {
		return false,fox.NewError("配置文件没有读取")
	}
	// 初始化AK，SK
	conf.ACCESS_KEY = t.Config["access_key"]
	conf.SECRET_KEY = t.Config["secret_key"]
	return true, nil
}
//设置回调URl
func (t *AuthorizeWeb)SetRedirectUri(str string) {
	t.RedirectUri = strings.TrimSpace(str)
}
//登录URL
func (t *AuthorizeWeb)GetAuthorizeUrl() string {
	if conf.ACCESS_KEY == "" {
		return "ACCESS_KEY 必须赋值"
	}
	if conf.SECRET_KEY == "" {
		return "SECRET_KEY 必须赋值"
	}
	//client_id：在开发者中心注册应用时获得的API Key。
	//redirect_uri：登录成功后浏览器回跳的URL。
	//response_type：服务端流程，此值固定为“code”。
	str := conf.WEB_URL
	str += "?client_id=" + url.QueryEscape(conf.ACCESS_KEY)
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
	str := conf.APP_URL
	str += "?client_id=" + url.QueryEscape(conf.ACCESS_KEY)
	str += "&client_secret=" + url.QueryEscape(conf.SECRET_KEY)
	str += "&grant_type=authorization_code"
	str += "&redirect_uri=" + url.QueryEscape(t.RedirectUri)
	str += "&code=" + url.QueryEscape(token)
	return str
}
//获取内容
func (t *AuthorizeWeb)Get(token string) (string, error) {
	//获取
	url := t.getAccessToken(token)
	fmt.Println("token url", url)
	//接口传输
	req := httplib.Get(url)
	//返回数据 及 判断
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}
//获取AccessToken
func (t *AuthorizeWeb)GetAccessToken(token string) (*entity.AccessToken, error) {
	//获取
	s, err := t.Get(token)
	if err != nil {
		return nil, err
	}
	fmt.Println("返回值", s)
	//反序列化 结构体
	var access *entity.AccessToken
	err = json.Unmarshal([]byte(s), &access)
	if err != nil {
		return nil, err
	}
	access.LastTime = time.Now()
	//缓存 序列化 json
	j,_:=json.Marshal(access)
	//存储
	err = t.PutAccessTokenCache(string(j))
	if err != nil {
		return nil, err
	}
	return access, nil
}
//获取token缓存
func (t *AuthorizeWeb)GetAccessTokenCache() (*entity.AccessToken, error) {
	//获取缓存
	str, err := t.GetCache("CSDN_AccessToken")
	if err != nil {
		return nil, err
	}
	fmt.Println("令牌",str)
	//反序列化 结构体
	var access *entity.AccessToken
	err = json.Unmarshal([]byte(str), &access)
	if err != nil {
		return nil,fox.NewError("token反序列化错误：" + err.Error())
	}
	fmt.Println("反序列化",access)
	return access, nil
}
//更新token缓存
func (t *AuthorizeWeb)PutAccessTokenCache(val interface{}) (error) {
	//设置缓存
	return t.PutCache("CSDN_AccessToken", val)
}
//获取缓存
func (t *AuthorizeWeb)GetCache(key string) (string, error) {
	//获取缓存
	tmp := cache.Get(key)
	fmt.Println("获取csdn 缓存",tmp)
	str := tmp.(string)
	if len(str) < 1 {
		return "",fox.NewError("CSDN Token 已过期，请重新用CSDN登陆")
	}
	return str, nil
}
//更新缓存
func (t *AuthorizeWeb)PutCache(key string, val interface{}) (error) {
	//设置缓存
	err := cache.Put(key, val, 86400 * time.Second)
	return err
}
