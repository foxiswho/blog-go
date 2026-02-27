package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service/attachment"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
)

// AttachmentListener 附件处理
// @Description:
type AttachmentListener struct {
	log *log2.Logger                               `autowire:"?"`
	dao *repositoryBasic.BasicAttachmentRepository `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *AttachmentListener) Run() error {
	syslog.Infof(context.Background(), syslog.TagAppDef, "eventBus.Register=%+v", constEventBusPg.BasicAttachmentCreate)
	eventBus.RegisterEvent(constEventBusPg.BasicAttachmentCreate).RegisterSubscribe(constEventBusPg.BasicAttachmentCreate, func(message any, _ core.EventArgs) {
		syslog.Infof(context.Background(), syslog.TagAppDef, "SchedulerEvent.event=%+v", message)
		dto := message.(entityBasic.BasicAttachmentEntity)
		if len(dto.File) > 0 {
			err := attachment.NewCreate(c.dao, dto).Processor()
			if nil != err {
				c.log.Error("", err)
			}
			message = nil
		}
	})
	return nil
}
