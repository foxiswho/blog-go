package blog

import (
	"github.com/astaxie/beego/httplib"
	"blog/app/csdn"
)
//获取博客系统分类
type Channel struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	ClientId    string `json:"client_id" App申请的Key`
}
//发送
func (t *Channel)Post() (string, error) {
	req := httplib.Post(csdn.BLOG_CHANNEL_URL)
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}