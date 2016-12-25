package admin

import (
	"fox/controllers"
	"fox/service/admin"
	"fmt"
	"strconv"
	"fox/util/Response"
	"fox/models"
)

type TypeController struct {
	BaseController
}
func (c *TypeController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("ListChild", c.ListChild)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /type [get]
func (c *TypeController)List() {
	var ser admin.Type
	data, err := ser.Query(0)
	fmt.Println(err)
	c.Data["data"] = data
	c.Data["title"] = "类别-列表"
	c.Data["HtmlHead"] = controllers.ExecuteTemplateHtml("admin/type/head.html", c.Data)
	c.TplName = "admin/type/list.html"
}
//子类
// @router /type/list_child/:id [get]
func (c *TypeController)ListChild() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var ser admin.Type
	data, err := ser.Query(int_id)
	fmt.Println(err)
	c.Data["info"] = ""
	if int_id > 0 {
		var model *admin.Type
		data, err := model.Read(int_id)
		if err == nil {
			c.Data["info"] = data["info"]
		}
	}
	c.Data["id"] = id
	c.Data["data"] = data
	c.Data["title"] = "类别-子类-列表"
	c.Data["HtmlHead"] = controllers.ExecuteTemplateHtml("admin/type/head.html", c.Data)
	c.TplName = "admin/type/list_child.html"
}
//添加
// @router /type/add [get]
// @router /type/add/:id [get]
func (c *TypeController)Add() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	c.Data["type_id"] = id
	c.Data["type_id_name"] = "无"
	c.Data["parent_id_name"] = "无"
	c.Data["info"] = &models.Type{TypeId:int_id,IsDefault:0,IsDel:0,IsSystem:0,IsShow:1,IsChild:0}
	if int_id > 0 {
		var model *admin.Type
		data, err := model.Read(int_id)
		if err == nil {
			var t models.Type
			t = data["info"].(models.Type)
			c.Data["type_id_name"] = t.Name
		}else{
			c.Data["info"] = &models.Type{TypeId:0,IsDefault:0,IsDel:0,IsSystem:0,IsShow:1,IsChild:0}
			c.Data["type_id"] = 0
		}
	}
	c.Data["_method"] = "post"
	c.Data["title"] = "类别-添加"
	c.TplName = "admin/type/get.html"
}
//保存
// @router /type [post]
func (c *TypeController)Post() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	model := models.Type{}
	//参数传递
	if err := c.ParseForm(&model); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	//创建
	var serv admin.Type
	id, err := serv.Create(&model)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		rsp.Success("")
	}
}
//编辑
// @router /type/:id [get]
func (c *TypeController)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var model *admin.Type
	data, err := model.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(c.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
		c.Data["info"] = data["info"]
		c.Data["type_id_name"] = data["type_id_name"]
		c.Data["parent_id_name"] = data["parent_id_name"]
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["title"] = "类别-修改"
		c.TplName = "admin/type/get.html"
	}
}
//更新
// @router /type/:id [put]
func (c *TypeController)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	model := models.Type{}
	if err := c.ParseForm(&model); err != nil {
		rsp.Error(err.Error())
	}
	//更新
	var ser *admin.Type
	_, err := ser.Update(int_id, &model)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
//检测名称重复
// @router /type/check_name [post]
func (c *TypeController)CheckName() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	int_id,_  := c.GetInt("type_id")
	id,_  := c.GetInt("id")
	name := c.GetString("name")
	//创建
	var serv admin.Type
	ok, err := serv.CheckNameTypeId(int_id,name,id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		rsp.Success("")
	}
}