package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicAccountApplyDenyListEntityRepository)).Init(func(s *BasicAccountApplyDenyListEntityRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicAccountApplyDenyListEntityRepository])).Init(func(s *support.BaseService[BasicAccountApplyDenyListEntityRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicAccountApplyDenyListEntityRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicAccountApplyDenyListEntity, int64]
}

func (c *BasicAccountApplyDenyListEntityRepository) FindByExprAndIdNot(name string, id string) (info *entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("expr=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
