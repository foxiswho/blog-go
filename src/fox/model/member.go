package model

import (
	"time"
)

type Member struct {
	Uid      int       `json:"uid" xorm:"not null pk autoincr INT(11)"`
	Mobile   string    `json:"mobile" xorm:"not null default '' index CHAR(11)"`
	Username string    `json:"username" xorm:"not null default '' index CHAR(30)"`
	Mail     string    `json:"mail" xorm:"not null default '' index CHAR(32)"`
	Password string    `json:"password" xorm:"not null default '' CHAR(32)"`
	Salt     string    `json:"salt" xorm:"not null default '' CHAR(6)"`
	RegIp    string    `json:"reg_ip" xorm:"not null default '' CHAR(15)"`
	RegTime  time.Time `json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	IsDel    int       `json:"is_del" xorm:"not null default 0 index TINYINT(1)"`
	GroupId  int       `json:"group_id" xorm:"not null default 410 index INT(11)"`
	TrueName string    `json:"true_name" xorm:"not null default '' VARCHAR(32)"`
	Name     string    `json:"name" xorm:"not null default '' VARCHAR(100)"`
}
