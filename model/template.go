package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type Template struct {
	TemplateId int       `form:"template_id" json:"template_id" xorm:"not null pk autoincr INT(11)"`
	Name       string    `form:"name" json:"name" xorm:"not null default '' VARCHAR(80)"`
	Mark       string    `form:"mark" json:"mark" xorm:"not null default '' VARCHAR(80)"`
	Title      string    `form:"title" json:"title" xorm:"not null default '' VARCHAR(255)"`
	Type       int       `form:"type" json:"type" xorm:"not null default 0 TINYINT(1)"`
	Use        int       `form:"use" json:"use" xorm:"not null default 0 TINYINT(2)"`
	Content    string    `form:"content" json:"content" xorm:"TEXT"`
	Remark     string    `form:"remark" json:"remark" xorm:"not null default '' VARCHAR(1024)"`
	TimeAdd    time.Time `form:"time_add" json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeUpdate time.Time `form:"time_update" json:"time_update" xorm:"TIMESTAMP"`
	CodeNum    int       `form:"code_num" json:"code_num" xorm:"not null default 0 TINYINT(3)"`
	Aid        int       `form:"aid" json:"aid" xorm:"not null default 0 INT(11)"`
}

//初始化
func NewTemplate() *Template {
	return new(Template)
}

//初始化列表
func (c *Template) newMakeDataArr() []Template {
	return make([]Template, 0)
}

//列表查询
func (c *Template) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Template) GetById(id int) (*Template, error) {
	m := NewTemplate()

	m.TemplateId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Template) Delete(id int) (int64, error) {
	m := NewTemplate()

	m.TemplateId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
