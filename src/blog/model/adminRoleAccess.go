package model

import (
	"fmt"
	"blog/fox"
	"blog/fox/db"
)

type AdminRoleAccess struct {
	Aid    int `json:"aid" xorm:"default 0 unique(aid_role_id) INT(11)"`
	RoleId int `json:"role_id" xorm:"default 0 unique(aid_role_id) INT(11)"`
}

//初始化
func NewAdminRoleAccess() *AdminRoleAccess {
	return new(AdminRoleAccess)
}

//初始化列表
func (c *AdminRoleAccess) newMakeDataArr() []AdminRoleAccess {
	return make([]AdminRoleAccess, 0)
}

//列表查询
func (c *AdminRoleAccess) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	session := db.Filter(q)
	count, err := session.Count(c)
	if err != nil {
		fmt.Println(err)
		return nil,fox.NewError( err.Error())
	}
	Query := db.Pagination(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}

	session = db.Filter(q)
	if orderBy != "" {
		session.OrderBy(orderBy)
	}
	session.Limit(limit, Query.Offset)
	if len(fields) == 0 {
		session.AllCols()
	}
	data := c.newMakeDataArr()
	err = session.Find(&data)
	if err != nil {
		fmt.Println(err)
		return nil,fox.NewError( err.Error())
	}
	Query.Data = make([]interface{}, len(data))
	for y, x := range data {
		Query.Data[y] = x
	}
	return Query, nil
}

// 获取 单条记录
func (c *AdminRoleAccess) GetById(id int) (*AdminRoleAccess, error) {
	m := NewAdminRoleAccess()

	o := db.NewDb()
	_, err := o.Get(m)
	if err == nil {
		return m, nil
	}
	return nil, err
}

// 删除 单条记录
func (c *AdminRoleAccess) Delete(id int) (int64, error) {
	m := NewAdminRoleAccess()

	o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
