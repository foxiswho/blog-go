package model

type AdminRoleAccess struct {
	Aid    int `json:"aid" xorm:"default 0 unique(aid_role_id) INT(11)"`
	RoleId int `json:"role_id" xorm:"default 0 unique(aid_role_id) INT(11)"`
}
