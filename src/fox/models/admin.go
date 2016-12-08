package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int       `orm:"column(aid);auto"`
	Username   string    `orm:"column(username);size(30)"`
	Password   string    `orm:"column(password);size(32)"`
	Mail       string    `orm:"column(mail);size(80)"`
	Salt       string    `orm:"column(salt);size(10)"`
	TimeAdd    time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"`
	TimeUpdate time.Time `orm:"column(time_update);type(timestamp);null"`
	Ip         string    `orm:"column(ip);size(15)"`
	JobNo      string    `orm:"column(job_no);size(15)"`
	NickName   string    `orm:"column(nick_name);size(50)"`
	TrueName   string    `orm:"column(true_name);size(50)"`
	Qq         string    `orm:"column(qq);size(50)"`
	Phone      string    `orm:"column(phone);size(50)"`
	Mobile     string    `orm:"column(mobile);size(20)"`
	IsDel      uint8     `orm:"column(is_del)"`
}

func (t *Admin) TableName() string {
	return "admin"
}

func init() {
	orm.RegisterModel(new(Admin))
}

// AddAdmin insert a new Admin into database and returns
// last inserted Id on success.
func AddAdmin(m *Admin) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdminById retrieves Admin by Id. Returns error if
// Id doesn't exist
func GetAdminById(id int) (v *Admin, err error) {
	o := orm.NewOrm()
	v = &Admin{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdmin retrieves all Admin matches certain condition. Returns empty list if
// no records exist
func GetAllAdmin(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Admin))
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

	var l []Admin
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

// UpdateAdmin updates Admin by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminById(m *Admin) (err error) {
	o := orm.NewOrm()
	v := Admin{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdmin deletes Admin by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdmin(id int) (err error) {
	o := orm.NewOrm()
	v := Admin{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Admin{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
