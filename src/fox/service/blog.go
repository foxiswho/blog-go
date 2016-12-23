package service

import (
	"fox/models"
	"fmt"
	"fox/util"
	"fox/util/datetime"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
)

type Blog struct {

}

func (this *Blog)Query() (data []interface{}, err error) {
	var query map[string]string
	var fields []string
	sortby :=[]string{"Id"}
	order :=[]string{"desc"}
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
func (this *Blog)Read(id int) (map[string]interface{}, error) {
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
	m["Content"]=string(blackfriday.MarkdownBasic([]byte(data.Content)))
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
//创建
func (this *Blog)Create(blog *models.Blog, stat *models.BlogStatistics) (int64, error) {

	fmt.Println("BLOG DATA:", blog)
	if len(blog.Title) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	if len(blog.Content) < 1 {
		return 0, &util.Error{Msg:"内容 不能为空"}
	}

	//时间
	if blog.TimeAdd.IsZero() {
		blog.TimeAdd = time.Now()
	}
	blog.TimeSystem = blog.TimeAdd

	//状态
	if blog.Status < 0 {
		blog.Status = 0
	}
	if blog.Status > 99 {
		blog.Status = 99
	}
	id, err := models.AddBlog(blog)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}

	stat.BlogId = int(id)
	stat.Id =stat.BlogId
	id2, err := models.AddBlogStatistics(stat)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	fmt.Println("BLOG DATA:", blog)
	fmt.Println("BlogId:", id)
	fmt.Println("BlogStatistics:", id2)
	return id, nil
}
//创建
func (this *Blog)Update(id int,blog *models.Blog, stat *models.BlogStatistics) (int, error) {
	if id < 1 {
		return 0, &util.Error{Msg:"ID 错误"}
	}

	_, err := models.GetBlogById(id)
	if err != nil {
		return 0, &util.Error{Msg:"数据不存在"}
	}
	if len(blog.Title) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	if len(blog.Content) < 1 {
		return 0, &util.Error{Msg:"内容 不能为空"}
	}
	fmt.Println("BLOG DATA:", blog)
	//时间
	if blog.TimeAdd.IsZero() {
		blog.TimeAdd = time.Now()
	}
	blog.TimeSystem = blog.TimeAdd

	//状态
	if blog.Status < 0 {
		blog.Status = 0
	}
	if blog.Status > 99 {
		blog.Status = 99
	}

	blog.Id=id
	_, err = this.UpdateBlogById(blog,"title","content","status","is_open","time_add","author","url_source","url_rewrite","url","thumb","sort","description","tag")
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}

	stat.BlogId = int(id)
	stat.Id =stat.BlogId
	_, err = this.UpdateBlogStatisticsById(stat,"seo_title","seo_keyword","seo_description")
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	fmt.Println("BLOG DATA:", blog)
	fmt.Println("BlogId:", id)
	return id, nil
}
//更新
func (this *Blog)UpdateBlogById(m *models.Blog, cols ...string) (num int64,err error) {
	o := orm.NewOrm()
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num,nil
	}
	return 0,err
}
//更新
func (this *Blog)UpdateBlogStatisticsById(m *models.BlogStatistics, cols ...string) (num int64,err error) {
	o := orm.NewOrm()
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num,nil
	}
	return 0,err
}