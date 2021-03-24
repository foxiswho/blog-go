package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
)

type AreaExt struct {
	ExtId           int    `form:"ext_id" json:"ext_id" xorm:"not null pk autoincr INT(11)"`
	Id              int    `form:"id" json:"id" xorm:"default 0 index(id) INT(11)"`
	Name            string `form:"name" json:"name" xorm:"default '' CHAR(50)"`
	NameEn          string `form:"name_en" json:"name_en" xorm:"default '' VARCHAR(100)"`
	ParentId        int    `form:"parent_id" json:"parent_id" xorm:"default 0 index(id) INT(11)"`
	Type            int    `form:"type" json:"type" xorm:"default 0 TINYINT(4)"`
	NameTraditional string `form:"name_traditional" json:"name_traditional" xorm:"default '' VARCHAR(50)"`
	Sort            int    `form:"sort" json:"sort" xorm:"default 0 INT(11)"`
	TypeName        string `form:"type_name" json:"type_name" xorm:"default '' VARCHAR(50)"`
	OtherName       string `form:"other_name" json:"other_name" xorm:"default '' VARCHAR(50)"`
}

//初始化
func NewAreaExt() *AreaExt {
	return new(AreaExt)
}

//初始化列表
func (c *AreaExt) newMakeDataArr() []AreaExt {
	return make([]AreaExt, 0)
}

//列表查询
func (c *AreaExt) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *AreaExt) GetById(id int) (*AreaExt, error) {
	m := NewAreaExt()

	m.ExtId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *AreaExt) Delete(id int) (int64, error) {
	m := NewAreaExt()

	m.ExtId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
