package table

import (
	"context"
	"time"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityApi"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityBasic"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityBlog"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityTc"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/configPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/dbMakePg"
	"go-spring.org/log"
	_ "go-spring.org/spring/gs"
	"gorm.io/gorm"
)

// AInitTable 初始化创建表
type AInitTable struct {
	ser      configPg.Server   `value:"${server}"`
	database configPg.Database `value:"${database}"`
	log      *log2.Logger      `autowire:"?"`
	db       *gorm.DB          `autowire:"?"`
}

func (b *AInitTable) Run(ctx context.Context) error {
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
		entityData = append(entityData, &entityBasic.BasicConfigEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigEventEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigEventFieldsEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigListEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigModelEntity{})
		entityData = append(entityData, &entityBasic.BasicConfigModelFieldsEntity{})
		entityData = append(entityData, &entityBasic.BasicDataDictionaryEntity{})
		entityData = append(entityData, &entityBasic.BasicDataSnapshotEntity{})
		entityData = append(entityData, &entityBasic.BasicModelRulesEntity{})
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
		entityData = append(entityData, &entityRam.RamJobFunctionEntity{})
		entityData = append(entityData, &entityRam.RamPositionEntity{})
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
		//插入 基础数据
		// b.sqlBasicDictionary()
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
	sql = append(sql, dbMakePg.MakeSequenceSql("basic_config_list", 100000))
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

func (b *AInitTable) sqlBasicDictionary() {
	sql := ""
	sql = "INSERT INTO \"public\".\"basic_data_dictionary\" (\"id\", \"no\", \"name\", \"name_fl\", \"code\", \"name_full\", \"state\", \"description\", \"create_at\", \"update_at\", \"create_by\", \"update_by\", \"sort\", \"type_unique_md5\", \"value\", \"extend\", \"range\", \"type_code\", \"tenant_no\", \"store_no\", \"owner_no\") VALUES\n(1, '1', '身份类型', NULL, 'typeIdentity', NULL, 1, NULL, '2026-04-28T05:10:52+08:00', '2026-05-04T08:10:30+08:00', NULL, NULL, 0, '7e3a158550d82e66c9f016977dcbec61', 'typeIdentity', NULL, NULL, '', NULL, NULL, NULL),\n(2, '202605032139333761667562', '普通', NULL, 'general', NULL, 1, NULL, '2026-05-03T13:39:33+08:00', '2026-05-04T08:11:36+08:00', NULL, NULL, 0, '958153f1b8b96ec4c4eb2147429105d9', 'general', NULL, NULL, 'typeIdentity', NULL, NULL, NULL),\n(3, '202605032140382912169335', '经理', NULL, 'manager', NULL, 1, NULL, '2026-05-03T13:40:38+08:00', '2026-05-04T08:11:50+08:00', NULL, NULL, 0, '1d0258c2440a8d19e716292b231e3190', 'manager', NULL, NULL, 'typeIdentity', NULL, NULL, NULL),\n(4, '202605032141103290247912', '性别', NULL, 'sex', NULL, 1, NULL, '2026-05-03T13:41:10+08:00', '2026-05-04T08:14:24+08:00', NULL, NULL, 0, '3c3662bcb661d6de679c636744c66b62', 'sex', NULL, NULL, '', NULL, NULL, NULL),\n(5, '5', '副经理', NULL, 'assistant_manager', NULL, 1, NULL, '2026-05-04T08:13:48+08:00', NULL, NULL, NULL, 0, 'adb153a9579779a6837a8205e07fa31f', 'assistant_manager', NULL, NULL, 'typeIdentity', NULL, NULL, NULL),\n(6, '6', '男', NULL, 'male', NULL, 1, NULL, '2026-05-04T08:17:08+08:00', NULL, NULL, NULL, 0, '07cf4f8f5d8b76282917320715dda2ad', 'male', NULL, NULL, 'sex', NULL, NULL, NULL),\n(7, '7', '女', NULL, 'female', NULL, 1, NULL, '2026-05-04T08:17:20+08:00', NULL, NULL, NULL, 0, '273b9ae535de53399c86a9b83148a8ed', 'female', NULL, NULL, 'sex', NULL, NULL, NULL),\n(8, '8', '未知', NULL, 'unknown', NULL, 1, NULL, '2026-05-04T08:17:30+08:00', NULL, NULL, NULL, 0, 'ad921d60486366258809553a3db49a4a', 'unknown', NULL, NULL, 'sex', NULL, NULL, NULL),\n(9, '9', '菜单类型', NULL, 'typeMenu', NULL, 1, NULL, '2026-05-15T01:16:09+08:00', NULL, NULL, NULL, 0, '0e6b09acf0d718df71494094446eb4c0', 'typeMenu', NULL, NULL, '', NULL, NULL, NULL),\n(10, '10', '目录', NULL, 'catalog', NULL, 1, NULL, '2026-05-15T01:16:46+08:00', NULL, NULL, NULL, 0, '46f22f2a56ddd091f4b2b2c35c5ca989', 'catalog', NULL, NULL, 'typeMenu', NULL, NULL, NULL),\n(11, '11', '菜单', NULL, 'menu', NULL, 1, NULL, '2026-05-15T01:18:15+08:00', '2026-05-15T01:19:32+08:00', NULL, NULL, 0, '8d6ab84ca2af9fccd4e4048694176ebf', 'menu', NULL, NULL, 'typeMenu', NULL, NULL, NULL),\n(12, '12', '内嵌', NULL, 'embedded', NULL, 1, NULL, '2026-05-15T01:18:33+08:00', NULL, NULL, NULL, 0, '605abe26d014c72e3df9deb267e73756', 'embedded', NULL, NULL, 'typeMenu', NULL, NULL, NULL),\n(13, '13', '按钮', NULL, 'button', NULL, 1, NULL, '2026-05-15T01:19:50+08:00', NULL, NULL, NULL, 0, 'ce50a09343724eb82df11390e2c1de18', 'button', NULL, NULL, 'typeMenu', NULL, NULL, NULL),\n(14, '14', '外链', NULL, 'link', NULL, 1, NULL, '2026-05-15T01:20:08+08:00', NULL, NULL, NULL, 0, '2a304a1348456ccd2234cd71a81bd338', 'link', NULL, NULL, 'typeMenu', NULL, NULL, NULL),\n(15, '15', '短信服务商', NULL, 'spCodeSms', NULL, 1, NULL, '2026-05-28T07:35:20+08:00', NULL, NULL, NULL, 0, '5923f12d736703f67a352b8a20c14d61', 'spCodeSms', NULL, NULL, '', NULL, NULL, NULL),\n(16, '16', '邮件服务商', NULL, 'spCodeMail', NULL, 1, NULL, '2026-05-28T07:35:40+08:00', NULL, NULL, NULL, 0, '4ffa88124b350a0d4c731e2bdd04721c', 'spCodeMail', NULL, NULL, '', NULL, NULL, NULL),\n(17, '17', '阿里云', NULL, 'aliyun', NULL, 1, NULL, '2026-05-28T07:35:58+08:00', NULL, NULL, NULL, 0, 'd41b8265807ce33038df2b3f7aa9fb10', 'aliyun', NULL, NULL, 'spCodeSms', NULL, NULL, NULL),\n(35, '202606121526467480058808', '终端类型', NULL, 'terminalCode', NULL, 1, NULL, '2026-06-12T07:26:45+08:00', NULL, NULL, NULL, 0, '2345a28058000459ae6173b7c66286bb', 'terminalCode', NULL, NULL, '', NULL, NULL, NULL),\n(36, '202606121527015741949205', '系统', NULL, 'system', NULL, 1, NULL, '2026-06-12T07:27:00+08:00', NULL, NULL, NULL, 0, '54b53072540eeeb8f8e9343e71f28176', 'system', NULL, NULL, 'terminalCode', NULL, NULL, NULL),\n(37, '202606121527254824084333', '管理', NULL, 'manage', NULL, 1, NULL, '2026-06-12T07:27:24+08:00', NULL, NULL, NULL, 0, '70682896e24287b0476eff2a14c148f0', 'manage', NULL, NULL, 'terminalCode', NULL, NULL, NULL),\n(38, '202606121528031104440704', '驾驶舱', NULL, 'cockpit', NULL, 1, NULL, '2026-06-12T07:28:01+08:00', NULL, NULL, NULL, 0, '356834670f264c2f9f1cda27bdb35f3b', 'cockpit', NULL, NULL, 'terminalCode', NULL, NULL, NULL);"
	err := b.db.Transaction(func(tx *gorm.DB) error {
		rs := tx.Exec(sql)
		if rs.Error != nil {
			log.Errorf(context.Background(), log.TagAppDef, "初始化序号异常:%+v", rs.Error)
			return rs.Error
		}
		log.Debugf(context.Background(), log.TagAppDef, "执行结果: %+v 行受影响", rs.RowsAffected)
		return nil
	})
	if err != nil {
		log.Errorf(context.Background(), log.TagAppDef, "创建表异常:%+v", err)
	}
}
