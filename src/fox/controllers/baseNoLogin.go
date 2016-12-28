package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fox/util/datetime"
	"fox/util/db"
)

type BaseNoLoginController struct {
	beego.Controller
	//Session *service.AdminSession //当前登录用户信息
}

//  框架中的扩展函数
func (this *BaseNoLoginController) Prepare() {
	this.Initialization()
}
// 初始化数据
func (this *BaseNoLoginController) Initialization() {
	this.Data["__public__"] = "/"
	this.Data["__static__"] = "/static/"
	this.Data["__theme__"] = "/static/post/"
	this.Data["blog_name"] = beego.AppConfig.String("blog_name")
	//orm.RunSyncdb("default", false, true)
}
//表单日期时间
func (this *BaseNoLoginController) GetDateTime(key string) (time.Time, bool) {
	date := this.GetString(key)
	if len(date) > 0 {
		date, err := datetime.FormatTimeStructLocation(date, datetime.Y_M_D_H_I_S)
		if err == nil {
			return date, true
		}
	}
	return time.Time{}, false
}
//表单日期时间
func (this *BaseNoLoginController) Error(key string) {
	this.Data["content"] = key
	this.TplName = "error/404.html"
}
//初始化数据库
func init() {
	//初始化
	db.Init();

	beego.Info("init orm start...")
}