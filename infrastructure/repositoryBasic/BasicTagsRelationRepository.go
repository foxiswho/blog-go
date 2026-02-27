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
	gs.Provide(new(BasicTagsRelationRepository)).Init(func(s *BasicTagsRelationRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicTagsRelationRepository])).Init(func(s *support.BaseService[BasicTagsRelationRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicTagsRelationRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicTagsRelationEntity, int64]
}

func (c *BasicTagsRelationRepository) FindAllByParentIdLink(code string) (info []entityBasic.BasicTagsRelationEntity, result bool) {
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
func (c *BasicTagsRelationRepository) FindByNameAndIdNot(name string, id int64) (info *entityBasic.BasicTagsRelationEntity, result bool) {
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
func (c *BasicTagsRelationRepository) FindByNameAndIdNotAndCategoryNot(name string, id int64, category string) (info *entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("name=?", name).Where("category=?", category).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicTagsRelationRepository) FindByCodeAndIdNot(name string, id int64) (info *entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("code=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicTagsRelationRepository) FindByCodeAndIdNotAndCategoryNot(name string, id int64, category string) (info *entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("code=?", name).Where("category=?", category).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicTagsRelationRepository) FindByCode(name string) (info *entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("code=?", name).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (b *BasicTagsRelationRepository) DeleteByIdsStringAndTypeSysNot(id []string, tp string) error {
	tx := b.Db().Where("type_sys!=?", tp).Where("id in ?").Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindAllByCategoryNoIn
//
//	@Description:
//	@receiver c
//	@param t
//	@param category
//	@return infos
//	@return result
func (c *BasicTagsRelationRepository) FindAllByCategoryNoIn(t entityBasic.BasicTagsRelationEntity, category []string) (infos []*entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where(t).Where("category_no in ?", category).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

// FindAllByCategoryRootIn
//
//	@Description:
//	@receiver c
//	@param t
//	@param category
//	@return infos
//	@return result
func (c *BasicTagsRelationRepository) FindAllByCategoryRootIn(t entityBasic.BasicTagsRelationEntity, category []string) (infos []*entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where(t).Where("category_root in ?", category).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

// FindAllByCodeInAndCategoryRoot
//
//	@Description:
//	@receiver c
//	@param t
//	@param category
//	@return infos
//	@return result
func (c *BasicTagsRelationRepository) FindAllByCodeInAndCategoryRoot(code []string, categoryRoot string) (infos []*entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("code in ?", code).Where("category_root=?", categoryRoot).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

// FindAllByTagNoInAndCategoryRoot
//
//	@Description:
//	@receiver c
//	@param t
//	@param category
//	@return infos
//	@return result
func (c *BasicTagsRelationRepository) FindAllByTagNoInAndCategoryRoot(code []string, categoryRoot string) (infos []*entityBasic.BasicTagsRelationEntity, result bool) {
	tx := c.Db().Where("tag_no in ?", code).Where("category_root=?", categoryRoot).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}
