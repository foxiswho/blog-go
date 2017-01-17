package admin

import (
	"blog/model"
	"blog/fox/db"
)
//扩展地区
type AreaExt struct {
	*model.AreaExt
}
//快速初始化
func NewAreaExtService() *AreaExt {
	return new(AreaExt)
}
//列表
func (c *AreaExt)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewAreaExt()
	data, err := mode.GetAll(q, fields, orderBy, page, 9999)
	if err != nil {
		return nil, err
	}
	return data, nil
}