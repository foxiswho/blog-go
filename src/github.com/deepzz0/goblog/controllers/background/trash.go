package background

import (
	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type TrashController struct {
	Common
}

func (this *TrashController) Get() {
	this.TplName = "manage/trash/trash.html"
	this.Data["Title"] = "回收箱 | " + models.Blogger.BlogName
	this.LeftBar("trash")
	this.Content()
}

func (this *TrashController) Content() {
	this.Data["DelTopics"] = models.TMgr.DeleteTopics
}

func (this *TrashController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag=%s", flag)
	switch flag {
	case "delete":
		this.doDelete(resp)
	case "restore":
		this.doRestore(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}

func (this *TrashController) doRestore(resp *helper.Response) {
	id, err := this.GetInt32("id")
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Tips(helper.WARNING, RS.RS_params_error)
		return
	}
	if topic := models.TMgr.GetWaitDelTopic(id); topic == nil {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.RS_not_found)
		return
	} else {
		if code := models.TMgr.RestoreTopic(topic); code != RS.RS_success {
			resp.Status = code
			resp.Tips(helper.WARNING, code)
			return
		}
	}
}

func (this *TrashController) doDelete(resp *helper.Response) {
	id, err := this.GetInt32("id")
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Tips(helper.WARNING, RS.RS_params_error)
		return
	}
	if topic := models.TMgr.GetWaitDelTopic(id); topic == nil {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.RS_not_found)
	} else {
		if err := models.TMgr.ImmeDelTopic(topic); err != nil {
			resp.Status = RS.RS_failed
			resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|" + err.Error()}
			return
		}
	}
}
