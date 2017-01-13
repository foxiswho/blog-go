package file

import (
	"blog/fox/str"
	"blog/fox/crypt"
	"github.com/astaxie/beego"
	"encoding/base64"
	"blog/fox"
	"blog/fox/array"
)
//令牌生成
//@maps 令牌数组
//
func TokeMake(maps map[string]interface{}) (string, error) {
	s, err := str.JsonEnCode(maps)
	if err != nil {
		return "", &fox.Error{Msg:"序列化失败：" + err.Error()}
	}
	key := []byte(beego.AppConfig.String("aes_key"))
	result, err := crypt.AesEncrypt([]byte(s), key)
	if err != nil {
		return "", &fox.Error{Msg:"加密失败：" + err.Error()}
	}
	b64:=base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	return b64.EncodeToString(result),nil
}
//令牌解密
//@str 加密的字符串
func TokenDeCode(str string) (map[string]interface{}, error) {
	if len(str) < 1 {
		return nil, &fox.Error{Msg:"字符串 不能为空"}
	}
	key := []byte(beego.AppConfig.String("aes_key"))
	b64:=base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	byt, err := b64.DecodeString(str)
	if err != nil {
		return nil, &fox.Error{Msg:"base64解码失败：" + err.Error()}
	}
	origData, err := crypt.AesDecrypt(byt, key)
	if err != nil {
		return nil, &fox.Error{Msg:"解密失败：" + err.Error()}
	}
	maps, err := array.StrToMap(string(origData))
	if err != nil {
		return nil, &fox.Error{Msg:"转换为map失败：" + err.Error()}
	}
	return maps, nil
}