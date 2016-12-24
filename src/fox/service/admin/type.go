package admin

import (
	"fmt"
	"fox/models"
	"strconv"
	"fox/util"
	"time"
	"github.com/astaxie/beego/orm"
	"fox/util/datetime"
)

type Type struct {

}

func (this *Type)Query(type_id int) (data []interface{}, err error) {
	var fields []string
	query := make(map[string]string)
	sortby := []string{"Id"}
	order := []string{"desc"}
	var offset int64
	var limit int64
	offset = 0
	limit = 2000
	query["type_id"] = strconv.Itoa(type_id)
	data, err = models.GetAllType(query, fields, sortby, order, offset, limit)
	//fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//创建
func (c *Type)Create(m *models.Type) (int64, error) {

	fmt.Println("DATA:", m)
	if len(m.Name) < 1 {
		return 0, &util.Error{Msg:"名称 不能为空"}
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
	id, err := models.AddType(m)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	fmt.Println("Id:", id)
	return id, nil
}
//更新
func (c *Type)Update(id int, m *models.Type) (int, error) {
	if id < 1 {
		return 0, &util.Error{Msg:"ID 错误"}
	}
	if id < 10000 {
		return 0, &util.Error{Msg:"系统数据 禁止修改"}
	}
	_, err := models.GetTypeById(id)
	if err != nil {
		return 0, &util.Error{Msg:"数据不存在"}
	}
	if len(m.Name) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	fmt.Println("DATA:", m)
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

	m.Id = id
	num, err := c.UpdateById(m, "name", "code", "mark", "type_id", "parent_id", "value", "is_del", "sort", "module", "is_default", "setting", "is_child", "is_system", "is_show")
	if err != nil {
		return 0, &util.Error{Msg:"更新错误：" + err.Error()}
	}
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	fmt.Println("nums:", num)
	return id, nil
}
//更新
func (c *Type)UpdateById(m *models.Type, cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	if num, err = o.Update(m, cols...); err == nil {
		fmt.Println("Number of records updated in database:", num)
		return num, nil
	}
	return 0, err
}
//删除
func (c *Type)Delete(id int) (bool, error) {
	if id < 1 {
		return false, &util.Error{Msg:"ID 错误"}
	}
	if id < 10000 {
		return false, &util.Error{Msg:"系统数据 禁止修改"}
	}
	err := models.DeleteType(id)
	if err != nil {
		fmt.Println("err:", err)
	}
	return true, nil
}
//详情
func (c *Type)Read(id int) (map[string]interface{}, error) {
	if id < 1 {
		return nil, &util.Error{Msg:"ID 错误"}
	}

	data, err := models.GetTypeById(id)
	if err != nil {
		return nil, &util.Error{Msg:"数据不存在"}
	}
	//整合
	m := make(map[string]interface{})
	m["info"] = *data
	//时间转换
	m["TimeAdd"] = datetime.Format(data.TimeAdd, datetime.Y_M_D_H_I_S)
	//类别
	m["type_id_name"] = "无"
	if data.TypeId > 0 {

		info, err := models.GetTypeById(data.TypeId)
		if err == nil {
			m["type_id_name"] = info.Name
		}
	}
	//上级
	m["parent_id_name"] = "无"
	if data.ParentId > 0 {

		info2, err := models.GetTypeById(data.ParentId)
		if err == nil {
			m["parent_id_name"] = info2.Name
		}
	}
	//fmt.Println(m)
	return m, err
}