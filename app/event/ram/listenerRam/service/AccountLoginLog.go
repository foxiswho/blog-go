package service

import (
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constHeaderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/sdk/ram/model/modRamAccount"
	"time"
)

type AccountLoginLog struct {
	log      *log2.Logger                                `autowire:"?"`
	acc      *repositoryRam.RamAccountRepository         `autowire:"?"`
	loginLog *repositoryRam.RamAccountLoginLogRepository `autowire:"?"`
	session  *repositoryRam.RamAccountSessionRepository  `autowire:"?"`
}

func NewAccountLoginLog(
	log *log2.Logger,
	acc *repositoryRam.RamAccountRepository,
	loginLog *repositoryRam.RamAccountLoginLogRepository,
	session *repositoryRam.RamAccountSessionRepository,
) *AccountLoginLog {
	return &AccountLoginLog{
		log:      log,
		acc:      acc,
		loginLog: loginLog,
		session:  session,
	}
}

// Processor
//
//	@Description: 处理
//	@receiver c
//	@param data
func (c *AccountLoginLog) Processor(data modRamAccount.LoginLogDto) error {
	acc := data.Account
	ua := ""
	if data.ExtraData != nil {
		if get, ok := data.ExtraData[constHeaderPg.HeaderUserAgent]; ok {
			ua = get.(string)
		}
	}
	now := time.Now()
	//登录时间
	{
		save := entityRam.RamAccountEntity{LoginTime: acc.LoginTime}
		if nil == save.LoginTime {
			save.LoginTime = &now
		}
		c.acc.Update(save, acc.ID)
	}
	//会话信息
	{
		session := entityRam.RamAccountSessionEntity{
			Ano:         acc.No,
			AppNo:       data.AppNo,
			Client:      data.Client,
			LoginSource: data.LoginSource,
			Os:          acc.Os,
			TenantNo:    acc.TenantNo,
			UserAgent:   ua,
			Ip:          data.Ip,
		}
		err, _ := c.session.Create(&session)
		if err != nil {
			c.log.Error("", err)
		}
	}
	//登录日志
	save := entityRam.RamAccountLoginLogEntity{
		AppNo:       data.AppNo,
		Ano:         acc.No,
		Client:      data.Client,
		LoginSource: data.LoginSource,
		Os:          acc.Os,
		TenantNo:    acc.TenantNo,
		UserAgent:   ua,
		Ip:          data.Ip,
	}
	err, _ := c.loginLog.Create(&save)
	if err != nil {
		c.log.Error("", err)
	}
	return err
}
