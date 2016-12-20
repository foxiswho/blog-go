package controllers

import (
	"fmt"
	"strconv"

	// "github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/models"
)

type CategoryController struct {
	Common
}

func (this *CategoryController) Get() {
	this.TplName = "groupTemplate.html"
	this.ListTopic()
}

func (this *CategoryController) ListTopic() {
	cat := this.Ctx.Input.Param(":cat")
	this.Leftbar(cat)
	category := models.Blogger.GetCategoryByID(cat)
	var name string = "Not Found."
	if category != nil && category.Extra != "" {
		name = category.Text
	}
	this.Data["Name"] = name
	this.Data["URL"] = fmt.Sprintf("/cat/%s", category.ID)
	pageStr := this.Ctx.Input.Param(":page")
	page := 1
	if temp, err := strconv.Atoi(pageStr); err == nil {
		page = temp
	}
	topics, remainpage := models.TMgr.GetTopicsByCatgory(cat, page)
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
			this.Data["UrlNewer"] = "/cat/" + cat + fmt.Sprintf("/p/%d", page-1)
		}
		if remainpage == 0 {
			this.Data["StyleOlder"] = "disabled"
			this.Data["UrlOlder"] = "#"
		} else {
			this.Data["StyleOlder"] = ""
			this.Data["UrlOlder"] = "/cat/" + cat + fmt.Sprintf("/p/%d", page+1)
		}
		this.Data["ListTopics"] = topics
	}
	this.Data["Title"] = name + " | " + models.Blogger.BlogName
	this.Data["Description"] = fmt.Sprintf("%s's blog,%s,%s,blog", models.Blogger.UserName, models.Blogger.Introduce, category.Text)
	this.Data["Keywords"] = fmt.Sprintf("blog category,%s,%s,%s", category.Text, models.Blogger.Introduce, models.Blogger.UserName)
}
