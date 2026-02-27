package controller

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/service"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
)

func init() {

}

// LogoutController 退出
// @Description:
type LogoutController struct {
	sv  *service.AccountLogoutService `autowire:"?"`
	log *log2.Logger                  `autowire:"?"`
}

// Logout 退出
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *LogoutController) Logout(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Logout(holderPg.GetContextAccount(ctx)))
}
