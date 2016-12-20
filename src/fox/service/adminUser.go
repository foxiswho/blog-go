package service

import (
	"fox/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"fox/util"
)

type AdminUser  struct {

}

func (this *AdminUser) Auth(account, password string) (admUser *models.Admin, err error) {
	if len(account) == 0 {
		return nil, &util.Error{Msg:"账号 不能为空"}
	}
	if len(password) == 0 {
		return nil, &util.Error{Msg:"密码 不能为空"}
	}
	admUser ,err = this.GetAdminByUserName(account)
	if err ==nil{
		if admUser.Id >0 {
			if !strings.EqualFold(password, admUser.Password) {
				return nil, &util.Error{Msg:"密码 错误"}
			}
			return admUser,nil
		}
	}
	return nil, &util.Error{Msg:"登陆失败，请稍后重试.."}
}

func (this *AdminUser) GetAdminByUserName(account string) (v *models.Admin, err error) {
	o := orm.NewOrm()
	v = &models.Admin{Username: account}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

