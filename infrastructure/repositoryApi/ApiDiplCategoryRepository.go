package repositoryApi

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(ApiDiplCategoryRepository)).Init(func(s *ApiDiplCategoryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[ApiDiplCategoryRepository])).Init(func(s *support.BaseService[ApiDiplCategoryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ApiDiplCategoryRepository struct {
	repositoryPg.BaseRepository[entityApi.ApiDiplCategoryEntity, int64]
}

func (c *ApiDiplCategoryRepository) FindAllByParentIdLink(code string) (info []*entityApi.ApiDiplCategoryEntity, result bool) {
	tx := c.Db().Where("id_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *ApiDiplCategoryRepository) FindAllByNoLink(code string) (infos []*entityApi.ApiDiplCategoryEntity, result bool) {
	tx := c.Db().Where("no_link like ?", "%"+code+"%").Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}
func (c *ApiDiplCategoryRepository) FindAllByCodeLinkAndTypeSys(code string, tpSys string) (info []*entityApi.ApiDiplCategoryEntity, result bool) {
	tx := c.Db().Where("type_sys = ?", tpSys).Where("no_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
