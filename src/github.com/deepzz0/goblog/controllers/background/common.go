package background

import (
	"github.com/astaxie/beego"
	"github.com/deepzz0/goblog/cache"
	// "github.com/deepzz0/go-com/log"
)

var SESSIONNAME = beego.AppConfig.String("sessionname")

type Common struct {
	beego.Controller
	index string
}

func (this *Common) Prepare() {
	this.Layout = "manage/adminlayout.html"
}
func (this *Common) LeftBar(index string) {
	this.Data["Choose"] = index
	this.Data["LeftBar"] = cache.Cache.BackgroundLeftBars
}
