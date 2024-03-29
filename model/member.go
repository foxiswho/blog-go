package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type Member struct {
	Uid      int       `form:"uid" json:"uid" xorm:"not null pk autoincr INT(11)"`
	Mobile   string    `form:"mobile" json:"mobile" xorm:"not null default '' index CHAR(11)"`
	Username string    `form:"username" json:"username" xorm:"not null default '' index CHAR(30)"`
	Mail     string    `form:"mail" json:"mail" xorm:"not null default '' index CHAR(32)"`
	Password string    `form:"password" json:"password" xorm:"not null default '' CHAR(32)"`
	Salt     string    `form:"salt" json:"salt" xorm:"not null default '' CHAR(6)"`
	RegIp    string    `form:"reg_ip" json:"reg_ip" xorm:"not null default '' CHAR(15)"`
	RegTime  time.Time `form:"reg_time" json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	IsDel    int       `form:"is_del" json:"is_del" xorm:"not null default 0 index TINYINT(1)"`
	GroupId  int       `form:"group_id" json:"group_id" xorm:"not null default 410 index INT(11)"`
	TrueName string    `form:"true_name" json:"true_name" xorm:"not null default '' VARCHAR(32)"`
	Name     string    `form:"name" json:"name" xorm:"not null default '' VARCHAR(100)"`
}

//初始化
func NewMember() *Member {
	return new(Member)
}

//初始化列表
func (c *Member) newMakeDataArr() []Member {
	return make([]Member, 0)
}

//列表查询
func (c *Member) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Member) GetById(id int) (*Member, error) {
	m := NewMember()

	m.Uid = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Member) Delete(id int) (int64, error) {
	m := NewMember()

	m.Uid = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
