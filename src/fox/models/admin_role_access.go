package models

type AdminRoleAccess struct {
	Aid    int `orm:"column(aid);null"`
	RoleId int `orm:"column(role_id);null"`
}
