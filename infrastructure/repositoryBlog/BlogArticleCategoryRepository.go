package repositoryBlog

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BlogArticleCategoryRepository)).Init(func(s *BlogArticleCategoryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogArticleCategoryRepository])).Init(func(s *support.BaseService[BlogArticleCategoryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogArticleCategoryRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogArticleCategoryEntity, int64]
}
