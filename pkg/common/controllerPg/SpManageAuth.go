package controllerPg

import "github.com/foxiswho/blog-go/middleware/authPg"

type SpManageAuth struct {
	Sp *authPg.GroupManageMiddlewareSp `autowire:""`
}
