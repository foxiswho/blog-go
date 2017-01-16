package config

import "github.com/astaxie/beego"

func String(key string) string {
	return beego.AppConfig.String(key)
}