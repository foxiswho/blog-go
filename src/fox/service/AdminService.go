package service

import (
	"fox/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"fox/util/Error"
)

type adminService struct {}


/**
登陆鉴权
*/
func (this *adminService) Authentication(account, pwd string) (adminUser *models.Admin, err error) {
	if len(account) < 1 {
		return nil, &Error.Error{"账号 不能为空"}
	}
	if len(pwd) < 1 {
		return nil, &Error.Error{"密码 不能为空"}
	}
	adminUser, err = this.GetAdminByAccount(account)
	if err != nil {
		return nil, &Error.Error{"登陆失败，请稍后重试"}
	}
	if !strings.EqualFold(pwd, adminUser.Password) {
		return nil, &Error.Error{"密码错误"}
	}
	return adminUser, nil
}

func (this *adminService) GetAdminByAccount(account string) (v *models.Admin, err error) {
	o := orm.NewOrm()
	v = &models.Admin{Username:account}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}