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
	gs.Provide(new(BasicTagsRepository)).Init(func(s *BasicTagsRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicTagsRepository])).Init(func(s *support.BaseService[BasicTagsRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicTagsRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicTagsEntity, int64]
}

func (c *BasicTagsRepository) FindAllByParentIdLink(code string) (info []entityBasic.BasicTagsEntity, result bool) {
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

func (c *BasicTagsRepository) FindByCode(name string) (info *entityBasic.BasicTagsEntity, result bool) {
	tx := c.Db().Where("code=?", name).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
