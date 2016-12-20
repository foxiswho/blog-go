package controllers

import (
	"fmt"
	"strconv"

	// "github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/models"
)

type HomeController struct {
	Common
}

func (this *HomeController) Get() {
	this.TplName = "homeTemplate.html"
	this.Data["Title"] = fmt.Sprintf("%s's Blog", models.Blogger.BlogName)
	this.Data["Description"] = fmt.Sprintf("%s's blog,%s.", models.Blogger.UserName, models.Blogger.Introduce)
	this.Data["Keywords"] = fmt.Sprintf("%s,%s,homepage,blog", models.Blogger.Introduce, models.Blogger.UserName)
	this.Leftbar("homepage")
	this.Verification()
	this.Home()
}
func (this *HomeController) Home() {
	this.Data["Tags"] = models.Blogger.Tags
	this.Data["Blogrolls"] = models.Blogger.Blogrolls
	this.Data["Archives"] = models.TMgr.Archives
	page := 1
	pageStr := this.Ctx.Input.Param(":page")
	if temp, err := strconv.Atoi(pageStr); err == nil {
		page = temp
	}
	topics, remainpage := models.TMgr.GetTopicsByPage(page)
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
			this.Data["UrlNewer"] = "/p/" + fmt.Sprint(page-1)
		}
		if remainpage == 0 {
			this.Data["StyleOlder"] = "disabled"
			this.Data["UrlOlder"] = "#"
		} else {
			this.Data["StyleOlder"] = ""
			this.Data["UrlOlder"] = "/p/" + fmt.Sprint(page+1)
		}
		this.Data["ListTopics"] = topics
	}
}
