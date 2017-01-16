package admin

import (
	"github.com/astaxie/beego"
	"blog/fox/str"
	"encoding/json"
	"fmt"
	"blog/fox"
	"blog/service/admin"
)
//后台 Session 处理
type AdminSession struct {
	beego.Controller            //继承 beego 控制器
	Session *admin.AdminSession //当前登录用户信息
}
// session 填充
func (c *AdminSession) SessionSet(session *admin.AdminSession) {
	SESSION_NAME := beego.AppConfig.String("session_name")
	//存入 Session
	str2, err := str.JsonEnCode(session)
	if err != nil {
		fmt.Println("session:" + err.Error())
	}
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
		return nil, &fox.Error{Msg:"Session 不存在"}
	}
	if ok && session == "" {
		return nil, &fox.Error{Msg:"Session 为空"}
	}
	Sess := admin.NewAdminSessionService()
	err := json.Unmarshal([]byte(session), &Sess)
	if err != nil {
		return nil, &fox.Error{Msg:"Session 序列号转换错误. " + err.Error()}
	}
	return Sess, nil
}
//删除
func (c *AdminSession) SessionDel() {
	SESSION_NAME := beego.AppConfig.String("session_name")
	c.DelSession(SESSION_NAME)
}
//错误信息
func (c *AdminSession) Error(key string) {
	c.Data["content"] = key
	if c.IsAjax() {
		type Msg struct {
			Info string `json:"info"`
			Code int    `json:"code"`
			Data interface{} `json:"data"`
		}
		msg := &Msg{}
		msg.Code = 0
		msg.Info = key
		msg.Data = c.Data["Data"]
		c.Data["json"] = msg
		c.ServeJSON()
	} else {
		c.TplName = "error/404.html"
	}
}
func (c *AdminSession) Success(key string) {
	c.Data["content"] = key
	if c.IsAjax() {
		type Msg struct {
			Info string `json:"info"`
			Code int    `json:"code"`
			Data interface{} `json:"data"`
		}
		msg := &Msg{}
		msg.Code = 1
		msg.Info = key
		msg.Data = c.Data["Data"]
		c.Data["json"] = msg
		c.ServeJSON()
	} else {
		c.TplName = "error/success.html"
	}
}