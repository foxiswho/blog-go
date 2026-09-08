package support

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/configPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPgI"
	"github.com/pangu-2/go-tools/tools/dbPg/genericPg"
	"gorm.io/gorm"
)

type IService[T any, ID genericPg.ID] interface {
}

type BaseService[T any] struct {
	dao *T `autowire:"?"`

	//从内部
	db *gorm.DB    `autowire:"?"`
	Pg configPg.Pg `value:"${pg}"`
}

func (b *BaseService[T]) New(ctx *gin.Context, arg ...interface{}) *T {
	e := new(T)
	base := interface{}(e).(repositoryPgI.IRepositoryBase)
	dbSet := false
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case *gorm.DB:
				base.SetCtx(ctx, result)
				dbSet = true
			}
		}
	} else if !dbSet && b.db != nil {
		base.SetCtx(ctx, b.db)
		dbSet = true
	}
	return e
}

func (b *BaseService[T]) Dao() *T {
	return b.dao
}

func (b *BaseService[T]) Config() configPg.Pg {
	return b.Pg
}

func (b *BaseService[T]) Db() *gorm.DB {
	return b.db
}

func (b *BaseService[T]) This() *BaseService[T] {
	return b
}
