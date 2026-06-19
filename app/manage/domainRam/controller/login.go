package controller

import (
	modRamLogin2 "github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamLogin"
	"github.com/foxiswho/blog-go/app/manage/domainRam/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(LoginController)).Name("ManageLoginController").Export(gs.As[routerPg.RouteRegistrar]())
}

// LoginController 登录
// @Description:
type LoginController struct {
	routerPg.RouteRegistrar
	sv  *service.AccountLoginService `autowire:"?"`
	log *log2.Logger                 `autowire:"?"`
}

// RegisterRoutes 注册路由
//
//	@Description:
//	@receiver c
//	@param e
func (c *LoginController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/pg2lq/auth/manage")
	group.POST("/login", c.Login)
	group.POST("/refresh", c.RefreshToken)
}

// Login 登陆
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *LoginController) Login(ctx *gin.Context) {
	var ct modRamLogin2.LoginCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Login(ctx, ct, appModulePg.Manage))
}

// RefreshToken 刷新
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *LoginController) RefreshToken(ctx *gin.Context) {
	var ct modRamLogin2.TokenRefreshCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.RefreshToken(ctx, ct))
}
