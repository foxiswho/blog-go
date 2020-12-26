package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
)

type BlogStatistics struct {
	StatisticsId   int    `json:"statistics_id" xorm:"not null pk autoincr INT(11)"`
	BlogId         int    `json:"blog_id" xorm:"not null default 0 index INT(11)"`
	Comment        int    `json:"comment" xorm:"not null default 0 INT(11)"`
	Read           int    `json:"read" xorm:"not null default 0 INT(11)"`
	SeoTitle       string `json:"seo_title" xorm:"not null default '' VARCHAR(255)"`
	SeoDescription string `json:"seo_description" xorm:"not null default '' VARCHAR(255)"`
	SeoKeyword     string `json:"seo_keyword" xorm:"not null default '' VARCHAR(255)"`
}

//初始化
func NewBlogStatistics() *BlogStatistics {
	return new(BlogStatistics)
}

//初始化列表
func (c *BlogStatistics) newMakeDataArr() []BlogStatistics {
	return make([]BlogStatistics, 0)
}

//列表查询
func (c *BlogStatistics) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *BlogStatistics) GetById(id int) (*BlogStatistics, error) {
	m := NewBlogStatistics()

	m.StatisticsId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *BlogStatistics) Delete(id int) (int64, error) {
	m := NewBlogStatistics()

	m.StatisticsId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
