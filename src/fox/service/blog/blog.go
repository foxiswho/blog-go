package blog

import (
	"fox/models"
	"fmt"
	"fox/util"
	"fox/util/datetime"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
	"strconv"
)
//博客模块ID
const TYPE_ID  = 10006
type Blog struct {

}

func (c *Blog)Query(cat_id int) (data []interface{}, err error) {
	query := map[string]string{}
	query["cat_id"] = strconv.Itoa(cat_id)
	var fields []string
	sortby := []string{"Id"}
	order := []string{"desc"}
	var offset int64
	var limit int64
	offset = 0
	limit = 20
	data, err = models.GetAllBlog(query, fields, sortby, order, offset, limit)
	//fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//详情
func (c *Blog)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil, &util.Error{Msg:"ID 错误"}
	}

	data, err := models.GetBlogById(id)
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
		_, err := tagSer.CreateFromTags(int(id),m.Tag, "")
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
	_, err = tagSer.CreateFromTags(id,m.Tag, info.Tag)
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
