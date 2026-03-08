package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/ram/service/accountDomainInit"
	"github.com/foxiswho/blog-go/pkg/log2"
)

// ZInitAccountAdmin
// @Description: 超管账号初始化
type ZInitAccountAdmin struct {
	log *log2.Logger          `autowire:"?"`
	sp  *accountDomainInit.Sp `autowire:"?"`
}

func (b *ZInitAccountAdmin) Run() error {
	b.log.Infof("[init].[超管账号初始化]===================")
	accountDomainInit.NewInitAccount(b.log, b.sp).Processor(context.Background())
	return nil
}
