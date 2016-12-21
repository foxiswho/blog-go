package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type BlogStatistics struct {
	Id             int    `orm:"column(statistics_id);auto"`
	BlogId         int    `orm:"column(blog_id)"`
	Comment        int    `orm:"column(comment)"`
	Read           int    `orm:"column(read)"`
	SeoTitle       string `orm:"column(seo_title);size(255)"`
	SeoDescription string `orm:"column(seo_description);size(255)"`
	SeoKeyword     string `orm:"column(seo_keyword);size(255)"`
}

func (t *BlogStatistics) TableName() string {
	return "blog_statistics"
}

func init() {
	orm.RegisterModel(new(BlogStatistics))
}

// AddBlogStatistics insert a new BlogStatistics into database and returns
// last inserted Id on success.
func AddBlogStatistics(m *BlogStatistics) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBlogStatisticsById retrieves BlogStatistics by Id. Returns error if
// Id doesn't exist
func GetBlogStatisticsById(id int) (v *BlogStatistics, err error) {
	o := orm.NewOrm()
	v = &BlogStatistics{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBlogStatistics retrieves all BlogStatistics matches certain condition. Returns empty list if
// no records exist
func GetAllBlogStatistics(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BlogStatistics))
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

	var l []BlogStatistics
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

// UpdateBlogStatistics updates BlogStatistics by Id and returns error if
// the record to be updated doesn't exist
func UpdateBlogStatisticsById(m *BlogStatistics) (err error) {
	o := orm.NewOrm()
	v := BlogStatistics{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBlogStatistics deletes BlogStatistics by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBlogStatistics(id int) (err error) {
	o := orm.NewOrm()
	v := BlogStatistics{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BlogStatistics{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
