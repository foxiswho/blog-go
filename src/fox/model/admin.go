package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
	"time"
)

type Admin struct {
	Aid        int       `json:"aid" xorm:"not null pk autoincr INT(11)" orm:"column(aid)"`
	Username   string    `json:"username" xorm:"not null default '' index CHAR(30)" orm:"column(username)"`
	Password   string    `json:"password" xorm:"not null default '' CHAR(32)" orm:"column(password)"`
	Mail       string    `json:"mail" xorm:"not null default '' VARCHAR(80)" orm:"column(mail)"`
	Salt       string    `json:"salt" xorm:"not null default '' VARCHAR(10)" orm:"column(salt)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" orm:"column(time_add)"`
	TimeUpdate time.Time `json:"time_update" xorm:"TIMESTAMP" orm:"column(time_update)"`
	Ip         string    `json:"ip" xorm:"not null default '' CHAR(15)" orm:"column(ip)"`
	JobNo      string    `json:"job_no" xorm:"not null default '' VARCHAR(15)" orm:"column(job_no)"`
	NickName   string    `json:"nick_name" xorm:"not null default '' VARCHAR(50)" orm:"column(nick_name)"`
	TrueName   string    `json:"true_name" xorm:"not null default '' VARCHAR(50)" orm:"column(true_name)"`
	Qq         string    `json:"qq" xorm:"not null default '' VARCHAR(50)" orm:"column(qq)"`
	Phone      string    `json:"phone" xorm:"not null default '' VARCHAR(50)" orm:"column(phone)"`
	Mobile     string    `json:"mobile" xorm:"not null default '' VARCHAR(20)" orm:"column(mobile)"`
	IsDel      int       `json:"is_del" xorm:"not null default 0 index TINYINT(1)" orm:"column(is_del)"`
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
func (c *Admin) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Page, error) {
	session := db.Filter(q)
	count, err := session.Count(c)
	if err != nil {
		fmt.Println(err)
		return nil, &util.Error{Msg: err.Error()}
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
		return nil, &util.Error{Msg: err.Error()}
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
