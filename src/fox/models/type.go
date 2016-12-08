package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Type struct {
	Id        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(100)"`
	Code      string    `orm:"column(code);size(32)"`
	Mark      string    `orm:"column(mark);size(32)"`
	TypeId    int       `orm:"column(type_id)"`
	ParentId  int       `orm:"column(parent_id)"`
	Value     int       `orm:"column(value)"`
	IsDel     int       `orm:"column(is_del)"`
	Sort      int       `orm:"column(sort)"`
	Remark    string    `orm:"column(remark);size(255);null"`
	TimeAdd   time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"`
	Aid       uint      `orm:"column(aid)"`
	Module    string    `orm:"column(module);size(50)"`
	IsDefault int8      `orm:"column(is_default)"`
	Setting   string    `orm:"column(setting);size(255);null"`
	IsChild   int8      `orm:"column(is_child)"`
	IsSystem  int8      `orm:"column(is_system)"`
	IsShow    int8      `orm:"column(is_show)"`
}

func (t *Type) TableName() string {
	return "type"
}

func init() {
	orm.RegisterModel(new(Type))
}

// AddType insert a new Type into database and returns
// last inserted Id on success.
func AddType(m *Type) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTypeById retrieves Type by Id. Returns error if
// Id doesn't exist
func GetTypeById(id int) (v *Type, err error) {
	o := orm.NewOrm()
	v = &Type{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllType retrieves all Type matches certain condition. Returns empty list if
// no records exist
func GetAllType(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Type))
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

	var l []Type
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

// UpdateType updates Type by Id and returns error if
// the record to be updated doesn't exist
func UpdateTypeById(m *Type) (err error) {
	o := orm.NewOrm()
	v := Type{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteType deletes Type by Id and returns error if
// the record to be deleted doesn't exist
func DeleteType(id int) (err error) {
	o := orm.NewOrm()
	v := Type{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Type{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
