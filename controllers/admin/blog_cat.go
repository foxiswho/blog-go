package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/model"
	"github.com/foxiswho/blog-go/service/blog"
	"github.com/foxiswho/blog-go/service/conf"
	"strconv"
)
//博客栏目控制器
type BlogCat struct {
	Base
}

func (c *BlogCat) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /blog/cat [get]
func (c *BlogCat)List() {
	where := make(map[string]interface{})
	where["type=?"] = conf.TYPE_CAT
	mod := model.NewBlog()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "blog_id desc", page, 999)
	//println(data)
	println(err)
	c.Data["data"] = data
	c.Data["title"] = "博客-分类-列表"
	c.TplName = "admin/blog/cat/list.html"
}
//编辑
// @router /blog/cat/:id [get]
func (c *BlogCat)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := blog.NewBlogService()
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Data["info"] = data["info"]
		c.Data["statistics"] = data["Statistics"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "博客-分类-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["type"] = conf.TYPE_CAT
		c.TplName = "admin/blog/cat/get.html"
	}
}
//添加
// @router /blog/cat/add [get]
func (c *BlogCat)Add() {
	c.Data["type"] = conf.TYPE_CAT
	c.Data["_method"] = "post"
	c.Data["title"] = "博客-分类-添加"
	c.TplName = "admin/blog/cat/get.html"
}
//保存
// @router /blog/cat [post]
func (c *BlogCat)Post() {
	blogModel := model.NewBlog()

	//参数传递
	if err := c.ParseForm(&blogModel); err != nil {
		c.Error(err.Error())
		return
	}
	blogModel.Type = conf.TYPE_CAT
	//日期
	date, ok := c.GetDateTime("time_add")
	if ok {
		blogModel.TimeAdd = date
	}
	//创建
	serv := blog.NewBlogCatService()
	id, err := serv.Create(blogModel)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//更新
// @router /blog/cat/:id [put]
func (c *BlogCat)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	blogModel := model.NewBlog()
	if err := c.ParseForm(&blogModel); err != nil {
		c.Error(err.Error())
		return
	}
	blogModel.Type = conf.TYPE_CAT
	//日期
	date, ok := c.GetDateTime("time_add")
	if ok {
		blogModel.TimeAdd = date
	}
	//更新
	ser := blog.NewBlogCatService()
	_, err := ser.Update(int_id, blogModel)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//删除
// @router /blog/cat/:id [delete]
func (c *BlogCat)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := blog.NewBlogCatService()
	_, err := ser.Delete(int_id)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
