package member

import (
	"time"
	"blog/model"
)
//前台session
type Session struct {
	Uid      int       `json:"uid" xorm:"not null pk autoincr INT(11)"`
	Mobile   string    `json:"mobile" xorm:"not null default '' index CHAR(11)"`
	Username string    `json:"username" xorm:"not null default '' index CHAR(30)"`
	Mail     string    `json:"mail" xorm:"not null default '' index CHAR(32)"`
	RegIp    string    `json:"reg_ip" xorm:"not null default '' CHAR(15)"`
	RegTime  time.Time `json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	GroupId  int       `json:"group_id" xorm:"not null default 410 index INT(11)"`
	TrueName string    `json:"true_name" xorm:"not null default '' VARCHAR(32)"`
	Name     string    `json:"name" xorm:"not null default '' VARCHAR(100)"`
	Ip       string       `json:"ip" xorm:"not null default '' CHAR(15)"`
}
//快速初始化
func NewSessionService() *Session {
	return new(Session)
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
func (c *Session) Convert(user *model.Member) *Session {
	//赋值
	Session := &Session{}
	Session.Username = user.Username
	Session.Uid = user.Uid
	Session.Mail = user.Mail
	Session.RegTime = user.RegTime
	Session.RegIp = user.RegIp
	//Session.Ip = user.Ip
	Session.TrueName = user.TrueName
	Session.Mobile = user.Mobile
	return Session
}
