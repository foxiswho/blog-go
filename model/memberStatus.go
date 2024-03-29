package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type MemberStatus struct {
	StatusId       int       `form:"status_id" json:"status_id" xorm:"not null pk autoincr INT(11)"`
	Uid            int       `form:"uid" json:"uid" xorm:"not null default 0 index INT(11)"`
	RegIp          string    `form:"reg_ip" json:"reg_ip" xorm:"not null default '' CHAR(15)"`
	RegTime        time.Time `form:"reg_time" json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	RegType        int       `form:"reg_type" json:"reg_type" xorm:"not null default 0 INT(11)"`
	RegAppId       int       `form:"reg_app_id" json:"reg_app_id" xorm:"not null default 1 INT(11)"`
	LastLoginIp    string    `form:"last_login_ip" json:"last_login_ip" xorm:"not null default '' CHAR(15)"`
	LastLoginTime  time.Time `form:"last_login_time" json:"last_login_time" xorm:"TIMESTAMP"`
	LastLoginAppId int       `form:"last_login_app_id" json:"last_login_app_id" xorm:"not null default 0 INT(11)"`
	Login          int       `form:"login" json:"login" xorm:"not null default 0 SMALLINT(5)"`
	IsMobile       int       `form:"is_mobile" json:"is_mobile" xorm:"not null default 0 TINYINT(1)"`
	IsEmail        int       `form:"is_email" json:"is_email" xorm:"not null default 0 TINYINT(1)"`
	AidAdd         int       `form:"aid_add" json:"aid_add" xorm:"not null default 0 INT(11)"`
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
		return nil, fox.NewError( err.Error())
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
		return nil, fox.NewError( err.Error())
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
