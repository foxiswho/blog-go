package model

import (
	"blog/fox"
	"blog/fox/db"
	"fmt"
	"time"
)

type Type struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"not null default '' CHAR(100)"`
	Code      string    `json:"code" xorm:"not null default '' CHAR(32)"`
	Mark      string    `json:"mark" xorm:"not null default '' index CHAR(32)"`
	TypeId    int       `json:"type_id" xorm:"not null default 0 index INT(11)"`
	ParentId  int       `json:"parent_id" xorm:"not null default 0 index INT(11)"`
	Value     int       `json:"value" xorm:"not null default 0 INT(10)"`
	Content   string    `json:"content" xorm:"not null default '' VARCHAR(255)"`
	IsDel     int       `json:"is_del" xorm:"not null default 0 index INT(11)"`
	Sort      int       `json:"sort" xorm:"not null default 0 index INT(11)"`
	Remark    string    `json:"remark" xorm:"VARCHAR(255)"`
	TimeAdd   time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Aid       int       `json:"aid" xorm:"not null default 0 INT(11)"`
	Module    string    `json:"module" xorm:"not null default '' CHAR(50)"`
	IsDefault int       `json:"is_default" xorm:"not null default 0 TINYINT(1)"`
	Setting   string    `json:"setting" xorm:"VARCHAR(255)"`
	IsChild   int       `json:"is_child" xorm:"not null default 0 TINYINT(1)"`
	IsSystem  int       `json:"is_system" xorm:"not null default 0 TINYINT(1)"`
	IsShow    int       `json:"is_show" xorm:"not null default 0 TINYINT(1)"`
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
func (c *Type) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	session := db.Filter(q)
	count, err := session.Count(c)
	if err != nil {
		fmt.Println(err)
		return nil,fox.NewError( err.Error())
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
		return nil,fox.NewError( err.Error())
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
