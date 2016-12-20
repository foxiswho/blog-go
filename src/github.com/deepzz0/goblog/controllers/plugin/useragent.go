package plugin

import (
	"fmt"
	"time"

	"github.com/deepzz0/go-com/useragent"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type UserAgent struct {
	Plugin
}

func (this *UserAgent) Get() {
	this.TplName = "plugin/useragent.html"
	this.Data["Title"] = "UserAgent Parser, Go | " + models.Blogger.BlogName
	this.Data["BlogName"] = models.Blogger.BlogName
	this.Data["Year"] = time.Now().Format("2006")
	this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
	this.Data["Keywords"] = fmt.Sprintf("useragent,golang,go,用户代理解析器, %s, %s", models.Blogger.Introduce, models.Blogger.BlogName)
	this.Data["Description"] = fmt.Sprintf("useragent,parser,%s,%s", models.Blogger.Introduce, models.Blogger.BlogName)
}

func (this *UserAgent) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	ua := this.GetString("ua")
	if ua == "" {
		resp.Status = RS.RS_params_error
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|参数错误。"}
		return
	}
	resp.Data = useragent.ParseByString(ua)
}
