package controller

import (
	"github.com/gin-gonic/gin"
	modRamLogin2 "github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/model/modRamLogin"
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/validatorPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/clientPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(AuthLoginController)).Name("ManageAuthLoginController").Export(gs.As[routerPg.RouteRegistrar]())
}

// AuthLoginController 登录
// @Description:
type AuthLoginController struct {
	routerPg.RouteRegistrar
	sv  *service.AccountLoginService `autowire:"?"`
	log *log2.Logger                 `autowire:"?"`
}

// RegisterRoutes 注册路由
//
//	@Description:
//	@receiver c
//	@param e
func (c *AuthLoginController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/xianfu/auth/manage")
	group.POST("/login", c.Login)
	group.POST("/refresh", c.RefreshToken)
}

// Login 登陆
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *AuthLoginController) Login(ctx *gin.Context) {
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
	ctx.JSON(200, c.sv.Login(ctx, ct, typeDomainPg.Manage))
}

// RefreshToken 刷新
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *AuthLoginController) RefreshToken(ctx *gin.Context) {
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
	ctx.JSON(200, c.sv.RefreshToken(ctx, ct, typeDomainPg.Manage, clientPg.Browser))
}
