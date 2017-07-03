package controllers

import (
	"blog/service/blog"
	"fmt"
)
//标签控制器
type Tag struct {
	BaseNoLogin
}

//标签 列表页面
// @router / [get]
func (c *Tag) GetAll() {
	//标签获取
	idStr := c.Ctx.Input.Param(":tag")
	//查询变量
	query := make(map[string]interface{})
	query["name"] = idStr
	//初始化
	mode := blog.NewBlogTagService()
	//分页
	page, _ := c.GetInt("page")
	//查询
	data, err := mode.GetAll(query, []string{}, "blog_id desc", page, 20)
	//错误
	if err != nil {
		//c.Data["data"] = err.Error()
		fmt.Println(err.Error())
		c.Error(err.Error())
		return
	} else {
		//变量赋值
		c.Data["data"] = data
	}
	//模版
	c.SetTpl("blog/index.html")
}
