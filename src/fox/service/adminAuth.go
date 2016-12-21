package service

import (
	"github.com/astaxie/beego"
	"fox/models"
	"strings"
	"strconv"
)

type AdminAuth struct {

}
//验证Session
func (this *AdminAuth)Validate(account string) (*AdminSession) {
	//查询用户
	var admin *AdminUser
	admUser, err := admin.GetAdminByUserName(account)
	if err != nil || admUser.Id < 0 {
		beego.Debug("Auth 验证错误：", err.Error())
		return nil
	}
	//赋值
	var session *AdminSession
	Session := session.Convert(admUser)
	beego.Debug("Auth 验证通过：", Session)
	return Session
}

/**
token 校验，判断是否登录
*/
func (this *AdminAuth)ValidateToken(token, currentIp string) (*models.Admin) {
	//解密
	//aes :=crypt.Aes{}
	//Dtoken, err := aes.Decrypt(token)
	//if err != nil {
	//	beego.Debug("token 解密失败")
	//	return nil
	//}
	Dtoken := "sdfs|sdfa|asdfa"
	array := strings.Split(Dtoken, "|")
	if len(array) != 3 {
		beego.Debug("token 校验失败")
		return nil
	}
	uid := array[0]
	ip := array[2]
	//IP发生变化 强制重新登录
	if !strings.EqualFold(ip, currentIp) {
		beego.Debug("IP发生变化 强制重新登录")
		return nil
	}
	int_id, _ := strconv.Atoi(uid)
	//查询用户
	var admin *AdminUser
	admUser, err := admin.GetAdminById(int_id)
	if err != nil || admUser.Id < 0 {
		beego.Debug("Auth 验证错误：", err.Error())
		return nil
	}
	return admUser
}
/**
校验角色权限
*/
func (this *AdminAuth) ValidateRole() (bool, error) {
	//TODO 待开发
	return true, nil
}
/**
校验 菜单权限
*/
func (this *AdminAuth) ValidateMenu() (bool, error) {
	//TODO 待开发
	return true, nil
}