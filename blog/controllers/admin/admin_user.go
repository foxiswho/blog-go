package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/blog/model"
	"github.com/foxiswho/blog-go/blog/service/admin"
	"strconv"
)
//管理员控制器
type AdminUser struct {
	Base
}
//路由自动化
func (c *AdminUser) URLMapping() {
	c.Mapping("CheckTitle", c.CheckTitle)
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Detail", c.Detail)
}
//检测名称重复
// @router /admin/check_title [post]
func (c *AdminUser)CheckTitle() {
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("wd")
	//初始化
	serv := admin.NewAdminUserService()
	//检测是否存在
	ok, err := serv.CheckUserNameById(name, id)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
//列表
// @router /admin [get]
func (c *AdminUser)List() {
	where := make(map[string]interface{})
	//初始化
	mod := admin.NewAdminUserService()
	//页码
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "aid desc", page, 20)
	if err != nil {
		println(err)
	}
	//模版赋值 及 模版文件
	c.Data["data"] = data
	c.Data["title"] = "管理员-列表"
	c.TplName = "admin/admin_user/list.html"
}
//编辑
// @router /admin/:id [get]
func (c *AdminUser)Get() {
	//获取ID 并转换为数值型
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//初始化
	ser := admin.NewAdminUserService()
	//获取该信息
	data, err := ser.Read(int_id)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		//模版赋值 及 模版文件
		c.Data["info"] = data["info"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "管理员-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.TplName = "admin/admin_user/get.html"
	}
}
//添加
// @router /admin/add [get]
func (c *AdminUser)Add() {
	//初始化
	mod := admin.NewAdminUserService()
	//初始化
	mod.Admin = &model.Admin{}
	//初始化
	mod.AdminStatus = &model.AdminStatus{}
	//模版赋值 及 模版文件
	c.Data["info"] = mod
	c.Data["_method"] = "post"
	c.Data["title"] = "管理员-添加"
	c.TplName = "admin/admin_user/get.html"
}
//保存
// @router /admin [post]
func (c *AdminUser)Post() {
	//初始化
	mod := model.NewAdmin()
	//初始化
	modExt := model.NewAdminStatus()
	//参数传递，表单值 自动保存到结构体对应的属性中,根据 tag 中的form
	if err := c.ParseForm(&mod); err != nil {
		//错误输出
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//参数传递，表单值 自动保存到结构体对应的属性中,根据 tag 中的form
	if err := c.ParseForm(&modExt); err != nil {
		//错误输出
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//更新人
	modExt.AidUpdate = c.Session.Aid
	//IP获取
	mod.Ip=c.Ctx.Input.IP()
	//初始化
	serv := admin.NewAdminUserService()
	//保存到数据库
	id, err := serv.Create(mod, modExt)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//查看
// @router /admin/detail/:id [get]
func (c *AdminUser)Detail() {
	//直接使用 编辑方法，此处调用都是一样的
	c.Get()
	//模版赋值 及 模版文件
	c.Data["title"] = "管理员-查看"
	c.TplName = "admin/admin_user/detail.html"
}
//更新
// @router /admin/:id [put]
func (c *AdminUser)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//初始化
	mod := model.NewAdmin()
	modExt := model.NewAdminStatus()
	//参数传递，表单值 自动保存到结构体对应的属性中,根据 tag 中的form
	if err := c.ParseForm(&mod); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//参数传递，表单值 自动保存到结构体对应的属性中,根据 tag 中的form
	if err := c.ParseForm(&modExt); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//初始化
	ser := admin.NewAdminUserService()
	//更新数据
	_, err := ser.Update(int_id, mod, modExt)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//删除
// @router /admin/:id [delete]
func (c *AdminUser)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//初始化
	ser := admin.NewAdminUserService()
	//数据库删除
	_, err := ser.Delete(int_id)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}


