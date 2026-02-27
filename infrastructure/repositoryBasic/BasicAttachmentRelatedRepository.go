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
	gs.Provide(new(BasicAttachmentRelatedRepository)).Init(func(s *BasicAttachmentRelatedRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicAttachmentRelatedRepository])).Init(func(s *support.BaseService[BasicAttachmentRelatedRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicAttachmentRelatedRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicAttachmentRelatedEntity, int64]
}

func (c *BasicAttachmentRelatedRepository) FindAllByIdLink(code string) (info []entityBasic.BasicAttachmentRelatedEntity, result bool, err error) {
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
