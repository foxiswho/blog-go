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
	gs.Provide(new(BlogArticleStatisticsRepository)).Init(func(s *BlogArticleStatisticsRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogArticleStatisticsRepository])).Init(func(s *support.BaseService[BlogArticleStatisticsRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogArticleStatisticsRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogArticleStatisticsEntity, int64]
}

func (c *BlogArticleStatisticsRepository) FindByArticleNo(no string) (info *entityBlog.BlogArticleStatisticsEntity, result bool) {
	tx := c.Db().Where("article_no=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BlogArticleStatisticsRepository) FindAllByArticleNoIn(no []string) (info []*entityBlog.BlogArticleStatisticsEntity, result bool) {
	tx := c.Db().Where("article_no in ?", no).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
