package config

import (
	"github.com/beego/beego/v2/core/config"
)

//此处 为以后 更换框架做准备
//获取配置 返回字符串
func String(key string) string {
	s, err := config.String(key)
	if err != nil {
		return ""
	}
	return s
}

//获取配置 返回布尔
func Bool(key string) (bool, error) {
	return config.Bool(key)
}

//获取配置
func GetSection(section string) (map[string]string, error) {
	return config.GetSection(section)
}

// GetConfig get the Appconfig
func GetConfig(returnType, key string, defaultVal interface{}) (value interface{}, err error) {
	section, err := config.GetSection(returnType)
	if err != nil {
		return nil, err
	}
	if section[key] != "" {
		return section[key], nil
	}
	return defaultVal, nil
}
