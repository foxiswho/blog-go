package model

type AdminRolePriv struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	RoleId int    `json:"role_id" xorm:"not null default 0 index(role_id_2) index SMALLINT(3)"`
	S      string `json:"s" xorm:"not null default '' index(role_id_2) CHAR(100)"`
	Data   string `json:"data" xorm:"not null default '' CHAR(50)"`
	Aid    int    `json:"aid" xorm:"not null default 0 INT(10)"`
	MenuId int    `json:"menu_id" xorm:"not null default 0 INT(10)"`
	Type   string `json:"type" xorm:"not null default 'url' CHAR(32)"`
}
