package rdsPg

import (
	"context"
	"errors"
	"reflect"

	"github.com/foxiswho/blog-go/pkg/consts/constBlogPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/redis/go-redis/v9"
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
func (t *BatchString) GetRdb() *redis.Client {
	return t.rdb
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

func (t *BatchString) GetAllByKeys(ctx context.Context, key []string) ([]interface{}, bool) {
	result, err := t.rdb.MGet(ctx, key...).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false
		}
		t.log.Error("获取缓存失败:", err)
		return nil, false
	}
	return result, true
}

func (t *BatchString) GetAllEvalByLua(ctx context.Context, key []string) ([]interface{}, bool) {
	resp, err := t.rdb.Eval(ctx, constBlogPg.ArticleCategoryLua, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false
		}
		t.log.Error("获取缓存失败:", err)
		return nil, false
	}
	// 解析返回结果
	result, ok := resp.([]interface{})
	if !ok {
		t.log.Error("获取缓存失败:返回结果格式错误，预期为数组类型:", err)
		return nil, false
	}
	return result, true
}
