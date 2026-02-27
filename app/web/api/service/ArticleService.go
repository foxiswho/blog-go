package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBlog/service/blogArticle"
	"github.com/foxiswho/blog-go/app/web/api/model/modBlogArticle"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(ArticleService)).Init(func(s *ArticleService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ArticleService struct {
	sv    *repositoryBlog.BlogArticleRepository           `autowire:"?"`
	catDb *repositoryBlog.BlogArticleCategoryRepository   `autowire:"?"`
	sata  *repositoryBlog.BlogArticleStatisticsRepository `autowire:"?"`
	sp    *blogArticle.Sp                                 `autowire:"?"`
	log   *log2.Logger                                    `autowire:"?"`
}

// Push
//
//	@Description: 推送文章连接
//	@receiver c
func (c *ArticleService) Push(ctx *gin.Context, ct modBlogArticle.PushCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if strPg.IsBlank(ct.CategoryNo) {
		return rt.ErrorMessage("请选择分类")
	}
	if strPg.IsBlank(ct.Title) {
		return rt.ErrorMessage("标题不能为空")
	}
	if strPg.IsBlank(ct.Url) {
		return rt.ErrorMessage("url地址不能为空")
	}
	info, result := c.catDb.FindByNo(ct.CategoryNo)
	if !result {
		return rt.ErrorMessage("分类不存在")
	}
	holder := holderPg.GetContextAccount(ctx)
	save := entityBlog.BlogArticleEntity{
		CategoryNo:  info.No,
		Name:        ct.Title,
		UrlSource:   ct.Url,
		Description: ct.Description,
	}
	save.No = noPg.No()
	save.Code = save.No
	save.TenantNo = holder.GetTenantNo()
	save.Ano = holder.GetAccountNo()
	err, _ := c.sv.Create(&save)
	if err != nil {
		c.log.Debugf("save err=%+v", err)
		return rt.ErrorMessage("保存失败：" + err.Error())
	}
	err, _ = c.sata.Create(&entityBlog.BlogArticleStatisticsEntity{ID: save.ID, ArticleNo: save.No})
	if err != nil {
		c.log.Debugf("save err=%+v", err)
	}
	return rt.Ok()
}
