package redisPg

import (
	"context"
	"fmt"
	"time"

	_ "github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	g "github.com/redis/go-redis/v9"
)

func init() {
	gs.Object(&Factory{})
	gs.Provide((*Factory).Open).
		//指定名称
		Name("RedisClient").
		//当指定类型/名称的 Bean 不存在时激活
		Condition(
			gs.OnProperty("pg.redis.enabled").HavingValue("true").MatchIfMissing(),
			// RedisClient 不存在
			gs.OnMissingBean[*g.Client]("RedisClient"))
}

// Factory 数据库工厂 redis
type Factory struct {
	pg  configPg.Pg  `value:"${pg}"`
	log *log2.Logger `autowire:"?"`
}

func (factory *Factory) Open() (*g.Client, error) {
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[redis]===================")
	config := factory.pg.Redis
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := g.NewClient(&g.Options{
		Addr:            address,
		Username:        config.Username,
		Password:        config.Password,
		DB:              config.Database,
		DialTimeout:     time.Duration(config.ConnectTimeout) * time.Millisecond,
		ReadTimeout:     time.Duration(config.ReadTimeout) * time.Millisecond,
		WriteTimeout:    time.Duration(config.WriteTimeout) * time.Millisecond,
		ConnMaxIdleTime: time.Duration(config.IdleTimeout) * time.Millisecond,
	})
	if config.Ping {
		if err := client.Ping(context.Background()).Err(); err != nil {
			return nil, err
		}
	}
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[redis]=>连接成功..===================")
	return client, nil
}
