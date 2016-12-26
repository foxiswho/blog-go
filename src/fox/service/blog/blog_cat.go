package blog

import (
	"fox/models"
	"fmt"
	"strconv"
	"fox/util"
	"time"
)

//博客分类属性 栏目ID
const CAT_ID = 10001

type BlogCat struct {

}

func (c *BlogCat)Query(cat_id int) (data []interface{}, err error) {
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
//创建
func (c *BlogCat)Create(m *models.Blog) (int64, error) {
	fmt.Println("DATA:", m)
	if len(m.Title) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeSystem = m.TimeAdd
	m.CatId = CAT_ID
	//状态
	m.Status = 99
	id, err := models.AddBlog(m)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	stat := &models.BlogStatistics{}
	stat.BlogId = int(id)
	stat.Id = stat.BlogId
	id2, err := models.AddBlogStatistics(stat)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	fmt.Println("Statistics:", id2)
	return id, nil
}
//更新
func (c *BlogCat)Update(id int, m *models.Blog) (int, error) {
	if id < 1 {
		return 0, &util.Error{Msg:"ID 错误"}
	}
	_, err := models.GetBlogById(id)
	if err != nil {
		return 0, &util.Error{Msg:"数据不存在"}
	}
	if len(m.Title) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	fmt.Println("DATA:", m)
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeSystem = m.TimeAdd
	m.CatId = CAT_ID
	//状态
	m.Status = 99

	m.Id = id
	var blogSer *Blog
	_, err = blogSer.UpdateById(m, "title", "remark", "status", "is_del", "time_add", "url_source", "url_rewrite", "url", "sort", )
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	stat := &models.BlogStatistics{}
	stat.BlogId = int(id)
	stat.Id = stat.BlogId
	_, err = blogSer.UpdateBlogStatisticsById(stat, "blog_id")
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//删除
func (c *BlogCat)Delete(id int) (bool, error) {
	if id < 1 {
		return false, &util.Error{Msg:"ID 错误"}
	}
	var blogSer *Blog
	info, err := models.GetBlogById(id)
	if err != nil {
		return false, err
	}
	if info.CatId != CAT_ID {
		return false, &util.Error{Msg:"不是栏目，不能删除"}
	}
	err = models.DeleteBlog(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	err = blogSer.DeleteBlogStatisticsByBlogId(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	return true, nil
}