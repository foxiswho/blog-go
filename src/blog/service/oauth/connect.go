package oauth

import (
	"blog/fox"
	"blog/fox/db"
	"blog/model"
	"fmt"
)

type Connect struct {

}

func NewConnect() *Connect {
	return new(Connect)
}
func (t *Connect)Admin(type_id int, val string, is_uid bool) (*model.Connect, error) {
	if len(val) < 1 {
		return nil, &fox.Error{Msg:"val 值不能为空"}
	}
	if type_id < 1 {
		return nil, &fox.Error{Msg:"type_id 值错误"}
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
	_,err := db.Filter(maps).Get(con)
	if err != nil {
		return nil, &fox.Error{Msg:"查询错误 " + err.Error()}
	}
	if con.Uid < 1 {
		return nil, &fox.Error{Msg:"未绑定"}
	}
	fmt.Println("con", con)
	return con, nil
}