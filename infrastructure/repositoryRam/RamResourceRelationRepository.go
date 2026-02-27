package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamResourceRelationRepository)).Init(func(s *RamResourceRelationRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceRelationRepository])).Init(func(s *support.BaseService[RamResourceRelationRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceRelationRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceRelationEntity, int64]
	//
}

func (c *RamResourceRelationRepository) DeleteByAuthorityId(code int64) {
	c.Db().Where("authority_id=?", code).Delete(&entityRam.RamResourceRelationEntity{})
}

func (c *RamResourceRelationRepository) FindByMark(code string) (info *entityRam.RamResourceRelationEntity, result bool) {
	tx := c.Db().Where("mark=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
