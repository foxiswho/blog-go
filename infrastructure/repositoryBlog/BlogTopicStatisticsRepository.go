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
	gs.Provide(new(BlogTopicStatisticsRepository)).Init(func(s *BlogTopicStatisticsRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogTopicStatisticsRepository])).Init(func(s *support.BaseService[BlogTopicStatisticsRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogTopicStatisticsRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogTopicStatisticsEntity, int64]
}

func (c *BlogTopicStatisticsRepository) FindByTopicNo(no string) (info *entityBlog.BlogTopicStatisticsEntity, result bool) {
	tx := c.Db().Where("topic_no=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BlogTopicStatisticsRepository) FindAllByTopicNoIn(no []string) (info []*entityBlog.BlogTopicStatisticsEntity, result bool) {
	tx := c.Db().Where("topic_no in ?", no).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
