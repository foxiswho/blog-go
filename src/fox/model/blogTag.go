package model

import (
	"time"
)

type BlogTag struct {
	TagId   int       `json:"tag_id" xorm:"not null pk autoincr INT(11)"`
	Name    string    `json:"name" xorm:"not null default '' CHAR(100)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Aid     int       `json:"aid" xorm:"not null default 0 INT(11)"`
	BlogId  int       `json:"blog_id" xorm:"not null default 0 INT(11)"`
}
