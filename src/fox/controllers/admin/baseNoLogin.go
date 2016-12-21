package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BaseNoLoginController struct {
	AdminSession
	//beego.Controller
	//Session *service.AdminSession //当前登录用户信息
}

//  框架中的扩展函数
func (this *BaseNoLoginController) Prepare() {
	this.Initialization()
}
// 初始化数据
func (this *BaseNoLoginController) Initialization()  {
	this.Data["__public__"] = "/"
	this.Data["__theme__"] = "/static/Hplus-v.4.1.0/"
	this.Data["blog_name"] = beego.AppConfig.String("blog_name")
	//orm.RunSyncdb("default", false, true)
}
//初始化数据库
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