package controller

import (
	"github.com/foxiswho/blog-go/app/core/cache/cacheRam"
	"github.com/foxiswho/blog-go/app/system/pub/model/modSysPub"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(PublicAuthPubPrivKeyController)).Name("SystemPublicAuthPubPrivKeyController").Export(gs.As[routerPg.RouteRegistrar]())
}

// PublicAuthPubPrivKeyController 公用 配置
type PublicAuthPubPrivKeyController struct {
	sv *cacheRam.CacheSessionPubPrive `autowire:"?"`
}

func (c *PublicAuthPubPrivKeyController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/pg2lq/pub")
	group.GET("/config", c.config)
}

// 获取 配置信息
func (c *PublicAuthPubPrivKeyController) config(ctx *gin.Context) {
	key := c.sv.LoginPubPriveKey(ctx)
	vo := modSysPub.NewAuthConfigPublic(key)
	vo.LoginEncrypt = true
	ctx.JSON(200, rg.OkData(vo))
}
