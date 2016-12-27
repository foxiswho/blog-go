package model

type MemberProfile struct {
	ProfileId   int    `json:"profile_id" xorm:"not null pk autoincr INT(11)"`
	Uid         int    `json:"uid" xorm:"not null default 0 index INT(11)"`
	Sex         int    `json:"sex" xorm:"not null default 1 TINYINT(1)"`
	Job         string `json:"job" xorm:"not null default '' VARCHAR(50)"`
	Qq          string `json:"qq" xorm:"not null default '' VARCHAR(20)"`
	Phone       string `json:"phone" xorm:"not null default '' VARCHAR(20)"`
	County      int    `json:"county" xorm:"not null default 1 INT(11)"`
	Province    int    `json:"province" xorm:"not null default 0 INT(11)"`
	City        int    `json:"city" xorm:"not null default 0 INT(11)"`
	District    int    `json:"district" xorm:"not null default 0 INT(11)"`
	Address     string `json:"address" xorm:"not null default '' VARCHAR(255)"`
	Wechat      string `json:"wechat" xorm:"not null default '' VARCHAR(20)"`
	RemarkAdmin string `json:"remark_admin" xorm:"TEXT"`
}
