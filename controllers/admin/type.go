package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/controllers"
	"github.com/foxiswho/blog-go/model"
	"github.com/foxiswho/blog-go/service/admin"
	"strconv"
)
//属性类别
type Type struct {
	Base
}

func (c *Type) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("ListChild", c.ListChild)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /type [get]
func (c *Type)List() {
	ser := admin.NewTypeService()
	data, err := ser.Query(0)
	if err != nil {
		c.Error(err.Error())
		return
	}
	c.Data["data"] = data
	c.Data["title"] = "类别-列表"
	c.Data["HtmlHead"] = controllers.ExecuteTemplateHtml("admin/type/head.html", c.Data)
	c.TplName = "admin/type/list.html"
}
//子类
// @router /type/list_child/:id [get]
func (c *Type)ListChild() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := admin.NewTypeService()
	data, err := ser.Query(int_id)
	if err != nil {
		c.Error(err.Error())
		return
	}
	c.Data["info"] = ""
	if int_id > 0 {
		ser := admin.NewTypeService()
		data, err := ser.Read(int_id)
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
func (c *Type)Add() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	mod := model.NewType()
	mod.TypeId = int_id
	mod.IsDefault = 0
	mod.IsDel = 0
	mod.IsSystem = 0
	mod.IsShow = 1
	mod.IsChild = 0
	c.Data["type_id"] = id
	c.Data["type_id_name"] = "无"
	c.Data["parent_id_name"] = "无"
	c.Data["info"] = mod
	if int_id > 0 {
		ser := admin.NewTypeService()
		data, err := ser.Read(int_id)
		if err == nil {
			var t model.Type
			t = data["info"].(model.Type)
			c.Data["type_id_name"] = t.Name
		} else {
			mod := model.NewType()
			mod.TypeId = 0
			mod.IsDefault = 0
			mod.IsDel = 0
			mod.IsSystem = 0
			mod.IsShow = 1
			mod.IsChild = 0
			c.Data["info"] = mod
			c.Data["type_id"] = 0
		}
	}
	c.Data["_method"] = "post"
	c.Data["title"] = "类别-添加"
	c.TplName = "admin/type/get.html"
}
//保存
// @router /type [post]
func (c *Type)Post() {
	mod := model.Type{}
	//参数传递
	if err := c.ParseForm(&mod); err != nil {
		c.Error(err.Error())
		return
	}
	//创建
	ser := admin.NewTypeService()
	id, err := ser.Create(&mod)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//编辑
// @router /type/:id [get]
func (c *Type)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := admin.NewTypeService()
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		c.Error(err.Error())
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
func (c *Type)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	mod := model.Type{}
	if err := c.ParseForm(&mod); err != nil {
		c.Error(err.Error())
	}
	//更新
	ser := admin.NewTypeService()
	_, err := ser.Update(int_id, &mod)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//检测名称重复
// @router /type/check_name [post]
func (c *Type)CheckName() {
	//ID 获取 格式化
	int_id, _ := c.GetInt("type_id")
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	//创建
	ser := admin.NewTypeService()
	ok, err := ser.CheckNameTypeId(int_id, name, id)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
