package repositoryPgI

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"gorm.io/gorm"
)

type IRepositoryBase interface {
	SetCtx(*gin.Context, ...interface{})
	SetCtxDbLog(*gin.Context, *gorm.DB, *log2.Logger) *gorm.DB
}
