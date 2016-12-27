package model

import (
	"time"
)

type Session struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uid        int       `json:"uid" xorm:"not null default 0 index(uid) INT(11)"`
	Ip         string    `json:"ip" xorm:"not null default '' CHAR(15)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ErrorCount int       `json:"error_count" xorm:"not null default 0 TINYINT(1)"`
	AppId      int       `json:"app_id" xorm:"not null default 0 INT(11)"`
	TypeLogin  int       `json:"type_login" xorm:"not null default 0 index(uid) INT(11)"`
	Md5        string    `json:"md5" xorm:"not null default '' CHAR(32)"`
	TypeClient int       `json:"type_client" xorm:"not null default 0 index(uid) INT(11)"`
}
