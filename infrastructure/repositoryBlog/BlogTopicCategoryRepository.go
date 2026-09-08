package repositoryBlog

import (
	"context"
	"reflect"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityBlog"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(BlogTopicCategoryRepository))

	gs.Provide(new(support.BaseService[BlogTopicCategoryRepository]))
}

type BlogTopicCategoryRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogTopicCategoryEntity, int64]
}
