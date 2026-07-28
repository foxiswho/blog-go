package eventBasicEvent

import (
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryBasic"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/cachePg/rdsPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	Log            *log2.Logger                                      `autowire:"?"`
	rdt            *rdsPg.BatchString                                `autowire:"?"`
	repModel       *repositoryBasic.BasicConfigModelRepository       `autowire:"?"`
	repModelFields *repositoryBasic.BasicConfigModelFieldsRepository `autowire:"?"`
	repEvent       *repositoryBasic.BasicConfigEventRepository       `autowire:"?"`
	repEventFields *repositoryBasic.BasicConfigEventFieldsRepository `autowire:"?"`
	repRules       *repositoryBasic.BasicModelRulesRepository        `autowire:"?"`
}
