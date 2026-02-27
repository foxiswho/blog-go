package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/app/manage/domainTc/model/dtoTenantDomain"
	"github.com/foxiswho/blog-go/infrastructure/entityTc"
	"github.com/foxiswho/blog-go/infrastructure/repositoryTc"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/strPg"
)

func init() {
	gs.Provide(new(TcTenantDomainCacheService)).Init(func(s *TcTenantDomainCacheService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type TcTenantDomainCacheService struct {
	log   *log2.Logger                           `autowire:"?"`
	sv    *repositoryTc.TcTenantDomainRepository `autowire:"?"`
	tenDb *repositoryTc.TcTenantRepository       `autowire:"?"`
	ca    *cacheTc.TenantDomainCache             `autowire:"?"`
}

// InitTenantDomain
//
//	@Description: 初始化加载域名缓存
//	@receiver c
func (c *TcTenantDomainCacheService) InitTenantDomain() {
	query := entityTc.TcTenantDomainEntity{
		State: enumStatePg.ENABLE.Index(),
	}
	//获取所有域名数据
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		idsTenantNo := make([]string, 0)
		mapTenant := make(map[string]*entityTc.TcTenantEntity)
		for _, item := range infos {
			if strPg.IsNotBlank(item.TenantNo) {
				idsTenantNo = append(idsTenantNo, item.TenantNo)
			}
		}
		//
		{
			if len(idsTenantNo) > 0 {
				info, result := c.tenDb.FindAllByNoIn(idsTenantNo)
				if result {
					for _, item := range info {
						mapTenant[item.No] = item
					}
				}
			}
		}
		//
		for _, item := range infos {
			syslog.Infof(context.Background(), syslog.TagAppDef, "domian=tenant,%+v=%+v", item.Code, item.TenantNo)
			c.ca.Domain.Store(item.Code, item.TenantNo)
			dto := dtoTenantDomain.Dto{
				TenantNo: item.TenantNo,
				Domain:   item.Code,
			}
			if strPg.IsNotBlank(item.TenantNo) {
				if info, ok := mapTenant[item.TenantNo]; ok {
					if strPg.IsNotBlank(info.Founder) {
						dto.Ano = info.Founder
					}
				}
			}
			c.ca.DomainData.Store(item.Code, dto)
		}
		infos = nil
		mapTenant = nil
	}
}
