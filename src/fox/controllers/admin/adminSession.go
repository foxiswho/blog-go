package admin

import (
	"github.com/astaxie/beego"
	"fox/util/str"
	"encoding/json"
	"fmt"
	"fox/util"
	"fox/service/admin"
)

type AdminSession struct {
	beego.Controller
	Session *admin.AdminSession //当前登录用户信息
}
// session 填充
func (c *AdminSession) SessionSet(session *admin.AdminSession) {
	SESSION_NAME := beego.AppConfig.String("session_name")
	//存入 Session
	str2 := str.JsonEnCode(session)
	//fmt.Println("str => ?", str2)
	c.SetSession(SESSION_NAME, str2)
}
//获取
func (c *AdminSession) SessionGet() (*admin.AdminSession, error) {
	SESSION_NAME := beego.AppConfig.String("session_name")
	session, ok := c.GetSession(SESSION_NAME).(string)
	fmt.Println("session:", session)
	fmt.Println("ok bool:", ok)
	if !ok {
		return nil, &util.Error{Msg:"Session 不存在"}
	}
	if ok && session == "" {
		return nil, &util.Error{Msg:"Session 为空"}
	}
	Sess :=admin.NewAdminSessionService()
	err := json.Unmarshal([]byte(session), &Sess)
	if err != nil {
		return nil, &util.Error{Msg:"Session 序列号转换错误. " + err.Error()}
	}
	return Sess, nil
}
//删除
func (c *AdminSession) SessionDel(){
	SESSION_NAME := beego.AppConfig.String("session_name")
	c.DelSession(SESSION_NAME)
}
//表单日期时间
func (c *AdminSession) Error(key string) {
	c.Data["content"] = key
	if c.IsAjax() {
		c.ServeJSON()
	}else{
		c.TplName = "error/404.html"
	}
}