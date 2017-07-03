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
//博客控制器
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
	//查询
	where := make(map[string]interface{})
	wd := c.GetString("wd")
	if len(wd) > 0 {
		where["title like ? "] = "%" + wd + "%"
	}
	where["type=?"] = conf.TYPE_ARTICLE
	//初始化
	mod := blog.NewBlogService()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "sort,blog_id desc", page, 20)
	if err != nil {
		c.Error(err.Error())
		return
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
	//初始化获取信息
	var ser *blog.Blog
	data, err := ser.Read(int_id)
	if err != nil {
		c.Error(err.Error())
		return
	} else {
		c.Data["info"] = data["info"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "博客-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["TYPE_ID"] = conf.TYPE_ID
		//初始化好 获取绑定 CSDN博客文章ID
		ser := blog.NewBlogSyncMappingService()
		mod, err := ser.GetAppId(conf.APP_CSDN, int_id)
		if err == nil {
			c.Data["app_csdn_id"] = mod.Id
		}
		//上传令牌 初始化
		maps := make(map[string]interface{})
		maps["type_id"] = conf.TYPE_ID
		maps["id"] = int_id
		maps["aid"] = c.Session.Aid
		cry, err := file.TokeMake(maps)
		if err != nil {
			fmt.Println("令牌加密错误：" + err.Error())
			c.Error(err.Error())
			return
		}
		c.Data["upload_token"] = cry
		c.TplName = "admin/blog/get.html"
	}
}
//添加
// @router /blog/add [get]
func (c *Blog)Add() {
	mod := blog.NewBlogService()
	//初始化状态
	blogMod := model.NewBlog()
	blogMod.Author = c.Site.GetString("author")
	blogMod.IsOpen = 1				//启用
	blogMod.Status = 99				//发布
	blogMod.IsRead = conf.READ_NOT	//未读
	blogMod.TypeId = conf.ORIGINAL	//原创
	blogMod.Type   = conf.TYPE_ARTICLE //文章
	mod.Blog = blogMod
	mod.BlogStatistics = &model.BlogStatistics{}
	//模版参数设置
	c.Data["info"] = mod
	c.Data["TYPE_ID"] = conf.TYPE_ID
	c.Data["_method"] = "post"
	c.Data["title"] = "博客-添加"
	//上传令牌设置
	maps := make(map[string]interface{})
	maps["type_id"] = conf.TYPE_ID
	maps["id"] = 0
	maps["aid"] = c.Session.Aid
	cry, err := file.TokeMake(maps)
	if err != nil {
		fmt.Println("令牌加密错误：" + err.Error())
		c.Error(err.Error())
		return
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
	//表单结构体绑定
	if err := url.ParseForm(c.Input(), blogModel); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//表单结构体绑定
	if err := url.ParseForm(c.Input(), blog_statistics); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//日期判断
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
		//附件归属更新
		admin.NewAttachmentService().UpdateByTypeIdId(conf.TYPE_ID, c.Session.Aid, id)
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//查看
// @router /blog/detail/:id [get]
func (c *Blog)Detail() {
	//直接使用编辑页数据
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
	//表单 与结构体绑定
	if err := url.ParseForm(c.Input(), blogMoel); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error("表单绑定错误"+err.Error())
		return
	}
	//表单与结构体绑定
	if err := url.ParseForm(c.Input(), blog_statistics); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error("表单绑定错误"+err.Error())
		return
	}
	//时间判断
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


