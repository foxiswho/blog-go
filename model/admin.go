package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type Admin struct {
	Aid        int       `form:"aid" json:"aid" xorm:"not null pk autoincr INT(11)"`
	Username   string    `form:"username" json:"username" xorm:"not null default '' index CHAR(30)"`
	Password   string    `form:"password" json:"password" xorm:"not null default '' CHAR(32)"`
	Mail       string    `form:"mail" json:"mail" xorm:"not null default '' VARCHAR(80)"`
	Salt       string    `form:"salt" json:"salt" xorm:"not null default '' VARCHAR(10)"`
	TimeAdd    time.Time `form:"time_add" json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeUpdate time.Time `form:"time_update" json:"time_update" xorm:"TIMESTAMP"`
	Ip         string    `form:"ip" json:"ip" xorm:"not null default '' CHAR(15)"`
	JobNo      string    `form:"job_no" json:"job_no" xorm:"not null default '' VARCHAR(15)"`
	NickName   string    `form:"nick_name" json:"nick_name" xorm:"not null default '' VARCHAR(50)"`
	TrueName   string    `form:"true_name" json:"true_name" xorm:"not null default '' VARCHAR(50)"`
	Qq         string    `form:"qq" json:"qq" xorm:"not null default '' VARCHAR(50)"`
	Phone      string    `form:"phone" json:"phone" xorm:"not null default '' VARCHAR(50)"`
	Mobile     string    `form:"mobile" json:"mobile" xorm:"not null default '' VARCHAR(20)"`
	IsDel      int       `form:"is_del" json:"is_del" xorm:"not null default 0 index TINYINT(1)"`
}

//初始化
func NewAdmin() *Admin {
	return new(Admin)
}

//初始化列表
func (c *Admin) newMakeDataArr() []Admin {
	return make([]Admin, 0)
}

//列表查询
func (c *Admin) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Admin) GetById(id int) (*Admin, error) {
	m := NewAdmin()

	m.Aid = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Admin) Delete(id int) (int64, error) {
	m := NewAdmin()

	m.Aid = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
