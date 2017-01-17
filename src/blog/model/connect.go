package model

import (
	"blog/fox"
	"blog/fox/db"
	"fmt"
	"time"
)

type Connect struct {
	ConnectId int       `json:"connect_id" xorm:"not null pk autoincr INT(11)"`
	TypeId    int       `json:"type_id" xorm:"not null default 0 index INT(11)"`
	Uid       int       `json:"uid" xorm:"not null default 0 index INT(11)"`
	OpenId    string    `json:"open_id" xorm:"not null default '' index VARCHAR(80)"`
	Token     string    `json:"token" xorm:"not null default '' VARCHAR(80)"`
	Type      int       `json:"type" xorm:"not null default 1 INT(11)"`
	TypeLogin int       `json:"type_login" xorm:"not null default 0 INT(11)"`
	TimeAdd   time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP <-"`
	Extend    string    `json:"extend" xorm:"VARCHAR(5000)"`
}

//初始化
func NewConnect() *Connect {
	return new(Connect)
}

//初始化列表
func (c *Connect) newMakeDataArr() []Connect {
	return make([]Connect, 0)
}

//列表查询
func (c *Connect) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Connect) GetById(id int) (*Connect, error) {
	m := NewConnect()

	m.ConnectId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Connect) Delete(id int) (int64, error) {
	m := NewConnect()

	m.ConnectId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
