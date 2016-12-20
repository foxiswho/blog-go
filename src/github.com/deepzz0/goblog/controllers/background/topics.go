package background

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type TopicsController struct {
	Common
}

func (this *TopicsController) Get() {
	this.TplName = "manage/topic/topicTemplate.html"
	this.Data["Title"] = "博文管理 | " + models.Blogger.BlogName
	this.LeftBar("topics")
	this.Content()
}

func (this *TopicsController) Content() {
	this.Data["Categories"] = models.Blogger.GetValidCategory()
	this.Data["ClassOlder"] = "disabled"
	this.Data["UrlOlder"] = "#"
	this.Data["ClassNewer"] = "disabled"
	this.Data["UrlNewer"] = "#"
	cat := this.GetString("cat")
	this.Data["ChooseCat"] = cat
	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}
	var pageTopics []*models.Topic
	var remainpage int
	if cat == "" || cat == "all" {
		pageTopics, remainpage = models.TMgr.GetTopicsByPage(page)
	} else {
		pageTopics, remainpage = models.TMgr.GetTopicsByCatgory(cat, page)
	}
	log.Debug(remainpage, page)
	if page > 1 {
		this.Data["ClassNewer"] = ""
		this.Data["UrlNewer"] = fmt.Sprintf("/admin/topics?cat=%s&p=%d", cat, page-1)
	}
	if remainpage > 0 {
		this.Data["ClassOlder"] = ""
		this.Data["UrlOlder"] = fmt.Sprintf("/admin/topics?cat=%s&p=%d", cat, page+1)
	}
	this.Data["Topics"] = pageTopics
	var style string
	for _, t := range models.Blogger.Tags {
		style += t.TagStyle()
	}
	this.Data["TagStyle"] = style
}

func (this *TopicsController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag = %s", flag)
	switch flag {
	case "save":
		this.saveTopic(resp)
	case "modify":
		this.getTopic(resp)
	case "delete":
		this.doDeleteTopic(resp)
	case "deleteall":
		this.doDeleteTopics(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}

func (this *TopicsController) saveTopic(resp *helper.Response) {
	operate := this.GetString("operate")
	if operate == "" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|操作类型错误。"}
		return
	}
	title := this.GetString("title")
	content := this.GetString("content")
	cat := this.GetString("cat")
	tags := this.GetString("tags")
	if title == "" || content == "" || cat == "" {
		log.Debugf("%s,%s,%s,%s, %s", title, content, cat, tags, operate)
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|请检查是否填写完整。"}
		return
	}
	if operate == "new" {
		topic := models.NewTopic()
		if strings.HasSuffix(title, "TAG:aboutme") {
			topic.ID = 1
			topic.Title = strings.Split(title, "-")[0]
		} else {
			if topic.ID == 1 {
				topic.ID = models.NextVal()
			}
			topic.Title = title
		}
		topic.Content = content
		if category := models.Blogger.GetCategoryByID(cat); category == nil {
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|查找不到该分类。"}
			return
		} else {
			topic.CategoryID = category.ID
		}
		if tags == "" {
			topic.TagIDs = make([]string, 0)
		} else {
			sliceTags := strings.Split(tags, ",")
			for _, tag := range sliceTags {
				topic.TagIDs = append(topic.TagIDs, tag)
			}
		}
		if err := models.TMgr.AddTopic(topic); err != nil {
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|" + err.Error()}
			return
		}
	} else {
		id, err := strconv.Atoi(operate)
		if err != nil {
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|修改文章ID解析失败。"}
			return
		}
		if t := models.TMgr.GetTopic(int32(id)); t == nil {
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|系统查找不到该文章ID。"}
			return
		} else {
			t.Title = title
			t.Content = content
			if err := models.TMgr.ModTopic(t, cat, tags); err != nil {
				resp.Status = RS.RS_failed
				resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|" + err.Error()}
				return
			}
		}
	}
}

func (this *TopicsController) getTopic(resp *helper.Response) {
	id, err := this.GetInt("id")
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|ID格式不正确。"}
		return
	}
	if topic, err := models.TMgr.LoadTopic(int32(id)); err != nil || topic == nil {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|系统未查询到该文章。"}
		return
	} else {
		resp.Data = topic
	}
}

func (this *TopicsController) doDeleteTopic(resp *helper.Response) {
	id, err := this.GetInt("id")
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "ID错误|走正常途径哦。"}
		return
	}
	err = models.TMgr.DelTopic(int32(id))
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "删除失败|" + err.Error()}
		return
	}
}

func (this *TopicsController) doDeleteTopics(resp *helper.Response) {
	ids := this.GetString("ids")
	log.Debugf("%s", ids)
	if ids == "" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "ID错误|走正常途径哦。"}
		return
	}
	sliceID := strings.Split(ids, ",")
	for _, v := range sliceID {
		id, err := strconv.Atoi(v)
		if err != nil {
			log.Error(err)
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "ID错误|走正常途径哦。"}
			return
		}
		err = models.TMgr.DelTopic(int32(id))
		if err != nil {
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "删除失败|" + err.Error()}
			return
		}
	}
}
