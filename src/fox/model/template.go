package model

import (
	"time"
)

type Template struct {
	TemplateId int       `json:"template_id" xorm:"not null pk autoincr INT(11)"`
	Name       string    `json:"name" xorm:"not null default '' VARCHAR(80)"`
	Mark       string    `json:"mark" xorm:"not null default '' VARCHAR(80)"`
	Title      string    `json:"title" xorm:"not null default '' VARCHAR(255)"`
	Type       int       `json:"type" xorm:"not null default 0 TINYINT(1)"`
	Use        int       `json:"use" xorm:"not null default 0 TINYINT(2)"`
	Content    string    `json:"content" xorm:"TEXT"`
	Remark     string    `json:"remark" xorm:"not null default '' VARCHAR(1024)"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	TimeUpdate time.Time `json:"time_update" xorm:"TIMESTAMP"`
	CodeNum    int       `json:"code_num" xorm:"not null default 0 TINYINT(3)"`
	Aid        int       `json:"aid" xorm:"not null default 0 INT(11)"`
}
