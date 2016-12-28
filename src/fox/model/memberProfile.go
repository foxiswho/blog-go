package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
)

type MemberProfile struct {
	ProfileId   int    `json:"profile_id" xorm:"not null pk autoincr INT(11)" orm:"column(profile_id)"`
	Uid         int    `json:"uid" xorm:"not null default 0 index INT(11)" orm:"column(uid)"`
	Sex         int    `json:"sex" xorm:"not null default 1 TINYINT(1)" orm:"column(sex)"`
	Job         string `json:"job" xorm:"not null default '' VARCHAR(50)" orm:"column(job)"`
	Qq          string `json:"qq" xorm:"not null default '' VARCHAR(20)" orm:"column(qq)"`
	Phone       string `json:"phone" xorm:"not null default '' VARCHAR(20)" orm:"column(phone)"`
	County      int    `json:"county" xorm:"not null default 1 INT(11)" orm:"column(county)"`
	Province    int    `json:"province" xorm:"not null default 0 INT(11)" orm:"column(province)"`
	City        int    `json:"city" xorm:"not null default 0 INT(11)" orm:"column(city)"`
	District    int    `json:"district" xorm:"not null default 0 INT(11)" orm:"column(district)"`
	Address     string `json:"address" xorm:"not null default '' VARCHAR(255)" orm:"column(address)"`
	Wechat      string `json:"wechat" xorm:"not null default '' VARCHAR(20)" orm:"column(wechat)"`
	RemarkAdmin string `json:"remark_admin" xorm:"TEXT" orm:"column(remark_admin)"`
}

//初始化
func NewMemberProfile() *MemberProfile {
	return new(MemberProfile)
}

//初始化列表
func (c *MemberProfile) newMakeDataArr() []MemberProfile {
	return make([]MemberProfile, 0)
}

//列表查询
func (c *MemberProfile) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Page, error) {
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
func (c *MemberProfile) GetById(id int) (*MemberProfile, error) {
	m := NewMemberProfile()

	m.ProfileId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *MemberProfile) Delete(id int) (int64, error) {
	m := NewMemberProfile()

	m.ProfileId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
