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
	gs.Provide(new(BlogTopicCategoryRepository)).Init(func(s *BlogTopicCategoryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogTopicCategoryRepository])).Init(func(s *support.BaseService[BlogTopicCategoryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogTopicCategoryRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogTopicCategoryEntity, int64]
}
