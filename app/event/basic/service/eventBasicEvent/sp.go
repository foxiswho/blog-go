package eventBasicEvent

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/spring-core/gs"
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
