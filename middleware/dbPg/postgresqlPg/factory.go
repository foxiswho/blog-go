package postgresqlPg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库工厂 Postgresql
type Factory struct {
	database configPg.Database `value:"${database}"`
	log      *log2.Logger      `autowire:"?"`
}

func (factory *Factory) CreateDB() (*gorm.DB, error) {
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[Postgresql]===================")
	syslog.Debugf(context.Background(), syslog.TagAppDef, "数据库连接地址： enabled:%+v,URL:=>%s", factory.database.Enabled, factory.database.URL)
	//fmt.Printf("数据库连接地址： enabled:%+v,URL:=>%s", factory.database.Enabled, factory.database.URL)
	if !factory.database.Enabled {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "未启用数据库： %s", factory.database.Enabled)
		return nil, nil
	}
	db, err := gorm.Open(postgres.Open(factory.database.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		syslog.Errorf(context.Background(), syslog.TagAppDef, "open gorm postgresql %s error: %v", factory.database.URL, err)
		panic(errors.New(err.Error()))
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		factory.log.Error("failed to get database connection")
	}
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大生存时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	//
	dbName := db.Migrator().CurrentDatabase()
	// check if dbPg exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", dbName)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return nil, rs.Error
	}
	// if not create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if rs := db.Exec(stmt); rs.Error != nil {
			return nil, rs.Error
		}

		// close dbPg connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func init() {
	gs.Object(&Factory{})
	gs.Provide((*Factory).CreateDB).
		//指定名称
		Name("GormDb").
		//当指定类型/名称的 Bean 不存在时激活
		Condition(
			gs.OnProperty("database.enabled").HavingValue("true").MatchIfMissing(),
			// GormDB 不存在
			gs.OnMissingBean[*gorm.DB]("GormDB"))
}
