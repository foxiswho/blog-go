package model

type AdminMenu struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name     string `json:"name" xorm:"not null default '' CHAR(100)"`
	ParentId int    `json:"parent_id" xorm:"not null default 0 index INT(11)"`
	S        string `json:"s" xorm:"not null default '' index CHAR(60)"`
	Data     string `json:"data" xorm:"not null default '' CHAR(100)"`
	Sort     int    `json:"sort" xorm:"not null default 0 index INT(11)"`
	Remark   string `json:"remark" xorm:"not null default '' VARCHAR(255)"`
	Type     string `json:"type" xorm:"not null default 'url' CHAR(32)"`
	Level    int    `json:"level" xorm:"not null default 0 TINYINT(3)"`
	Level1Id int    `json:"level1_id" xorm:"not null default 0 INT(11)"`
	Md5      string `json:"md5" xorm:"not null default '' CHAR(32)"`
	IsShow   int    `json:"is_show" xorm:"not null default 1 TINYINT(1)"`
	IsUnique int    `json:"is_unique" xorm:"not null default 0 TINYINT(1)"`
}
