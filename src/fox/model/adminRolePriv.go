package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
)

type AdminRolePriv struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(10)" orm:"column(id)"`
	RoleId int    `json:"role_id" xorm:"not null default 0 index index(role_id_2) SMALLINT(3)" orm:"column(role_id)"`
	S      string `json:"s" xorm:"not null default '' index(role_id_2) CHAR(100)" orm:"column(s)"`
	Data   string `json:"data" xorm:"not null default '' CHAR(50)" orm:"column(data)"`
	Aid    int    `json:"aid" xorm:"not null default 0 INT(10)" orm:"column(aid)"`
	MenuId int    `json:"menu_id" xorm:"not null default 0 INT(10)" orm:"column(menu_id)"`
	Type   string `json:"type" xorm:"not null default 'url' CHAR(32)" orm:"column(type)"`
}

//初始化
func NewAdminRolePriv() *AdminRolePriv {
	return new(AdminRolePriv)
}

//初始化列表
func (c *AdminRolePriv) newMakeDataArr() []AdminRolePriv {
	return make([]AdminRolePriv, 0)
}

//列表查询
func (c *AdminRolePriv) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Page, error) {
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
func (c *AdminRolePriv) GetById(id int) (*AdminRolePriv, error) {
	m := NewAdminRolePriv()

	m.Id = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *AdminRolePriv) Delete(id int) (int64, error) {
	m := NewAdminRolePriv()

	m.Id = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
