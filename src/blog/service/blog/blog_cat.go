package blog

import (
	"fmt"
	"blog/fox"
	"time"
	"blog/model"
	"blog/fox/db"
	"blog/service/conf"
)
//分类
type BlogCat struct {

}
//快速初始化
func NewBlogCatService() *BlogCat{
	return new(BlogCat)
}
//指定ID 数据列表
func (c *BlogCat)Query(cat_id int) ( *db.Paginator,  error) {
	query := make(map[string]interface{})
	query["cat_id"] = cat_id
	fields := []string{}
	mode := model.NewBlog()
	data, err := mode.GetAll(query, fields, "blog_id desc", 1, 999)
	//fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//创建
func (c *BlogCat)Create(m *model.Blog) (int, error) {
	fmt.Println("DATA:", m)
	if len(m.Title) < 1 {
		return 0,fox.NewError("标题 不能为空")
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeSystem = m.TimeAdd
	m.CatId = conf.TYPE_CAT
	//状态
	m.Status = 99
	o := db.NewDb()
	affected, err := o.Insert(m)
	if err != nil {
		return 0,fox.NewError("创建错误：" + err.Error())
	}
	stat := model.NewBlogStatistics()
	stat.BlogId = m.BlogId
	stat.StatisticsId = stat.BlogId
	id2, err := o.Insert(stat)
	if err != nil {
		return 0,fox.NewError("创建错误：" + err.Error())
	}
	if m.Tag != "" {
		var tagSer *BlogTag
		_, err := tagSer.CreateFromTags(m.BlogId, m.Tag, "")
		fmt.Println("TAG:", err)
	}
	fmt.Println("DATA:", m)
	fmt.Println("affected:", affected)
	fmt.Println("Statistics:", id2)
	return m.BlogId, nil
}
//更新
func (c *BlogCat)Update(id int, m *model.Blog) (int, error) {
	if id < 1 {
		return 0,fox.NewError("ID 错误")
	}
	mode := model.NewBlog()
	_, err := mode.GetById(id)
	if err != nil {
		return 0,fox.NewError("数据不存在")
	}
	if len(m.Title) < 1 {
		return 0,fox.NewError("标题 不能为空")
	}
	fmt.Println("DATA:", m)
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeSystem = m.TimeAdd
	m.CatId = conf.TYPE_CAT
	//状态
	m.Status = 99

	o := db.NewDb()
	num, err := o.Id(id).Update(m)
	if err != nil {
		return 0,fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println(num)
	//
	stat := model.NewBlogStatistics()
	stat.BlogId = id
	num2, err := o.Id(id).Update(stat)
	if err != nil {
		return 0,fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println(num2)
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//删除
func (c *BlogCat)Delete(id int) (bool, error) {
	if id < 1 {
		return false,fox.NewError("ID 错误")
	}
	mode := model.NewBlog()
	info, err := mode.GetById(id)
	if err != nil {
		return false, err
	}
	if info.CatId != conf.TYPE_CAT {
		return false,fox.NewError("不是栏目，不能删除")
	}
	num, err := mode.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num:", num)
	stat := model.NewBlogStatistics()
	num2, err := stat.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num2:", num2)
	return true, nil
}