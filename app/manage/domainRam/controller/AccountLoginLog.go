package controller

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamAccountLoginLog"
	"github.com/foxiswho/blog-go/app/manage/domainRam/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"

	"github.com/foxiswho/blog-go/pkg/model"
)

func init() {

}

// AccountLoginLogController 团队
// @Description:
type AccountLoginLogController struct {
	controllerPg.SpManageAuth
	sv  *service.RamAccountLoginLogService `autowire:"?"`
	log *log2.Logger                       `autowire:"?"`
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *AccountLoginLogController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
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
	ctx.JSON(200, c.sv.PhysicalDeletion(ctx, ct.Ids))
}

// Query 查询列表
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *AccountLoginLogController) Query(ctx *gin.Context) {
	var ct modRamAccountLoginLog.QueryCt
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
	ctx.JSON(200, c.sv.Query(ctx, ct))
}
