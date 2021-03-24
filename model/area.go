package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
)

type Area struct {
	Id              int    `form:"id" json:"id" xorm:"not null pk autoincr INT(11)"`
	Name            string `form:"name" json:"name" xorm:"default '' CHAR(50)"`
	NameEn          string `form:"name_en" json:"name_en" xorm:"default '' VARCHAR(100)"`
	ParentId        int    `form:"parent_id" json:"parent_id" xorm:"default 0 index INT(11)"`
	Type            int    `form:"type" json:"type" xorm:"default 0 TINYINT(4)"`
	NameTraditional string `form:"name_traditional" json:"name_traditional" xorm:"default '' VARCHAR(50)"`
	Sort            int    `form:"sort" json:"sort" xorm:"default 0 INT(11)"`
}

//初始化
func NewArea() *Area {
	return new(Area)
}

//初始化列表
func (c *Area) newMakeDataArr() []Area {
	return make([]Area, 0)
}

//列表查询
func (c *Area) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Area) GetById(id int) (*Area, error) {
	m := NewArea()

	m.Id = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Area) Delete(id int) (int64, error) {
	m := NewArea()

	m.Id = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
