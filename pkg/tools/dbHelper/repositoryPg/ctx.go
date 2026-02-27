package repositoryPg

import (
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPgI"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OptionArg struct {
	Ctx *gin.Context
	Log *log2.Logger
	db  *gorm.DB
}

func (c *OptionArg) Db() *gorm.DB {
	return c.db
}

func (c *OptionArg) SetDb(db *gorm.DB) {
	c.db = db
}

type Option func(*OptionArg)

// CtxDb ctx 存入db
func CtxDb(opts ...Option) *gorm.DB {
	arg := OptionArg{}
	for _, opt := range opts {
		opt(&arg)
	}
	return arg.db.WithContext(holderPg.SetContextValue(arg.Ctx))
}

// CtxDbFun 上下文传入 dbPg
func CtxDbFun(ctx *gin.Context, db *gorm.DB) *gorm.DB {
	return db.WithContext(holderPg.SetContextValue(ctx))
}

type IRepositoryRange interface {
	repositoryPgI.IRepositoryBase
}

func GetOption(ctx *gin.Context) Option {
	return func(arg *OptionArg) {
		arg.Ctx = ctx
	}
}
