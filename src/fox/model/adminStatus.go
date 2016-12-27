package model

import (
	"time"
)

type AdminStatus struct {
	StatusId   int       `json:"status_id" xorm:"not null pk autoincr INT(11)"`
	Aid        int       `json:"aid" xorm:"not null default 0 INT(11)"`
	LoginTime  time.Time `json:"login_time" xorm:"TIMESTAMP"`
	LoginIp    string    `json:"login_ip" xorm:"not null default '' CHAR(20)"`
	Login      int       `json:"login" xorm:"not null default 0 INT(11)"`
	AidAdd     int       `json:"aid_add" xorm:"not null default 0 INT(11)"`
	AidUpdate  int       `json:"aid_update" xorm:"not null default 0 INT(11)"`
	TimeUpdate time.Time `json:"time_update" xorm:"TIMESTAMP"`
}
