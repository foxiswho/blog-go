package member

import (
	"blog/model"
	"time"
	"blog/fox/db"
	"blog/fox"
	"fmt"
	"blog/fox/str"
	"strings"
	"blog/fox/datetime"
	UtilAuth "blog/fox/auth"
)
//用户
type Member struct {
	*model.Member
	*model.MemberProfile
	*model.MemberStatus
}
//快速初始化
func NewMemberService() *Member {
	return new(Member)
}
//根据用户名查找
func (c *Member) GetByUserName(account string) (*model.Member, error) {
	mod := model.NewMember()
	//mod.Username = account
	o := db.NewDb()
	ok, err := o.Where("username=?", account).Get(mod)
	if err == nil && ok {
		if mod.Uid == 0 {
			return nil,fox.NewError("用户不存在")
		}
		return mod, nil
	}
	fmt.Println(err)
	return nil, err
}
//根据ID查找
func (c *Member) GetById(id int) (*model.Member, error) {
	if id <= 0 {
		return nil,fox.NewError("id 错误")
	}
	mod := model.NewMember()
	user, err := mod.GetById(id)
	if err == nil {
		if user.Uid != 0 {
			return user, nil
		} else {
			return nil,fox.NewError("账号 不存在")
		}
	}
	fmt.Println("获取错误", err)
	return nil,fox.NewError("获取失败，请稍后重试..")
}
//密码更新
func (c *Member)UpdatePassword(pwd string, uid int) (bool, error) {
	t := len(pwd)
	if t < 1 {
		return false,fox.NewError("密码不能为空")
	}
	if t < 6 {
		return false,fox.NewError("密码 最小长度为 6 个字符")
	}
	if t >= 32 {
		return false,fox.NewError("密码 不能超过32个字符")
	}
	if uid < 1 {
		return false,fox.NewError("用户 UID 错误")
	}
	mod := model.NewMember()
	user, err := mod.GetById(uid)
	if err == nil {
		if user.Uid != 0 {
			pwd := UtilAuth.PasswordSalt(pwd, user.Salt)
			if strings.EqualFold(pwd, user.Password) {
				return false,fox.NewError("新密码与旧密码相同")
			}
			mod.Uid = uid
			mod.Password = pwd
			c.UpdateById(mod, "password")
			return true, nil
		} else {
			return false,fox.NewError("账号 不存在")
		}
	}
	return false,fox.NewError("账号 不存在")
}
//更新
func (c *Member)UpdateById(m *model.Member, cols ...interface{}) (num int64, err error) {
	o := db.NewDb()
	if num, err = o.Id(m.Uid).Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}
func (c *Member)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewMember()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	ids := make([]int, data.TotalCount)
	for i, x := range data.Data {
		r := x.(model.Member)
		ids[i] = r.Uid
	}
	//fmt.Println(ids)
	stat := make([]model.MemberStatus, 0)
	o := db.NewDb()
	err = o.In("uid", ids).Find(&stat)
	if err != nil {
		stat = []model.MemberStatus{}
		fmt.Println(err)
	}
	for i, x := range data.Data {
		row := &Member{}
		tmp := x.(model.Member)
		row.Member = &tmp
		row.MemberStatus = &model.MemberStatus{}
		for _, v := range stat {
			//fmt.Println(v)
			if (v.Uid == tmp.Uid) {
				row.LastLoginAppId = v.LastLoginAppId
				row.LastLoginIp = v.LastLoginIp
				row.MemberStatus.Login = v.Login
				row.AidAdd = v.AidAdd
				row.MemberStatus.RegTime = v.RegTime
				row.MemberStatus.RegIp = v.RegIp
				row.RegType = v.RegType
				row.RegAppId = v.RegAppId
			}
		}
		//fmt.Println("===",row.AdminStatus)
		data.Data[i] = &row
	}

	return data, nil
}
//详情
func (c *Member)CheckUserNameById(str string, id int) (bool, error) {
	if str == "" {
		return false,fox.NewError("名称 不能为空")
	}
	mode := new(model.Member)
	where := make(map[string]interface{})
	where["username"] = str
	if id > 0 {
		where["uid!=?"] = id
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
	return false,fox.NewError("已存在")
}
//详情
func (c *Member)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil,fox.NewError("ID 错误")
	}
	mode := model.NewMember()
	data, err := mode.GetById(id)
	if err != nil {
		fmt.Println(err)
		return nil,fox.NewError("数据不存在")
	}
	//整合
	info := NewMemberService()
	info.Member = data
	//整合
	m := make(map[string]interface{})
	//时间转换
	m["TimeAdd"] = datetime.Format(data.RegTime, datetime.Y_M_D_H_I_S)
	//主键ID值和blog_id值一样所以这里直接取值
	Stat := model.NewMemberStatus()
	StatData, err := Stat.GetById(id)
	if err == nil {
		info.MemberStatus = StatData
	} else {
		//错误屏蔽
		err = nil
		//初始化赋值
		info.MemberStatus = Stat
	}
	m["info"] = info
	//fmt.Println(m)
	return m, err
}
//创建
func (c *Member)Create(m *model.Member, stat *model.MemberStatus) (int, error) {
	fmt.Println("DATA:", m)
	if len(m.Username) < 1 {
		return 0,fox.NewError("用户名 不能为空")
	}
	t := len(m.Password)
	if t < 1 {
		return 0,fox.NewError("密码 不能为空")
	}
	if t < 6 {
		return 0,fox.NewError("密码 最小长度为 6 个字符")
	}
	if t >= 32 {
		return 0,fox.NewError("密码 不能超过32个字符")
	}
	if len(m.Mail) > 0 && !UtilAuth.CheckMail(m.Mail) {
		return 0,fox.NewError("邮箱 格式不正确")
	}
	if len(m.Mobile) > 0 && !UtilAuth.CheckMail(m.Mobile) {
		return 0,fox.NewError("手机号 格式不正确")
	}
	//时间
	if m.RegTime.IsZero() {
		m.RegTime = time.Now()
	}
	m.RegTime = time.Now()
	//干扰码生成
	m.Salt = str.RandSalt()
	//加密后密码
	m.Password = UtilAuth.PasswordSalt(m.Password, m.Salt)
	o := db.NewDb()
	affected, err := o.Insert(m)
	if err != nil {
		return 0,fox.NewError("创建错误1：" + err.Error())
	}
	stat.Uid = m.Uid
	stat.StatusId = stat.Uid
	id2, err := o.Insert(stat)
	if err != nil {
		return 0,fox.NewError("创建错误2：" + err.Error())
	}
	fmt.Println("DATA:", m)
	fmt.Println("affected:", affected)
	fmt.Println("Id:", m.Uid)
	fmt.Println("Statistics:", id2)
	return m.Uid, nil
}
//更新
func (c *Member)Update(id int, m *model.Member, stat *model.MemberStatus) (int, error) {
	if id < 1 {
		return 0,fox.NewError("ID 错误")
	}
	mode := model.NewMember()
	_, err := mode.GetById(id)
	if err != nil {
		return 0,fox.NewError("数据不存在")
	}
	if len(m.Username) < 1 {
		return 0,fox.NewError("用户名 不能为空")
	}
	if m.Password != "" {
		t := len(m.Password)
		if t < 1 {
			return 0,fox.NewError("密码 不能为空")
		}
		if t < 6 {
			return 0,fox.NewError("密码 最小长度为 6 个字符")
		}
		if t >= 32 {
			return 0,fox.NewError("密码 不能超过32个字符")
		}
		//干扰码生成
		m.Salt = str.RandSalt()
		//加密后密码
		m.Password = UtilAuth.PasswordSalt(m.Password, m.Salt)
	}
	if len(m.Mail) > 0 && !UtilAuth.CheckMail(m.Mail) {
		return 0,fox.NewError("邮箱 格式不正确")
	}
	if len(m.Mobile) > 0 && !UtilAuth.CheckMail(m.Mobile) {
		return 0,fox.NewError("手机号 格式不正确")
	}
	fmt.Println("DATA:", m)
	//时间
	if m.RegTime.IsZero() {
		m.RegTime = time.Now()
	}
	o := db.NewDb()
	num, err := o.Id(id).Update(m)
	if err != nil {
		return 0,fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println("============", num)
	//
	stat.StatusId = id
	o = db.NewDb()
	num2, err := o.Id(id).Update(stat)
	if err != nil {
		return 0,fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println(num2)
	//fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//删除
func (c *Member)Delete(id int) (bool, error) {
	if id < 1 {
		return false,fox.NewError("ID 错误")
	}
	if id == 1 {
		return false,fox.NewError("超级管理员 禁止删除")
	}
	mode := model.NewMember()
	num, err := mode.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num:", num)
	num2, err := c.MemberStatus.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num2:", num2)
	return true, nil
}