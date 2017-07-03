package model

import (
	"fmt"
	"blog/fox"
	"blog/fox/db"
	"time"
)

type Blog struct {
	BlogId      int       `json:"blog_id" xorm:"not null pk autoincr INT(11)"`
	Aid         int       `json:"aid" xorm:"not null default 0 INT(11)"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 index(is_del) TINYINT(1)"`
	IsOpen      int       `json:"is_open" xorm:"not null default 1 index(is_del) TINYINT(1)"`
	Status      int       `json:"status" xorm:"not null default 0 index(is_del) INT(11)"`
	TimeSystem  time.Time `json:"time_system" xorm:"TIMESTAMP"`
	TimeUpdate  time.Time `json:"time_update" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeAdd     time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Title       string    `json:"title" xorm:"not null default '' VARCHAR(255)"`
	Author      string    `json:"author" xorm:"not null default '' VARCHAR(255)"`
	Url         string    `json:"url" xorm:"not null default '' VARCHAR(255)"`
	UrlSource   string    `json:"url_source" xorm:"not null default '' VARCHAR(255)"`
	UrlRewrite  string    `json:"url_rewrite" xorm:"not null default '' index CHAR(255)"`
	Description string    `json:"description" xorm:"not null default '' VARCHAR(255)"`
	Content     string    `json:"content" xorm:"TEXT"`
	Type        int       `json:"type" xorm:"not null default 0 index(type) INT(11)"`
	ModuleId    int       `json:"module_id" xorm:"not null default 0 index(module_id) INT(11)"`
	SourceId    int       `json:"source_id" xorm:"not null default 0 index(source_id) INT(11)"`
	TypeId      int       `json:"type_id" xorm:"not null default 0 index(is_del) INT(11)"`
	CatId       int       `json:"cat_id" xorm:"not null default 0 index(is_del) INT(11)"`
	Tag         string    `json:"tag" xorm:"not null default '' VARCHAR(255)"`
	Thumb       string    `json:"thumb" xorm:"not null default '' VARCHAR(255)"`
	IsRelevant  int       `json:"is_relevant" xorm:"not null default 0 TINYINT(1)"`
	IsJump      int       `json:"is_jump" xorm:"not null default 0 TINYINT(1)"`
	IsComment   int       `json:"is_comment" xorm:"not null default 1 TINYINT(1)"`
	IsRead      int       `json:"is_read" xorm:"not null default 10014 INT(11)"`
	Sort        int       `json:"sort" xorm:"not null default 0 index(is_del) INT(11)"`
	Remark      string    `json:"remark" xorm:"not null default '' VARCHAR(255)"`
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
		return nil, fox.NewError(err.Error())
	}
	//fmt.Println("count",int(count))
	//fmt.Println("page",page)
	//fmt.Println("limit",limit)
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
		return nil, fox.NewError(err.Error())
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
	ok, err := o.Get(m)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil,fox.NewError("数据不存在")
	}
	return m, nil
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
