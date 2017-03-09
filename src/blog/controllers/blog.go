package controllers

import (
	"strconv"

	"regexp"
	"blog/service/blog"
	"fmt"
	"blog/service/conf"
)
//首页 博客控制器
type Blog struct {
	BaseNoLogin
}


//文章详情
// @router /article/:id [get]
func (c *Blog) Get() {
	idStr := c.Ctx.Input.Param(":id")
	//初始化
	ser := blog.NewBlogService()
	var err error
	var read map[string]interface{}
	//正则匹配 该idStr 数字
	if ok, _ := regexp.Match(`^\d+$`, []byte(idStr)); ok {
		//字符串转换为数值类型
		id, _ := strconv.Atoi(idStr)
		//获取该ID信息
		read, err = ser.Read(id)
	} else {
		//根据自定义URL，获取该信息，
		read, err = ser.ReadByUrlRewrite(idStr)
	}
	//错误输出
	if err != nil {
		c.Error(err.Error())
		//直接返回
		return
	} else {
		//赋值
		tmp := read["info"]
		//interface变量转换为结构体
		B := tmp.(*blog.Blog)
		if B.Blog.IsOpen != 1 {
			c.Error("信息 不存在")
			return
		}
		if B.Blog.Status != 99 {
			c.Error("信息 不存在")
			return
		}
		fmt.Println(B)
		//更新浏览次数
		_, err := ser.UpdateRead(B.Blog.BlogId)
		if err != nil {
			fmt.Println(err)
		}
		//模版变量赋值
		c.Data["title"] = read["title"]
		c.Data["info"] = read["info"]
		c.Data["TimeAdd"] = read["TimeAdd"]
		c.Data["Content"] = read["Content"]
	}
	//模版
	c.TplName = "blog/get.html"
}

//博客首页，列表页面
// @router / [get]
func (c *Blog) GetAll() {
	//查询变量
	query := make(map[string]interface{})
	query["type=?"] = conf.TYPE_ARTICLE
	query["is_open=?"] = 1
	query["status=?"] = 99
	//初始化
	mode := blog.NewBlogService()
	//分页
	id := c.Ctx.Input.Param(":page")
	//字符串转换为数值
	page, _ := strconv.Atoi(id)
	//查询
	data, err := mode.GetAll(query, []string{}, "blog_id desc", page, 10)
	//错误输出
	if err != nil {
		fmt.Println("err", err)
		c.Error(err.Error())
		return
	} else {
		//模版变量
		c.Data["data"] = data
	}
	//模版
	c.TplName = "blog/index.html"

}
