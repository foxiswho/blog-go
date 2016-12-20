// Package controllers provides ...
package controllers

import (
	"fmt"

	"github.com/deepzz0/goblog/models"
)

type ArchivesController struct {
	Common
}

func (this *ArchivesController) Get() {
	this.TplName = "groupTemplate.html"
	this.Leftbar("")
	this.ListTopic()
}

func (this *ArchivesController) ListTopic() {
	date := fmt.Sprintf("%s/%s", this.Ctx.Input.Param(":year"), this.Ctx.Input.Param(":month"))
	this.Data["StyleOlder"] = "disabled"
	this.Data["UrlOlder"] = "#"
	this.Data["StyleNewer"] = "disabled"
	this.Data["UrlNewer"] = "#"
	this.Data["Name"] = "归档：" + date
	this.Data["URL"] = fmt.Sprintf("/archives/%s", date)
	this.Data["ListTopics"] = models.TMgr.GetTopicsArchives(date)
	this.Data["Title"] = fmt.Sprintf("归档: %s | %s", date, models.Blogger.BlogName)
	this.Data["Description"] = fmt.Sprintf("archives title,%s,%s,blog", models.Blogger.Introduce, models.Blogger.UserName)
	this.Data["Keywords"] = fmt.Sprintf("archives title,find,%s,%s", models.Blogger.Introduce, models.Blogger.UserName)
}
