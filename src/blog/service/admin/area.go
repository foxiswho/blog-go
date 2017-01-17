package admin

import (
	"blog/model"
	"blog/fox/db"
)
//地区
type Area struct {
	*model.Area
	*model.AreaExt
}
//快速初始化地区
func NewAreaService() *Area {
	return new(Area)
}
//列表
func (c *Area)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewArea()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	//ids := make([]int, data.TotalCount)
	//for i, x := range data.Data {
	//	r := x.(model.Area)
	//	ids[i] = r.Id
	//}
	////fmt.Println(ids)
	//stat := make([]model.AreaExt, 0)
	//o := db.NewDb()
	//err = o.In("aid", ids).Find(&stat)
	//if err != nil {
	//	stat = []model.AreaExt{}
	//	fmt.Println(err)
	//}
	return data, nil
}