package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BlogTag struct {
	Id      int       `orm:"column(tag_id);auto"  json:"tag_id" form:"-"`
	Name    string    `orm:"column(name);size(100)"  json:"name" form:"name"`
	TimeAdd time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"  json:"time_add" form:"time_add"`
	Aid     int       `orm:"column(aid)"  json:"aid" form:"aid"`
	BlogId  int       `orm:"column(blog_id)"  json:"blog_id" form:"blog_id"`
}

func (t *BlogTag) TableName() string {
	return "blog_tag"
}

func init() {
	orm.RegisterModel(new(BlogTag))
}

// AddBlogTag insert a new BlogTag into database and returns
// last inserted Id on success.
func AddBlogTag(m *BlogTag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBlogTagById retrieves BlogTag by Id. Returns error if
// Id doesn't exist
func GetBlogTagById(id int) (v *BlogTag, err error) {
	o := orm.NewOrm()
	v = &BlogTag{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBlogTag retrieves all BlogTag matches certain condition. Returns empty list if
// no records exist
func GetAllBlogTag(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BlogTag))
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

	var l []BlogTag
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

// UpdateBlogTag updates BlogTag by Id and returns error if
// the record to be updated doesn't exist
func UpdateBlogTagById(m *BlogTag) (err error) {
	o := orm.NewOrm()
	v := BlogTag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBlogTag deletes BlogTag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBlogTag(id int) (err error) {
	o := orm.NewOrm()
	v := BlogTag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BlogTag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
