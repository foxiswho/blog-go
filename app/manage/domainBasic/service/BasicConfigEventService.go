package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicConfigEvent"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
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
	gs.Provide(new(BasicConfigEventService)).Init(func(s *BasicConfigEventService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicConfigEventService struct {
	log       *log2.Logger                                      `autowire:"?"`
	repoEvent *repositoryBasic.BasicConfigEventRepository       `autowire:"?"`
	repoField *repositoryBasic.BasicConfigEventFieldsRepository `autowire:"?"`
}

// Create 创建事件
func (s *BasicConfigEventService) Create(ctx *gin.Context, ct modBasicConfigEvent.CreateCt) (rt rg.Rs[string]) {
	if strPg.IsBlank(ct.Name) {
		return rt.ErrorMessage("中文名称不能为空")
	}
	if strPg.IsBlank(ct.Field) {
		return rt.ErrorMessage("英文标识不能为空")
	}
	if strPg.IsBlank(ct.Model) {
		return rt.ErrorMessage("模型英文标识不能为空")
	}

	// 检查事件是否已存在 (EventNo = Field)
	var existEvent entityBasic.BasicConfigEventEntity
	// Field 对应 EventNo 或者 Field 本身，这里假设 BasicConfigEventEntity.Field 存储事件的英文标识
	// 根据用户描述：Field 英文标识
	tx := s.repoEvent.Db().Where("field = ?", ct.Field).First(&existEvent)
	if tx.RowsAffected > 0 {
		return rt.ErrorMessage("事件标识已存在")
	}

	holder := holderPg.GetContextAccount(ctx)

	// 1. 创建事件头
	eventEntity := entityBasic.BasicConfigEventEntity{
		Name:      ct.Name,
		Field:     ct.Field, // 事件标识
		Model:     ct.Model, // 关联的模型
		ModelNo:   ct.Model, // 关联的模型编号
		ModuleSub: ct.ModuleSub,
		// TypeSys:      ct.TypeSys, // Entity中没有该字段
		Description: ct.Description,
		No:          noPg.No(), // 生成唯一编号
		TenantNo:    holder.GetTenantNo(),
		CreateBy:    holder.GetAccount().Account,
	}

	// Entity 中没有 TypeCategory, 只有 TypeSys 和 ModuleSub
	// KindUnique = ModelNo + Field (EventNo)
	eventEntity.KindUnique = eventEntity.ModelNo + eventEntity.Field

	err, _ := s.repoEvent.Create(&eventEntity)
	if err != nil {
		s.log.Errorf("Create event error: %v", err)
		return rt.ErrorMessage("创建事件失败")
	}

	// 2. 创建事件字段
	for _, fieldCt := range ct.Fields {
		fieldEntity := entityBasic.BasicConfigEventFieldsEntity{
			EventNo:         eventEntity.Field,
			ModelNo:         eventEntity.ModelNo,
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
		// FieldKindUnique = EventNo + Field
		fieldEntity.FieldKindUnique = fieldEntity.EventNo + fieldEntity.Field

		err, _ := s.repoField.Create(&fieldEntity)
		if err != nil {
			s.log.Errorf("Create event field error: %v", err)
			continue
		}
	}

	return rg.OkData(numberPg.Int64ToString(eventEntity.ID))
}
