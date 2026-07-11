package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/basic/model/modBasicAttachment"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/basic/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/authPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
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
	group.POST("/upload-ownerDel", c.UploadDelByOwner)
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
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.UploadLink(ctx, ct))
}

func (c *FilePubController) Query(ctx *gin.Context) {
	var ct modBasicAttachment.QueryCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Query(ctx, ct))
}

func (c *FilePubController) UploadDetail(ctx *gin.Context) {
	var ct modBasicAttachment.DetailByFileOwnerCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.UpdateDetail(ctx, ct))
}

func (c *FilePubController) UploadAddByFileOwner(ctx *gin.Context) {
	var ct modBasicAttachment.AddByFileOwnerCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.UpdateAddByFileOwner(ctx, ct))
}

func (c *FilePubController) UploadDelByOwner(ctx *gin.Context) {
	var ct modBasicAttachment.DelFileOwnerCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.DelByOwner(ctx, ct))
}
