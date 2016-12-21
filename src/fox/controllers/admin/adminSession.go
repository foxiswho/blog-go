package admin

import (
	"github.com/astaxie/beego"
	"fox/util/str"
	"encoding/json"
	"fmt"
	"fox/util"
	"fox/service"
)

type AdminSession struct {
	beego.Controller
	Session *service.AdminSession //当前登录用户信息
}
// session 填充
func (this *AdminSession) SessionSet(session *service.AdminSession) {
	SESSION_NAME := beego.AppConfig.String("session_name")
	//存入 Session
	str2 := str.JsonEnCode(session)
	//fmt.Println("str => ?", str2)
	this.SetSession(SESSION_NAME, str2)
}
//获取
func (this *AdminSession) SessionGet() (*service.AdminSession, error) {
	SESSION_NAME := beego.AppConfig.String("session_name")
	session, ok := this.GetSession(SESSION_NAME).(string)
	fmt.Println("session:", session)
	fmt.Println("ok bool:", ok)
	if !ok {
		return nil, &util.Error{Msg:"Session 不存在"}
	}
	if ok && session == "" {
		return nil, &util.Error{Msg:"Session 为空"}
	}
	var Sess *service.AdminSession
	err := json.Unmarshal([]byte(session), &Sess)
	if err != nil {
		return nil, &util.Error{Msg:"Session 序列号转换错误. " + err.Error()}
	}
	return Sess, nil
}
//删除
func (this *AdminSession) SessionDel(){
	SESSION_NAME := beego.AppConfig.String("session_name")
	this.DelSession(SESSION_NAME)
}