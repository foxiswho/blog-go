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
	gs.Provide(new(BasicAreaRepository)).Init(func(s *BasicAreaRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicAreaRepository])).Init(func(s *support.BaseService[BasicAreaRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicAreaRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicAreaEntity, int64]
}

func (c *BasicAreaRepository) FindAllByIdLink(code string) (info []entityBasic.BasicAreaEntity, result bool, err error) {
	tx := c.Db().Where("id_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}
