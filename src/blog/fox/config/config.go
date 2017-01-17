package config

import "github.com/astaxie/beego"

func String(key string) string {
	return beego.AppConfig.String(key)
}
func Bool(key string) (bool,error) {
	return beego.AppConfig.Bool(key)
}
func GetSection(section string) (map[string]string, error) {
	return beego.AppConfig.GetSection(section)
}
// GetConfig get the Appconfig
func GetConfig(returnType, key string, defaultVal interface{}) (value interface{}, err error) {
	return beego.GetConfig(returnType,key,defaultVal)
}