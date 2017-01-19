package http

import (
	"fmt"
	"strings"
	"blog/fox/str"
	"net/http"
	"github.com/astaxie/beego/context"
)
/**
获取IP
 */
//func getClientIp() string {
//	request := http.Request{}
//	ip := request.Header.Get("Remote_addr")
//	if ip == "" {
//		ip = "127.0.0.1"
//	}
//	fmt.Println(ip)
//	if strings.Contains(ip, ":") {
//		ip = str.Substr(ip, 0, strings.Index(ip, ":"))
//	}
//	fmt.Println(ip)
//	return ip
//}
//func ip() string {
//	ctx := context.BeegoInput{}
//	return ctx.Context.Input.IP()
//}