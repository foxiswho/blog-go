package model

import (
	"blog/fox"
	"blog/fox/db"
	"fmt"
	"time"
)

type Session struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uid        int       `json:"uid" xorm:"not null default 0 index(uid) INT(11)"`
	Ip         string    `json:"ip" xorm:"not null default '' CHAR(15)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ErrorCount int       `json:"error_count" xorm:"not null default 0 TINYINT(1)"`
	AppId      int       `json:"app_id" xorm:"not null default 0 INT(11)"`
	TypeLogin  int       `json:"type_login" xorm:"not null default 0 index(uid) INT(11)"`
	Md5        string    `json:"md5" xorm:"not null default '' CHAR(32)"`
	TypeClient int       `json:"type_client" xorm:"not null default 0 index(uid) INT(11)"`
}

//初始化
func NewSession() *Session {
	return new(Session)
}

//初始化列表
func (c *Session) newMakeDataArr() []Session {
	return make([]Session, 0)
}

//列表查询
func (c *Session) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Session) GetById(id int) (*Session, error) {
	m := NewSession()

	m.Id = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Session) Delete(id int) (int64, error) {
	m := NewSession()

	m.Id = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
