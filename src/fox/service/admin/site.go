package admin

import (
	"fox/util/db"
	"fox/service"
)

type Site struct {

}
func NewSiteService()*Site{
	return new(Site)
}
//列表
func (t *Site)Query() (*db.Paginator, error) {
	return NewTypeService().Query(service.SITE_ID)
}