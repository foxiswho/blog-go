package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
)

type AdminMenu struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)" orm:"column(id)"`
	Name     string `json:"name" xorm:"not null default '' CHAR(100)" orm:"column(name)"`
	ParentId int    `json:"parent_id" xorm:"not null default 0 index INT(11)" orm:"column(parent_id)"`
	S        string `json:"s" xorm:"not null default '' index CHAR(60)" orm:"column(s)"`
	Data     string `json:"data" xorm:"not null default '' CHAR(100)" orm:"column(data)"`
	Sort     int    `json:"sort" xorm:"not null default 0 index INT(11)" orm:"column(sort)"`
	Remark   string `json:"remark" xorm:"not null default '' VARCHAR(255)" orm:"column(remark)"`
	Type     string `json:"type" xorm:"not null default 'url' CHAR(32)" orm:"column(type)"`
	Level    int    `json:"level" xorm:"not null default 0 TINYINT(3)" orm:"column(level)"`
	Level1Id int    `json:"level1_id" xorm:"not null default 0 INT(11)" orm:"column(level1_id)"`
	Md5      string `json:"md5" xorm:"not null default '' CHAR(32)" orm:"column(md5)"`
	IsShow   int    `json:"is_show" xorm:"not null default 1 TINYINT(1)" orm:"column(is_show)"`
	IsUnique int    `json:"is_unique" xorm:"not null default 0 TINYINT(1)" orm:"column(is_unique)"`
}

//初始化
func NewAdminMenu() *AdminMenu {
	return new(AdminMenu)
}

//初始化列表
func (c *AdminMenu) newMakeDataArr() []AdminMenu {
	return make([]AdminMenu, 0)
}

//列表查询
func (c *AdminMenu) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *AdminMenu) GetById(id int) (*AdminMenu, error) {
	m := NewAdminMenu()

	m.Id = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *AdminMenu) Delete(id int) (int64, error) {
	m := NewAdminMenu()

	m.Id = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
