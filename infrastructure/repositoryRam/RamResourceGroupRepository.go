package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamResourceGroupRepository)).Init(func(s *RamResourceGroupRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceGroupRepository])).Init(func(s *support.BaseService[RamResourceGroupRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceGroupRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceGroupEntity, int64]
	//
}

func (c *RamResourceGroupRepository) FindByNameAndIdNot(name string, id int64) (info *entityRam.RamResourceGroupEntity, result bool) {
	tx := c.Db().Where("name=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamResourceGroupRepository) FindByParentId(code int64) (info *entityRam.RamResourceGroupEntity, result bool) {
	tx := c.Db().Where("parent_id=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamResourceGroupRepository) FindAllByIdLink(code string) (info []*entityRam.RamResourceGroupEntity, result bool) {
	tx := c.Db().Where("id_link like ?", "%|"+code+"|%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamResourceGroupRepository) FindByTypeCategoryAndTypeValue(typeCategory, typeValue string) (info *entityRam.RamResourceGroupEntity, result bool) {
	tx := c.Db().Where("type_category=?", typeCategory).Where("type_value=?", typeValue).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// CountByParentIdString
//
//	@Description: 统计
//	@receiver c
//	@param pid
//	@return info
//	@return result
func (c *RamResourceGroupRepository) CountByParentIdString(pid string) (total int64, result bool) {
	tx := c.Db().Where("parent_id= ? ", pid).Count(&total)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return 0, false
	}
	if 0 == tx.RowsAffected {
		return 0, false
	}
	return total, true
}

// CountByIdLinkAndNotParentId
//
//	@Description: 统计
//	@receiver c
//	@param pid
//	@return info
//	@return result
func (c *RamResourceGroupRepository) CountByIdLinkAndNotParentId(id, pid string) (total int64, result bool) {
	tx := c.Db().Where("parent_id != ? ", pid).Where("id_link like ?", "%"+id+"%").Count(&total)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return 0, false
	}
	if 0 == tx.RowsAffected {
		return 0, false
	}
	return total, true
}
