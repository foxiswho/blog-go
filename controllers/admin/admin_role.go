package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/model"
	"github.com/foxiswho/blog-go/service/admin"
	"github.com/foxiswho/blog-go/service/conf"
	"strconv"
)
//角色控制器
type AdminRole struct {
	Base
}
//路由自动
func (c *AdminRole) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /admin_role [get]
func (c *AdminRole)List() {
	//初始化
	ser := admin.NewTypeService()
	//查询
	data, err := ser.Query(conf.ADMIN_ROLE)
	//错误检测
	if err != nil {
		fmt.Println(err)
		c.Error(err.Error())
		return
	}
	//变量赋值
	c.Data["data"] = data
	c.Data["title"] = "角色-列表"
	//模版
	c.TplName = "admin/admin_role/list.html"
}
//添加
// @router /admin_role/add [get]
func (c *AdminRole)Add() {
	//结构体初始化，并给予初始值
	mod := model.NewType()
	mod.TypeId = conf.ADMIN_ROLE
	mod.IsDefault = 0
	mod.IsDel = 0
	mod.IsSystem = 0
	mod.IsShow = 1
	mod.IsChild = 0
	//模版赋值
	c.Data["type_id"] = conf.ADMIN_ROLE
	c.Data["type_id_name"] = "无"
	c.Data["parent_id_name"] = "无"
	c.Data["info"] = mod
	c.Data["_method"] = "post"
	c.Data["title"] = "角色-添加"
	//模版
	c.TplName = "admin/admin_role/get.html"
}
//保存
// @router /admin_role [post]
func (c *AdminRole)Post() {
	//初始化
	mod := model.Type{}
	//参数传递，表单值 自动保存到结构体对应的属性中,根据 tag 中的form
	if err := c.ParseForm(&mod); err != nil {
		//错误显示
		c.Error(err.Error())
		return
	}
	mod.TypeId = conf.ADMIN_ROLE
	//创建，初始化
	ser := admin.NewTypeService()
	//保存到数据库创建记录
	id, err := ser.Create(&mod)
	//错误检测
	if err != nil {
		c.Error(err.Error())
		return
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//编辑
// @router /admin_role/:id [get]
func (c *AdminRole)Get() {
	//获取ID 和字符串ID转为数值型
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//初始化
	ser := admin.NewTypeService()
	//获取该ID记录
	data, err := ser.Read(int_id)
	//错误检测和输出
	if err != nil {
		c.Error(err.Error())
	} else {
		//模版变量赋值
		c.Data["info"] = data["info"]
		c.Data["type_id_name"] = data["type_id_name"]
		c.Data["parent_id_name"] = data["parent_id_name"]
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["title"] = "角色-修改"
		//模版
		c.TplName = "admin/admin_role/get.html"
	}
}
//更新
// @router /admin_role/:id [put]
func (c *AdminRole)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//初始化
	mod := model.Type{}
	//参数传递，表单值 自动保存到结构体对应的属性中,根据 tag 中的form
	if err := c.ParseForm(&mod); err != nil {
		//错误输出
		c.Error(err.Error())
		return
	}
	mod.TypeId = conf.ADMIN_ROLE
	//初始化
	ser := admin.NewTypeService()
	//更新
	_, err := ser.Update(int_id, &mod)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//检测名称重复
// @router /admin_role/check_name [post]
func (c *AdminRole)CheckName() {
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	//初始化
	ser := admin.NewTypeService()
	//根据ID检测是否重复
	ok, err := ser.CheckNameTypeId(conf.ADMIN_ROLE, name, id)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
//删除
// @router /admin_role/:id [delete]
func (c *AdminRole)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//初始化
	ser := admin.NewTypeService()
	//删除该记录
	_, err := ser.DeleteAndTypeId(int_id, conf.ADMIN_ROLE)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
