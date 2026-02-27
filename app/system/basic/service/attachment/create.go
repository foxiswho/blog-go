package attachment

import (
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
)

// Create 附件创建
// @Description:
type Create struct {
	dao    *repositoryBasic.BasicAttachmentRepository `autowire:"?"`
	entity entityBasic.BasicAttachmentEntity
}

// NewCreate 附件创建
//
//	@Description:
//	@param dao
//	@param dto
//	@return *Create
func NewCreate(dao *repositoryBasic.BasicAttachmentRepository, entity entityBasic.BasicAttachmentEntity) *Create {
	return &Create{dao: dao, entity: entity}
}

// Processor 处理
//
//	@Description:
//	@receiver c
//	@param ctx
//	@return error
func (c *Create) Processor() error {
	c.entity.ID = 0
	c.dao.Create(&c.entity)
	return nil
}
