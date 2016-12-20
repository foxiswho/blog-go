package service

import (
	"fox/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"fox/util"
	"github.com/astaxie/beego"
	"fmt"
	UtilAuth "fox/util/Auth"
)

type AdminUser  struct {

}
//登录验证
func (this *AdminUser) Auth(account, password string) (admUser *models.Admin, err error) {
	if len(account) == 0 {
		return nil, &util.Error{Msg:"账号 不能为空"}
	}
	if len(password) == 0 {
		return nil, &util.Error{Msg:"密码 不能为空"}
	}
	admUser, err = this.GetAdminByUserName(account)
	if err == nil {
		if admUser.Id!=0 {
			password := UtilAuth.PasswordSalt(password, admUser.Salt)
			if !strings.EqualFold(password, admUser.Password) {
				return nil, &util.Error{Msg:"密码 错误"}
			}
			return admUser, nil
		} else {
			return nil, &util.Error{Msg:"账号 不存在"}
		}
	} else if err == orm.ErrNoRows {
		return nil, &util.Error{Msg:"账号 不存在"}
	}
	beego.Info(err)
	return nil, &util.Error{Msg:"登陆失败，请稍后重试.."}
}
//根据用户名查找
func (this *AdminUser) GetAdminByUserName(account string) (v *models.Admin, err error) {
	o := orm.NewOrm()
	v = &models.Admin{Username: account}

	if err = o.Read(v,"Username"); err == nil {
		return v, nil
	}
	fmt.Println(err)
	return nil, err
}