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
	o := db.NewDb()
	//UID
	if is_uid {
		err := o.Where("type_id=? and uid=?", type_id, val).Find(con)
		if err != nil {
			return nil, &fox.Error{Msg:"查询错误 "+err.Error()}
		}
		if con.Uid < 1 {
			return nil, &fox.Error{Msg:"uid 值错误,请联系管理员处理"}
		}
	} else {
		//open_id
		err := o.Where("type_id=? and open_id=?", type_id, val).Find(con)
		if err != nil {
			return nil, &fox.Error{Msg:"查询错误 "+err.Error()}
		}
		if con.Uid < 1 {
			return nil, &fox.Error{Msg:"uid 值错误,请联系管理员处理"}
		}
	}
	fmt.Println("con",con)
	return con, nil
}