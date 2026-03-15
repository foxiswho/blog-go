package repositoryPg

import (
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"gorm.io/gorm"
)

type OptionParams struct {
	Ctx       *gin.Context
	Log       *log2.Logger
	Db        *gorm.DB
	Condition Condition
	Pageable  *pagePg.Pageable
}

type OptionPg func(arg *OptionParams)

func WithOptionPg(opt OptionPg) OptionPg {
	return opt
}
func WithDb(db *gorm.DB) OptionPg {
	return func(arg *OptionParams) {
		arg.Db = db
	}
}
func WithCtx(ctx *gin.Context) OptionPg {
	return func(arg *OptionParams) {
		arg.Ctx = ctx
	}
}

// WithPageable 设置分页参数
func WithPageable(pageable pagePg.Pageable) OptionPg {
	return func(arg *OptionParams) {
		arg.Pageable = &pageable
	}
}

func Scopes(page *pagePg.Pageable) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		size := int64(0)
		offset := int(10)
		if nil != page {
			if page.PageSize > 0 {
				size = page.PageSize
			}
			if page.PageNum > 0 {
				offset = int((page.PageNum - 1) * size)
			}
		}
		return db.Offset(offset).Limit(int(size))
	}
}
