package repositoryBlog

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BlogCollectCategoryRepository)).Init(func(s *BlogCollectCategoryRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogCollectCategoryRepository])).Init(func(s *support.BaseService[BlogCollectCategoryRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogCollectCategoryRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogCollectCategoryEntity, int64]
}
