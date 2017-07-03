package admin

import (
	"fmt"
	"blog/fox"
	"time"
	"blog/fox/datetime"
	"blog/model"
	"blog/fox/db"
)
//类别，属性
type Type struct {

}
//快速初始化
func NewTypeService() *Type {
	return new(Type)
}
//获取 指定id 数据列表
func (c *Type)Query(type_id int) (*db.Paginator, error) {
	fields := []string{}
	query := make(map[string]interface{})
	query["type_id"] = type_id
	mod := model.NewType()
	data, err := mod.GetAll(query, fields, "id asc", 0, 9999)
	//fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//创建
func (c *Type)Create(m *model.Type) (int, error) {

	fmt.Println("DATA:", m)
	if len(m.Name) < 1 {
		return 0,fox.NewError("名称 不能为空")
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	//删除状态
	if m.IsDel < 0 {
		m.IsDel = 0
	}
	if m.IsDel > 1 {
		m.IsDel = 1
	}
	//默认
	if m.IsDefault < 0 {
		m.IsDel = 0
	}
	if m.IsDefault > 1 {
		m.IsDel = 1
	}
	o := db.NewDb()
	affected, err := o.Insert(m)
	if err != nil {
		return 0,fox.NewError("创建错误：" + err.Error())
	}
	fmt.Println("affected:", affected)
	fmt.Println("Id:", m.Id)
	return m.Id, nil
}
//更新
func (c *Type)Update(id int, m *model.Type) (int, error) {
	if id < 1 {
		return 0,fox.NewError("ID 错误")
	}
	if id < 10000 {
		return 0,fox.NewError("系统数据 禁止修改")
	}
	_, err := m.GetById(id)
	if err != nil {
		return 0,fox.NewError("数据不存在")
	}
	if len(m.Name) < 1 {
		return 0,fox.NewError("标题 不能为空")
	}
	//fmt.Println("DATA:", m)
	//删除状态
	if m.IsDel < 0 {
		m.IsDel = 0
	}
	if m.IsDel > 1 {
		m.IsDel = 1
	}
	//默认
	if m.IsDefault < 0 {
		m.IsDel = 0
	}
	if m.IsDefault > 1 {
		m.IsDel = 1
	}
	o := db.NewDb()
	m.Id = id
	num, err := o.Id(id).Where("id>10000").Update(m)
	if err != nil {
		return 0,fox.NewError("更新错误：" + err.Error())
	}
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	fmt.Println("nums:", num)
	return id, nil
}
//更新
func (c *Type)UpdateById(m *model.Type, cols ...interface{}) (num int64, err error) {
	o := db.NewDb()
	o.Where("id>10000")
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}
//删除
func (c *Type)Delete(id int) (bool, error) {
	if id < 1 {
		return false,fox.NewError("ID 错误")
	}
	if id < 10000 {
		return false,fox.NewError("系统数据 禁止修改")
	}
	mod := model.NewType()
	num, err := mod.Delete(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num:", num)
	return true, nil
}//删除
func (c *Type)DeleteAndTypeId(id, type_id int) (bool, error) {
	if id < 1 {
		return false,fox.NewError("ID 错误")
	}
	if id < 10000 {
		return false,fox.NewError("系统数据 禁止修改")
	}
	if type_id < 1 {
		return false,fox.NewError("类型 错误")
	}
	mod := model.NewType()
	o := db.NewDb()
	num, err := o.Id(id).Where("type_id=?", type_id).Delete(mod)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("num:", num)
	return true, nil
}
//详情
func (c *Type)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil,fox.NewError("ID 错误")
	}
	mod := model.NewType()
	data, err := mod.GetById(id)
	if err != nil {
		return nil,fox.NewError("数据不存在")
	}
	//整合
	m := make(map[string]interface{})
	m["info"] = *data
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)
	//类别
	m["type_id_name"] = "无"
	if data.TypeId > 0 {
		info, err := mod.GetById(data.TypeId)
		if err == nil {
			m["type_id_name"] = info.Name
		}
	}
	//上级
	m["parent_id_name"] = "无"
	if data.ParentId > 0 {

		info2, err := mod.GetById(data.ParentId)
		if err == nil {
			m["parent_id_name"] = info2.Name
		}
	}
	//fmt.Println(m)
	return m, err
}
//检测名称是否存在
func (c *Type)CheckNameTypeId(type_id int, str string, id int) (bool, error) {
	if str == "" {
		return false,fox.NewError("名称 不能为空")
	}
	mod := model.NewType()
	where := make(map[string]interface{})
	where["type_id=?"] = type_id
	where["name=?"] = str
	if id > 0 {
		where["id !=?"] = id
	}
	//总数
	count, err := db.Filter(where).Count(mod)
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
