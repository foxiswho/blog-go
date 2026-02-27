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

func (c *BlogTopicCategoryRepository) FindAllByParentIdLink(code string) (info []*entityBlog.BlogTopicCategoryEntity, result bool) {
	tx := c.Db().Where("id_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BlogTopicCategoryRepository) FindAllByNoLink(code string) (infos []*entityBlog.BlogTopicCategoryEntity, result bool) {
	tx := c.Db().Where("no_link like ?", "%"+code+"%").Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}
func (c *BlogTopicCategoryRepository) FindAllByCodeLinkAndTypeSys(code string, tpSys string) (info []*entityBlog.BlogTopicCategoryEntity, result bool) {
	tx := c.Db().Where("type_sys = ?", tpSys).Where("no_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
