package blog

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/foxiswho/blog-go/app/csdn/conf"
)
//获取博主的自定义分类
type List struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	Status string `json:"status" 否 文章状态，取值范围：enabled|draft，默认enabled`
	Page int `json:"page" 否 当前页码，默认1`
	Size int `json:"size" 否 每页条数，默认15`
}
//发送
func (t *List)Post() (string, error) {
	//接口传输数据
	req := httplib.Post(conf.BLOG_LIST_URL)
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}
