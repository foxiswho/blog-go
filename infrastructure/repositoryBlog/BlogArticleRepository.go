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
	gs.Provide(new(BlogArticleRepository)).Init(func(s *BlogArticleRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogArticleRepository])).Init(func(s *support.BaseService[BlogArticleRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogArticleRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogArticleEntity, int64]
}
