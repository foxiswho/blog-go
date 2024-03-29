package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type Type struct {
	Id        int       `form:"id" json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `form:"name" json:"name" xorm:"not null default '' CHAR(100)"`
	Code      string    `form:"code" json:"code" xorm:"not null default '' CHAR(32)"`
	Mark      string    `form:"mark" json:"mark" xorm:"not null default '' index CHAR(32)"`
	TypeId    int       `form:"type_id" json:"type_id" xorm:"not null default 0 index INT(11)"`
	ParentId  int       `form:"parent_id" json:"parent_id" xorm:"not null default 0 index INT(11)"`
	Value     int       `form:"value" json:"value" xorm:"not null default 0 INT(10)"`
	Content   string    `form:"content" json:"content" xorm:"not null default '' VARCHAR(255)"`
	IsDel     int       `form:"is_del" json:"is_del" xorm:"not null default 0 index INT(11)"`
	Sort      int       `form:"sort" json:"sort" xorm:"not null default 0 index INT(11)"`
	Remark    string    `form:"remark" json:"remark" xorm:"VARCHAR(255)"`
	TimeAdd   time.Time `form:"time_add" json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Aid       int       `form:"aid" json:"aid" xorm:"not null default 0 INT(11)"`
	Module    string    `form:"module" json:"module" xorm:"not null default '' CHAR(50)"`
	IsDefault int       `form:"is_default" json:"is_default" xorm:"not null default 0 TINYINT(1)"`
	Setting   string    `form:"setting" json:"setting" xorm:"VARCHAR(255)"`
	IsChild   int       `form:"is_child" json:"is_child" xorm:"not null default 0 TINYINT(1)"`
	IsSystem  int       `form:"is_system" json:"is_system" xorm:"not null default 0 TINYINT(1)"`
	IsShow    int       `form:"is_show" json:"is_show" xorm:"not null default 0 TINYINT(1)"`
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
