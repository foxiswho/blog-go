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
	gs.Provide(new(BlogTopicRelationRepository)).Init(func(s *BlogTopicRelationRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BlogTopicRelationRepository])).Init(func(s *support.BaseService[BlogTopicRelationRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogTopicRelationRepository struct {
	repositoryPg.BaseRepository[entityBlog.BlogTopicRelationEntity, int64]
}

func (c *BlogTopicRelationRepository) FindAllByTopicNo(no string) (info []*entityBlog.BlogTopicRelationEntity, result bool) {
	tx := c.Db().Where("topic_no=?", no).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BlogTopicRelationRepository) FindByTopicNoAndArticleNo(topicNo, no string) (info *entityBlog.BlogTopicRelationEntity, result bool) {
	tx := c.Db().Where("topic_no=?", topicNo).Where("article_no=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BlogTopicRelationRepository) FindAllByArticleNo(no string) (info []*entityBlog.BlogTopicRelationEntity, result bool) {
	tx := c.Db().Where("article_no=?", no).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
