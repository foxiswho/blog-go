package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Template struct {
	Id         int       `orm:"column(template_id);auto"`
	Name       string    `orm:"column(name);size(80)"`
	Mark       string    `orm:"column(mark);size(80)"`
	Title      string    `orm:"column(title);size(255)"`
	Type       int8      `orm:"column(type)"`
	Use        uint8     `orm:"column(use)"`
	Content    string    `orm:"column(content);null"`
	Remark     string    `orm:"column(remark);size(1024)"`
	TimeAdd    time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"`
	TimeUpdate time.Time `orm:"column(time_update);type(timestamp);null"`
	CodeNum    uint8     `orm:"column(code_num)"`
	Aid        int       `orm:"column(aid)"`
}

func (t *Template) TableName() string {
	return "template"
}

func init() {
	orm.RegisterModel(new(Template))
}

// AddTemplate insert a new Template into database and returns
// last inserted Id on success.
func AddTemplate(m *Template) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTemplateById retrieves Template by Id. Returns error if
// Id doesn't exist
func GetTemplateById(id int) (v *Template, err error) {
	o := orm.NewOrm()
	v = &Template{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTemplate retrieves all Template matches certain condition. Returns empty list if
// no records exist
func GetAllTemplate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Template))
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

	var l []Template
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

// UpdateTemplate updates Template by Id and returns error if
// the record to be updated doesn't exist
func UpdateTemplateById(m *Template) (err error) {
	o := orm.NewOrm()
	v := Template{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTemplate deletes Template by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTemplate(id int) (err error) {
	o := orm.NewOrm()
	v := Template{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Template{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
