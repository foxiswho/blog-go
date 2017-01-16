package admin

import (
)
//博客同步
type BlogSync struct {
	Base
}

func (c *BlogSync) URLMapping() {
	c.Mapping("List", c.List)
}
//列表
// @router /auth_csdn [get]
func (c *BlogSync)List() {

	//web:=csdn.NewAuthorizeWeb()
	//ok,err:=web.SetConfig()
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("status:",ok);
	//web.SetRedirectUri("http://www.foxwho.com:8080/admin/auth_token")
	//
	//c.Data["url"] = web.GetAuthorizeUrl()
	c.TplName = "admin/auth/list.html"
}
// @router /blog_sync/go [get]
func (c *BlogSync)Go() {
	//blog_id,err:=c.GetInt("blog_id")
	//if err!=nil{
	//	c.Error("blog_id 不能为空")
	//}
	//type_id,err:=c.GetInt("type_id")
	//if err!=nil{
	//	c.Error("type_id 不能为空")
	//}
	//id,err:=c.GetInt("id")
	//if err!=nil{
	//	c.Error("id 不能为空")
	//}

	c.TplName = "admin/auth/get_token.html"
}
