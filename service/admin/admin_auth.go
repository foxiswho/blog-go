package admin

import (
	"github.com/foxiswho/blog-go/fox/log"
	"github.com/foxiswho/blog-go/model"
	"strconv"
	"strings"
)
//后台权限验证
type AdminAuth struct {

}

func NewAdminAuthService() *AdminAuth {
	return new(AdminAuth)
}
//验证Session
func (c *AdminAuth)Validate(account string) (*AdminSession) {
	//查询用户
	var admin *AdminUser
	admUser, err := admin.GetAdminByUserName(account)
	if err != nil || admUser.Aid < 0 {
		log.Debug("Auth 验证错误：", err.Error())
		return nil
	}
	//赋值
	session := NewAdminSessionService()
	Session := session.Convert(admUser)
	log.Debug("Auth 验证通过：", Session)
	return Session
}

/**
token 校验，判断是否登录
*/
func (c *AdminAuth)ValidateToken(token, currentIp string) (*model.Admin) {
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
		log.Debug("token 校验失败")
		return nil
	}
	uid := array[0]
	ip := array[2]
	//IP发生变化 强制重新登录
	if !strings.EqualFold(ip, currentIp) {
		log.Debug("IP发生变化 强制重新登录")
		return nil
	}
	int_id, _ := strconv.Atoi(uid)
	//查询用户
	var admin *AdminUser
	admUser, err := admin.GetAdminById(int_id)
	if err != nil || admUser.Aid < 0 {
		log.Debug("Auth 验证错误：", err.Error())
		return nil
	}
	return admUser
}
/**
校验角色权限
*/
func (c *AdminAuth) ValidateRole() (bool, error) {
	//TODO 待开发
	return true, nil
}
/**
校验 菜单权限
*/
func (c *AdminAuth) ValidateMenu() (bool, error) {
	//TODO 待开发
	return true, nil
}
