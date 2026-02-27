package service

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/foxiswho/blog-go/app/manage/domainApi/model/modApiDiplAccessKey"
	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/infrastructure/repositoryApi"
	"github.com/foxiswho/blog-go/middleware/components/cachePg/cacheDiplPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"gorm.io/gorm"
)

func init() {
	gs.Provide(new(ApiDiplAccessKeyService)).Init(func(s *ApiDiplAccessKeyService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// ApiDiplAccessKeyService 团队
// @Description:
type ApiDiplAccessKeyService struct {
	sv  *repositoryApi.ApiDiplAccessKeyRepository `autowire:"?"`
	app *repositoryApi.ApiDiplRepository          `autowire:"?"`
	log *log2.Logger                              `autowire:"?"`
}

// MakeNewRecord
//
//	@Description: 生成新记录
//	@receiver c
//	@param ctx
//	@param ct
//	@return rt
func (c *ApiDiplAccessKeyService) MakeNewRecord(ctx *gin.Context, ct model.BaseIdCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	save := entityApi.ApiDiplAccessKeyEntity{}
	save.DiplNo = strings.TrimSpace(ct.Id)
	if strPg.IsBlank(save.DiplNo) {
		return rt.ErrorMessage("应用 编号 错误")
	}
	//当前时间增加 5年
	add := time.Now().AddDate(10, 0, 0)
	holder := holderPg.GetContextAccount(ctx)
	save.State = enumStatePg.ENABLE.IndexInt8()
	save.Name = noPg.No()
	save.TenantNo = holder.GetTenantNo()
	save.Key = strPg.GetNanoIdNumber(20)
	save.Secret = strPg.GetNanoid(20)
	save.ExpiryDate = &add
	save.KindUnique = userPg.SaltMake(save.Key, save.Secret+save.ExpiryDate.String())
	err, _ := c.sv.Create(&save)
	if err != nil {
		c.log.Error("", err)
		return rt.ErrorMessage("保存失败")
	}
	//加入缓存
	sha := cacheDiplPg.HashSha(save.Key, save.Secret)
	obj := cacheDiplPg.DiplCo{
		HashSha:  sha,
		No:       save.DiplNo,
		TenantNo: save.TenantNo,
		Key:      save.Key,
		Secret:   save.Secret,
	}
	cacheDiplPg.Set(save.Key, obj)
	return rt.Ok()
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	ct2 := model.BaseStateIdsCt[string]{
		Ids: ct.Ids,
	}
	ct2.State = enumStatePg.ENABLE.IndexInt64()
	return c.State(ctx, ct2)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	ct2 := model.BaseStateIdsCt[string]{
		Ids: ct.Ids,
	}
	ct2.State = enumStatePg.DISABLE.IndexInt64()
	return c.State(ctx, ct2)
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) State(ctx *gin.Context, ct model.BaseStateIdsCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	ids := ct.Ids
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	state, ok := enumStatePg.IsExistInt64(ct.State)
	if !ok {
		return rt.ErrorMessage("类型不正确")
	}
	if !state.IsEnableDisable() {
		return rt.ErrorMessage("状态错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	if nil == ct.Extend {
		return rt.ErrorMessage("扩展参数错误")
	}
	no := ""
	{
		if obj, ok := ct.Extend["no"]; ok {
			no = obj.(string)
		} else {
			return rt.ErrorMessage("扩展参数错误")
		}
	}
	no = strings.TrimSpace(no)
	if strPg.IsBlank(no) {
		return rt.ErrorMessage("扩展参数错误")
	}
	for _, info := range finds {
		if info.State != state.IndexInt8() {
			r.UpdateAllByDiplNoAndNoSetState(no, numberPg.Int64ToString(info.ID), state.IndexInt8())
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) StateEnableDisable(ctx *gin.Context, ct model.BaseStateIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct)
}

// LogicalDeletion 逻辑删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	if c.sv.Config().Data.Delete {
		for _, info := range finds {
			c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		}
		repository.DeleteByIdsString(ids, repositoryPg.GetOption(ctx))
	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityApi.ApiDiplAccessKeyEntity{State: reverse.IndexInt8()}, info.ID)
			}
		}
	}
	return rt.Ok()
}

// LogicalRecovery 逻辑删除恢复
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		enum := enumStatePg.State(info.State)
		//  取消 弃置 批量删除，反转 为对应的 有效 停用 停用
		if ok, reverse := enum.ReverseCancelLayAside(); ok {
			repository.Update(entityApi.ApiDiplAccessKeyEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
		//删除缓存
		cacheDiplPg.Remove(info.Key)
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
func (c *ApiDiplAccessKeyService) Query(ctx *gin.Context, ct modApiDiplAccessKey.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modApiDiplAccessKey.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityApi.ApiDiplAccessKeyEntity
	copier.Copy(&query, &ct)
	slice := make([]modApiDiplAccessKey.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityApi.ApiDiplAccessKeyEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityApi.ApiDiplAccessKeyEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		p.Condition = r.DbModel().Order("create_at desc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
	}, repositoryPg.GetOption(ctx))
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modApiDiplAccessKey.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modApiDiplAccessKey.Vo
			copier.Copy(&vo, &item)
			//
			vo.Hash = item.Key + ":" + cacheDiplPg.HashSha(item.Key, item.Secret)
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

// SelectPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ApiDiplAccessKeyService) SelectPublic(ctx *gin.Context, ct modApiDiplAccessKey.QueryCt) (rt rg.Rs[[]modApiDiplAccessKey.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityApi.ApiDiplAccessKeyEntity
	copier.Copy(&query, &ct)
	if strPg.IsBlank(ct.DiplNo) {
		query.DiplNo = "-1"
	}
	var con repositoryPg.Condition = func(db *gorm.DB) *gorm.DB {
		db = db.Order("create_at desc")
		return db
	}
	slice := make([]modApiDiplAccessKey.Vo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, con, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modApiDiplAccessKey.Vo
			copier.Copy(&vo, &item)
			//
			vo.Hash = item.Key + ":" + cacheDiplPg.HashSha(item.Key, item.Secret)
			//c.log.Debugf("vo.Hash=%+v", vo.Hash)
			//
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}
