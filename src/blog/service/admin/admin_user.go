package admin

import (
	"strings"
	"blog/fox"
	"fmt"
	UtilAuth "blog/fox/auth"
	"blog/model"
	"blog/fox/db"
	"blog/fox/datetime"
	"time"
	"blog/fox/str"
)
//管理员
type AdminUser  struct {
	*model.Admin
	*model.AdminStatus
}
//快速初始化好管理员
func NewAdminUserService() *AdminUser {
	return new(AdminUser)
}
//登录验证
func (c *AdminUser) Auth(account, password string) (*AdminSession, error) {
	//登录验证
	admUser, err := c.Login(account, password)
	//登录成功
	if err == nil {
		//转换为session
		AdminSession := NewAdminSessionService()
		Session := AdminSession.Convert(admUser)
		return Session, nil
	}
	return nil, err
}
//登录验证
func (c *AdminUser) Login(account, password string) (admUser *model.Admin, err error) {
	if len(account) == 0 {
		return nil, fox.NewError("账号 不能为空")
	}
	if len(password) == 0 {
		return nil, fox.NewError("密码 不能为空")
	}
	admUser, err = c.GetAdminByUserName(account)
	if err == nil {
		if admUser.Aid != 0 {
			password := UtilAuth.PasswordSalt(password, admUser.Salt)
			//fmt.Println(password)
			//fmt.Println(admUser.Password)
			if !strings.EqualFold(password, admUser.Password) {
				return nil, fox.NewError("密码 错误")
			}
			return admUser, nil
		} else {
			return nil, fox.NewError("账号 不存在")
		}
	}
	fmt.Println("登陆 err", err)
	return nil, fox.NewError("登陆失败，请稍后重试..")
}
//根据用户名查找
func (c *AdminUser) GetAdminByUserName(account string) (*model.Admin, error) {
	mod := model.NewAdmin()
	//mod.Username = account
	o := db.NewDb()
	ok, err := o.Where("username=?", account).Get(mod)
	if err == nil && ok {
		if mod.Aid == 0 {
			return nil, fox.NewError("用户 不存在")
		}
		return mod, nil
	}
	fmt.Println(err)
	return nil, err
}
//根据ID查找
func (c *AdminUser) GetAdminById(id int) (*model.Admin, error) {
	if id <= 0 {
		return nil, fox.NewError("id 错误")
	}
	mod := model.NewAdmin()
	admUser, err := mod.GetById(id)
	if err == nil {
		if admUser.Aid != 0 {
			return admUser, nil
		} else {
			return nil, fox.NewError("账号 不存在")
		}
	}
	fmt.Println("获取 err", err)
	return nil, fox.NewError("获取失败，请稍后重试..")
}
//密码更新
func (c *AdminUser)UpdatePassword(pwd string, uid int) (bool, error) {
	t := len(pwd)
	if t < 1 {
		return false, fox.NewError("密码不能为空")
	}
	if t < 6 {
		return false, fox.NewError("密码 最小长度为 6 个字符")
	}
	if t >= 32 {
		return false, fox.NewError("密码 不能超过32个字符")
	}
	if uid < 1 {
		return false, fox.NewError("用户 UID 错误")
	}
	mod := model.NewAdmin()
	admUser, err := mod.GetById(uid)
	if err == nil {
		if admUser.Aid != 0 {
			pwd := UtilAuth.PasswordSalt(pwd, admUser.Salt)
			if strings.EqualFold(pwd, admUser.Password) {
				return false, fox.NewError("新密码与旧密码相同")
			}
			mod.Aid = uid
			mod.Password = pwd
			o := db.NewDb()
			num, err := o.Id(mod.Aid).Update(mod)
			if err != nil {
				return false, err
			}
			fmt.Println("更新条数", num)
			return true, nil
		} else {
			return false, fox.NewError("账号 不存在")
		}
	}
	return false, fox.NewError("账号 不存在")
}
//列表
func (c *AdminUser)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewAdmin()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	ids := make([]int, data.TotalCount)
	for i, x := range data.Data {
		r := x.(model.Admin)
		ids[i] = r.Aid
	}
	//fmt.Println(ids)
	stat := make([]model.AdminStatus, 0)
	o := db.NewDb()
	err = o.In("aid", ids).Find(&stat)
	if err != nil {
		stat = []model.AdminStatus{}
		fmt.Println(err)
	}
	for i, x := range data.Data {
		row := &AdminUser{}
		tmp := x.(model.Admin)
		row.Admin = &tmp
		row.AdminStatus = &model.AdminStatus{}
		for _, v := range stat {
			//fmt.Println(v)
			if (v.Aid == tmp.Aid) {
				row.LoginTime = v.LoginTime
				row.LoginIp = v.LoginIp
				row.AdminStatus.Login = v.Login
				row.AidAdd = v.AidAdd
				row.AidUpdate = v.AidUpdate
				row.AdminStatus.TimeUpdate = v.TimeUpdate
				//fmt.Println(">>>>",row.AdminStatus)
			}
		}
		//fmt.Println("===",row.AdminStatus)
		data.Data[i] = &row
	}

	return data, nil
}
//根据ID获取检测用户名是否存在
func (c *AdminUser)CheckUserNameById(str string, id int) (bool, error) {
	if str == "" {
		return false, fox.NewError("名称 不能为空")
	}
	mode := new(model.Admin)
	where := make(map[string]interface{})
	where["username"] = str
	if id > 0 {
		where["aid!=?"] = id
	}
	count, err := db.Filter(where).Count(mode)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(count)
	if count == 0 {
		return true, nil
	}
	return false, fox.NewError("已存在")
}
//详情
func (c *AdminUser)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil, fox.NewError("ID 错误")
	}
	mode := model.NewAdmin()
	data, err := mode.GetById(id)
	if err != nil {
		fmt.Println(err)
		return nil, fox.NewError("数据不存在")
	}
	//整合
	info := NewAdminUserService()
	info.Admin = data
	//整合
	m := make(map[string]interface{})
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)
	//主键ID值和blog_id值一样所以这里直接取值
	Stat := model.NewAdminStatus()
	StatData, err := Stat.GetById(id)
	if err == nil {
		info.AdminStatus = StatData
	} else {
		//错误屏蔽
		err = nil
		//初始化赋值
		info.AdminStatus = Stat
	}
	m["info"] = info
	//fmt.Println(m)
	return m, err
}
//创建
func (c *AdminUser)Create(m *model.Admin, stat *model.AdminStatus) (int, error) {
	fmt.Println("DATA:", m)
	if len(m.Username) < 1 {
		return 0, fox.NewError("用户名 不能为空")
	}
	t := len(m.Password)
	if t < 1 {
		return 0, fox.NewError("密码 不能为空")
	}
	if t < 6 {
		return 0, fox.NewError("密码 最小长度为 6 个字符")
	}
	if t >= 32 {
		return 0, fox.NewError("密码 不能超过32个字符")
	}
	if len(m.Mail) > 0 && !UtilAuth.CheckMail(m.Mail) {
		return 0, fox.NewError("邮箱 格式不正确")
	}
	if len(m.Mobile) > 0 && !UtilAuth.CheckMail(m.Mobile) {
		return 0, fox.NewError("手机号 格式不正确")
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	m.TimeUpdate = time.Now()
	//干扰码生成
	m.Salt = str.RandSalt()
	//加密后密码
	m.Password = UtilAuth.PasswordSalt(m.Password, m.Salt)
	o := db.NewDb()
	affected, err := o.Insert(m)
	if err != nil {
		return 0, fox.NewError("创建错误1：" + err.Error())
	}
	stat.Aid = m.Aid
	stat.Aid = stat.Aid
	id2, err := o.Insert(stat)
	if err != nil {
		return 0, fox.NewError("创建错误2：" + err.Error())
	}
	fmt.Println("DATA:", m)
	fmt.Println("affected:", affected)
	fmt.Println("Id:", m.Aid)
	fmt.Println("Statistics:", id2)
	return m.Aid, nil
}
//更新
func (c *AdminUser)Update(id int, m *model.Admin, stat *model.AdminStatus) (int, error) {
	if id < 1 {
		return 0, fox.NewError("ID 错误")
	}
	mode := model.NewAdmin()
	_, err := mode.GetById(id)
	if err != nil {
		return 0, fox.NewError("数据不存在")
	}
	if len(m.Username) < 1 {
		return 0, fox.NewError("用户名 不能为空")
	}
	if m.Password != "" {
		t := len(m.Password)
		if t < 1 {
			return 0, fox.NewError("密码 不能为空")
		}
		if t < 6 {
			return 0, fox.NewError("密码 最小长度为 6 个字符")
		}
		if t >= 32 {
			return 0, fox.NewError("密码 不能超过32个字符")
		}
		//干扰码生成
		m.Salt = str.RandSalt()
		//加密后密码
		m.Password = UtilAuth.PasswordSalt(m.Password, m.Salt)
	}
	if len(m.Mail) > 0 && !UtilAuth.CheckMail(m.Mail) {
		return 0, fox.NewError("邮箱 格式不正确")
	}
	if len(m.Mobile) > 0 && !UtilAuth.CheckMail(m.Mobile) {
		return 0, fox.NewError("手机号 格式不正确")
	}
	fmt.Println("DATA:", m)
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	o := db.NewDb()
	num, err := o.Id(id).Update(m)
	if err != nil {
		return 0, fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println("============", num)
	//
	stat.StatusId = id
	o = db.NewDb()
	num2, err := o.Id(id).Update(stat)
	if err != nil {
		return 0, fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println(num2)
	//fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//删除
func (c *AdminUser)Delete(id int) (bool, error) {
	if id < 1 {
		return false, fox.NewError("ID 错误")
	}
	if id == 1 {
		return false, fox.NewError("超级管理员 禁止删除")
	}
	mode := model.NewAdmin()
	num, err := mode.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num:", num)
	num2, err := c.AdminStatus.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num2:", num2)
	return true, nil
}