package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
	"time"
)

type MemberStatus struct {
	StatusId       int       `json:"status_id" xorm:"not null pk autoincr INT(11)" orm:"column(status_id)"`
	Uid            int       `json:"uid" xorm:"not null default 0 index INT(11)" orm:"column(uid)"`
	RegIp          string    `json:"reg_ip" xorm:"not null default '' CHAR(15)" orm:"column(reg_ip)"`
	RegTime        time.Time `json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP" orm:"column(reg_time)"`
	RegType        int       `json:"reg_type" xorm:"not null default 0 INT(11)" orm:"column(reg_type)"`
	RegAppId       int       `json:"reg_app_id" xorm:"not null default 1 INT(11)" orm:"column(reg_app_id)"`
	LastLoginIp    string    `json:"last_login_ip" xorm:"not null default '' CHAR(15)" orm:"column(last_login_ip)"`
	LastLoginTime  time.Time `json:"last_login_time" xorm:"TIMESTAMP" orm:"column(last_login_time)"`
	LastLoginAppId int       `json:"last_login_app_id" xorm:"not null default 0 INT(11)" orm:"column(last_login_app_id)"`
	Login          int       `json:"login" xorm:"not null default 0 SMALLINT(5)" orm:"column(login)"`
	IsMobile       int       `json:"is_mobile" xorm:"not null default 0 TINYINT(1)" orm:"column(is_mobile)"`
	IsEmail        int       `json:"is_email" xorm:"not null default 0 TINYINT(1)" orm:"column(is_email)"`
	AidAid         int       `json:"aid_aid" xorm:"not null default 0 INT(11)" orm:"column(aid_aid)"`
}

//初始化
func NewMemberStatus() *MemberStatus {
	return new(MemberStatus)
}

//初始化列表
func (c *MemberStatus) newMakeDataArr() []MemberStatus {
	return make([]MemberStatus, 0)
}

//列表查询
func (c *MemberStatus) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	session := db.Filter(q)
	count, err := session.Count(c)
	if err != nil {
		fmt.Println(err)
		return nil, &util.Error{Msg: err.Error()}
	}
	Query := db.Pagination(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}

	session = db.Filter(q)
	if orderBy != "" {
		session.OrderBy(orderBy)
	}
	session.Limit(limit, Query.Offset)
	if len(fields) == 0 {
		session.AllCols()
	}
	data := c.newMakeDataArr()
	err = session.Find(&data)
	if err != nil {
		fmt.Println(err)
		return nil, &util.Error{Msg: err.Error()}
	}
	Query.Data = make([]interface{}, len(data))
	for y, x := range data {
		Query.Data[y] = x
	}
	return Query, nil
}

// 获取 单条记录
func (c *MemberStatus) GetById(id int) (*MemberStatus, error) {
	m := NewMemberStatus()

	m.StatusId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *MemberStatus) Delete(id int) (int64, error) {
	m := NewMemberStatus()

	m.StatusId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
