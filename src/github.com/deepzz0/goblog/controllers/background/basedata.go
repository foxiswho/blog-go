package background

import (
	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type DataController struct {
	Common
}

func (this *DataController) Get() {
	this.TplName = "manage/basedata/basedataTemplate.html"
	this.Data["Title"] = "基础数据 | " + models.Blogger.BlogName
	this.LeftBar("data")
	this.Content()
}

func (this *DataController) Content() {

}

func (this *DataController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag = %s", flag)
	switch flag {
	case "base":
		this.Base(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}

func (this *DataController) Base(resp *helper.Response) {
	resp.Data = models.ManageData
}
