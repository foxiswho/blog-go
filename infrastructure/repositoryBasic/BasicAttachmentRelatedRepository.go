package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicAttachmentRelatedRepository)).Init(func(s *BasicAttachmentRelatedRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicAttachmentRelatedRepository])).Init(func(s *support.BaseService[BasicAttachmentRelatedRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicAttachmentRelatedRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicAttachmentRelatedEntity, int64]
}
