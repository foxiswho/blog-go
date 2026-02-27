package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modPublic"
	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamAccount"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/passwordTypePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/numberPg"

	"reflect"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(NewRamAccountPublicService).Init(func(s *RamAccountPublicService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamAccountPublicService 账户公共动作
// @Description:
type RamAccountPublicService struct {
	sv    *repositoryRam.RamAccountRepository              `autowire:"?"`
	aAuth *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	log   *log2.Logger                                     `autowire:"?"`
}

func NewRamAccountPublicService() *RamAccountPublicService {
	return new(RamAccountPublicService)
}

// Public 登陆用户信息
func (c *RamAccountPublicService) Public(holder holderPg.HolderPg) (rt rg.Rs[modRamAccount.AccountPub]) {
	c.log.Infof("holder=%+v", holder)
	c.log.Infof("HolderData=%+v", holder.HolderData)
	if nil == holder.HolderData {
		return rt.ErrorMessage("账号登陆失败")
	}
	data := rt.Data
	account := holder.GetAccount()
	copier.Copy(&data, &account)
	data.RealName = account.Name
	data.Avatar = ""
	data.Username = account.Account
	data.UserId = numberPg.Int64ToString(account.ID)
	data.Departments = make([]string, 0)
	if len(account.Os.Departments) > 0 {
		data.Departments = account.Os.Departments
	}
	rt.Data = data
	return rt.Ok()
}

// UpdatePassword 修改密码
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAccountPublicService) UpdatePassword(ctx *gin.Context, ct modPublic.PasswordCt) (rt rg.Rs[string]) {
	if "" == ct.PasswordNew {
		return rt.ErrorMessage("密码不能为空")
	}
	holder := holderPg.GetContextAccount(ctx)
	account := holder.GetAccount()
	r := c.sv
	info, b := r.FindById(account.ID)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	r2 := c.aAuth
	passwd, result := r2.FindByTypePasswordANo(info.No)
	if !result {
		return rt.ErrorMessage("数据不存在")
	}
	entity := entityRam.RamAccountAuthorizationEntity{}
	entity.ExtraData = strPg.GetNanoid(8)
	entity.Value = userPg.PasswordSalt(ct.PasswordNew, entity.ExtraData)
	if nil == passwd {
		entity.Ano = info.No
		entity.TenantNo = info.TenantNo
		entity.Type = passwordTypePg.Password.String()
		r2.Create(&entity)
	} else {
		r2.Update(entity, passwd.ID)
	}

	return rt.Ok()
}
