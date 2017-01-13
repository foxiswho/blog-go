package admin

import (
	"blog/controllers"
	"blog/service/admin"
	"fmt"
	"strconv"
	"blog/util/Response"
	"blog/model"
	"blog/util/url"
	"blog/service"
)

type MemberGroup struct {
	Base
}

func (c *MemberGroup) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /member_group [get]
func (c *MemberGroup)List() {
	ser := admin.NewTypeService()
	data, err := ser.Query(service.MEMBER_GROUP)
	fmt.Println(err)
	c.Data["data"] = data
	c.Data["title"] = "用户组-列表"
	c.Data["HtmlHead"] = controllers.ExecuteTemplateHtml("admin/type/head.html", c.Data)
	c.TplName = "admin/member_group/list.html"
}
//添加
// @router /member_group/add [get]
func (c *MemberGroup)Add() {
	mod := model.NewType()
	mod.TypeId = service.MEMBER_GROUP
	mod.IsDefault = 0
	mod.IsDel = 0
	mod.IsSystem = 0
	mod.IsShow = 1
	mod.IsChild = 0
	c.Data["type_id"] = service.MEMBER_GROUP
	c.Data["type_id_name"] = "无"
	c.Data["parent_id_name"] = "无"
	c.Data["info"] = mod
	c.Data["_method"] = "post"
	c.Data["title"] = "用户组-添加"
	c.TplName = "admin/member_group/get.html"
}
//保存
// @router /member_group [post]
func (c *MemberGroup)Post() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	mod := model.NewType()
	//参数传递
	if err := url.ParseForm(c.Input(), mod); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	mod.TypeId = service.MEMBER_GROUP
	//创建
	ser := admin.NewTypeService()
	id, err := ser.Create(mod)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		rsp.Success("")
	}
}
//编辑
// @router /member_group/:id [get]
func (c *MemberGroup)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := admin.NewTypeService()
	data, err := ser.Read(int_id)
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
		c.Data["title"] = "用户组-修改"
		c.TplName = "admin/member_group/get.html"
	}
}
//更新
// @router /member_group/:id [put]
func (c *MemberGroup)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	mod := model.NewType()
	if err := url.ParseForm(c.Input(), mod); err != nil {
		rsp.Error(err.Error())
	}
	mod.TypeId = service.MEMBER_GROUP
	//更新
	ser := admin.NewTypeService()
	_, err := ser.Update(int_id, mod)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
//检测名称重复
// @router /member_group/check_name [post]
func (c *MemberGroup)CheckName() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	//创建
	ser := admin.NewTypeService()
	ok, err := ser.CheckNameTypeId(service.MEMBER_GROUP, name, id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		rsp.Success("")
	}
}
//删除
// @router /member_group/:id [delete]
func (c *MemberGroup)Delete() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := admin.NewTypeService()
	_, err := ser.DeleteAndTypeId(int_id,service.MEMBER_GROUP)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
