package model

import (
	"time"
)

type Type struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"not null default '' CHAR(100)"`
	Code      string    `json:"code" xorm:"not null default '' CHAR(32)"`
	Mark      string    `json:"mark" xorm:"not null default '' index CHAR(32)"`
	TypeId    int       `json:"type_id" xorm:"not null default 0 index INT(11)"`
	ParentId  int       `json:"parent_id" xorm:"not null default 0 index INT(11)"`
	Value     int       `json:"value" xorm:"not null default 0 INT(10)"`
	IsDel     int       `json:"is_del" xorm:"not null default 0 index INT(11)"`
	Sort      int       `json:"sort" xorm:"not null default 0 index INT(11)"`
	Remark    string    `json:"remark" xorm:"VARCHAR(255)"`
	TimeAdd   time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Aid       int       `json:"aid" xorm:"not null default 0 INT(11)"`
	Module    string    `json:"module" xorm:"not null default '' CHAR(50)"`
	IsDefault int       `json:"is_default" xorm:"not null default 0 TINYINT(1)"`
	Setting   string    `json:"setting" xorm:"VARCHAR(255)"`
	IsChild   int       `json:"is_child" xorm:"not null default 0 TINYINT(1)"`
	IsSystem  int       `json:"is_system" xorm:"not null default 0 TINYINT(1)"`
	IsShow    int       `json:"is_show" xorm:"not null default 0 TINYINT(1)"`
}
