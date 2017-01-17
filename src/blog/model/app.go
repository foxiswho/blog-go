package model

import (
	"blog/fox"
	"blog/fox/db"
	"fmt"
	"time"
)

type App struct {
	AppId   int       `json:"app_id" xorm:"not null pk autoincr INT(11)"`
	TypeId  int       `json:"type_id" xorm:"not null default 0 unique INT(11)"`
	Name    string    `json:"name" xorm:"not null default '' VARCHAR(100)"`
	Mark    string    `json:"mark" xorm:"not null default '' CHAR(32)"`
	Setting string    `json:"setting" xorm:"VARCHAR(5000)"`
	Remark  string    `json:"remark" xorm:"VARCHAR(255)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	IsDel   int       `json:"is_del" xorm:"not null default 0 INT(11)"`
}

//初始化
func NewApp() *App {
	return new(App)
}

//初始化列表
func (c *App) newMakeDataArr() []App {
	return make([]App, 0)
}

//列表查询
func (c *App) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *App) GetById(id int) (*App, error) {
	m := NewApp()

	m.AppId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *App) Delete(id int) (int64, error) {
	m := NewApp()

	m.AppId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
