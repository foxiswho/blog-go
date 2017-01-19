package admin

import (
	"blog/model"
	"fmt"
)
//博客标签
type BlogTag struct {
	Base
}

func (c *BlogTag) URLMapping() {
	c.Mapping("List", c.List)
}
//列表
// @router /blog/tag [get]
func (c *BlogTag)List() {
	//查询
	query := make(map[string]interface{})
	fields := []string{}
	str := c.GetString("wd")
	if str != "" {
		query["name"] = str
	}
	page, _ := c.GetInt("page")
	//初始化
	mode := model.NewBlogTag()
	data, err := mode.GetAll(query, fields, "tag_id desc",page, 20)
	if err!=nil{
		fmt.Println(err.Error())
		c.Error(err.Error())
		return
	}
	if c.IsAjax(){
		c.Data["json"]=data
		c.ServeJSON()
	}else{
		//println(data)
		c.Data["data"] = data
		c.Data["title"] = "博客-TAG-列表"
		c.TplName = "admin/blog/tag/list.html"
	}
}