package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamAccount"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/passwordTypePg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"

	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(NewRamAccountPasswordService).Init(func(s *RamAccountPasswordService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamAccountPasswordService 密码修改
// @Description:
type RamAccountPasswordService struct {
	sv    *repositoryRam.RamAccountRepository              `autowire:"?"`
	aAuth *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	log   *log2.Logger                                     `autowire:"?"`
}

func NewRamAccountPasswordService() *RamAccountPasswordService {
	return new(RamAccountPasswordService)
}

// UpdatePassword 修改密码
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAccountPasswordService) UpdatePassword(ctx *gin.Context, ct modRamAccount.PasswordCt, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	c.log.Infof("tp=%+v,ct=%+v", tp, ct)
	if ct.ID == 0 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.PasswordNew {
		return rt.ErrorMessage("密码不能为空")
	}
	r := c.sv
	info, b := r.FindByIdAndTypeDomain(ct.ID.ToInt64(), tp.ToTypeDomain().String())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}

	r2 := c.aAuth

	passwd, err := r2.FindByTypePasswordANo(info.No)
	if !err {
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
