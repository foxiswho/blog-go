package controllers

import (
	"strconv"

	"regexp"
	"fox/service/blog"
	"fmt"
	"fox/model"
)

type BlogController struct {
	BaseNoLoginController
}


// GetOne ...
// @Title Get One
// @Description get Blog by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Blog
// @Failure 403 :id is empty
// @router /article/:id [get]
func (c *BlogController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	var ser *blog.Blog
	var err error
	var read map[string]interface{}
	if ok, _ := regexp.Match(`^\d+$`, []byte(idStr)); ok {
		id, _ := strconv.Atoi(idStr)
		read, err = ser.Read(id)
	} else {
		read, err = ser.ReadByUrlRewrite(idStr)
	}
	if err != nil {
		c.Error(err.Error())
		return
	} else {
		c.Data["info"] = read["info"]
		c.Data["statistics"] = read["Statistics"]
		c.Data["TimeAdd"] = read["TimeAdd"]
		c.Data["Content"] = read["Content"]
	}
	c.TplName = "blog/get.html"
}

// GetAll ...
// @Title Get All
// @Description get Blog
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Blog
// @Failure 403
// @router / [get]
func (c *BlogController) GetAll() {
	fields := []string{}
	orderBy := "blog_id desc"
	query := make(map[string]interface{})
	query["type=?"] = blog.TYPE_ARTICLE
	mode := model.NewBlog()
	//分页
	page, _ := c.GetInt("page")
	data, err := mode.GetAll(query, fields, orderBy, page, 20)
	fmt.Println("err", err)
	fmt.Println("data", data)
	if err != nil {
		//c.Data["data"] = err.Error()
		fmt.Println(err.Error())
	} else {
		c.Data["data"] = data
	}
	c.TplName = "blog/index.html"
	//fmt.Println("==========")
	//db:=db.NewDb()
	//bb:=make([]model.Blog,0)
	//err=db.Where("blog_id =?",54).Find(&bb)
	//fmt.Println(err)
	//fmt.Println(bb)
	//fmt.Println("==========")
	//b2:=new(model.Blog)
	//var ok bool
	//ok,err=db.Get(b2)
	//fmt.Println(err)
	//fmt.Println(ok)
	//fmt.Println(b2)
	//fmt.Println("==========")
	//b22:=new(model.Blog)
	//err=db.Find(&b22)
	//fmt.Println(err)
	//fmt.Println(b22)
	//where :=make(map[string]interface{})
	//where["blog_id in (?)"]=[]string{"53","54"}
	////where["blog_id in (?)"]=[]int{53,54}
	//o:=db.Filter(where).OrderBy("blog_id asc")
	//err=o.Find(&bb)
	//fmt.Println(err)
	//fmt.Println(bb)
	//q := make(map[string]interface{})
	//q["content like ?"] = "%jpeg%"
	//b := model.NewBlog()
	//pag, err := b.GetAll(q, []string{}, "blog_id", 1, 20)
	////var b *blog.Blog
	////pag,err:=b.GetAll(q,[]string{},"blog_id",page,20)
	//fmt.Println(err)
	//fmt.Println(pag)

}