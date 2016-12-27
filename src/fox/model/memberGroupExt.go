package model

type MemberGroupExt struct {
	GroupId int `json:"group_id" xorm:"not null pk autoincr INT(10)"`
}
