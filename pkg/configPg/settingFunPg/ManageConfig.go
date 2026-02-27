package settingFunPg

import (
	"context"
	"encoding/json"

	"github.com/foxiswho/blog-go/pkg/configPg/settingPg"
	"github.com/go-spring/log"
)

// ManageConfigKey
//
//	@Description: 获取配置key
//	@param key
//	@param val
//	@return interface{}
func ManageConfigKey(key string, val ...any) (interface{}, bool) {
	return key, true
}

// ManageConfigStructMapping
//
//	@Description:  管理配置结构映射
//	@param key
//	@param data
//	@return any
//	@return bool
func ManageConfigStructMapping(key string, data []byte) (any, bool) {
	var info settingPg.ManageConfig
	err := json.Unmarshal(data, &info)
	if nil != err {
		log.Errorf(context.Background(), log.TagAppDef, "json.error=%+v", err.Error())
		return nil, false
	}
	return key, true
}
