package admin

import (
	"fox/util/Response"
	"strconv"
	"fmt"
	"fox/model"
	"fox/util/url"
	"fox/service/admin"
)

type AdminUser struct {
	Base
}

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
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("wd")
	//创建
	serv := admin.NewAdminUserService()
	ok, err := serv.CheckUserNameById(name, id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		rsp.Success("")
	}
}
//列表
// @router /admin [get]
func (c *AdminUser)List() {
	where := make(map[string]interface{})
	mod := admin.NewAdminUserService()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "aid desc", page, 20)
	if err != nil {
		println(err)
	}
	c.Data["data"] = data
	c.Data["title"] = "管理员-列表"
	c.TplName = "admin/admin_user/list.html"
}
//编辑
// @router /admin/:id [get]
func (c *AdminUser)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := admin.NewAdminUserService()
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(c.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
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
	mod := admin.NewAdminUserService()
	mod.Admin = &model.Admin{}
	mod.AdminStatus = &model.AdminStatus{}
	c.Data["info"] = mod
	c.Data["_method"] = "post"
	c.Data["title"] = "管理员-添加"
	c.TplName = "admin/admin_user/get.html"
}
//保存
// @router /admin [post]
func (c *AdminUser)Post() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	mod := model.NewAdmin()
	//参数传递
	modExt := model.NewAdminStatus()
	if err := url.ParseForm(c.Input(), mod); err != nil {
		fmt.Println("ParseForm-err:", err)
		rsp.Error(err.Error())
	}
	if err := url.ParseForm(c.Input(), modExt); err != nil {
		fmt.Println("ParseForm-err:", err)
		rsp.Error(err.Error())
	}
	//更新人
	modExt.AidUpdate = c.Session.Aid
	mod.Ip=c.Ctx.Input.IP()
	//创建
	serv := admin.NewAdminUserService()
	id, err := serv.Create(mod, modExt)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		rsp.Success("")
	}
}
//查看
// @router /admin/detail/:id [get]
func (c *AdminUser)Detail() {
	c.Get()
	c.Data["title"] = "管理员-查看"
	c.TplName = "admin/admin_user/detail.html"
}
//更新
// @router /admin/:id [put]
func (c *AdminUser)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	mod := model.NewAdmin()
	modExt := model.NewAdminStatus()
	if err := url.ParseForm(c.Input(), mod); err != nil {
		fmt.Println("ParseForm-err:", err)
		rsp.Error(err.Error())
	}
	if err := url.ParseForm(c.Input(), modExt); err != nil {
		fmt.Println("ParseForm-err:", err)
		rsp.Error(err.Error())
	}
	//更新
	ser := admin.NewAdminUserService()
	_, err := ser.Update(int_id, mod, modExt)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
//删除
// @router /admin/:id [delete]
func (c *AdminUser)Delete() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := admin.NewAdminUserService()
	_, err := ser.Delete(int_id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}


