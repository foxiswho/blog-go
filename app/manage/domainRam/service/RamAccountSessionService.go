package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamAccountSession"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/strPg"

	"reflect"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamAccountSessionService)).Init(func(s *RamAccountSessionService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamAccountSessionService 团队
// @Description:
type RamAccountSessionService struct {
	sv    *repositoryRam.RamAccountSessionRepository `autowire:"?"`
	accDb *repositoryRam.RamAccountRepository        `autowire:"?"`
	log   *log2.Logger                               `autowire:"?"`
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAccountSessionService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	cn := c.sv
	finds, b := cn.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	idsNew := make([]int64, 0)
	for _, info := range finds {
		c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		idsNew = append(idsNew, info.ID)
	}
	if len(idsNew) > 0 {
		cn.DeleteByIds(idsNew, repositoryPg.GetOption(ctx))
	}
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAccountSessionService) Query(ctx *gin.Context, ct modRamAccountSession.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamAccountSession.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamAccountSessionEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamAccountSession.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamAccountSessionEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamAccountSessionEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		p.Condition = r.DbModel().Order("update_at desc")
		//自定义查询
		//if "" != ct.Wd {
		//	p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		//}
	}, repositoryPg.GetOption(ctx))
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamAccountSession.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		mapAcc := make(map[string]*entityRam.RamAccountEntity)
		idsAcc := make([]string, 0)
		for _, item := range page.Data {
			if strPg.IsNotBlank(item.Ano) {
				idsAcc = append(idsAcc, item.Ano)
			}
		}
		// 账号
		{
			if len(idsAcc) > 0 {
				acc, b := c.accDb.FindAllByNoIn(idsAcc)
				if b {
					for _, item := range acc {
						mapAcc[item.No] = item
					}
				}
			}
		}
		//字段赋值
		for _, item := range page.Data {
			var vo modRamAccountSession.Vo
			copier.Copy(&vo, &item)
			//
			vo.ExtraData = make(map[string]any)
			if strPg.IsNotBlank(item.Ano) {
				if acc, ok := mapAcc[item.Ano]; ok {
					vo.ExtraData["account"] = acc.Account
					vo.ExtraData["cc"] = acc.Cc
					vo.ExtraData["code"] = acc.Code
					vo.ExtraData["description"] = acc.Description
					vo.ExtraData["mail"] = acc.Mail
					vo.ExtraData["phone"] = acc.Phone
					vo.ExtraData["registerTime"] = acc.RegisterTime
				}
			}
			//
			slice = append(slice, vo)
		}
		pg.Data = slice
		pg.Pageable = page.Pageable
		rt.Data = pg
		return rt.Ok()
	}
	return rt.Ok()
}
