package entity

import (
	"github.com/astaxie/beego/httplib"
	"fox/app/csdn"
)
//发表/修改文章
type Article struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	Id          int `json:"id" 否 文章id，修改文章的时候需要`
	Title       string `json:"title" 是 文章标题`
	Type        string `json:"type" 是 文章类型（original|report|translated）`
	Description string `json:"description" 否 文章简介`
	Content     string `json:"content" 是 文章内容`
	Categories  string `json:"content" 否 自定义类别（英文逗号分割）`
	Tags        string `json:"content" 否 文章标签（英文逗号分割）`
	Ip          string `json:"content" 否 用户ip`
}