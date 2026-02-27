package configModel

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type Sp struct {
	log       *log2.Logger                                      `autowire:"?"`
	repModel  *repositoryBasic.BasicConfigModelRepository       `autowire:"?"`
	repField  *repositoryBasic.BasicConfigModelFieldsRepository `autowire:"?"`
	repRule   *repositoryBasic.BasicConfigModelRulesRepository  `autowire:"?"`
	repModule *repositoryBasic.BasicModuleRepository            `autowire:"?"`
}
