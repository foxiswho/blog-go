package ram

import (
	"github.com/foxiswho/blog-go/app/event/ram/service/accountDomainInit"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(new(accountDomainInit.Sp))
}
