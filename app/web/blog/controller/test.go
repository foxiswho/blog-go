package controller

import (
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/utils/webPg"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumApprovedPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/sdk/blog/key/blogKeyPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/strPg"
)

// TestController test
type TestController struct {
	log *log2.Logger                          `autowire:"?"`
	sv  *repositoryBlog.BlogArticleRepository `autowire:"?"`
}

func (c *TestController) Cache(ctx *gin.Context) {
	//err := articleBlogEvent.NewStartInit(c.log).Processor(context.Background())
	//if err != nil {
	//	c.log.Error("error:", err)
	//}
	// 模版
	ctx.JSON(200, gin.H{"data": "ok"})
}

func (c *TestController) FindAllPage(ctx *gin.Context) {
	var ct modBlogArticle.QueryCt
	ctx.Bind(&ct)
	var query entityBlog.BlogArticleEntity
	copier.Copy(&query, &ct)
	tenantNo := webPg.GetTenantNo(ctx)
	if strPg.IsNotBlank(tenantNo) {
		query.TenantNo = tenantNo
	}
	//启用
	query.State = enumStatePg.ENABLE.Index()
	//审批通过
	query.PlatformApproved = enumApprovedPg.ApprovedStateApproved.Index()
	slice := make([]modBlogArticle.Vo, 0)
	r := c.sv
	page, err := r.FindAllPage(ctx, query, repositoryPg.WithOptionPg[entityBlog.BlogArticleEntity](func(arg *repositoryPg.OptionParams[entityBlog.BlogArticleEntity]) {
		arg.PageOption = func(c *pagePg.Paginator[*entityBlog.BlogArticleEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
	}), repositoryPg.WithOptionPg(func(arg *repositoryPg.OptionParams[any]) {

	}))
	func(p *pagePg.PageCondition[*entityBlog.BlogArticleEntity]) {
		p.PageOption = func(c *pagePg.Paginator[*entityBlog.BlogArticleEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		//自定义查询
		p.Condition = r.DbModel().Order("create_at desc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
	}
	// 模版
	ctx.JSON(200, gin.H{"data": "ok"})
}
