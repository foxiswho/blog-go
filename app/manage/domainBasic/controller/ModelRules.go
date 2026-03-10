package controller

import (
	"fmt"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicModelRules"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {

}

// ModelRulesController 国家
// @Description:
type ModelRulesController struct {
	Sp *authPg.GroupManageMiddlewareSp `autowire:""`
	sv *service.BasicModelRulesService `autowire:"?"`
}

// CreateUpdate 创建
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) CreateUpdate(ctx *gin.Context) {
	var ct modBasicModelRules.CreateUpdateDataCt
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
	ctx.JSON(200, c.sv.CreateUpdate(ctx, ct))
}

// Delete 逻辑删除
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) Delete(ctx *gin.Context) {
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
	ctx.JSON(200, c.sv.LogicalDeletion(ctx, ct.Ids))
}

// Recovery 逻辑删除恢复
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) Recovery(ctx *gin.Context) {
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
	ctx.JSON(200, c.sv.LogicalRecovery(ctx, ct.Ids))
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) PhysicalDeletion(ctx *gin.Context) {
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

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) Detail(ctx *gin.Context) {
	param := ctx.Param("id")

	fmt.Println(param)
	if "" == param {
		ctx.JSON(200, rg.Error[string]("id不能为空"))
		return
	}
	//ctx.JSON(200, c.sv.Detail(ctx, strPg.ToInt64(param)))
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) Enable(ctx *gin.Context) {
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
	ctx.JSON(200, c.sv.Enable(ctx, ct))
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) Disable(ctx *gin.Context) {
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
	ctx.JSON(200, c.sv.Disable(ctx, ct))
}

// State 状态
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) State(ctx *gin.Context) {
	var ct model.BaseStateIdsCt[string]
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
	state, ok := enumStatePg.IsExistInt64(ct.State)
	if !ok {
		ctx.JSON(200, rg.Error[string]("类型不正确"))
		return
	}
	ctx.JSON(200, c.sv.StateEnableDisable(ctx, ct.Ids, state))
}

// Query 查询列表
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) Query(ctx *gin.Context) {
	var ct modBasicModelRules.QueryCt
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

func (c *ModelRulesController) AllByValueNo(ctx *gin.Context) {
	var ct modBasicModelRules.QueryPublicCt
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
	ct.State = enumStatePg.ENABLE.IndexPg()
	ctx.JSON(200, c.sv.AllByValueNo(ctx, ct))
}

func (c *ModelRulesController) SelectNodeAllPublic(ctx *gin.Context) {
	var ct modBasicModelRules.QueryPublicCt
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
	ct.State = enumStatePg.ENABLE.IndexPg()
	ctx.JSON(200, c.sv.SelectNodeAllPublic(ctx, ct))
}

// ExistName 查重
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) ExistName(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
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
	ctx.JSON(200, c.sv.ExistName(ctx, ct))
}

// ExistCode 查重
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ModelRulesController) ExistCode(ctx *gin.Context) {
	var ct model.BaseExistWdCt[string]
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
	ctx.JSON(200, c.sv.ExistCode(ctx, ct))
}
