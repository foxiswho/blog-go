package model

import (
	"time"
)

type Log struct {
	LogId      int       `json:"log_id" xorm:"not null pk INT(11)"`
	Id         int       `json:"id" xorm:"not null default 0 index INT(11)"`
	Aid        int       `json:"aid" xorm:"not null default 0 index INT(11)"`
	Uid        int       `json:"uid" xorm:"not null default 0 index INT(11)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Mark       string    `json:"mark" xorm:"not null default '' CHAR(32)"`
	Data       string    `json:"data" xorm:"TEXT"`
	No         string    `json:"no" xorm:"not null default '' index CHAR(50)"`
	TypeLogin  int       `json:"type_login" xorm:"not null default 0 index INT(11)"`
	TypeClient int       `json:"type_client" xorm:"not null default 0 index INT(11)"`
	Ip         string    `json:"ip" xorm:"not null default '' CHAR(20)"`
	Msg        string    `json:"msg" xorm:"VARCHAR(255)"`
}
