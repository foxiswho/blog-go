package rdsPg

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-viper/mapstructure/v2"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"github.com/redis/go-redis/v9"
	"reflect"
)

func init() {
	gs.Provide(new(Get)).Init(func(s *Get) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// Get
// @Description: 获取缓存
type Get struct {
	log *log2.Logger  `autowire:"?"`
	rdb *redis.Client `autowire:""`
}

func NewGet(log *log2.Logger, rdb *redis.Client) *Get {
	return &Get{log: log, rdb: rdb}
}

// GetJson
//
//	@Description:
//	@receiver c
//	@param ctx
//	@param key
//	@param v
//	@return rt
func (c *Get) GetJson(ctx context.Context, key string) (rt rg.Rs[string]) {
	result, err := c.rdb.JSONGet(ctx, key).Result()
	if err != nil {
		c.log.Warnf("获取缓存失败: %+v", err.Error())
		if errors.Is(err, redis.Nil) {
			return rt.ErrorMessage("缓存不存在")
		}
		return rt.ErrorMessage("获取缓存错误")
	}
	return rt.OkData(result)
}

// GetJsonStruct
//
//	@Description:
//	@receiver c
//	@param ctx
//	@param key
//	@param v
//	@return rt
func (c *Get) GetJsonStruct(ctx context.Context, key string, v interface{}) (rt rg.Rs[string]) {
	result, err := c.rdb.JSONGet(ctx, key).Result()
	if err == redis.Nil {
		return rt.ErrorMessage("缓存不存在")
	}
	if err != nil {
		c.log.Warnf("获取缓存失败: %+v", err.Error())
		if errors.Is(err, redis.Nil) {
			return rt.ErrorMessage("缓存不存在")
		}
		return rt.ErrorMessage("获取缓存错误")
	}
	c.log.Debugf("缓存[key]=%+v,[data]=%+v", key, result)
	err = json.Unmarshal([]byte(result), &v)
	if err != nil {
		c.log.Errorf("缓存序列化失败: %+v", err.Error())
		return rt.ErrorMessage("缓存序列化失败")
	}
	return rt.OkData(result)
}

// GetJsonMapStruct
//
//	@Description:
//	@receiver c
//	@param ctx
//	@param key
//	@param v
//	@return rt
func (c *Get) GetJsonMapStruct(ctx context.Context, key string, v any) (rt rg.Rs[string]) {
	result, err := c.rdb.JSONGet(ctx, key).Result()
	if err == redis.Nil {
		return rt.ErrorMessage("缓存不存在")
	}
	if err != nil {
		c.log.Warnf("获取缓存失败: %+v", err.Error())
		if errors.Is(err, redis.Nil) {
			return rt.ErrorMessage("缓存不存在")
		}
		return rt.ErrorMessage("获取缓存错误")
	}
	c.log.Debugf("缓存[key]=%+v,[data]=%+v", key, result)
	var tmp map[string]interface{}
	err = json.Unmarshal([]byte(result), &tmp)
	if err != nil {
		c.log.Errorf("缓存序列化失败: %+v", err.Error())
		return rt.ErrorMessage("缓存序列化失败")
	}
	c.log.Debugf("缓存[key]=%+v,[map[string]interface{}]=%+v", key, tmp)
	//map 转为 struct
	if err = mapstructure.Decode(tmp, &v); err != nil {
		c.log.Errorf("map 转 struct err=%+v", err)
		return rt.ErrorMessage("缓存序列化失败")
	}
	return rt.OkData(result)
}

// GetString
//
//	@Description:
//	@receiver c
//	@param ctx
//	@param key
//	@return rt
func (c *Get) GetString(ctx context.Context, key string) (rt rg.Rs[string]) {
	result, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		c.log.Warnf("获取缓存失败: %+v", err.Error())
		if errors.Is(err, redis.Nil) {
			return rt.ErrorMessage("缓存不存在或已过期")
		}
		return rt.ErrorMessage("获取缓存错误")
	}
	return rt.OkData(result)
}

// GetStringToJson
//
//	@Description:
//	@receiver c
//	@param ctx
//	@param key
//	@param v   为指针对象
//	@return rt
func (c *Get) GetStringToJson(ctx context.Context, key string, v any) (rt rg.Rs[string]) {
	result, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		c.log.Warnf("获取缓存失败: %+v", err.Error())
		if errors.Is(err, redis.Nil) {
			return rt.ErrorMessage("缓存不存在或已过期")
		}
		return rt.ErrorMessage("获取缓存错误")
	}
	c.log.Debugf("缓存[key]=%+v,[data]=%+v", key, result)
	err = json.Unmarshal([]byte(result), v)
	if err != nil {
		c.log.Errorf("缓存序列化失败: %+v", err.Error())
		return rt.ErrorMessage("缓存序列化失败")
	}
	c.log.Debugf("缓存[key]=%+v,[data]=%+v", key, v)
	return rt.OkData(result)
}
