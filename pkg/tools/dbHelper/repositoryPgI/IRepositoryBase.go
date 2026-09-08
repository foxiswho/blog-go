package repositoryPgI

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRepositoryBase interface {
	SetCtx(*gin.Context, ...any)
	SetCtxDbLog(*gin.Context, *gorm.DB) *gorm.DB
}
