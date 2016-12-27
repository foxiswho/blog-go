package model

import (
	"time"
)

type MemberStatus struct {
	StatusId       int       `json:"status_id" xorm:"not null pk autoincr INT(11)"`
	Uid            int       `json:"uid" xorm:"not null default 0 index INT(11)"`
	RegIp          string    `json:"reg_ip" xorm:"not null default '' CHAR(15)"`
	RegTime        time.Time `json:"reg_time" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	RegType        int       `json:"reg_type" xorm:"not null default 0 INT(11)"`
	RegAppId       int       `json:"reg_app_id" xorm:"not null default 1 INT(11)"`
	LastLoginIp    string    `json:"last_login_ip" xorm:"not null default '' CHAR(15)"`
	LastLoginTime  time.Time `json:"last_login_time" xorm:"TIMESTAMP"`
	LastLoginAppId int       `json:"last_login_app_id" xorm:"not null default 0 INT(11)"`
	Login          int       `json:"login" xorm:"not null default 0 SMALLINT(5)"`
	IsMobile       int       `json:"is_mobile" xorm:"not null default 0 TINYINT(1)"`
	IsEmail        int       `json:"is_email" xorm:"not null default 0 TINYINT(1)"`
	AidAid         int       `json:"aid_aid" xorm:"not null default 0 INT(11)"`
}
