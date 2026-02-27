package controllerPg

import "github.com/foxiswho/blog-go/middleware/authPg"

type SpSystemAuth struct {
	Sp *authPg.GroupSystemMiddlewareSp `autowire:""`
}
