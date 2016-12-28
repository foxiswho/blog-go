package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
	"time"
)

type Type struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)" orm:"column(id)"`
	Name      string    `json:"name" xorm:"not null default '' CHAR(100)" orm:"column(name)"`
	Code      string    `json:"code" xorm:"not null default '' CHAR(32)" orm:"column(code)"`
	Mark      string    `json:"mark" xorm:"not null default '' index CHAR(32)" orm:"column(mark)"`
	TypeId    int       `json:"type_id" xorm:"not null default 0 index INT(11)" orm:"column(type_id)"`
	ParentId  int       `json:"parent_id" xorm:"not null default 0 index INT(11)" orm:"column(parent_id)"`
	Value     int       `json:"value" xorm:"not null default 0 INT(10)" orm:"column(value)"`
	IsDel     int       `json:"is_del" xorm:"not null default 0 index INT(11)" orm:"column(is_del)"`
	Sort      int       `json:"sort" xorm:"not null default 0 index INT(11)" orm:"column(sort)"`
	Remark    string    `json:"remark" xorm:"VARCHAR(255)" orm:"column(remark)"`
	TimeAdd   time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" orm:"column(time_add)"`
	Aid       int       `json:"aid" xorm:"not null default 0 INT(11)" orm:"column(aid)"`
	Module    string    `json:"module" xorm:"not null default '' CHAR(50)" orm:"column(module)"`
	IsDefault int       `json:"is_default" xorm:"not null default 0 TINYINT(1)" orm:"column(is_default)"`
	Setting   string    `json:"setting" xorm:"VARCHAR(255)" orm:"column(setting)"`
	IsChild   int       `json:"is_child" xorm:"not null default 0 TINYINT(1)" orm:"column(is_child)"`
	IsSystem  int       `json:"is_system" xorm:"not null default 0 TINYINT(1)" orm:"column(is_system)"`
	IsShow    int       `json:"is_show" xorm:"not null default 0 TINYINT(1)" orm:"column(is_show)"`
}

//初始化
func NewType() *Type {
	return new(Type)
}

//初始化列表
func (c *Type) newMakeDataArr() []Type {
	return make([]Type, 0)
}

//列表查询
func (c *Type) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Page, error) {
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
func (c *Type) GetById(id int) (*Type, error) {
	m := NewType()

	m.Id = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Type) Delete(id int) (int64, error) {
	m := NewType()

	m.Id = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
