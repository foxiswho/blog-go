package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fox/util/datetime"
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
	beego.Info("init orm start...")
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC

	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_name := beego.AppConfig.String("db_name")
	db_type := beego.AppConfig.String("db_type")
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8&loc=Asia%2FShanghai"
	// set default database
	orm.RegisterDataBase("default", db_type, dsn, 30)
}
