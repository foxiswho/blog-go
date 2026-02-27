package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modPublic"
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

// PublicController 用户公共动作
// @Description:
type PublicController struct {
	controllerPg.SpSystemAuth
	sv  *service.RamAccountPublicService `autowire:"?"`
	log *log2.Logger                     `autowire:"?"`
}

// Public 用户详情
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *PublicController) Public(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Public(holderPg.GetContextAccount(ctx)))
}

// UpdatePassword 修改密码
func (c *PublicController) UpdatePassword(ctx *gin.Context) {
	var ct modPublic.PasswordCt
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
	ctx.JSON(200, c.sv.UpdatePassword(ctx, ct))
}
