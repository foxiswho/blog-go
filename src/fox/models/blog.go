package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Blog struct {
	Id          int       `orm:"column(blog_id);auto"  json:"blog_id" form:"-"`
	Aid         int       `orm:"column(aid)"  json:"aid" form:"-"`
	IsDel       int8      `orm:"column(is_del)"  json:"is_del" form:"is_del"`
	IsOpen      int8      `orm:"column(is_open)"  json:"is_open" form:"is_open"`
	Status      int       `orm:"column(status)"  json:"status" form:"status"`
	TimeSystem  time.Time `orm:"column(time_system);type(timestamp);null"  json:"time_system" form:"-"`
	TimeUpdate  time.Time `orm:"column(time_update);type(timestamp);null;auto_now"  json:"time_update" form:"-"`
	TimeAdd     time.Time `orm:"column(time_add);type(timestamp);null;auto_now_add"  json:"time_add" form:"-"`
	Title       string    `orm:"column(title);size(255)"  json:"title" form:"title"`
	Author      string    `orm:"column(author);size(255)"  json:"author" form:"author"`
	Url         string    `orm:"column(url);size(255)"  json:"url" form:"url"`
	UrlSource   string    `orm:"column(url_source);size(255)"  json:"url_source" form:"url_source"`
	UrlRewrite  string    `orm:"column(url_rewrite);size(255)"  json:"url_rewrite" form:"url_rewrite"`
	Description string    `orm:"column(description);size(255)"  json:"description" form:"description"`
	Content     string    `orm:"column(content);null"  json:"content" form:"content"`
	TypeId      int       `orm:"column(type_id)"  json:"type_id" form:"type_id"`
	CatId       int       `orm:"column(cat_id)"  json:"cat_id" form:"cat_id"`
	Tag         string    `orm:"column(tag);size(255)"  json:"tag" form:"tag"`
	Thumb       string    `orm:"column(thumb);size(255)"  json:"thumb" form:"thumb"`
	IsRelevant  int8      `orm:"column(is_relevant)"  json:"is_relevant" form:"is_relevant"`
	IsJump      int8      `orm:"column(is_jump)"  json:"is_jump" form:"is_jump"`
	IsComment   int8      `orm:"column(is_comment)"  json:"is_comment" form:"is_comment"`
	Sort        int       `orm:"column(sort)"  json:"sort" form:"sort"`
	Remark      string    `orm:"column(remark);size(255)"  json:"remark" form:"remark"`
}

func (t *Blog) TableName() string {
	return "blog"
}

func init() {
	orm.RegisterModel(new(Blog))
}

// AddBlog insert a new Blog into database and returns
// last inserted Id on success.
func AddBlog(m *Blog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBlogById retrieves Blog by Id. Returns error if
// Id doesn't exist
func GetBlogById(id int) (v *Blog, err error) {
	o := orm.NewOrm()
	v = &Blog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBlog retrieves all Blog matches certain condition. Returns empty list if
// no records exist
func GetAllBlog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Blog))
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

	var l []Blog
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

// UpdateBlog updates Blog by Id and returns error if
// the record to be updated doesn't exist
func UpdateBlogById(m *Blog) (err error) {
	o := orm.NewOrm()
	v := Blog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBlog deletes Blog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBlog(id int) (err error) {
	o := orm.NewOrm()
	v := Blog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Blog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
