package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type AdminStatus struct {
	StatusId   int       `form:"status_id" json:"status_id" xorm:"not null pk autoincr INT(11)"`
	Aid        int       `form:"aid" json:"aid" xorm:"not null default 0 INT(11)"`
	LoginTime  time.Time `form:"login_time" json:"login_time" xorm:"TIMESTAMP <-" `
	LoginIp    string    `form:"login_ip" json:"login_ip" xorm:"not null default '' CHAR(20)"`
	Login      int       `form:"login" json:"login" xorm:"not null default 0 INT(11)"`
	AidAdd     int       `form:"aid_add" json:"aid_add" xorm:"not null default 0 INT(11)"`
	AidUpdate  int       `form:"aid_update" json:"aid_update" xorm:"not null default 0 INT(11)"`
	TimeUpdate time.Time `form:"time_update" json:"time_update" xorm:"TIMESTAMP <-"`
	Remark     string    `form:"remark" json:"remark" xorm:"not null default '' VARCHAR(255)"`
}

//初始化
func NewAdminStatus() *AdminStatus {
	return new(AdminStatus)
}

//初始化列表
func (c *AdminStatus) newMakeDataArr() []AdminStatus {
	return make([]AdminStatus, 0)
}

//列表查询
func (c *AdminStatus) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *AdminStatus) GetById(id int) (*AdminStatus, error) {
	m := NewAdminStatus()

	m.StatusId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *AdminStatus) Delete(id int) (int64, error) {
	m := NewAdminStatus()

	m.StatusId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
