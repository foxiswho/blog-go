package repositoryRam

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(RamAppCategoryRepository)).Init(func(s *RamAppCategoryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAppCategoryRepository])).Init(func(s *support.BaseService[RamAppCategoryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAppCategoryRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAppCategoryEntity, int64]
}

func (c *RamAppCategoryRepository) FindAllByParentIdLink(code string) (info []*entityRam.RamAppCategoryEntity, result bool) {
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
func (c *RamAppCategoryRepository) FindAllByNoLink(code string) (infos []*entityRam.RamAppCategoryEntity, result bool) {
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
func (c *RamAppCategoryRepository) FindAllByCodeLinkAndTypeSys(code string, tpSys string) (info []*entityRam.RamAppCategoryEntity, result bool) {
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
