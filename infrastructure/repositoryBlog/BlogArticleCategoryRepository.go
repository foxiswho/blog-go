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

func (c *BlogArticleCategoryRepository) FindAllByParentIdLink(code string) (info []*entityBlog.BlogArticleCategoryEntity, result bool) {
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
func (c *BlogArticleCategoryRepository) FindAllByNoLink(code string) (infos []*entityBlog.BlogArticleCategoryEntity, result bool) {
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
func (c *BlogArticleCategoryRepository) FindAllByCodeLinkAndTypeSys(code string, tpSys string) (info []*entityBlog.BlogArticleCategoryEntity, result bool) {
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
