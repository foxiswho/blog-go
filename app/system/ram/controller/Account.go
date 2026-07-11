package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/model/modRamAccount"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/authPg"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/serverPg/ginServer"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/common/controllerPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/model"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(AccountController).SetAppModule(appModulePg.System)).Name("SystemAccountController").Export(gs.As[routerPg.RouteRegistrar]())
}

type AccountController struct {
	routerPg.RouteRegistrar
	controllerPg.SpSystemAuth
	sv        *service.RamAccountService         `autowire:"?"`
	ap        *service.RamAccountPasswordService `autowire:"?"`
	appModule appModulePg.AppModule
	log       *log2.Logger `autowire:"?"`
}

func (c *AccountController) SetAppModule(appModule appModulePg.AppModule) *AccountController {
	c.appModule = appModule
	return c
}

func (c *AccountController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/pg2lq/sys/ram/account", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/enable", c.Enable)
	group.POST("/disable", c.Disable)
	group.GET("/detail/:id", c.Detail)
	group.POST("/delete", c.Delete)
	group.POST("/recovery", c.Recovery)
	group.POST("/physicalDeletion", c.PhysicalDeletion)
	group.POST("/query", c.Query)
	group.POST("/updatePassword", c.UpdatePassword)
	group.POST("/createUpdateAccount", c.CreateUpdateAccount)
	group.POST("/createUpdate", c.CreateUpdate)
	group.POST("/existAccount", c.ExistAccount)
	group.POST("/existPhone", c.ExistPhone)
	group.POST("/existMail", c.ExistMail)
	group.POST("/existCode", c.ExistCode)
	group.POST("/existIdentityCode", c.ExistIdentityCode)
	group.POST("/existRealName", c.ExistRealName)
}

func (c *AccountController) Detail(ctx *gin.Context) {
	param := ctx.Param("id")
	if "" == param {
		ctx.JSON(200, rg.Error[string]("id不能为空"))
		return
	}
	ctx.JSON(200, c.sv.Detail(ctx, param, c.appModule))
}

func (c *AccountController) Enable(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Enable(ctx, ct, c.appModule))
}

func (c *AccountController) Disable(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Disable(ctx, ct, c.appModule))
}

func (c *AccountController) Delete(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.LogicalDeletion(ctx, ct.Ids, c.appModule))
}

func (c *AccountController) Recovery(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.LogicalRecovery(ctx, ct.Ids, c.appModule))
}

func (c *AccountController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.PhysicalDeletion(ctx, ct.Ids, c.appModule))
}

func (c *AccountController) UpdatePassword(ctx *gin.Context) {
	var ct modRamAccount.PasswordCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.ap.UpdatePassword(ctx, ct, c.appModule))
}

func (c *AccountController) Query(ctx *gin.Context) {
	var ct modRamAccount.QueryCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Query(ctx, ct, c.appModule))
}

func (c *AccountController) CreateUpdate(ctx *gin.Context) {
	var ct modRamAccount.CreateUpdateCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.CreateUpdate(ctx, ct, c.appModule))
}

func (c *AccountController) CreateUpdateAccount(ctx *gin.Context) {
	var ct modRamAccount.CreateUpdateAccountCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.CreateUpdateAccount(ctx, ct, c.appModule))
}

func (c *AccountController) ExistAccount(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.ExistAccount(ctx, ct, c.appModule))
}

func (c *AccountController) ExistPhone(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.ExistPhone(ctx, ct, c.appModule))
}

func (c *AccountController) ExistMail(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.ExistMail(ctx, ct, c.appModule))
}

func (c *AccountController) ExistCode(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.ExistCode(ctx, ct, c.appModule))
}

func (c *AccountController) ExistIdentityCode(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.ExistIdentityCode(ctx, ct, c.appModule))
}

func (c *AccountController) ExistRealName(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.ExistRealName(ctx, ct, c.appModule))
}
