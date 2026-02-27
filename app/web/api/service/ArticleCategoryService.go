package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/web/api/model/modBlogArticleCategory"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(ArticleCategoryService)).Init(func(s *ArticleCategoryService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
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
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBlogArticleCategory.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeNo{
				Key:      item.No,
				Id:       item.No,
				No:       item.No,
				Label:    item.Name,
				ParentNo: item.ParentNo,
				ParentId: item.ParentNo,
				Extend:   vo,
			}
			//编码
			if enumParameterPg.NodeQueryById.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
				code.ParentId = item.ParentId
			}
			slice = append(slice, code)
		}
		rt.Data = slice
	}
	return rt.Ok()
}
