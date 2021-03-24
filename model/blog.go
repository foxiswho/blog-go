package model

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox"
	"github.com/foxiswho/blog-go/fox/db"
	"time"
)

type Blog struct {
	BlogId      int       `form:"blog_id" json:"blog_id" xorm:"not null pk autoincr INT(11)"`
	Aid         int       `form:"aid" json:"aid" xorm:"not null default 0 INT(11)"`
	IsDel       int       `form:"is_del" json:"is_del" xorm:"not null default 0 index(is_del) TINYINT(1)"`
	IsOpen      int       `form:"is_open" json:"is_open" xorm:"not null default 1 index(is_del) TINYINT(1)"`
	Status      int       `form:"status" json:"status" xorm:"not null default 0 index(is_del) INT(11)"`
	TimeSystem  time.Time `form:"time_system" json:"time_system" xorm:"TIMESTAMP"`
	TimeUpdate  time.Time `form:"time_update" json:"time_update" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeAdd     time.Time `form:"time_add" json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Title       string    `form:"title" json:"title" xorm:"not null default '' VARCHAR(255)"`
	Author      string    `form:"author" json:"author" xorm:"not null default '' VARCHAR(255)"`
	Url         string    `form:"url" json:"url" xorm:"not null default '' VARCHAR(255)"`
	UrlSource   string    `form:"url_source" json:"url_source" xorm:"not null default '' VARCHAR(255)"`
	UrlRewrite  string    `form:"url_rewrite" json:"url_rewrite" xorm:"not null default '' index CHAR(255)"`
	Description string    `form:"description" json:"description" xorm:"not null default '' VARCHAR(255)"`
	Content     string    `form:"content" json:"content" xorm:"TEXT"`
	Type        int       `form:"type" json:"type" xorm:"not null default 0 index(type) INT(11)"`
	ModuleId    int       `form:"module_id" json:"module_id" xorm:"not null default 0 index(module_id) INT(11)"`
	SourceId    int       `form:"source_id" json:"source_id" xorm:"not null default 0 index(source_id) INT(11)"`
	TypeId      int       `form:"type_id" json:"type_id" xorm:"not null default 0 index(is_del) INT(11)"`
	CatId       int       `form:"cat_id" json:"cat_id" xorm:"not null default 0 index(is_del) INT(11)"`
	Tag         string    `form:"tag" json:"tag" xorm:"not null default '' VARCHAR(255)"`
	Thumb       string    `form:"thumb" json:"thumb" xorm:"not null default '' VARCHAR(255)"`
	IsRelevant  int       `form:"is_relevant" json:"is_relevant" xorm:"not null default 0 TINYINT(1)"`
	IsJump      int       `form:"is_jump" json:"is_jump" xorm:"not null default 0 TINYINT(1)"`
	IsComment   int       `form:"is_comment" json:"is_comment" xorm:"not null default 1 TINYINT(1)"`
	IsRead      int       `form:"is_read" json:"is_read" xorm:"not null default 10014 INT(11)"`
	Sort        int       `form:"sort" json:"sort" xorm:"not null default 0 index(is_del) INT(11)"`
	Remark      string    `form:"remark" json:"remark" xorm:"not null default '' VARCHAR(255)"`
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
		return nil, fox.NewError("数据不存在")
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