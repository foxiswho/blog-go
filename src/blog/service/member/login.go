package member

import "blog/fox"
//登录
type Login struct {
	Username    string
	Password    string
	Ip          string
	AppId       int
	LoginTypeId int
}
//用户名密码
func (c *Login)Login(username, password string) {
	c.Username = username
	c.Password = password
}
//登录方式
func (c *Login)LoginType(id int) {
	c.LoginTypeId = id
}
//登录Ip
func (c *Login)SetIp(ip string) {
	c.Ip = ip
}
//登录应用
func (c *Login)App(id int) {
	c.AppId = id
}
//检测
func (c *Login)check() (*Login,error) {
	if c.Username == "" {
		return nil,fox.NewError("用户名不能为空")
	}
	return nil,fox.NewError("用户名不能为空")
}