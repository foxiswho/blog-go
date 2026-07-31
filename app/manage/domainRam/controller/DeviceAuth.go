package controller

import (
	"github.com/gin-gonic/gin"
	modRamDeviceAuth "github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/model/modRamDeviceAuth"
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/validatorPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(DeviceAuthController)).Name("ManageDeviceAuthController").Export(gs.As[routerPg.RouteRegistrar]())
}

// DeviceAuthController 设备授权（RFC 8628）
// @Description:
type DeviceAuthController struct {
	routerPg.RouteRegistrar
	sv  *service.DeviceAuthService `autowire:"?"`
	log *log2.Logger               `autowire:"?"`
}

// RegisterRoutes 注册路由
func (c *DeviceAuthController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/xianfu/auth/manage/device")
	group.POST("/start", c.Start)
	group.POST("/poll", c.Poll)
	group.POST("/approve", c.Approve)
	group.POST("/cancel", c.Cancel)
	group.POST("/complete", c.Complete)
}

// Start 发起设备授权
func (c *DeviceAuthController) Start(ctx *gin.Context) {
	var ct modRamDeviceAuth.DeviceAuthRequestCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.StartDeviceAuth(ctx, ct))
}

// Poll 设备轮询状态
func (c *DeviceAuthController) Poll(ctx *gin.Context) {
	var ct modRamDeviceAuth.DevicePollCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.PollDeviceStatus(ctx, ct))
}

// Approve 用户授权
func (c *DeviceAuthController) Approve(ctx *gin.Context) {
	var ct modRamDeviceAuth.DeviceApproveCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.ApproveDeviceAuth(ctx, ct))
}

// Cancel 取消授权
func (c *DeviceAuthController) Cancel(ctx *gin.Context) {
	var ct modRamDeviceAuth.DeviceCancelCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.CancelDeviceAuth(ctx, ct))
}

// Complete 完成授权
func (c *DeviceAuthController) Complete(ctx *gin.Context) {
	var ct modRamDeviceAuth.DeviceCompleteCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.CompleteDeviceAuth(ctx, ct))
}
