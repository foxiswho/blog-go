package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
)

func main() {
	//匹配模式【这里指匹配 默认首页】处理对应 函数
	http.HandleFunc("/", sayHelloName)
	// 创建一个服务,IP 端口，监听, 编译运行后，浏览器访问 localhost:9999 即可在命令行后看出打印结果日志
	error := http.ListenAndServe("0.0.0.0:9999", nil)
	if error != nil {
		//日志
		log.Fatal("ListenAndServe: ", error)
	}
}
/***
处理函数
 */
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	//解析表单
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "这是一个web测试页面")
	fmt.Fprintf(w, "\n")
}