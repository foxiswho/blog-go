package controllers

import (
	"fmt"
	"strconv"

	"github.com/deepzz0/goblog/models"
	// "github.com/deepzz0/go-com/log"
)

type TagController struct {
	Common
}

func (this *TagController) Get() {
	this.TplName = "groupTemplate.html"
	this.Leftbar("")
	this.ListTopic()
}

func (this *TagController) ListTopic() {
	tagName := this.Ctx.Input.Param(":tag")
	tag := models.Blogger.GetTagByID(tagName)
	this.Data["Name"] = "Not Found."
	if tag != nil {
		this.Data["Name"] = tag.ID
		this.Data["URL"] = fmt.Sprintf("/tag/%s", tag.ID)
		page := 1
		tagName := this.Ctx.Input.Param(":tag")
		pageStr := this.Ctx.Input.Param(":page")
		if temp, err := strconv.Atoi(pageStr); err == nil {
			page = temp
		}
		topics, remainpage := models.TMgr.GetTopicsByTag(tagName, page)
		if remainpage == -1 {
			this.Data["StyleOlder"] = "disabled"
			this.Data["UrlOlder"] = "#"
			this.Data["StyleNewer"] = "disabled"
			this.Data["UrlNewer"] = "#"
		} else {
			if page == 1 {
				this.Data["StyleNewer"] = "disabled"
				this.Data["UrlNewer"] = "#"
			} else {
				this.Data["StyleNewer"] = ""
				this.Data["UrlNewer"] = "/tag/" + tagName + fmt.Sprintf("/p/%d", page-1)
			}
			if remainpage == 0 {
				this.Data["StyleOlder"] = "disabled"
				this.Data["UrlOlder"] = "#"
			} else {
				this.Data["StyleOlder"] = ""
				this.Data["UrlOlder"] = "/tag/" + tagName + fmt.Sprintf("/p/%d", page+1)
			}
			this.Data["ListTopics"] = topics
		}
	}
	this.Data["Title"] = tagName + " | " + models.Blogger.BlogName
	this.Data["Description"] = fmt.Sprintf("tag,%s,%s,blog", models.Blogger.Introduce, models.Blogger.UserName)
	this.Data["Keywords"] = fmt.Sprintf("tag,%s,%s", models.Blogger.Introduce, models.Blogger.UserName)
}
