package controller

import (
	"github.com/foxiswho/blog-go/app/system/basic/model/modBasicAttachment"
	"github.com/foxiswho/blog-go/app/system/basic/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(FilePubController)).Name("SystemFilePubController").Export(gs.As[routerPg.RouteRegistrar]())
}

type FilePubController struct {
	routerPg.RouteRegistrar
	Sp  *authPg.GroupSystemMiddlewareSp `autowire:""`
	sv  *service.BasicAttachmentService `autowire:""`
	log *log2.Logger                    `autowire:"?"`
}

func (c *FilePubController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/pg2lq/sys/basic/filePub", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/upload", c.Upload)
	group.POST("/upload-more", c.UploadMore)
	group.POST("/upload-link", c.UploadLink)
	group.POST("/upload-list", c.Query)
	group.POST("/upload-detail", c.UploadDetail)
	group.POST("/upload-addByFileOwner", c.UploadAddByFileOwner)
	group.POST("/query", c.Query)
}

func (c *FilePubController) Upload(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Upload(ctx))
}

func (c *FilePubController) UploadMore(ctx *gin.Context) {
	ctx.JSON(200, c.sv.UploadMore(ctx))
}

func (c *FilePubController) UploadLink(ctx *gin.Context) {
	var ct modBasicAttachment.WebUrlCt
	if err := ctx.ShouldBind(&ct); err != nil {
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

func (c *FilePubController) Query(ctx *gin.Context) {
	var ct modBasicAttachment.QueryCt
	if err := ctx.ShouldBind(&ct); err != nil {
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

func (c *FilePubController) UploadDetail(ctx *gin.Context) {
	var ct modBasicAttachment.DetailByFileOwnerCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.UpdateDetail(ctx, ct))
}

func (c *FilePubController) UploadAddByFileOwner(ctx *gin.Context) {
	var ct modBasicAttachment.AddByFileOwnerCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.UpdateAddByFileOwner(ctx, ct))
}
