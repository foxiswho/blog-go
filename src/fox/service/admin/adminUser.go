package admin

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"fox/util"
	"github.com/astaxie/beego"
	"fmt"
	UtilAuth "fox/util/Auth"
	"fox/model"
	"fox/util/db"
)

type AdminUser  struct {

}
func NewAdminUserService() *AdminUser{
	return new(AdminUser)
}
//登录验证
func (c *AdminUser) Auth(account, password string) (*AdminSession, error) {
	//登录验证
	admUser, err := c.Login(account, password)
	//登录成功
	if err == nil {
		//转换为session
		AdminSession:=NewAdminSessionService()
		Session := AdminSession.Convert(admUser)
		return Session, nil
	}
	return nil, err
}
//登录验证
func (c *AdminUser) Login(account, password string) (admUser *model.Admin, err error) {
	if len(account) == 0 {
		return nil, &util.Error{Msg:"账号 不能为空"}
	}
	if len(password) == 0 {
		return nil, &util.Error{Msg:"密码 不能为空"}
	}
	admUser, err = c.GetAdminByUserName(account)
	if err == nil {
		if admUser.Aid != 0 {
			password := UtilAuth.PasswordSalt(password, admUser.Salt)
			//fmt.Println(password)
			//fmt.Println(admUser.Password)
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
func (c *AdminUser) GetAdminByUserName(account string) (*model.Admin, error) {
	mod := model.NewAdmin()
	//mod.Username = account
	o := db.NewDb()
	ok,err := o.Where("username=?",account).Get(mod)
	if err == nil && ok{
		if mod.Aid ==0{
			return nil,&util.Error{Msg:"用户不存在"}
		}
		return mod, nil
	}
	fmt.Println(err)
	return nil, err
}
//根据ID查找
func (c *AdminUser) GetAdminById(id int) (*model.Admin, error) {
	if id <= 0 {
		return nil, &util.Error{Msg:"id 错误"}
	}
	mod := model.NewAdmin()
	admUser, err := mod.GetById(id)
	if err == nil {
		if admUser.Aid != 0 {
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
func (c *AdminUser)UpdatePassword(pwd string, uid int) (bool, error) {
	t := len(pwd)
	if t < 1 {
		return false, &util.Error{Msg:"密码不能为空"}
	}
	if t < 6 {
		return false, &util.Error{Msg:"密码 最小长度为 6 个字符"}
	}
	if t >= 32 {
		return false, &util.Error{Msg:"密码 不能超过32个字符"}
	}
	if uid < 1 {
		return false, &util.Error{Msg:"用户 UID 错误"}
	}
	mod := model.NewAdmin()
	admUser, err := mod.GetById(uid)
	if err == nil {
		if admUser.Aid != 0 {
			pwd := UtilAuth.PasswordSalt(pwd, admUser.Salt)
			if strings.EqualFold(pwd, admUser.Password) {
				return false, &util.Error{Msg:"新密码与旧密码相同"}
			}
			mod.Aid = uid
			mod.Password = pwd
			c.UpdateAdminById(mod, "password")
			return true, nil
		} else {
			return false, &util.Error{Msg:"账号 不存在"}
		}
	}
	return false, &util.Error{Msg:"账号 不存在"}
}
//更新
func (c *AdminUser)UpdateAdminById(m *model.Admin, cols ...interface{}) (num int64, err error) {
	o := db.NewDb()
	if num, err = o.Id(m.Aid).Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}