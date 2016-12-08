package crypt

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/astaxie/beego"
)

func EnCode(msg string) string {
	h := md5.New()
	coding := base64.NewEncoding(beego.AppConfig.String("base64key"))
	h.Write([]byte(msg)) // 需要加密的字符串为 123456
	key := []byte(beego.AppConfig.String("md5key"))
	cipherStr := h.Sum([]byte(key))

	return coding.EncodeToString(cipherStr)
}

