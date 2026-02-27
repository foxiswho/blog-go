package controller

import (
	"github.com/foxiswho/blog-go/app/system/basic/model/modBasicAttachment"
	"github.com/foxiswho/blog-go/app/system/basic/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

// AttachmentController 附件上传
// @Description:
type AttachmentController struct {
	controllerPg.SpSystemAuth
	sv  *service.BasicAttachmentService `autowire:""`
	log *log2.Logger                    `autowire:"?"`
}

// Upload
//
//	@Description: 但文件上传
//	@receiver c
//	@param ctx
func (c *AttachmentController) Upload(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Upload(ctx))
}

// UploadMore
//
//	@Description:  多文件上传
//	@receiver c
//	@param ctx
func (c *AttachmentController) UploadMore(ctx *gin.Context) {
	ctx.JSON(200, c.sv.UploadMore(ctx))
}

// UploadLink
//
//	@Description:  多url文件上传
//	@receiver c
//	@param ctx
func (c *AttachmentController) UploadLink(ctx *gin.Context) {
	var ct modBasicAttachment.WebUrlCt
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
	ctx.JSON(200, c.sv.UploadLink(ctx, ct))
}

// Query 查询列表
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *AttachmentController) Query(ctx *gin.Context) {
	var ct modBasicAttachment.QueryCt
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
