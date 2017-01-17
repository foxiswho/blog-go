package admin

import (
	"blog/fox/str"
	"encoding/json"
	"fmt"
	"blog/fox"
	"blog/service/admin"
	"blog/fox/config"
)
//后台 Session 处理控制器
type AdminSession struct {
	fox.Controller              //继承 控制器
	Session *admin.AdminSession //当前登录用户信息
}
// session 填充
func (c *AdminSession) SessionSet(session *admin.AdminSession) {
	//获取 配置文件中 Session 名字
	SESSION_NAME := config.String("session_name")
	//结构体序列化为json,为存入 Session做准备
	str2, err := str.JsonEnCode(session)
	//错误检测
	if err != nil {
		fmt.Println("session:" + err.Error())
		c.Error(err.Error())
		return
	}
	//存入
	c.SetSession(SESSION_NAME, str2)
}
//获取
func (c *AdminSession) SessionGet() (*admin.AdminSession, error) {
	//获取 配置文件中 Session 名字
	SESSION_NAME := config.String("session_name")
	//获取值为SESSION_NAME 的session，并转换为字符串型
	session, ok := c.GetSession(SESSION_NAME).(string)
	//打印
	fmt.Println("session:", session)
	fmt.Println("ok bool:", ok)
	//判断
	if !ok {
		return nil, fox.NewError("Session 不存在")
	}
	//判断
	if ok && session == "" {
		return nil, fox.NewError("Session 为空")
	}
	//初始化
	Sess := admin.NewAdminSessionService()
	//字符串 转换为 session 结构体并赋值
	err := json.Unmarshal([]byte(session), &Sess)
	//错误检测
	if err != nil {
		return nil, fox.NewError("Session 序列号转换错误. " + err.Error())
	}
	return Sess, nil
}
//删除
func (c *AdminSession) SessionDel() {
	//获取 配置文件中 Session 名字
	SESSION_NAME := config.String("session_name")
	//删除 此session
	c.DelSession(SESSION_NAME)
}
//错误信息
func (c *AdminSession) Error(key string, def ...map[string]interface{}) {
	//检测是否为 ajax传输过来，输出相应格式的数据
	if c.IsAjax() {
		//JSON格式
		c.ErrorJson(key, def...)
	} else {
		//页面格式
		c.Data["content"] = key
		c.TplName = "error/404.html"
	}
}
//成功
func (c *AdminSession) Success(key string, def ...map[string]interface{}) {
	//检测是否为 ajax传输过来，输出相应格式的数据
	if c.IsAjax() {
		//JSON格式
		c.SuccessJson(key, def...)
	} else {
		//页面
		c.Data["content"] = key
		c.TplName = "error/success.html"
	}
}
//@key 内容
//@def 第一个为map 形式数据变量
//错误信息JSON
func (c *AdminSession) ErrorJson(key string, def ...map[string]interface{}) {
	m := make(map[string]interface{})
	if len(def) > 0 {
		m = def[0]
	}
	m["info"] = key
	m["code"] = 0
	c.Data["json"] = m
	//输出 json 内的内容
	c.ServeJSON()
}
//@key 内容
//@def 第一个为map 形式数据变量
//成功JSON
func (c *AdminSession) SuccessJson(key string, def ...map[string]interface{}) {
	m := make(map[string]interface{})
	if len(def) > 0 {
		m = def[0]
	}
	m["info"] = key
	m["code"] = 1
	c.Data["json"] = m
	//输出 json 内的内容
	c.ServeJSON()
}