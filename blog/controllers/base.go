package controllers

import (
	"bytes"
	"github.com/beego/beego/v2/server/web"
	"html/template"
)
//模版编译
//@sectionTpl 模版路径
//@Data 模版替换内容
func ExecuteTemplateHtml(sectionTpl string, Data map[interface{}]interface{}) string {
	if sectionTpl ==""{
		return ""
	}
	var buf bytes.Buffer
	buf.Reset()
	err := web.ExecuteTemplate(&buf, sectionTpl, Data)
	if err != nil {
		return ""
	}
	return string(template.HTML(buf.String()))
}
