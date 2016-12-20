package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
	// "github.com/deepzz0/go-com/log"
)

type TopicController struct {
	Common
}

func (this *TopicController) Get() {
	this.TplName = "topicTemplate.html"
	this.Leftbar("")
	this.Topic()
}

func (this *TopicController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	resp.Data = "Not Found."
	id := this.Ctx.Input.Param(":id")
	ID, err := strconv.Atoi(id)
	if err == nil {
		topic := models.TMgr.GetTopic(int32(ID))
		if topic != nil {
			resp.Data = string(topic.Content)
		}
	}
}

func (this *TopicController) Topic() {
	id := this.Ctx.Input.Param(":id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		this.Data["IsFalse"] = true
		return
	}
	topic := models.TMgr.GetTopic(int32(ID))
	if topic == nil {
		this.Data["IsFalse"] = true
		return
	}
	topic.PV++
	this.Data["IsFalse"] = false
	this.Data["Title"] = topic.Title + " | " + models.Blogger.BlogName
	this.Data["Topic"] = topic
	this.Data["Domain"] = Domain
	this.Data["Description"] = fmt.Sprintf("%s,%s,%s,blog", topic.Title, models.Blogger.Introduce, models.Blogger.UserName)
	this.Data["Keywords"] = fmt.Sprintf("%s,%s,%s,%s,%s", topic.Title, topic.CategoryID, strings.Join(topic.TagIDs, ","), models.Blogger.Introduce, models.Blogger.UserName)
}
