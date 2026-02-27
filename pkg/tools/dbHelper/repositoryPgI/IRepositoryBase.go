package repositoryPgI

import (
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRepositoryBase interface {
	SetCtx(*gin.Context, ...interface{})
	SetCtxDbLog(*gin.Context, *gorm.DB, *log2.Logger) *gorm.DB
}
