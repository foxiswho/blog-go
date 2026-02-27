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
	gs.Provide(new(BlogCollectRepository)).Init(func(s *BlogCollectRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogCollectRepository])).Init(func(s *support.BaseService[BlogCollectRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogCollectRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogCollectEntity, int64]
}

func (c *BlogCollectRepository) FindAllByUrlSourceMd5(code string) (infos []*entityBlog.BlogCollectEntity, result bool) {
	tx := c.Db().Where("url_source_md5 = ?", code).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}
