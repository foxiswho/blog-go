package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MemberProfile struct {
	Id          int    `orm:"column(profile_id);auto"`
	Uid         int    `orm:"column(uid)"`
	Sex         uint8  `orm:"column(sex)"`
	Job         string `orm:"column(job);size(50)"`
	Qq          string `orm:"column(qq);size(20)"`
	Phone       string `orm:"column(phone);size(20)"`
	County      uint   `orm:"column(county)"`
	Province    uint   `orm:"column(province)"`
	City        uint   `orm:"column(city)"`
	District    uint   `orm:"column(district)"`
	Address     string `orm:"column(address);size(255)"`
	Wechat      string `orm:"column(wechat);size(20)"`
	RemarkAdmin string `orm:"column(remark_admin);null"`
}

func (t *MemberProfile) TableName() string {
	return "member_profile"
}

func init() {
	orm.RegisterModel(new(MemberProfile))
}

// AddMemberProfile insert a new MemberProfile into database and returns
// last inserted Id on success.
func AddMemberProfile(m *MemberProfile) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberProfileById retrieves MemberProfile by Id. Returns error if
// Id doesn't exist
func GetMemberProfileById(id int) (v *MemberProfile, err error) {
	o := orm.NewOrm()
	v = &MemberProfile{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMemberProfile retrieves all MemberProfile matches certain condition. Returns empty list if
// no records exist
func GetAllMemberProfile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MemberProfile))
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

	var l []MemberProfile
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

// UpdateMemberProfile updates MemberProfile by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberProfileById(m *MemberProfile) (err error) {
	o := orm.NewOrm()
	v := MemberProfile{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMemberProfile deletes MemberProfile by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMemberProfile(id int) (err error) {
	o := orm.NewOrm()
	v := MemberProfile{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MemberProfile{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
