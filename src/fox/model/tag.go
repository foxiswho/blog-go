package model

import (
	"time"
)

type Tag struct {
	TagId   int       `json:"tag_id" xorm:"not null pk autoincr INT(11)"`
	Name    string    `json:"name" xorm:"not null default '' CHAR(50)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
