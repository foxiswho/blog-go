package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
)

type MemberProfile struct {
	ProfileId   int    `form:"profile_id" json:"profile_id" xorm:"not null pk autoincr INT(11)"`
	Uid         int    `form:"uid" json:"uid" xorm:"not null default 0 index INT(11)"`
	Sex         int    `form:"sex" json:"sex" xorm:"not null default 1 TINYINT(1)"`
	Job         string `form:"job" json:"job" xorm:"not null default '' VARCHAR(50)"`
	Qq          string `form:"qq" json:"qq" xorm:"not null default '' VARCHAR(20)"`
	Phone       string `form:"phone" json:"phone" xorm:"not null default '' VARCHAR(20)"`
	County      int    `form:"county" json:"county" xorm:"not null default 1 INT(11)"`
	Province    int    `form:"province" json:"province" xorm:"not null default 0 INT(11)"`
	City        int    `form:"city" json:"city" xorm:"not null default 0 INT(11)"`
	District    int    `form:"district" json:"district" xorm:"not null default 0 INT(11)"`
	Address     string `form:"address" json:"address" xorm:"not null default '' VARCHAR(255)"`
	Wechat      string `form:"wechat" json:"wechat" xorm:"not null default '' VARCHAR(20)"`
	RemarkAdmin string `form:"remark_admin" json:"remark_admin" xorm:"TEXT"`
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
func (c *MemberProfile) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
