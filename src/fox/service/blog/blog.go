package blog

import (
	"fox/models"
	"fmt"
	"fox/util"
	UtilOrm "fox/util/orm"
	"fox/util/datetime"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
	"strconv"
	"errors"
	"reflect"
	"strings"
)
//博客模块ID
const TYPE_ID = 10006
const ORIGINAL = 10003

type Blog struct {

}

func (c *Blog)Query(cat_id, page int) (*UtilOrm.Page, error) {
	query := map[string]string{}
	query["cat_id"] = strconv.Itoa(cat_id)
	var fields []string
	sortby := []string{"Id"}
	order := []string{"desc"}
	var limit int
	limit = 20
	data, err := GetAllBlog(query, fields, sortby, order, page, limit)
	if err == nil {
		return data, nil
	}
	return nil, err
}
//详情
func (c *Blog)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil, &util.Error{Msg:"ID 错误"}
	}

	data, err := models.GetBlogById(id)
	if err != nil {
		fmt.Println(err)
		return nil, &util.Error{Msg:"数据不存在"}
	}
	//整合
	m := make(map[string]interface{})
	m["Blog"] = *data
	m["Content"] = string(blackfriday.MarkdownBasic([]byte(data.Content)))
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)

	var Statistics *models.BlogStatistics
	Statistics, err = models.GetBlogStatisticsById(id)
	if err == nil {
		m["Statistics"] = Statistics
	} else {
		//错误屏蔽
		err = nil
		//初始化赋值
		m["Statistics"] = models.BlogStatistics{}
	}

	//fmt.Println(m)
	return m, err
}
//详情
func (c *Blog)ReadByUrlRewrite(id string) (map[string]interface{}, error) {
	if id == "" {
		return nil, &util.Error{Msg:"URL 错误"}
	}

	data, err := GetBlogByUrlRewrite(id)
	if err != nil {
		return nil, &util.Error{Msg:"数据不存在"}
	}
	//整合
	m := make(map[string]interface{})
	m["Blog"] = *data
	m["Content"] = string(blackfriday.MarkdownBasic([]byte(data.Content)))
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)

	var Statistics *models.BlogStatistics
	Statistics, err = models.GetBlogStatisticsById(data.Id)
	if err == nil {
		m["Statistics"] = Statistics
	} else {
		//错误屏蔽
		err = nil
		//初始化赋值
		m["Statistics"] = models.BlogStatistics{}
	}

	//fmt.Println(m)
	return m, err
}

//创建
func (c *Blog)Create(m *models.Blog, stat *models.BlogStatistics) (int64, error) {

	fmt.Println("DATA:", m)
	if len(m.Title) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	if len(m.Content) < 1 {
		return 0, &util.Error{Msg:"内容 不能为空"}
	}
	if m.TypeId < 10003 {
		return 0, &util.Error{Msg:"请选择类别"}
	}
	if m.TypeId > 10005 {
		return 0, &util.Error{Msg:"类别 错误"}
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeSystem = m.TimeAdd

	//状态
	if m.Status < 0 {
		m.Status = 0
	}
	if m.Status > 99 {
		m.Status = 99
	}
	id, err := models.AddBlog(m)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}

	stat.BlogId = int(id)
	stat.Id = stat.BlogId
	id2, err := models.AddBlogStatistics(stat)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	if m.Tag != "" {
		var tagSer *BlogTag
		_, err := tagSer.CreateFromTags(int(id), m.Tag, "")
		fmt.Println("TAG:", err)
	}

	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	fmt.Println("Statistics:", id2)
	return id, nil
}
//更新
func (c *Blog)Update(id int, m *models.Blog, stat *models.BlogStatistics) (int, error) {
	if id < 1 {
		return 0, &util.Error{Msg:"ID 错误"}
	}

	info, err := models.GetBlogById(id)
	if err != nil {
		return 0, &util.Error{Msg:"数据不存在"}
	}
	if len(m.Title) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	if len(m.Content) < 1 {
		return 0, &util.Error{Msg:"内容 不能为空"}
	}
	if m.TypeId < 10003 {
		return 0, &util.Error{Msg:"请选择类别"}
	}
	if m.TypeId > 10005 {
		return 0, &util.Error{Msg:"类别 错误"}
	}
	fmt.Println("DATA:", m)
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeSystem = m.TimeAdd

	//状态
	if m.Status < 0 {
		m.Status = 0
	}
	if m.Status > 99 {
		m.Status = 99
	}

	m.Id = id
	_, err = c.UpdateById(m, "title", "content", "status", "is_open", "time_add", "author", "url_source", "url_rewrite", "url", "thumb", "sort", "description", "tag")
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}

	stat.BlogId = id
	stat.Id = stat.BlogId
	_, err = c.UpdateBlogStatisticsById(stat, "blog_id", "seo_title", "seo_keyword", "seo_description")
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	//标签 创建和删除
	var tagSer *BlogTag
	_, err = tagSer.CreateFromTags(id, m.Tag, info.Tag)
	fmt.Println("TAG:", err)
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//更新
func (c *Blog)UpdateById(m *models.Blog, cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}
//更新
func (c *Blog)UpdateBlogStatisticsById(m *models.BlogStatistics, cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}
//删除
func (c *Blog)Delete(id int) (bool, error) {
	if id < 1 {
		return false, &util.Error{Msg:"ID 错误"}
	}
	err := models.DeleteBlog(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	err = c.DeleteBlogStatisticsByBlogId(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	return true, nil
}
func (c *Blog)DeleteBlogStatisticsByBlogId(id int) (err error) {
	o := orm.NewOrm()
	v := models.BlogStatistics{BlogId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&models.BlogStatistics{BlogId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
//根据自定义伪静态查询
func GetBlogByUrlRewrite(id string) (v *models.Blog, err error) {
	o := orm.NewOrm()
	v = &models.Blog{UrlRewrite: id}
	if err = o.Read(v, "url_rewrite"); err == nil {
		return v, nil
	}
	return nil, err
}
//详情
func (c *Blog)CheckTitleById(cat_id int, str string, id int) (bool, error) {
	if str == "" {
		return false, &util.Error{Msg:"名称 不能为空"}
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Blog))
	qs = qs.Filter("cat_id", cat_id)
	qs = qs.Filter("title", str)
	if id > 0 {
		qs = qs.Filter("blog_id__nq", id)
	}
	count, err := qs.Count()
	fmt.Println(count)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if count == 0 {
		return true, nil
	}
	return false, &util.Error{Msg:"已存在"}
}

// GetAllBlog retrieves all Blog matches certain condition. Returns empty list if
// no records exist
func GetAllBlog(query map[string]string, fields []string, sortby []string, order []string,
page int, limit int) (*UtilOrm.Page, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Blog))
	var l []models.Blog
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	count, err := qs.Count()
	if err != nil {
		fmt.Println(err)
	}
	if page < 1 {
		page = 1
	}
	Query := UtilOrm.PageUtil(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}
	//fmt.Println("Query",Query)
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
	fmt.Println("Offset", Query)
	fmt.Println("Offset", Query.Offset)

	var ml []interface{}
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, int64(Query.Offset)).All(&l, fields...); err == nil {
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
		Query.Data = ml
		return Query, nil
	}
	return nil, err
}