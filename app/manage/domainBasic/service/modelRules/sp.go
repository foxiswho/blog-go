package modelRules

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type Sp struct {
	log            *log2.Logger                                      `autowire:"?"`
	repModel       *repositoryBasic.BasicConfigModelRepository       `autowire:"?"`
	repModelFields *repositoryBasic.BasicConfigModelFieldsRepository `autowire:"?"`
	repEvent       *repositoryBasic.BasicConfigEventRepository       `autowire:"?"`
	repEventFields *repositoryBasic.BasicConfigEventFieldsRepository `autowire:"?"`
	repRules       *repositoryBasic.BasicModelRulesRepository        `autowire:"?"`
}
