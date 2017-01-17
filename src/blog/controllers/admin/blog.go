package admin

import (
	"strconv"
	"blog/service/blog"
	"fmt"
	"blog/model"
	"blog/fox/url"
	"blog/fox/file"
	"blog/service/admin"
	"blog/service/conf"
)

type Blog struct {
	Base
}

func (c *Blog) URLMapping() {
	c.Mapping("CheckTitle", c.CheckTitle)
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Detail", c.Detail)
}
//检测名称重复
// @router /blog/check_title [post]
func (c *Blog)CheckTitle() {
	//ID 获取 格式化
	int_id, _ := c.GetInt("cat_id")
	id, _ := c.GetInt("id")
	name := c.GetString("title")
	//创建
	var serv blog.Blog
	ok, err := serv.CheckTitleById(int_id, name, id)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
//列表
// @router /blog [get]
func (c *Blog)List() {
	where := make(map[string]interface{})
	wd := c.GetString("wd")
	if len(wd) > 0 {
		where["title like ? "] = "%" + wd + "%"
	}
	where["type=?"] = conf.TYPE_ARTICLE
	mod := blog.NewBlogService()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "blog_id desc", page, 20)
	if err!=nil{
		fmt.Println(err)
	}
	c.Data["data"] = data
	c.Data["wd"] = wd
	c.Data["title"] = "博客-列表"
	c.TplName = "admin/blog/list.html"
}
//编辑
// @router /blog/:id [get]
func (c *Blog)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var ser *blog.Blog
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Data["info"] = data["info"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "博客-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["TYPE_ID"] = conf.TYPE_ID
		ser := blog.NewBlogSyncMappingService()
		mod, err := ser.GetAppId(conf.APP_CSDN, int_id)
		if err == nil {
			c.Data["app_csdn_id"] = mod.Id
		}
		maps := make(map[string]interface{})
		maps["type_id"] = conf.TYPE_ID
		maps["id"] = int_id
		maps["aid"] = c.Session.Aid
		cry, err := file.TokeMake(maps)
		if err != nil {
			fmt.Println("令牌加密错误：" + err.Error())
		}
		c.Data["upload_token"] = cry
		c.TplName = "admin/blog/get.html"
	}
}
//添加
// @router /blog/add [get]
func (c *Blog)Add() {
	mod := blog.NewBlogService()
	blog := model.NewBlog()
	blog.Author = c.Site.GetString("AUTHOR")
	blog.IsOpen = 1
	blog.Status = 99
	mod.Blog = blog
	mod.BlogStatistics = &model.BlogStatistics{}
	mod.TypeId = conf.ORIGINAL
	c.Data["info"] = mod
	c.Data["TYPE_ID"] = conf.TYPE_ID
	c.Data["_method"] = "post"
	c.Data["title"] = "博客-添加"
	maps := make(map[string]interface{})
	maps["type_id"] = conf.TYPE_ID
	maps["id"] = 0
	maps["aid"] = c.Session.Aid
	cry, err := file.TokeMake(maps)
	if err != nil {
		fmt.Println("令牌加密错误：" + err.Error())
	}
	//fmt.Println("令牌加密："+cry)
	//c1,err:=file.TokenDeCode(cry)
	//fmt.Println("令牌解密：")
	//fmt.Println(c1)
	//fmt.Println(err)
	c.Data["upload_token"] = cry
	c.TplName = "admin/blog/get.html"
}
//保存
// @router /blog [post]
func (c *Blog)Post() {
	blogModel := model.NewBlog()
	//参数传递
	blog_statistics := model.NewBlogStatistics()
	if err := url.ParseForm(c.Input(), blogModel); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	if err := url.ParseForm(c.Input(), blog_statistics); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	if blogModel.TimeAdd.IsZero() {
		//日期
		date, ok := c.GetDateTime("time_add")
		if ok {
			blogModel.TimeAdd = date
		}
	}
	//创建
	serv := blog.NewBlogService()
	id, err := serv.Create(blogModel, blog_statistics)
	if err != nil {
		c.Error(err.Error())
	} else {
		admin.NewAttachmentService().UpdateByTypeIdId(conf.TYPE_ID, c.Session.Aid, id)
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//查看
// @router /blog/detail/:id [get]
func (c *Blog)Detail() {
	c.Get()
	c.Data["title"] = "博客-查看"
	c.TplName = "admin/blog/detail.html"
}
//更新
// @router /blog/:id [put]
func (c *Blog)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	blogMoel := model.NewBlog()
	blog_statistics := model.NewBlogStatistics()
	if err := url.ParseForm(c.Input(), blogMoel); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	if err := url.ParseForm(c.Input(), blog_statistics); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	if blogMoel.TimeAdd.IsZero() {
		//日期
		date, ok := c.GetDateTime("time_add")
		if ok {
			blogMoel.TimeAdd = date
		}
	}
	//更新
	ser := blog.NewBlogService()
	_, err := ser.Update(int_id, blogMoel, blog_statistics)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//删除
// @router /blog/:id [delete]
func (c *Blog)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := blog.NewBlogService()
	_, err := ser.Delete(int_id)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}


