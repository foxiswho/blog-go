package blog

import (
	"fmt"
	"fox/util"
	"fox/util/datetime"
	"time"
	"github.com/russross/blackfriday"
	"fox/util/db"
	"fox/model"
)

const (
	//博客模块ID
	TYPE_ID = 10006
	//原创
	ORIGINAL = 10003
	//栏目 博客分类属性 栏目ID
	TYPE_CAT = 10001
	//文章
	TYPE_ARTICLE = 0
)

//

type Blog struct {

}
func NewBlogService() *Blog{
	return new(Blog)
}
//详情
func (c *Blog)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil, &util.Error{Msg:"ID 错误"}
	}
	mode := model.NewBlog()
	data, err := mode.GetById(id)
	fmt.Println(data)
	if err != nil {
		fmt.Println(err)
		return nil, &util.Error{Msg:"数据不存在"}
	}
	//整合
	m := make(map[string]interface{})
	m["info"] = *data
	m["Content"] = string(blackfriday.MarkdownBasic([]byte(data.Content)))
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)
	//主键ID值和blog_id值一样所以这里直接取值
	Statistics := model.NewBlogStatistics()
	StatisticsData, err := Statistics.GetById(id)
	if err == nil {
		m["Statistics"] = StatisticsData
	} else {
		//错误屏蔽
		err = nil
		//初始化赋值
		m["Statistics"] = Statistics
	}
	//fmt.Println(m)
	return m, err
}
//详情
func (c *Blog)ReadByUrlRewrite(id string) (map[string]interface{}, error) {
	if id == "" {
		return nil, &util.Error{Msg:"URL 错误"}
	}

	data, err := c.GetBlogByUrlRewrite(id)
	if err != nil {
		return nil, &util.Error{Msg:"数据不存在"}
	}
	//整合
	m := make(map[string]interface{})
	m["Blog"] = *data
	m["Content"] = string(blackfriday.MarkdownBasic([]byte(data.Content)))
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)

	Statistics := model.NewBlogStatistics()
	StatisticsData, err := Statistics.GetById(data.BlogId)
	if err == nil {
		m["Statistics"] = StatisticsData
	} else {
		//错误屏蔽
		err = nil
		//初始化赋值
		m["Statistics"] = Statistics
	}

	//fmt.Println(m)
	return m, err
}

//创建
func (c *Blog)Create(m *model.Blog, stat *model.BlogStatistics) (int64, error) {

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
	o := db.NewDb()
	id, err := o.Insert(m)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	stat.BlogId = int(id)
	stat.StatisticsId = stat.BlogId
	id2, err := o.Insert(stat)
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
func (c *Blog)Update(id int, m *model.Blog, stat *model.BlogStatistics) (int, error) {
	if id < 1 {
		return 0, &util.Error{Msg:"ID 错误"}
	}
	mode := model.NewBlog()
	info, err := mode.GetById(id)
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

	o := db.NewDb()
	num, err := o.Id(id).Update(m)
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	fmt.Println("============",num)
	//
	stat.BlogId = id
	o = db.NewDb()
	num2, err := o.Id(id).Update(stat)
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	fmt.Println(num2)
	//标签 创建和删除
	var tagSer *BlogTag
	_, err = tagSer.CreateFromTags(id, m.Tag, info.Tag)
	fmt.Println("TAG:", err)
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//更新
func (c *Blog)UpdateById(m *model.Blog, cols ...interface{}) (num int64, err error) {
	o := db.NewDb()
	if num, err = o.Id(m.BlogId).Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}
//更新
func (c *Blog)UpdateBlogStatisticsById(m *model.BlogStatistics, cols ...interface{}) (num int64, err error) {
	o := db.NewDb()
	if num, err = o.Id(m.StatisticsId).Update(m, cols...); err == nil {
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
	mode := model.NewBlog()
	num, err := mode.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num:", num)
	num2, err := c.DeleteBlogStatisticsByBlogId(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num2:", num2)
	return true, nil
}
func (c *Blog)DeleteBlogStatisticsByBlogId(id int) (int64, error) {
	o := db.NewDb()
	mode := model.NewBlogStatistics()
	mode.BlogId = id
	num, err := o.Delete(mode)
	if err == nil {
		return num, nil
	}
	return num, err
}
//根据自定义伪静态查询
func (c *Blog)GetBlogByUrlRewrite(id string) (v *model.Blog, err error) {
	o := db.NewDb()
	v = new(model.Blog)
	if err = o.Find(v); err == nil {
		return v, nil
	}
	return nil, err
}
//详情
func (c *Blog)CheckTitleById(cat_id int, str string, id int) (bool, error) {
	if str == "" {
		return false, &util.Error{Msg:"名称 不能为空"}
	}
	mode := new(model.Blog)
	where := make(map[string]interface{})
	where["cat_id"] = cat_id
	where["title"] = str
	if id > 0 {
		where["blog_id!=?"] = id
	}
	count, err := db.Filter(where).Count(mode)

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(count)
	if count == 0 {
		return true, nil
	}
	return false, &util.Error{Msg:"已存在"}
}