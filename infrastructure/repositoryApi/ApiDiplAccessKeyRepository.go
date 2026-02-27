package repositoryApi

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(ApiDiplAccessKeyRepository)).Init(func(s *ApiDiplAccessKeyRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[ApiDiplAccessKeyRepository])).Init(func(s *support.BaseService[ApiDiplAccessKeyRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ApiDiplAccessKeyRepository struct {
	repositoryPg.BaseRepository[entityApi.ApiDiplAccessKeyEntity, int64]
}

func (c *ApiDiplAccessKeyRepository) FindByTenantNoAndDiplNo(no, DiplNo string) (info *entityApi.ApiDiplAccessKeyEntity, result bool) {
	tx := c.Db().Where("tenant_no=?", no).Where("dipl_no=?", DiplNo).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *ApiDiplAccessKeyRepository) UpdateAllByDiplNoAndNoSetState(DiplNo, id string, state int8) (sum int64, result bool) {
	tx := c.Db().Where("dipl_no=?", DiplNo).Where("id=?", id).Updates(entityApi.ApiDiplAccessKeyEntity{State: state})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return 0, false
	}
	if 0 == tx.RowsAffected {
		return 0, false
	}
	return tx.RowsAffected, true
}
