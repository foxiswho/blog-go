package controller

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
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
	
	// 模版
	ctx.JSON(200, gin.H{"data": "ok"})
}
