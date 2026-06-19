package controller

import (
	"fmt"

	"github.com/foxiswho/blog-go/app/system/basic/model/modBasicCountry"
	"github.com/foxiswho/blog-go/app/system/basic/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(CountryController)).Name("SystemCountryController").Export(gs.As[routerPg.RouteRegistrar]())
}

type CountryController struct {
	routerPg.RouteRegistrar
	Sp *authPg.GroupSystemMiddlewareSp `autowire:""`
	sv *service.BasicCountryService    `autowire:"?"`
}

func (c *CountryController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/pg2lq/sys/basic/country", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/createUpdate", c.CreateUpdate)
	group.GET("/detail/:id", c.Detail)
	group.POST("/enable", c.Enable)
	group.POST("/disable", c.Disable)
	group.POST("/delete", c.Delete)
	group.POST("/recovery", c.Recovery)
	group.POST("/physicalDeletion", c.PhysicalDeletion)
	group.POST("/query", c.Query)
	group.POST("/selectNodeAll", c.SelectNodeAll)
	group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
	group.POST("/selectPublicCountryCode", c.SelectPublicCountryCode)
	group.POST("/selectNodePublicCountryCode", c.SelectNodePublicCountryCode)
	group.POST("/existName", c.ExistName)
	group.POST("/existCode", c.ExistCode)
	group.POST("/existCountryCode", c.ExistCountryCode)
}

func (c *CountryController) CreateUpdate(ctx *gin.Context) {
	var ct modBasicCountry.CreateUpdateCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	if ct.ID.ToInt64() > 0 {
		ctx.JSON(200, c.sv.Update(ctx, ct))
	} else {
		ctx.JSON(200, c.sv.Create(ctx, ct))
	}
}

func (c *CountryController) Delete(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.LogicalDeletion(ctx, ct.Ids))
}

func (c *CountryController) Recovery(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.LogicalRecovery(ctx, ct.Ids))
}

func (c *CountryController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
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

func (c *CountryController) Detail(ctx *gin.Context) {
	param := ctx.Param("id")

	fmt.Println(param)
	if "" == param {
		ctx.JSON(200, rg.Error[string]("id不能为空"))
		return
	}
	ctx.JSON(200, c.sv.Detail(ctx, strPg.ToInt64(param)))
}

func (c *CountryController) Enable(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Enable(ctx, ct))
}

func (c *CountryController) Disable(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Disable(ctx, ct))
}

func (c *CountryController) State(ctx *gin.Context) {
	var ct model.BaseStateIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	state, ok := enumStatePg.IsExistInt64(ct.State)
	if !ok {
		ctx.JSON(200, rg.Error[string]("类型不正确"))
		return
	}
	ctx.JSON(200, c.sv.StateEnableDisable(ctx, ct.Ids, state))
}

func (c *CountryController) Query(ctx *gin.Context) {
	var ct modBasicCountry.QueryCt
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

func (c *CountryController) SelectPublic(ctx *gin.Context) {
	var ct modBasicCountry.QueryCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	if ct.State.ToInt8() <= 0 {
		ct.State = enumStatePg.ENABLE.IndexPg()
	}
	ctx.JSON(200, c.sv.SelectPublic(ctx, ct))
}

func (c *CountryController) SelectNodePublic(ctx *gin.Context) {
	var ct modBasicCountry.QueryPublicCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ct.State = enumStatePg.ENABLE.IndexPg()
	ctx.JSON(200, c.sv.SelectNodePublic(ctx, ct))
}
func (c *CountryController) SelectNodeAll(ctx *gin.Context) {
	var ct modBasicCountry.QueryPublicCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.SelectNodeAllPublic(ctx, ct))
}
func (c *CountryController) SelectNodeAllPublic(ctx *gin.Context) {
	var ct modBasicCountry.QueryPublicCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ct.State = enumStatePg.ENABLE.IndexPg()
	ctx.JSON(200, c.sv.SelectNodeAllPublic(ctx, ct))
}

func (c *CountryController) SelectPublicCountryCode(ctx *gin.Context) {
	var ct modBasicCountry.QueryPublicCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ct.State = enumStatePg.ENABLE.IndexPg()
	ctx.JSON(200, c.sv.SelectPublicCountryCode(ctx, ct))
}

func (c *CountryController) SelectNodePublicCountryCode(ctx *gin.Context) {
	var ct modBasicCountry.QueryPublicCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ct.State = enumStatePg.ENABLE.IndexPg()
	ctx.JSON(200, c.sv.SelectNodePublicCountryCode(ctx, ct))
}

func (c *CountryController) ExistName(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.ExistName(ctx, ct))
}

func (c *CountryController) ExistCode(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.ExistCode(ctx, ct))
}

func (c *CountryController) ExistCountryCode(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.ExistCountryCode(ctx, ct))
}
