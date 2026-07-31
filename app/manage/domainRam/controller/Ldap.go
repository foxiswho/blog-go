package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/model/modRamLdap"
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/validatorPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(LdapController)).Name("ManageLdapController").Export(gs.As[routerPg.RouteRegistrar]())
}

// LdapController LDAP 管理
type LdapController struct {
	routerPg.RouteRegistrar
	sv  *service.LdapService `autowire:"?"`
	log *log2.Logger         `autowire:"?"`
}

// RegisterRoutes 注册路由
func (c *LdapController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/xianfu/manage/ram/ldap")
	group.POST("/test", c.TestConnection)
	group.POST("/search", c.SearchUsers)
	group.POST("/sync", c.SyncUsers)
	group.POST("/login", c.Login)
}

// TestConnection 测试 LDAP 连接
func (c *LdapController) TestConnection(ctx *gin.Context) {
	var ct modRamLdap.LdapTestCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.TestConnection(ctx, ct))
}

// SearchUsers 搜索 LDAP 用户
func (c *LdapController) SearchUsers(ctx *gin.Context) {
	var ct modRamLdap.LdapSearchCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.SearchUsers(ctx, ct))
}

// SyncUsers 同步 LDAP 用户
func (c *LdapController) SyncUsers(ctx *gin.Context) {
	var ct modRamLdap.LdapSyncCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.SyncUsers(ctx, ct))
}

// Login LDAP 登录
func (c *LdapController) Login(ctx *gin.Context) {
	var ct modRamLdap.LdapLoginCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Login(ctx, ct))
}
