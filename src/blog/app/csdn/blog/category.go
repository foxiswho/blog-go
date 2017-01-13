package blog

import (
	"github.com/astaxie/beego/httplib"
	"blog/app/csdn"
)
//获取博主的自定义分类
type Category struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
}
//发送
func (t *Category)Post() (string, error) {
	req := httplib.Post(csdn.BLOG_CATEGORY_URL)
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}