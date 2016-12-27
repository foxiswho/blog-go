package model

import (
	"time"
)

type Connect struct {
	ConnectId int       `json:"connect_id" xorm:"not null pk autoincr INT(11)"`
	Uid       int       `json:"uid" xorm:"not null default 0 index INT(11)"`
	OpenId    string    `json:"open_id" xorm:"not null default '' index VARCHAR(80)"`
	Token     string    `json:"token" xorm:"not null default '' VARCHAR(80)"`
	Type      int       `json:"type" xorm:"not null default 1 INT(11)"`
	TypeLogin int       `json:"type_login" xorm:"not null default 0 INT(11)"`
	TimeAdd   time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
