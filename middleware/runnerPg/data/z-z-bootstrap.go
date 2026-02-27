package data

import (
	"fmt"
	"github.com/foxiswho/blog-go/pkg/configPg"
)

// ZzBootstrap
// @Description: 显示服务启动信息
type ZzBootstrap struct {
	ser configPg.Server `value:"${server}"`
}

func (b *ZzBootstrap) Run() error {
	fmt.Println()
	fmt.Printf("host: %+v\n", "localhost")
	fmt.Printf("port: %+v\n", b.ser.Port)
	fmt.Printf("url: http://localhost:%+v\n", b.ser.Port)
	fmt.Println()
	return nil
}
