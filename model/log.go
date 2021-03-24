package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type Log struct {
	LogId      int       `form:"log_id" json:"log_id" xorm:"not null pk autoincr INT(11)"`
	Id         int       `form:"id" json:"id" xorm:"not null default 0 index INT(11)"`
	Aid        int       `form:"aid" json:"aid" xorm:"not null default 0 index INT(11)"`
	Uid        int       `form:"uid" json:"uid" xorm:"not null default 0 index INT(11)"`
	TimeAdd    time.Time `form:"time_add" json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Mark       string    `form:"mark" json:"mark" xorm:"not null default '' CHAR(32)"`
	Data       string    `form:"data" json:"data" xorm:"TEXT"`
	No         string    `form:"no" json:"no" xorm:"not null default '' index CHAR(50)"`
	TypeLogin  int       `form:"type_login" json:"type_login" xorm:"not null default 0 index INT(11)"`
	TypeClient int       `form:"type_client" json:"type_client" xorm:"not null default 0 index INT(11)"`
	Ip         string    `form:"ip" json:"ip" xorm:"not null default '' CHAR(20)"`
	Msg        string    `form:"msg" json:"msg" xorm:"VARCHAR(255)"`
}

//初始化
func NewLog() *Log {
	return new(Log)
}

//初始化列表
func (c *Log) newMakeDataArr() []Log {
	return make([]Log, 0)
}

//列表查询
func (c *Log) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Log) GetById(id int) (*Log, error) {
	m := NewLog()

	m.LogId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Log) Delete(id int) (int64, error) {
	m := NewLog()

	m.LogId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
