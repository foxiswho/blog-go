package listenerRam

import (
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/event/ram/listenerRam/service"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/sdk/ram/model/modRamAccount"
	"github.com/pangu-2/go-tools/tools/strPg"
)

// RamListener ram相关
type RamListener struct {
	log      *log2.Logger                                `autowire:"?"`
	acc      *repositoryRam.RamAccountRepository         `autowire:"?"`
	loginLog *repositoryRam.RamAccountLoginLogRepository `autowire:"?"`
	session  *repositoryRam.RamAccountSessionRepository  `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *RamListener) Run() error {
	//账号 登录日志
	eventBus.RegisterEvent(constEventBusPg.RamAccountLoginLog).RegisterSubscribe(constEventBusPg.RamAccountLoginLog, func(message any, _ core.EventArgs) {
		c.log.Infof("SchedulerEvent[账号.登录日志].event=%+v", message)
		dto := message.(modRamAccount.LoginLogDto)
		if strPg.IsNotBlank(dto.Ano) {
			err := service.NewAccountLoginLog(c.log, c.acc, c.loginLog, c.session).Processor(dto)
			if nil != err {
				c.log.Error("", err)
			}
			message = nil
		}
	})
	return nil
}
