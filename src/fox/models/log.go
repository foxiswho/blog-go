package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id         int       `orm:"column(log_id);pk"`
	Id_RENAME  int       `orm:"column(id)"`
	Aid        int       `orm:"column(aid)"`
	Uid        int       `orm:"column(uid)"`
	TimeAdd    time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"`
	Mark       string    `orm:"column(mark);size(32)"`
	Data       string    `orm:"column(data);null"`
	No         string    `orm:"column(no);size(50)"`
	TypeLogin  int       `orm:"column(type_login)"`
	TypeClient int       `orm:"column(type_client)"`
	Ip         string    `orm:"column(ip);size(20)"`
	Str        string    `orm:"column(str);size(255);null"`
}

func (t *Log) TableName() string {
	return "log"
}

func init() {
	orm.RegisterModel(new(Log))
}

// AddLog insert a new Log into database and returns
// last inserted Id on success.
func AddLog(m *Log) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLogById retrieves Log by Id. Returns error if
// Id doesn't exist
func GetLogById(id int) (v *Log, err error) {
	o := orm.NewOrm()
	v = &Log{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLog retrieves all Log matches certain condition. Returns empty list if
// no records exist
func GetAllLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Log))
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

	var l []Log
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

// UpdateLog updates Log by Id and returns error if
// the record to be updated doesn't exist
func UpdateLogById(m *Log) (err error) {
	o := orm.NewOrm()
	v := Log{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLog deletes Log by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLog(id int) (err error) {
	o := orm.NewOrm()
	v := Log{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Log{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
