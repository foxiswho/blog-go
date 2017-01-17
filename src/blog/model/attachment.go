package model

import (
	"fmt"
	"blog/fox"
	"blog/fox/db"
	"time"
)

type Attachment struct {
	AttachmentId int       `json:"attachment_id" xorm:"not null pk autoincr INT(10)"`
	Module       string    `json:"module" xorm:"not null default '' index CHAR(32)"`
	Mark         string    `json:"mark" xorm:"not null default '' index CHAR(60)"`
	TypeId       int       `json:"type_id" xorm:"not null default 0 INT(5)"`
	Name         string    `json:"name" xorm:"not null default '' CHAR(50)"`
	NameOriginal string    `json:"name_original" xorm:"not null default '' VARCHAR(255)"`
	Path         string    `json:"path" xorm:"not null default '' CHAR(200)"`
	Size         int       `json:"size" xorm:"not null default 0 INT(10)"`
	Ext          string    `json:"ext" xorm:"not null default '' CHAR(10)"`
	IsImage      int       `json:"is_image" xorm:"not null default 0 TINYINT(1)"`
	IsThumb      int       `json:"is_thumb" xorm:"not null default 0 TINYINT(1)"`
	Downloads    int       `json:"downloads" xorm:"not null default 0 INT(8)"`
	TimeAdd      time.Time `json:"time_add" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP <-"`
	Ip           string    `json:"ip" xorm:"not null default '' CHAR(15)"`
	Status       int       `json:"status" xorm:"not null default 0 index TINYINT(2)"`
	Md5          string    `json:"md5" xorm:"not null default '' index CHAR(32)"`
	Sha1         string    `json:"sha1" xorm:"not null default '' CHAR(40)"`
	Id           int       `json:"id" xorm:"not null default 0 index INT(10)"`
	Aid          int       `json:"aid" xorm:"not null default 0 index INT(10)"`
	Uid          int       `json:"uid" xorm:"not null default 0 index INT(10)"`
	IsShow       int       `json:"is_show" xorm:"not null default 1 index TINYINT(1)"`
	Http         string    `json:"http" xorm:"not null default '' VARCHAR(100)"`
}

//初始化
func NewAttachment() *Attachment {
	return new(Attachment)
}

//初始化列表
func (c *Attachment) newMakeDataArr() []Attachment {
	return make([]Attachment, 0)
}

//列表查询
func (c *Attachment) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
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
func (c *Attachment) GetById(id int) (*Attachment, error) {
	m := NewAttachment()

	m.AttachmentId = id

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *Attachment) Delete(id int) (int64, error) {
	m := NewAttachment()

	m.AttachmentId = id

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
