package model

import (
	"blog/fox"
	"blog/fox/db"
	"fmt"
	"time"
)

type BlogSyncQueue struct {
	QueueId    int       `json:"queue_id" xorm:"not null pk autoincr INT(11)"`
	BlogId     int       `json:"blog_id" xorm:"not null default 0 INT(11)"`
	TypeId     int       `json:"type_id" xorm:"not null default 0 INT(11)"`
	Status     int       `json:"status" xorm:"not null default 0 TINYINT(3)"`
	TimeUpdate time.Time `json:"time_update" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Msg        string    `json:"msg" xorm:"not null default '' VARCHAR(255)"`
	MapId      int       `json:"map_id" xorm:"not null default 0 INT(11)"`
}

//初始化
func NewBlogSyncQueue() *BlogSyncQueue {
	return new(BlogSyncQueue)
}

//初始化列表
func (c *BlogSyncQueue) newMakeDataArr() []BlogSyncQueue {
	return make([]BlogSyncQueue, 0)
}

//列表查询
func (c *BlogSyncQueue) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *BlogSyncQueue) GetById(id int) (*BlogSyncQueue, error) {
	m := NewBlogSyncQueue()

	m.QueueId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *BlogSyncQueue) Delete(id int) (int64, error) {
	m := NewBlogSyncQueue()

	m.QueueId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
