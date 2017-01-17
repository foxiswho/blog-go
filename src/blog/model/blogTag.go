package model

import (
	"fmt"
	"blog/fox"
	"blog/fox/db"
	"time"
)

type BlogTag struct {
	TagId   int       `json:"tag_id" xorm:"not null pk autoincr INT(11)"`
	Name    string    `json:"name" xorm:"not null default '' CHAR(100)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP <-"`
	Aid     int       `json:"aid" xorm:"not null default 0 INT(11)"`
	BlogId  int       `json:"blog_id" xorm:"not null default 0 INT(11)"`
}

//初始化
func NewBlogTag() *BlogTag {
	return new(BlogTag)
}

//初始化列表
func (c *BlogTag) newMakeDataArr() []BlogTag {
	return make([]BlogTag, 0)
}

//列表查询
func (c *BlogTag) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *BlogTag) GetById(id int) (*BlogTag, error) {
	m := NewBlogTag()

	m.TagId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *BlogTag) Delete(id int) (int64, error) {
	m := NewBlogTag()

	m.TagId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
