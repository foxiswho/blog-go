package listenerRam

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/ram/listenerRam/service"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constEventBusPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/sdk/ram/model/modRamAccount"
	"github.com/pangu-2/go-tools/tools/strPg"
	"go-spring.org/log"
	_ "go-spring.org/spring/gs"
)

// RamListener ram相关
type RamListener struct {
	acc      *repositoryRam.RamAccountRepository         `autowire:"?"`
	loginLog *repositoryRam.RamAccountLoginLogRepository `autowire:"?"`
	session  *repositoryRam.RamAccountSessionRepository  `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *RamListener) Run(ctx context.Context) error {
	//账号 登录日志
	eventBus.RegisterEvent(constEventBusPg.RamAccountLoginLog).RegisterSubscribe(constEventBusPg.RamAccountLoginLog, func(message any, _ core.EventArgs) {
		log.Infof(ctx, log.TagAppDef, "SchedulerEvent[账号.登录日志].event=%+v", message)
		dto := message.(modRamAccount.LoginLogDto)
		if strPg.IsNotBlank(dto.Ano) {
			err := service.NewAccountLoginLog(c.acc, c.loginLog, c.session).Processor(context.Background(), dto)
			if nil != err {
				log.Errorf(ctx, log.TagAppDef, "", err)
			}
			message = nil
		}
	})
	return nil
}
