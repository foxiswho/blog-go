package model

import (
	"fmt"
	"fox/util"
	"fox/util/db"
	"time"
)

type Blog struct {
	BlogId      int       `json:"blog_id" xorm:"not null pk autoincr INT(11)" orm:"column(blog_id)"`
	Aid         int       `json:"aid" xorm:"not null default 0 INT(11)" orm:"column(aid)"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 index(is_del) TINYINT(1)" orm:"column(is_del)"`
	IsOpen      int       `json:"is_open" xorm:"not null default 1 index(is_del) TINYINT(1)" orm:"column(is_open)"`
	Status      int       `json:"status" xorm:"not null default 0 index(is_del) INT(11)" orm:"column(status)"`
	TimeSystem  time.Time `json:"time_system" xorm:"TIMESTAMP" orm:"column(time_system)"`
	TimeUpdate  time.Time `json:"time_update" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" orm:"column(time_update)"`
	TimeAdd     time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP" orm:"column(time_add)"`
	Title       string    `json:"title" xorm:"not null default '' VARCHAR(255)" orm:"column(title)"`
	Author      string    `json:"author" xorm:"not null default '' VARCHAR(255)" orm:"column(author)"`
	Url         string    `json:"url" xorm:"not null default '' VARCHAR(255)" orm:"column(url)"`
	UrlSource   string    `json:"url_source" xorm:"not null default '' VARCHAR(255)" orm:"column(url_source)"`
	UrlRewrite  string    `json:"url_rewrite" xorm:"not null default '' index CHAR(255)" orm:"column(url_rewrite)"`
	Description string    `json:"description" xorm:"not null default '' VARCHAR(255)" orm:"column(description)"`
	Content     string    `json:"content" xorm:"TEXT" orm:"column(content)"`
	Type        int       `json:"type" xorm:"not null default 0 INT(11)" orm:"column(type)"`
	TypeId      int       `json:"type_id" xorm:"not null default 0 index(is_del) INT(11)" orm:"column(type_id)"`
	CatId       int       `json:"cat_id" xorm:"not null default 0 index(is_del) INT(11)" orm:"column(cat_id)"`
	Tag         string    `json:"tag" xorm:"not null default '' VARCHAR(255)" orm:"column(tag)"`
	Thumb       string    `json:"thumb" xorm:"not null default '' VARCHAR(255)" orm:"column(thumb)"`
	IsRelevant  int       `json:"is_relevant" xorm:"not null default 0 TINYINT(1)" orm:"column(is_relevant)"`
	IsJump      int       `json:"is_jump" xorm:"not null default 0 TINYINT(1)" orm:"column(is_jump)"`
	IsComment   int       `json:"is_comment" xorm:"not null default 1 TINYINT(1)" orm:"column(is_comment)"`
	Sort        int       `json:"sort" xorm:"not null default 0 index(is_del) INT(11)" orm:"column(sort)"`
	Remark      string    `json:"remark" xorm:"not null default '' VARCHAR(255)" orm:"column(remark)"`
}

//初始化
func NewBlog() *Blog {
	return new(Blog)
}

//初始化列表
func (c *Blog) newMakeDataArr() []Blog {
	return make([]Blog, 0)
}

//列表查询
func (c *Blog) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Blog) GetById(id int) (*Blog, error) {
	m := NewBlog()

	m.BlogId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Blog) Delete(id int) (int64, error) {
	m := NewBlog()

	m.BlogId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
