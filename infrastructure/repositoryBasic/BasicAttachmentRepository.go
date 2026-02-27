package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicAttachmentRepository)).Init(func(s *BasicAttachmentRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicAttachmentRepository])).Init(func(s *support.BaseService[BasicAttachmentRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicAttachmentRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicAttachmentEntity, int64]
}

func (c *BasicAttachmentRepository) FindAllByModuleValue(module, value string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("module=?", module).Where("value=?", value).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) FindAllByModuleValueIn(module string, value []string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("module=?", module).Where("value in ?", value).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) FindByModuleTypeValue(module, typ, value string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("module=?", module).Where("type=?", typ).Where("value=?", value).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) DeleteByModuleTypeValue(module, typ, value string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("module=?", module).Where("type=?", typ).Where("value=?", value).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) FindByMark(mark string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("mark=?", mark).Order("sort ASC,id desc").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) DeleteByMark(mark string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("mark=?", mark).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) DeleteByNoAndFileOwner(no []string, fileOwner string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("no in ?", no).Where("file_owner=?", fileOwner).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) DeleteByIdAndFileOwner(no []string, fileOwner string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.Db().Where("id in ?", no).Where("file_owner=?", fileOwner).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) UpdateByNoAndFileOwnerSetState13(no []string, fileOwner string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.DbModel().Where("no in ?", no).Where("file_owner=?", fileOwner).Update("state", 13)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) UpdateByIdAndFileOwnerSetState13(no []string, fileOwner string) (info []*entityBasic.BasicAttachmentEntity, result bool) {
	tx := c.DbModel().Where("id in ?", no).Where("file_owner=?", fileOwner).Update("state", 13)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicAttachmentRepository) UpdateByIdSetFileOwner(no []string, fileOwner string) (result bool) {
	tx := c.DbModel().Where("id in ?", no).Update("file_owner", fileOwner)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return false
	}
	if 0 == tx.RowsAffected {
		return false
	}
	return true
}
