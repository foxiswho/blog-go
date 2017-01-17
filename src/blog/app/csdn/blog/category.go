package blog

import (
	"github.com/astaxie/beego/httplib"
	"blog/fox"
	"time"
	"blog/app/csdn/conf"
)
//获取博主的自定义分类
type Category struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
}
func (t *Category)Check() (error) {
	if len(t.AccessToken) < 1 {
		return fox.NewError("access_token 不能为空")
	}
	return nil
}
//发送
func (t *Category)Post() (string, error) {
	err := t.Check()
	if err != nil {
		return "", err
	}
	req := httplib.Post(conf.BLOG_CATEGORY_URL)
	//超时
	req.SetTimeout(100 * time.Second, 30 * time.Second)
	//参数
	req.Param("access_token", t.AccessToken)
	//返回
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}