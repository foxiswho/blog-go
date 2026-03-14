package repositoryPg

import (
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

type IRepositoryRange interface {
	repositoryPgI.IRepositoryBase
}

func WithCtxOption(ctx *gin.Context) Option {
	return func(arg *OptionArg) {
		arg.Ctx = ctx
	}
}
func WithOption(opt Option) Option {
	return opt
}
