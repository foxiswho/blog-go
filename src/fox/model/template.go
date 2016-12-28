package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
	"time"
)

type Template struct {
	TemplateId int       `json:"template_id" xorm:"not null pk autoincr INT(11)" orm:"column(template_id)"`
	Name       string    `json:"name" xorm:"not null default '' VARCHAR(80)" orm:"column(name)"`
	Mark       string    `json:"mark" xorm:"not null default '' VARCHAR(80)" orm:"column(mark)"`
	Title      string    `json:"title" xorm:"not null default '' VARCHAR(255)" orm:"column(title)"`
	Type       int       `json:"type" xorm:"not null default 0 TINYINT(1)" orm:"column(type)"`
	Use        int       `json:"use" xorm:"not null default 0 TINYINT(2)" orm:"column(use)"`
	Content    string    `json:"content" xorm:"TEXT" orm:"column(content)"`
	Remark     string    `json:"remark" xorm:"not null default '' VARCHAR(1024)" orm:"column(remark)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" orm:"column(time_add)"`
	TimeUpdate time.Time `json:"time_update" xorm:"TIMESTAMP" orm:"column(time_update)"`
	CodeNum    int       `json:"code_num" xorm:"not null default 0 TINYINT(3)" orm:"column(code_num)"`
	Aid        int       `json:"aid" xorm:"not null default 0 INT(11)" orm:"column(aid)"`
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
func (c *Template) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Page, error) {
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
