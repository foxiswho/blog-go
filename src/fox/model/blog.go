package model

import (
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
	TypeId      int       `json:"type_id" xorm:"not null default 0 index(is_del) INT(11)"`
	CatId       int       `json:"cat_id" xorm:"not null default 0 index(is_del) INT(11)"`
	Tag         string    `json:"tag" xorm:"not null default '' VARCHAR(255)"`
	Thumb       string    `json:"thumb" xorm:"not null default '' VARCHAR(255)"`
	IsRelevant  int       `json:"is_relevant" xorm:"not null default 0 TINYINT(1)"`
	IsJump      int       `json:"is_jump" xorm:"not null default 0 TINYINT(1)"`
	IsComment   int       `json:"is_comment" xorm:"not null default 1 TINYINT(1)"`
	Sort        int       `json:"sort" xorm:"not null default 0 index(is_del) INT(11)"`
	Remark      string    `json:"remark" xorm:"not null default '' VARCHAR(255)"`
}
