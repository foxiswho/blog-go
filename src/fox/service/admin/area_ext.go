package admin

import (
	"fox/model"
	"fox/util/db"
)
type AreaExt struct {
	*model.AreaExt
}

func NewAreaExtService() *AreaExt {
	return new(AreaExt)
}
func (c *AreaExt)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewAreaExt()
	data, err := mode.GetAll(q, fields, orderBy, page, 9999)
	if err != nil {
		return nil, err
	}
	return data, nil
}