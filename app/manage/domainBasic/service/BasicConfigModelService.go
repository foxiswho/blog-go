package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicConfigModel"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/configModelPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(BasicConfigModelService)).Init(func(s *BasicConfigModelService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicConfigModelService struct {
	log       *log2.Logger                                      `autowire:"?"`
	repoModel *repositoryBasic.BasicConfigModelRepository       `autowire:"?"`
	repoField *repositoryBasic.BasicConfigModelFieldsRepository `autowire:"?"`
	repoRule  *repositoryBasic.BasicConfigModelRulesRepository  `autowire:"?"`
}

// Create 创建模型
func (s *BasicConfigModelService) Create(ctx *gin.Context, ct modBasicConfigModel.CreateCt) (rt rg.Rs[string]) {
	if strPg.IsBlank(ct.Name) {
		return rt.ErrorMessage("中文名称不能为空")
	}
	if strPg.IsBlank(ct.Model) {
		return rt.ErrorMessage("英文标识不能为空")
	}

	// 检查模型是否已存在
	// 假设 Model 字段是唯一的英文标识
	// 需要在 Repository 中实现 FindByModel 或者使用 Where 查询
	// 这里使用 Where 查询
	var existModel entityBasic.BasicConfigModelEntity
	tx := s.repoModel.Db().Where("model = ?", ct.Model).First(&existModel)
	if tx.RowsAffected > 0 {
		return rt.ErrorMessage("模型标识已存在")
	}

	holder := holderPg.GetContextAccount(ctx)

	// 1. 创建模型头
	modelEntity := entityBasic.BasicConfigModelEntity{
		Name:          ct.Name,
		Model:         ct.Model,
		TypeSys:       ct.TypeSys,
		ModelCategory: ct.TypeCategory,
		ModuleSub:     ct.ModuleSub,
		Description:   ct.Description,
		No:            noPg.No(), // 生成编号
		TenantNo:      holder.GetTenantNo(),
		CreateBy:      holder.GetAccount().Account,
	}
	if strPg.IsNotBlank(modelEntity.TypeSys) {
		modelEntity.TypeSys = typeSysPg.General.String()
	}
	if modelEntity.ModelCategory == "" {
		modelEntity.ModelCategory = configModelPg.ModelCategoryTable.String()
	}

	err, _ := s.repoModel.Create(&modelEntity)
	if err != nil {
		s.log.Errorf("Create model error: %v", err)
		return rt.ErrorMessage("创建模型失败")
	}

	// 2. 创建字段和规则
	for _, fieldCt := range ct.Fields {
		fieldEntity := entityBasic.BasicConfigModelFieldsEntity{
			ModelNo:         modelEntity.Model, // 使用 Model 标识作为关联
			Name:            fieldCt.Name,
			Field:           fieldCt.Field,
			Show:            fieldCt.Show,
			Binary:          fieldCt.Binary,
			DefaultValue:    fieldCt.DefaultValue,
			ValueType:       fieldCt.ValueType,
			FormCode:        fieldCt.FormCode,
			ParameterSource: fieldCt.ParameterSource,
			TenantNo:        holder.GetTenantNo(),
			CreateBy:        holder.GetAccount().Account,
		}
		// 生成字段唯一标识 KindUnique = model_no + field
		fieldEntity.KindUnique = modelEntity.Model + fieldEntity.Field

		err, _ := s.repoField.Create(&fieldEntity)
		if err != nil {
			s.log.Errorf("Create field error: %v", err)
			// 继续创建其他字段，或者回滚（这里简化处理，暂不回滚）
			continue
		}

		// 3. 创建规则
		if strPg.IsNotBlank(fieldCt.RuleMode) {
			ruleEntity := entityBasic.BasicConfigModelRulesEntity{
				ModelNo:     modelEntity.Model,
				Field:       fieldCt.Field,
				RuleMode:    fieldCt.RuleMode,
				Description: fieldCt.Name + "规则",
				TenantNo:    holder.GetTenantNo(),
				CreateBy:    holder.GetAccount().Account,
			}
			s.repoRule.Create(&ruleEntity)
		}
	}

	return rg.OkData(numberPg.Int64ToString(modelEntity.ID))
}
