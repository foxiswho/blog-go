package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MemberStatus struct {
	Id             int       `orm:"column(status_id);auto"`
	Uid            int       `orm:"column(uid)"`
	RegIp          string    `orm:"column(reg_ip);size(15)"`
	RegTime        time.Time `orm:"column(reg_time);type(timestamp);auto_now_add"`
	RegType        int       `orm:"column(reg_type)"`
	RegAppId       int       `orm:"column(reg_app_id)"`
	LastLoginIp    string    `orm:"column(last_login_ip);size(15)"`
	LastLoginTime  time.Time `orm:"column(last_login_time);type(timestamp);null"`
	LastLoginAppId int       `orm:"column(last_login_app_id)"`
	Login          int16     `orm:"column(login)"`
	IsMobile       uint8     `orm:"column(is_mobile)"`
	IsEmail        uint8     `orm:"column(is_email)"`
	AidAid         uint      `orm:"column(aid_aid)"`
}

func (t *MemberStatus) TableName() string {
	return "member_status"
}

func init() {
	orm.RegisterModel(new(MemberStatus))
}

// AddMemberStatus insert a new MemberStatus into database and returns
// last inserted Id on success.
func AddMemberStatus(m *MemberStatus) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberStatusById retrieves MemberStatus by Id. Returns error if
// Id doesn't exist
func GetMemberStatusById(id int) (v *MemberStatus, err error) {
	o := orm.NewOrm()
	v = &MemberStatus{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberStatus retrieves all MemberStatus matches certain condition. Returns empty list if
// no records exist
func GetAllMemberStatus(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberStatus))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MemberStatus
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMemberStatus updates MemberStatus by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberStatusById(m *MemberStatus) (err error) {
	o := orm.NewOrm()
	v := MemberStatus{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberStatus deletes MemberStatus by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberStatus(id int) (err error) {
	o := orm.NewOrm()
	v := MemberStatus{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberStatus{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
