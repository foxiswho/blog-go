package model

import (
	"blog/fox"
	"blog/fox/db"
	"fmt"
	"time"
)

type BlogSyncMapping struct {
	MapId      int       `json:"map_id" xorm:"not null pk autoincr INT(11)"`
	BlogId     int       `json:"blog_id" xorm:"not null default 0 INT(11)"`
	TypeId     int       `json:"type_id" xorm:"not null default 0 INT(11)"`
	Id         string    `json:"id" xorm:"not null default '' VARCHAR(64)"`
	TimeUpdate time.Time `json:"time_update" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP <-"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP <-"`
	Mark       string    `json:"mark" xorm:"not null default '' CHAR(32)"`
	IsSync     int       `json:"is_sync" xorm:"not null default 0 TINYINT(1)"`
	Extend     string    `json:"extend" xorm:"VARCHAR(5000)"`
}

//初始化
func NewBlogSyncMapping() *BlogSyncMapping {
	return new(BlogSyncMapping)
}

//初始化列表
func (c *BlogSyncMapping) newMakeDataArr() []BlogSyncMapping {
	return make([]BlogSyncMapping, 0)
}

//列表查询
func (c *BlogSyncMapping) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *BlogSyncMapping) GetById(id int) (*BlogSyncMapping, error) {
	m := NewBlogSyncMapping()

	m.MapId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *BlogSyncMapping) Delete(id int) (int64, error) {
	m := NewBlogSyncMapping()

	m.MapId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
