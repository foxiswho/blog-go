package admin

import (
	"github.com/astaxie/beego"
	"fox/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BaseNoLoginController struct {
	beego.Controller
	adminUser *models.Admin //当前登录用户信息
}

func (this *BaseNoLoginController) Prepare() {
	this.Data["__public__"] = "/"
	this.Data["__theme__"] = "/static/Hplus-v.4.1.0/"

	//orm.RunSyncdb("default", false, true)
}
func init() {
	beego.Info("init orm start...")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_name := beego.AppConfig.String("db_name")
	db_type := beego.AppConfig.String("db_type")
	dsn:=db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_name+"?charset=utf8"
	// set default database
	orm.RegisterDataBase("default", db_type,dsn , 30)
}