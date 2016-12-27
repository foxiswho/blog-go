package model

import (
	"time"
)

type Admin struct {
	Aid        int       `json:"aid" xorm:"not null pk autoincr INT(11)"`
	Username   string    `json:"username" xorm:"not null default '' index CHAR(30)"`
	Password   string    `json:"password" xorm:"not null default '' CHAR(32)"`
	Mail       string    `json:"mail" xorm:"not null default '' VARCHAR(80)"`
	Salt       string    `json:"salt" xorm:"not null default '' VARCHAR(10)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeUpdate time.Time `json:"time_update" xorm:"TIMESTAMP"`
	Ip         string    `json:"ip" xorm:"not null default '' CHAR(15)"`
	JobNo      string    `json:"job_no" xorm:"not null default '' VARCHAR(15)"`
	NickName   string    `json:"nick_name" xorm:"not null default '' VARCHAR(50)"`
	TrueName   string    `json:"true_name" xorm:"not null default '' VARCHAR(50)"`
	Qq         string    `json:"qq" xorm:"not null default '' VARCHAR(50)"`
	Phone      string    `json:"phone" xorm:"not null default '' VARCHAR(50)"`
	Mobile     string    `json:"mobile" xorm:"not null default '' VARCHAR(20)"`
	IsDel      int       `json:"is_del" xorm:"not null default 0 index TINYINT(1)"`
}
