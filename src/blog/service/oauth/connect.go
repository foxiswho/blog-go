package oauth

import (
	"blog/fox"
	"blog/fox/db"
	"blog/model"
	"fmt"
)
//第三方
type Connect struct {

}
//快速初始化好
func NewConnect() *Connect {
	return new(Connect)
}
//后台绑定 验证
func (t *Connect)Admin(type_id int, val string, is_uid bool) (*model.Connect, error) {
	if len(val) < 1 {
		return nil, fox.NewError("val 值不能为空:")
	}
	if type_id < 1 {
		return nil, fox.NewError("type_id 值错误:")
	}
	con := model.NewConnect()
	fmt.Println(type_id, val)
	maps := make(map[string]interface{})
	maps["type_id"] = type_id
	//不是UID登陆
	if is_uid {
		maps["uid"] = val
	} else {
		maps["open_id"] = val
	}
	_, err := db.Filter(maps).Get(con)
	if err != nil {
		return nil, fox.NewError("查询错误:" + err.Error())
	}
	if con.Uid < 1 {
		return nil, fox.NewError("未绑定")
	}
	fmt.Println("con", con)
	return con, nil
}
//@type_id 应用
//@uid     用户或者管理员ID
//@open_id 应用 open_id
//@token   应用token
//@cat     登录应用
//@type_login 登录模块;302前台还是后台301
//绑定
func (t *Connect)Binding(type_id, uid int, open_id, token string, cat, type_login int) error {
	if type_id < 1 {
		return fox.NewError("绑定 应用错误")
	}
	if uid < 1 {
		return fox.NewError("绑定 用户错误")
	}
	if len(open_id) < 1 {
		return fox.NewError("绑定 open_id错误")
	}
	if type_login < 301 {
		return fox.NewError("绑定 前后台错误")
	}
	if type_login > 302 {
		return fox.NewError("绑定 前后台错误")
	}
	if cat < 1 {
		return fox.NewError("绑定 登陆应用错误")
	}
	con := model.NewConnect()
	maps := make(map[string]interface{})
	maps["type_id"] = type_id
	maps["uid"] = uid
	maps["open_id"] = open_id
	_, err := db.Filter(maps).Get(con)
	if err != nil {
		return fox.NewError("查询错误:" + err.Error())
	}
	if con.ConnectId > 0 {
		return fox.NewError("已存在 禁止重复绑定")
	}
	//重新初始化
	con = model.NewConnect()
	con.OpenId = open_id
	con.Token = token
	con.TypeId = type_id
	con.Uid = uid
	con.Type = cat
	con.TypeLogin = type_login
	//保存数据库
	if _, err := db.NewDb().Insert(con); err != nil {
		return fox.NewError("保存错误:" + err.Error())
	}
	return nil
}