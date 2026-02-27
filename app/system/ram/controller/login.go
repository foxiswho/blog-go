package controller

import (
	modRamLogin2 "github.com/foxiswho/blog-go/app/system/ram/model/modRamLogin"
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {

}

// LoginController 登录
// @Description:
type LoginController struct {
	sv  *service.AccountLoginService `autowire:"?"`
	log *log2.Logger                 `autowire:"?"`
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
	ctx.JSON(200, c.sv.Login(ctx, ct, appModulePg.System))
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
