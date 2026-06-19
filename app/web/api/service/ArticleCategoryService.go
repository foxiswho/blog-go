package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/web/api/model/modBlogArticleCategory"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(ArticleCategoryService)).Init(func(s *ArticleCategoryService) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ArticleCategoryService struct {
	sv *repositoryBlog.BlogArticleCategoryRepository `autowire:"?"`
}

// SelectNodeAllPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ArticleCategoryService) SelectNodeAllPublic(ctx *gin.Context, ct modBlogArticleCategory.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityBlog.BlogArticleCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(ctx, query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBlogArticleCategory.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeNo{
				Value:    item.No,
				Label:    item.Name,
				ParentNo: item.ParentNo,
				Extend:   vo,
			}
			slice = append(slice, code)
		}
		rt.Data = slice
	}
	return rt.Ok()
}
