package controllers

import (
	"fmt"

	"github.com/deepzz0/goblog/models"
)

type SearchController struct {
	Common
}

func (this *SearchController) Get() {
	this.TplName = "groupTemplate.html"
	this.Leftbar("")
	this.ListTopic()
}

func (this *SearchController) ListTopic() {
	search := this.GetString("title")
	this.Data["StyleOlder"] = "disabled"
	this.Data["UrlOlder"] = "#"
	this.Data["StyleNewer"] = "disabled"
	this.Data["UrlNewer"] = "#"
	this.Data["Name"] = "Search：" + search
	this.Data["URL"] = fmt.Sprintf("/search?title=%s", search)
	this.Data["ListTopics"] = models.TMgr.GetTopicsSearch(search)
	this.Data["Title"] = fmt.Sprintf("Search: %s | %s", search, models.Blogger.BlogName)
	this.Data["Description"] = fmt.Sprintf("search title,%s,%s,blog", models.Blogger.Introduce, models.Blogger.UserName)
	this.Data["Keywords"] = fmt.Sprintf("search title,find,%s,%s", models.Blogger.Introduce, models.Blogger.UserName)
}
