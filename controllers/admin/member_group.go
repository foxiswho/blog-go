package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/controllers"
	"github.com/foxiswho/blog-go/model"
	"github.com/foxiswho/blog-go/service/admin"
	"github.com/foxiswho/blog-go/service/conf"
	"strconv"
)
//用户组
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
	data, err := ser.Query(conf.MEMBER_GROUP)
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
	mod.TypeId = conf.MEMBER_GROUP
	mod.IsDefault = 0
	mod.IsDel = 0
	mod.IsSystem = 0
	mod.IsShow = 1
	mod.IsChild = 0
	c.Data["type_id"] = conf.MEMBER_GROUP
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
	mod := model.Type{}
	//参数传递
	if err := c.ParseForm(&mod); err != nil {
		c.Error(err.Error())
		return
	}
	mod.TypeId = conf.MEMBER_GROUP
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
// @router /member_group/:id [get]
func (c *MemberGroup)Get() {
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
		c.Data["title"] = "用户组-修改"
		c.TplName = "admin/member_group/get.html"
	}
}
//更新
// @router /member_group/:id [put]
func (c *MemberGroup)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	mod := model.Type{}
	if err := c.ParseForm(&mod); err != nil {
		c.Error(err.Error())
		return
	}
	mod.TypeId = conf.MEMBER_GROUP
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
// @router /member_group/check_name [post]
func (c *MemberGroup)CheckName() {
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	//创建
	ser := admin.NewTypeService()
	ok, err := ser.CheckNameTypeId(conf.MEMBER_GROUP, name, id)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
//删除
// @router /member_group/:id [delete]
func (c *MemberGroup)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := admin.NewTypeService()
	_, err := ser.DeleteAndTypeId(int_id, conf.MEMBER_GROUP)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
