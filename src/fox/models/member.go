package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Member struct {
	Id       int       `orm:"column(uid);auto"`
	Mobile   string    `orm:"column(mobile);size(11)"`
	Username string    `orm:"column(username);size(30)"`
	Email    string    `orm:"column(email);size(32)"`
	Password string    `orm:"column(password);size(32)"`
	Salt     string    `orm:"column(salt);size(6)"`
	RegIp    string    `orm:"column(reg_ip);size(15)"`
	RegTime  time.Time `orm:"column(reg_time);type(timestamp);auto_now_add"`
	IsDel    uint8     `orm:"column(is_del)"`
	GroupId  uint8     `orm:"column(group_id)"`
	TrueName string    `orm:"column(true_name);size(32)"`
}

func (t *Member) TableName() string {
	return "member"
}

func init() {
	orm.RegisterModel(new(Member))
}

// AddMember insert a new Member into database and returns
// last inserted Id on success.
func AddMember(m *Member) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMemberById retrieves Member by Id. Returns error if
// Id doesn't exist
func GetMemberById(id int) (v *Member, err error) {
	o := orm.NewOrm()
	v = &Member{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMember retrieves all Member matches certain condition. Returns empty list if
// no records exist
func GetAllMember(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Member))
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

	var l []Member
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

// UpdateMember updates Member by Id and returns error if
// the record to be updated doesn't exist
func UpdateMemberById(m *Member) (err error) {
	o := orm.NewOrm()
	v := Member{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMember deletes Member by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMember(id int) (err error) {
	o := orm.NewOrm()
	v := Member{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Member{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
