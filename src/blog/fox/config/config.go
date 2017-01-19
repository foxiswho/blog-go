package config

import "github.com/astaxie/beego"
//此处 为以后 更换框架做准备
//获取配置 返回字符串
func String(key string) string {
	return beego.AppConfig.String(key)
}
//获取配置 返回布尔
func Bool(key string) (bool,error) {
	return beego.AppConfig.Bool(key)
}
//获取配置
func GetSection(section string) (map[string]string, error) {
	return beego.AppConfig.GetSection(section)
}
// GetConfig get the Appconfig
func GetConfig(returnType, key string, defaultVal interface{}) (value interface{}, err error) {
	return beego.GetConfig(returnType,key,defaultVal)
}