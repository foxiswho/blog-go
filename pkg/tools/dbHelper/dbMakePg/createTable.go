package dbMakePg

import (
	"context"
	"time"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/tablePg"
	"github.com/go-spring/log"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type CreateTable struct {
	Database  configPg.Database `value:"database"`
	Log       *log2.Logger      `autowire:"?"`
	db        *gorm.DB
	tableList []string
}

// DbOpen
//
//	@Description: 打开数据库
//	@receiver c
//	@return rt
func (c *CreateTable) DbOpen() (rt rg.Rs[string]) {
	db, err := gorm.Open(postgres.Open(c.Database.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(200)
		sqlDB.SetConnMaxLifetime(0)
	}
	c.db = db
	if err != nil {
		log.Errorf(context.Background(), log.TagAppDef, "open gorm postgresql %s error: %v", c.Database.URL, err)
		return rt.ErrorMessage(err.Error())
	}
	return rt.Ok()
}

// SetDb
//
//	@Description: 打开数据库
//	@receiver c
//	@return rt
func (c *CreateTable) SetDb(db *gorm.DB) {
	c.db = db
}

func (c *CreateTable) TableCreateOne(entity interface{}) (rt rg.Rs[string]) {
	isExists := c.db.Migrator().HasTable(entity)
	if isExists {
		c.Log.Infof("表存在")
	} else {
		err2 := c.db.AutoMigrate(entity)
		if err2 != nil {
			c.Log.Errorf("创建表异常", err2)
			return
		}
	}

	return rt.Ok()
}

// 生成表，创建表注释
func (c *CreateTable) dbRun(tmp interface{}, tableName, tableComment string) {
	// 判断表是否已创建
	if !slice.Contain(c.tableList, tableName) {
		c.Log.Infof("创建表 %s", tableName)
		MakeTable(c.db, tmp, tableName, tableComment)

		c.Log.Infof("创建表 %s [完成]", tableName)
	} else {
		c.Log.Infof("表已存在 %s", tableName)
	}
}

// 生成表，创建表注释
func (c *CreateTable) dbRunByDb(db *gorm.DB, tmp interface{}, tableName, tableComment string) {
	// 判断表是否已创建
	if !slice.Contain(c.tableList, tableName) {
		log.Infof(context.Background(), log.TagAppDef, "创建表 %s", tableName)
		MakeTable(db, tmp, tableName, tableComment)

		log.Infof(context.Background(), log.TagAppDef, "创建表 %s [完成]", tableName)
	} else {
		log.Infof(context.Background(), log.TagAppDef, "表已存在 %s", tableName)
	}
}

// TableCreateAll
//
//	@Description: 批量创建表
//	@receiver c
//	@param dst
//	@return rt
func (c *CreateTable) TableCreateAll(dst []interface{}) (rt rg.Rs[string]) {
	c.tableList = GetTablesByTime3Minute(c.db)
	log.Debugf(context.Background(), log.TagAppDef, "已存在的表: %+v", c.tableList)
	i := 0
	for _, item := range dst {
		name := tablePg.GetTableName(item)
		log.Debugf(context.Background(), log.TagAppDef, "创建表: %s", name)
		comment := tablePg.GetTableComment(item)
		if strPg.IsNotBlank(name) {
			c.dbRun(item, name, comment)
		}
		if i%3 == 0 {
			time.Sleep(10000)
		}
		i++
	}

	return rt.Ok()
}

// TableCreateAllByTransaction
//
//	@Description: 批量创建表
//	@receiver c
//	@param dst
//	@return rt
func (c *CreateTable) TableCreateAllByTransaction(dst []interface{}) (rt rg.Rs[string]) {
	c.tableList = GetTablesByTime3Minute(c.db)
	log.Debugf(context.Background(), log.TagAppDef, "已存在的表: %+v", c.tableList)
	//
	err := c.db.Transaction(func(tx *gorm.DB) error {
		i := 0
		for _, item := range dst {
			name := tablePg.GetTableName(item)
			log.Debugf(context.Background(), log.TagAppDef, "创建表: %s", name)
			comment := tablePg.GetTableComment(item)
			if strPg.IsNotBlank(name) {
				c.dbRunByDb(tx, item, name, comment)
			}
			if i%3 == 0 {
				time.Sleep(10000)
			}
			i++
		}

		return nil
	})
	if err != nil {
		log.Errorf(context.Background(), log.TagAppDef, "创建表异常 %+v", err)
	}

	return rt.Ok()
}
