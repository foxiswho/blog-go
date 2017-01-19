package admin

import (
	"strconv"
	"fmt"
	"blog/model"
	"blog/fox/url"
	"blog/service/member"
)
//用户
type Member struct {
	Base
}

func (c *Member) URLMapping() {
	c.Mapping("CheckTitle", c.CheckTitle)
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Detail", c.Detail)
}
//检测名称重复
// @router /member/check_title [post]
func (c *Member)CheckTitle() {
	//ID 获取 格式化
	id, _ := c.GetInt("id")
	name := c.GetString("wd")
	//创建
	serv := member.NewMemberService()
	ok, err := serv.CheckUserNameById(name, id)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		c.Success("操作成功")
	}
}
//列表
// @router /member [get]
func (c *Member)List() {
	//查询
	where := make(map[string]interface{})
	mod := member.NewMemberService()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "uid desc", page, 20)
	if err != nil {
		println(err)
	}
	c.Data["data"] = data
	c.Data["title"] = "用户-列表"
	c.TplName = "admin/member/list.html"
}
//编辑
// @router /member/:id [get]
func (c *Member)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser := member.NewMemberService()
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Data["info"] = data["info"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "用户-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.TplName = "admin/member/get.html"
	}
}
//添加
// @router /member/add [get]
func (c *Member)Add() {
	mod := member.NewMemberService()
	mod.Member = &model.Member{}
	mod.MemberStatus = &model.MemberStatus{}
	c.Data["info"] = mod
	c.Data["_method"] = "post"
	c.Data["title"] = "用户-添加"
	c.TplName = "admin/member/get.html"
}
//保存
// @router /member [post]
func (c *Member)Post() {
	mod := model.NewMember()
	//参数传递
	modExt := model.NewMemberStatus()
	if err := url.ParseForm(c.Input(), mod); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	if err := url.ParseForm(c.Input(), modExt); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//更新人
	//modExt.AidUpdate = c.Session.Aid
	//mod.Ip=c.Ctx.Input.IP()
	//创建
	serv := member.NewMemberService()
	id, err := serv.Create(mod, modExt)
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		c.Success("操作成功")
	}
}
//查看
// @router /member/detail/:id [get]
func (c *Member)Detail() {
	c.Get()
	c.Data["title"] = "用户-查看"
	c.TplName = "admin/member/detail.html"
}
//更新
// @router /member/:id [put]
func (c *Member)Put() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	mod := model.NewMember()
	modExt := model.NewMemberStatus()
	if err := url.ParseForm(c.Input(), mod); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	if err := url.ParseForm(c.Input(), modExt); err != nil {
		fmt.Println("ParseForm-err:", err)
		c.Error(err.Error())
		return
	}
	//更新
	ser := member.NewMemberService()
	_, err := ser.Update(int_id, mod, modExt)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
//删除
// @router /member/:id [delete]
func (c *Member)Delete() {
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	ser := member.NewMemberService()
	_, err := ser.Delete(int_id)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}


