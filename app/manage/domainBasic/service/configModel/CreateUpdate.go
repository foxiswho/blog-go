package configModel

import (
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicConfigModel"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/configModelPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/enum/state/yesNoPg/yesNoIntPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model/modelBasePg"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

type CreateUpdate struct {
	Sp     Sp `autowire:"?"`
	ct     modBasicConfigModel.CreateUpdateCt
	model  *entityBasic.BasicConfigModelEntity
	log    *log2.Logger `autowire:"?"`
	module *entityBasic.BasicModuleEntity
	fields []*entityBasic.BasicConfigModelFieldsEntity
}

func (c *CreateUpdate) Process(ctx *gin.Context) (rt rg.Rs[string]) {
	c.fields = make([]*entityBasic.BasicConfigModelFieldsEntity, 0)
	return c.verify(ctx)
}

func (c *CreateUpdate) verify(ctx *gin.Context) (rt rg.Rs[string]) {
	header := c.ct.Header
	//
	err := copier.Copy(c.model, &header)
	if err != nil {
		c.log.Infof("copier.Copy error: %+v", err)
	}
	c.model.Name = strings.TrimSpace(c.model.Name)
	c.model.Model = strings.TrimSpace(c.model.Model)
	c.model.ModuleSub = strings.TrimSpace(c.model.ModuleSub)
	c.model.ModelCategory = strings.TrimSpace(c.model.ModelCategory)
	//
	if strPg.IsBlank(c.model.Name) {
		return rt.ErrorMessage("模型中文名称不能为空")
	}
	if strPg.IsBlank(c.model.Model) {
		return rt.ErrorMessage("模型英文标识不能为空")
	}
	if strPg.IsBlank(c.model.ModuleSub) {
		return rt.ErrorMessage("类型种类不能为空")
	}
	if _, ok := configModelPg.IsExistModelCategory(c.model.ModelCategory); !ok {
		return rt.ErrorMessage("类型种类不存在")
	}
	if strPg.IsBlank(header.ModuleSub) {
		return rt.ErrorMessage("子模块不能为空")
	}
	//子模块是否存在
	{
		b := false
		c.module, b = c.Sp.repModule.FindByCode(header.ModuleSub)
		if !b {
			return rt.ErrorMessage("子模块不存在")
		}
	}
	//
	//
	if nil != c.ct.Body && len(c.ct.Body) > 0 {
		errs := make([]modelBasePg.ItemResult, 0)
		rt.Extend = make(map[string]interface{})
		i := int64(0)
		//
		uniqueField := make(map[string]bool)
		//
		for _, item := range c.ct.Body {
			i++

			var field entityBasic.BasicConfigModelFieldsEntity
			err := copier.Copy(&field, &item)
			if err != nil {
				c.log.Infof("copier.Copy error: %+v", err)
				errs = append(errs, modelBasePg.ItemResult{
					Col:   i,
					Field: "col",
					Msg:   "整行格式转换错误",
				})
				continue
			}
			field.Name = strings.TrimSpace(field.Name)
			field.Field = strings.TrimSpace(field.Field)
			if strPg.IsBlank(field.Name) {
				errs = append(errs, modelBasePg.ItemResult{
					Col:   i,
					Field: "field",
					Msg:   "字段中文名称不能为空",
				})
				continue
			}
			if strPg.IsBlank(field.Field) {
				errs = append(errs, modelBasePg.ItemResult{
					Col:   i,
					Field: "field",
					Msg:   "字段英文标识不能为空",
				})
				continue
			}
			if _, ok := uniqueField[field.Field]; ok {
				errs = append(errs, modelBasePg.ItemResult{
					Col:   i,
					Field: "field",
					Msg:   "字段英文标识重复",
				})
				continue
			}
			uniqueField[field.Field] = true
			field.Show = yesNoIntPg.No.IndexInt8()
			if yesNoIntPg.Yes.IsEqual(item.Show.ToInt8()) {
				field.Show = yesNoIntPg.Yes.IndexInt8()
			}
			field.Binary = yesNoIntPg.No.IndexInt8()
			if yesNoIntPg.Yes.IsEqual(item.Binary.ToInt8()) {
				field.Binary = yesNoIntPg.Yes.IndexInt8()
			}
			c.fields = append(c.fields, &field)
		}
	}
	//
	//
	if header.Id.ToInt64() <= 0 {
		c.model.State = enumStatePg.ENABLE.Index()
		c.model.TypeSys = typeSysPg.General.String()
		c.model.No = noPg.No()
		c.model.Sort = 0
		err, _ := c.Sp.repModel.Create(c.model)
		if err != nil {
			return rt.ErrorMessage("保存失败 " + err.Error())
		}
		//

	} else {
		info, result := c.Sp.repModel.FindByIdString(header.Id.ToString())
		if !result {
			return rt.ErrorMessage("模型不存在")
		}
		var save entityBasic.BasicConfigModelEntity
		{
			err := copier.Copy(&save, info)
			if err != nil {
				c.log.Infof("copier.Copy error: %+v", err)
			}
		}
		{
			err := copier.Copy(&save, &header)
			if err != nil {
				c.log.Infof("copier.Copy error: %+v", err)
			}
		}
		err := c.Sp.repModel.Update(save, info.ID)
		if err != nil {
			c.log.Infof("copier.Copy error: %+v", err)
		}
	}

	return rt.Ok()
}
