package table

import (
	"context"
	"time"

	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/entityTc"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/dbMakePg"
	"github.com/go-spring/log"
	"gorm.io/gorm"
)

// AInitTable 初始化创建表
type AInitTable struct {
	ser      configPg.Server   `value:"${server}"`
	database configPg.Database `value:"${database}"`
	log      *log2.Logger      `autowire:"?"`
	db       *gorm.DB          `autowire:"?"`
}

func (b *AInitTable) Run() error {
	log.Infof(context.Background(), log.TagAppDef, "初始化表=>表不存在时,则进行初始化")
	entityData := make([]interface{}, 0)
	{
		entityData = append(entityData, &entityApi.ApiDiplEntity{})
		entityData = append(entityData, &entityApi.ApiDiplAccessKeyEntity{})
		entityData = append(entityData, &entityApi.ApiDiplCategoryEntity{})
	}
	{
		entityData = append(entityData, &entityBasic.BasicAccountApplyDenyListEntity{})
		entityData = append(entityData, &entityBasic.BasicAreaEntity{})
		entityData = append(entityData, &entityBasic.BasicAttachmentEntity{})
		entityData = append(entityData, &entityBasic.BasicCountryEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigEventEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigEventFieldsEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigModelEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigModelFieldsEntity{})
		entityData = append(entityData, &entityBasic.BasicDataDictionaryEntity{})
		entityData = append(entityData, &entityBasic.BasicDataSnapshotEntity{})
		entityData = append(entityData, &entityBasic.BasicModuleEntity{})
		entityData = append(entityData, &entityBasic.BasicTagsEntity{})
		entityData = append(entityData, &entityBasic.BasicTagsCategoryEntity{})
		entityData = append(entityData, &entityBasic.BasicTagsRelationEntity{})
	}
	{
		entityData = append(entityData, &entityBlog.BlogArticleEntity{})
		entityData = append(entityData, &entityBlog.BlogArticleCategoryEntity{})
		entityData = append(entityData, &entityBlog.BlogArticleStatisticsEntity{})
		entityData = append(entityData, &entityBlog.BlogCollectEntity{})
		entityData = append(entityData, &entityBlog.BlogCollectCategoryEntity{})
		entityData = append(entityData, &entityBlog.BlogTopicEntity{})
		entityData = append(entityData, &entityBlog.BlogTopicCategoryEntity{})
		entityData = append(entityData, &entityBlog.BlogTopicRelationEntity{})
		entityData = append(entityData, &entityBlog.BlogTopicStatisticsEntity{})
	}
	//
	{
		entityData = append(entityData, &entityRam.RamAccountEntity{})
		entityData = append(entityData, &entityRam.RamAccountAuthorizationEntity{})
		entityData = append(entityData, &entityRam.RamAccountDenyListEntity{})
		entityData = append(entityData, &entityRam.RamAccountDeviceEntity{})
		entityData = append(entityData, &entityRam.RamAccountLoginLogEntity{})
		entityData = append(entityData, &entityRam.RamAccountSessionEntity{})
		entityData = append(entityData, &entityRam.RamAccountSessionAccessKeyEntity{})
		entityData = append(entityData, &entityRam.RamAppEntity{})
		entityData = append(entityData, &entityRam.RamAppAccessKeyEntity{})
		entityData = append(entityData, &entityRam.RamAppCategoryEntity{})
		entityData = append(entityData, &entityRam.RamChannelEntity{})
		entityData = append(entityData, &entityRam.RamDepartmentEntity{})
		entityData = append(entityData, &entityRam.RamFavoritesEntity{})
		entityData = append(entityData, &entityRam.RamGroupEntity{})
		entityData = append(entityData, &entityRam.RamLevelEntity{})
		entityData = append(entityData, &entityRam.RamMenuEntity{})
		entityData = append(entityData, &entityRam.RamMenuRelationEntity{})
		entityData = append(entityData, &entityRam.RamPositionEntity{})
		entityData = append(entityData, &entityRam.RamPostEntity{})
		entityData = append(entityData, &entityRam.RamResourceEntity{})
		entityData = append(entityData, &entityRam.RamResourceAuthorityEntity{})
		entityData = append(entityData, &entityRam.RamResourceGroupEntity{})
		entityData = append(entityData, &entityRam.RamResourceGroupRelationEntity{})
		entityData = append(entityData, &entityRam.RamResourceMenuEntity{})
		entityData = append(entityData, &entityRam.RamResourceRelationEntity{})
		entityData = append(entityData, &entityRam.RamRoleEntity{})
		entityData = append(entityData, &entityRam.RamTeamEntity{})
	}
	//
	{
		entityData = append(entityData, &entityTc.TcLevelEntity{})
		entityData = append(entityData, &entityTc.TcTenantEntity{})
		entityData = append(entityData, &entityTc.TcTenantDomainEntity{})
	}

	//初始化创建表
	sv := &dbMakePg.CreateTable{
		Database: b.database,
		Log:      b.log,
	}
	rt := sv.DbOpen()
	if rt.SuccessIs() {
		sv.TableCreateAllByTransaction(entityData)
		//
		log.Infof(context.Background(), log.TagAppDef, "[init].[主键序号保留].")
		b.seqEdit()
		log.Infof(context.Background(), log.TagAppDef, "初始化表 successfully")
	} else {
		log.Errorf(context.Background(), log.TagAppDef, "初始化表异常", rt.Error())
	}
	sv = nil
	return nil
}

// seqEdit
//
//	@Description: 序号修改 初始值
//	@receiver b
func (b *AInitTable) seqEdit() {
	sql := make([]string, 0)
	sql = append(sql, dbMakePg.MakeSequenceSql("api_dipl", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_account_apply_deny_list", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_area", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_country", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_data_dictionary", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_tags", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_tags_category", 100000))
	//
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_account", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_account_authorization", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_app", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_department", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_group", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_level", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_position", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_post", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("ram_team", 100000))
	//
	sql = append(sql, dbMakePg.MakeSequenceSql("tc_tenant", 100000))
	sql = append(sql, dbMakePg.MakeSequenceSql("tc_tenant_domain", 100000))
	//
	err := b.db.Transaction(func(tx *gorm.DB) error {
		for _, raw := range sql {
			rs := tx.Exec(raw)
			if rs.Error != nil {
				log.Errorf(context.Background(), log.TagAppDef, "初始化序号异常:%+v", rs.Error)
			}
			log.Debugf(context.Background(), log.TagAppDef, "执行结果: %+v 行受影响", rs.RowsAffected)
			time.Sleep(time.Microsecond * 10)
		}
		return nil
	})
	if err != nil {
		log.Errorf(context.Background(), log.TagAppDef, "创建表异常:%+v", err)
	}
}
