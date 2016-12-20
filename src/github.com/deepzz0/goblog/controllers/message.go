package controllers

import (
	"fmt"

	"github.com/deepzz0/goblog/models"
)

type MessageController struct {
	Common
}

func (this *MessageController) Get() {
	this.TplName = "messageTemplate.html"
	this.Data["Title"] = "留言面板 | " + models.Blogger.BlogName
	this.Leftbar("message")
	this.Content()
}

func (this *MessageController) Content() {
	this.Data["ID"] = "99999"
	this.Data["URL"] = "/message"
	this.Data["Name"] = "留言面板"
	this.Data["Domain"] = Domain
	this.Data["Description"] = fmt.Sprintf("message board,%s,%s,blog", models.Blogger.Introduce, models.Blogger.UserName)
	this.Data["Keywords"] = fmt.Sprintf("message,%s,%s", models.Blogger.Introduce, models.Blogger.UserName)
}
