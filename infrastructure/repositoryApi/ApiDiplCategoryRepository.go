package repositoryApi

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(ApiDiplCategoryRepository)).Init(func(s *ApiDiplCategoryRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[ApiDiplCategoryRepository])).Init(func(s *support.BaseService[ApiDiplCategoryRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ApiDiplCategoryRepository struct {
	repositoryPg.BaseRepository[entityApi.ApiDiplCategoryEntity, int64]
}

func (c *ApiDiplCategoryRepository) FindAllByParentIdLink(ctx context.Context, code string) (info []*entityApi.ApiDiplCategoryEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("id_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *ApiDiplCategoryRepository) FindAllByCodeLinkAndTypeSys(ctx context.Context, code string, tpSys string) (info []*entityApi.ApiDiplCategoryEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("type_sys = ?", tpSys).Where("no_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
