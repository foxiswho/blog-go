package admin

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
func (this *AdminUser) Auth(account, password string) (*AdminSession, error) {
	//登录验证
	admUser, err := this.Login(account, password)
	//登录成功
	if err == nil {
		//转换为session
		var AdminSession *AdminSession
		Session := AdminSession.Convert(admUser)
		return Session, nil
	}
	return nil, err
}
//登录验证
func (this *AdminUser) Login(account, password string) (admUser *models.Admin, err error) {
	if len(account) == 0 {
		return nil, &util.Error{Msg:"账号 不能为空"}
	}
	if len(password) == 0 {
		return nil, &util.Error{Msg:"密码 不能为空"}
	}
	admUser, err = this.GetAdminByUserName(account)
	if err == nil {
		if admUser.Id != 0 {
			password := UtilAuth.PasswordSalt(password, admUser.Salt)
			fmt.Println(password)
			if !strings.EqualFold(password, admUser.Password) {
				return nil, &util.Error{Msg:"密码 错误"}
			}
			return admUser, nil
		} else {
			return nil, &util.Error{Msg:"账号 不存在"}
		}
	} else if err == orm.ErrNoRows {
		return nil, &util.Error{Msg:"账号 不存在."}
	}
	beego.Info(err)
	return nil, &util.Error{Msg:"登陆失败，请稍后重试.."}
}
//根据用户名查找
func (this *AdminUser) GetAdminByUserName(account string) (v *models.Admin, err error) {
	o := orm.NewOrm()
	v = &models.Admin{Username: account}

	if err = o.Read(v, "Username"); err == nil {
		return v, nil
	}
	fmt.Println(err)
	return nil, err
}
//根据ID查找
func (this *AdminUser) GetAdminById(id int) (admUser *models.Admin, err error) {
	if id <= 0 {
		return nil, &util.Error{Msg:"id 错误"}
	}
	admUser, err = models.GetAdminById(id)
	if err == nil {
		if admUser.Id != 0 {
			return admUser, nil
		} else {
			return nil, &util.Error{Msg:"账号 不存在"}
		}
	} else if err == orm.ErrNoRows {
		return nil, &util.Error{Msg:"账号 不存在"}
	}
	beego.Info(err)
	return nil, &util.Error{Msg:"获取失败，请稍后重试.."}
}
//密码更新
func (this *AdminUser)UpdatePassword(pwd string, uid int) (bool, error) {
	c := len(pwd)
	if c < 1 {
		return false, &util.Error{Msg:"密码不能为空"}
	}
	if c < 6 {
		return false, &util.Error{Msg:"密码 最小长度为 6 个字符"}
	}
	if c >= 32 {
		return false, &util.Error{Msg:"密码 不能超过32个字符"}
	}
	if uid < 1 {
		return false, &util.Error{Msg:"用户 UID 错误"}
	}
	admUser, err := models.GetAdminById(uid)
	if err == nil {
		if admUser.Id != 0 {
			pwd := UtilAuth.PasswordSalt(pwd, admUser.Salt)
			if strings.EqualFold(pwd, admUser.Password) {
				return false, &util.Error{Msg:"新密码与旧密码相同"}
			}
			admin := &models.Admin{Id: uid, Password:pwd}
			this.UpdateAdminById(admin,"password")
			return true, nil
		} else {
			return false, &util.Error{Msg:"账号 不存在"}
		}
	}
	return false, &util.Error{Msg:"账号 不存在"}
}
//更新
func (this *AdminUser)UpdateAdminById(m *models.Admin, cols ...string) (num int64,err error) {
	o := orm.NewOrm()
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num,nil
	}
	return 0,err
}