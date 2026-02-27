package rdsPg

import (
	"context"
	"errors"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/redis/go-redis/v9"
	"reflect"
)

func init() {
	gs.Provide(new(BatchString)).Init(func(s *BatchString) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BatchString struct {
	rdb *redis.Client `autowire:"?"`
	log *log2.Logger  `autowire:"?"`
}

func NewBatchString(
	log *log2.Logger,
	rdb *redis.Client,
) *BatchString {
	return &BatchString{
		log: log,
		rdb: rdb,
	}
}

func (t *BatchString) SetPipeline(ctx context.Context, keysValues map[string]interface{}) {
	pipeline := t.rdb.Pipeline()
	for key, value := range keysValues {
		pipeline.Set(ctx, key, value, 0) // 0 表示无过期时间
	}
	// 执行批量操作
	_, err := pipeline.Exec(ctx)
	if err != nil {
		t.log.Error("批量操作失败:", err)
		return
	}
}

func (t *BatchString) Get(ctx context.Context, key string) (string, bool) {
	result, err := t.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", false
		}
		t.log.Error("获取缓存失败:", err)
		return "", false
	}
	return result, true
}
