package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Attachment struct {
	Id           int       `orm:"column(attachment_id);auto"`
	Module       string    `orm:"column(module);size(32)"`
	Mark         string    `orm:"column(mark);size(60)"`
	TypeId       uint      `orm:"column(type_id)"`
	Name         string    `orm:"column(name);size(50)"`
	NameOriginal string    `orm:"column(name_original);size(255)"`
	Path         string    `orm:"column(path);size(200)"`
	Size         uint      `orm:"column(size)"`
	Ext          string    `orm:"column(ext);size(10)"`
	IsImage      uint8     `orm:"column(is_image)"`
	IsThumb      uint8     `orm:"column(is_thumb)"`
	Downloads    uint      `orm:"column(downloads)"`
	TimeAdd      time.Time `orm:"column(time_add);type(timestamp);auto_now_add"`
	Ip           string    `orm:"column(ip);size(15)"`
	Status       uint8     `orm:"column(status)"`
	Md5          string    `orm:"column(md5);size(32)"`
	Sha1         string    `orm:"column(sha1);size(40)"`
	Id_RENAME    uint      `orm:"column(id)"`
	Aid          uint      `orm:"column(aid)"`
	Uid          uint      `orm:"column(uid)"`
	IsShow       uint8     `orm:"column(is_show)"`
	Http         string    `orm:"column(http);size(100)"`
}

func (t *Attachment) TableName() string {
	return "attachment"
}

func init() {
	orm.RegisterModel(new(Attachment))
}

// AddAttachment insert a new Attachment into database and returns
// last inserted Id on success.
func AddAttachment(m *Attachment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAttachmentById retrieves Attachment by Id. Returns error if
// Id doesn't exist
func GetAttachmentById(id int) (v *Attachment, err error) {
	o := orm.NewOrm()
	v = &Attachment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAttachment retrieves all Attachment matches certain condition. Returns empty list if
// no records exist
func GetAllAttachment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Attachment))
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

	var l []Attachment
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

// UpdateAttachment updates Attachment by Id and returns error if
// the record to be updated doesn't exist
func UpdateAttachmentById(m *Attachment) (err error) {
	o := orm.NewOrm()
	v := Attachment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAttachment deletes Attachment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAttachment(id int) (err error) {
	o := orm.NewOrm()
	v := Attachment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Attachment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
