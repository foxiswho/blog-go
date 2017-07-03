package site

import (
	"blog/service/admin"
	"strings"
	"strconv"
)

/**
	页面尾部数据 增加 自定义尾部
 */
func GetPageTemplate(id int, content string) (string) {
	//页面尾部操作
	site := admin.NewSiteService()
	config := site.SiteConfig()
	if config["this_page_url"] == "yes" {
		str := config["this_page_template"]
		str = strings.Replace(str, "{$id}", strconv.Itoa(id), -1)     //ID
		str = strings.Replace(str, "{$author}", strconv.Itoa(id), -1) //作者
		content = strings.TrimSpace(content) + "  \n" + str
	}
	return content
}
