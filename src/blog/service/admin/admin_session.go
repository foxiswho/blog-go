package admin

import (
	"time"
	"blog/model"
)
//后台session
type AdminSession struct {
	Aid        int       `orm:"column(aid);auto"`
	Username   string    `orm:"column(username);size(30)"`
	Mail       string    `orm:"column(mail);size(80)"`
	TimeAdd    time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"`
	TimeUpdate time.Time `orm:"column(time_update);type(timestamp);null"`
	Ip         string    `orm:"column(ip);size(15)"`
	JobNo      string    `orm:"column(job_no);size(15)"`
	NickName   string    `orm:"column(nick_name);size(50)"`
	TrueName   string    `orm:"column(true_name);size(50)"`
	Qq         string    `orm:"column(qq);size(50)"`
	Phone      string    `orm:"column(phone);size(50)"`
	Mobile     string    `orm:"column(mobile);size(20)"`
	Role_id    map[int]model.AdminRoleAccess //扩展角色
}
//快速初始化
func NewAdminSessionService() *AdminSession{
	return new(AdminSession)
}
/*
// session 填充
func (this *AdminSession) Set(session *AdminSession) {
	var admin *beego.Controller
	SESSION_NAME := beego.AppConfig.String("session_name")
	//存入 Session
	str2 := str.JsonEnCode(session)
	//fmt.Println("str => ?", str2)
	admin.SetSession(SESSION_NAME, str2)
}
//获取
func (this *AdminSession) Get() (*AdminSession, error) {
	var admin *beego.Controller
	SESSION_NAME := beego.AppConfig.String("session_name")
	session, ok := admin.GetSession(SESSION_NAME).(string)
	fmt.Println("session:", session)
	fmt.Println("ok bool:", ok)
	if !ok {
		return nil, &util.Error{Msg:"Session 不存在"}
	}
	if ok && session == "" {
		return nil, &util.Error{Msg:"Session 为空"}
	}
	var Session *AdminSession
	err := json.Unmarshal([]byte(session), &Session)
	if err != nil {
		return nil, &util.Error{Msg:"Session 序列号转换错误. " + err.Error()}
	}
	return Session, nil
}
//删除
func (this *AdminSession) Del(){
	var admin *beego.Controller
	SESSION_NAME := beego.AppConfig.String("session_name")
	admin.DelSession(SESSION_NAME)
}*/
//转换
func (c *AdminSession) Convert(admUser *model.Admin) *AdminSession {
	//赋值
	Session := &AdminSession{}
	Session.Username = admUser.Username
	Session.Aid = admUser.Aid
	Session.Mail = admUser.Mail
	Session.TimeAdd = admUser.TimeAdd
	Session.Ip = admUser.Ip
	Session.NickName = admUser.NickName
	Session.TrueName = admUser.TrueName
	Session.Qq = admUser.Qq
	Session.Phone = admUser.Phone
	Session.Mobile = admUser.Mobile
	return Session
}
