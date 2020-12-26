package api

import (
	"fmt"
	"github.com/foxiswho/blog-go/controllers"
	"github.com/foxiswho/blog-go/model"
	"github.com/foxiswho/blog-go/service/admin"
	"github.com/foxiswho/blog-go/service/blog"
	"github.com/foxiswho/blog-go/service/conf"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//
type Blog struct {
	controllers.BaseNoLogin
}

func (c *Blog) Create() {
	//接口秘钥验证
	site := admin.NewSiteService()
	config := site.SiteConfig()
	// markdown 换行
	token := c.GetString("token")
	token = strings.TrimSpace(token)
	fmt.Println("更新标志或接口秘钥=", config["app_api_token"],token)
	if !strings.EqualFold(config["app_api_token"], token) {
		c.Error("更新标志或接口秘钥 错误")
		return
	}

	blogModel := model.NewBlog()
	//参数传递
	blog_statistics := model.NewBlogStatistics()
	//表单结构体绑定
	if err := c.ParseForm(&blogModel); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error("表单绑定错误:" + err.Error())
		return
	}
	//表单结构体绑定
	if err := c.ParseForm(&blog_statistics); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error("表单绑定错误:" + err.Error())
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
	blogModel.Content = "博客收藏"
	is_life, _ := c.GetInt("is_life")
	if is_life == 1 {
		//生活
		blogModel.ModuleId = conf.MODULE_ID_WORK_OTHER
	} else {
		blogModel.ModuleId = conf.MODULE_ID_WORK
	}
	is_open, _ := c.GetInt("is_open")
	if is_open == 1 {
		//显示
		blogModel.IsOpen = 1
	} else {
		blogModel.IsOpen = 0
	}
	blogModel.Status = 99               //发布
	blogModel.IsRead = conf.READ_FINISH //已看
	blogModel.TypeId = 10005            //转载
	blogModel.SourceId = conf.APP_API   //接口
	blogModel.Url = blogModel.UrlSource
	//创建
	serv := blog.NewBlogService()
	id, err := serv.Create(blogModel, blog_statistics)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
