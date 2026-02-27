package service

import (
	"context"
	"github.com/foxiswho/blog-go/app/system/tc/model/modTcAccount"
	"github.com/foxiswho/blog-go/app/system/tc/service/tcAccount"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/entityTc"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryTc"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/enum/state/yesNoPg/yesNoIntPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/strPg2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/slicePg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"reflect"
)

func init() {
	gs.Provide(NewTcTenantAccountService).Init(func(s *TcTenantAccountService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// TcTenantAccountService 账户，账号
// @Description:
type TcTenantAccountService struct {
	sv         *repositoryRam.RamAccountRepository              `autowire:"?"`
	aAuth      *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	dep        *repositoryRam.RamDepartmentRepository           `autowire:"?"`
	role       *repositoryRam.RamRoleRepository                 `autowire:"?"`
	team       *repositoryRam.RamTeamRepository                 `autowire:"?"`
	levelDb    *repositoryRam.RamLevelRepository                `autowire:"?"`
	groupDb    *repositoryRam.RamGroupRepository                `autowire:"?"`
	positionDb *repositoryRam.RamPositionRepository             `autowire:"?"`
	postDb     *repositoryRam.RamPostRepository                 `autowire:"?"`
	tenDb      *repositoryTc.TcTenantRepository                 `autowire:"?"`
	sp         *tcAccount.Sp                                    `autowire:"?"`
	log        *log2.Logger                                     `autowire:"?"`
}

func NewTcTenantAccountService() *TcTenantAccountService {
	return new(TcTenantAccountService)
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *TcTenantAccountService) Detail(ctx *gin.Context, id string, tp appModulePg.AppModule) (rt rg.Rs[modRamAccount.DetailVo]) {
	detail := tcAccount.NewDetail(c.log, c.sp, tp)
	return detail.Process(ctx, id)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE, tp)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE), tp)
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) State(ctx *gin.Context, ids []string, state enumStatePg.State, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringInAndTypeDomain(ids, tp.ToTypeDomain().String())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		//  founder 不可禁用
		if yesNoIntPg.Yes.IsEqual(info.Founder) {
			continue
		}
		if info.State != state.IndexInt8() {
			r.Update(entityRam.RamAccountEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if !state.IsEnableDisable() {
		return rt.ErrorMessage("状态错误")
	}
	return c.State(ctx, ids, state, tp)
}

// LogicalDeletion 逻辑删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) LogicalDeletion(ctx *gin.Context, ids []string, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringInAndTypeDomain(ids, tp.ToTypeDomain().String())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	// 数据 点击删除时是否直接删除
	if c.sv.Config().Data.Delete {
		idsNow := make([]int64, 0)
		for _, info := range finds {
			//  founder 不可禁用
			if yesNoIntPg.Yes.IsEqual(info.Founder) {
				continue
			}
			idsNow = append(idsNow, info.ID)
			c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		}
		if len(idsNow) > 0 {
			r.DeleteByIds(idsNow)
		}

	} else {
		for _, info := range finds {
			//  founder 不可禁用
			if yesNoIntPg.Yes.IsEqual(info.Founder) {
				continue
			}
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				r.Update(entityRam.RamAccountEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *TcTenantAccountService) LogicalRecovery(ctx *gin.Context, ids []string, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringInAndTypeDomain(ids, tp.ToTypeDomain().String())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		//  founder 不可禁用
		if yesNoIntPg.Yes.IsEqual(info.Founder) {
			continue
		}
		enum := enumStatePg.State(info.State)
		//  取消 弃置 批量删除，反转 为对应的 有效 停用 停用
		if ok, reverse := enum.ReverseCancelLayAside(); ok {
			r.Update(entityRam.RamAccountEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) PhysicalDeletion(ctx *gin.Context, ids []string, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringInAndTypeDomain(ids, tp.ToTypeDomain().String())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	idsNow := make([]int64, 0)
	for _, info := range finds {
		//  founder 不可禁用
		if yesNoIntPg.Yes.IsEqual(info.Founder) {
			continue
		}
		c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
	}
	if len(idsNow) > 0 {
		r.DeleteByIds(idsNow)
	}
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) Query(ctx *gin.Context, ct modRamAccount.QueryCt, tp appModulePg.AppModule) (rt rg.Rs[pagePg.PaginatorPg[modRamAccount.Vo]]) {
	var query entityRam.RamAccountEntity
	copier.Copy(&query, &ct)
	// 创始人
	query.Founder = enumStatePg.ENABLE.Index()
	//
	slice := make([]modRamAccount.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	depDb := c.dep
	roleDb := c.role
	groupDb := c.groupDb
	levelDb := c.levelDb
	teamDb := c.team
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamAccountEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamAccountEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
		}
		p.Condition = r.DbModel().Order("create_at desc")
		p.Condition.Where("type_domain= ?", tp.ToTypeDomain().String())
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("account like ?", "%"+ct.Wd+"%")
		}
		//部门
		if nil != ct.Departments && len(ct.Departments) > 0 {
			depInfo, result := depDb.FindAllByNoLinkArr(ct.Departments)
			if result {
				sqlDb := r.Db()
				for i, obj := range depInfo {
					if 0 == i {
						sqlDb = sqlDb.Or("os->'departments' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
					} else {
						sqlDb = sqlDb.Or("os->'departments' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
					}
				}
				p.Condition.Where(sqlDb)
			} else {
				p.Condition.Where("os->'departments' @> ? ", strPg2.StrToArrayJsonExpr("0"))
			}
		}
		//角色
		if nil != ct.Roles && len(ct.Roles) > 0 {
			depInfo, result := roleDb.FindAllByNoIn(ct.Roles)
			if result {
				sqlDb := r.Db()
				for i, obj := range depInfo {
					if 0 == i {
						sqlDb = sqlDb.Or("os->'roles' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
					} else {
						sqlDb = sqlDb.Or("os->'roles' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
					}
				}
				p.Condition.Where(sqlDb)
			} else {
				p.Condition.Where("os->'roles' @> ? ", strPg2.StrToArrayJsonExpr("0"))
			}
		}
		//级别
		{
			if nil != ct.Levels && len(ct.Levels) > 0 {
				depInfo, result := levelDb.FindAllByNoIn(ct.Levels)
				if result {
					sqlDb := r.Db()
					for i, obj := range depInfo {
						if 0 == i {
							sqlDb = sqlDb.Or("os->'levels' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
						} else {
							sqlDb = sqlDb.Or("os->'levels' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
						}
					}
					p.Condition.Where(sqlDb)
				} else {
					p.Condition.Where("os->'levels' @> ? ", strPg2.StrToArrayJsonExpr("0"))
				}
			}
		}
		//组
		{
			if nil != ct.Groups && len(ct.Groups) > 0 {
				depInfo, result := groupDb.FindAllByNoIn(ct.Groups)
				if result {
					sqlDb := r.Db()
					for i, obj := range depInfo {
						if 0 == i {
							sqlDb = sqlDb.Or("os->'groups' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
						} else {
							sqlDb = sqlDb.Or("os->'groups' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
						}
					}
					p.Condition.Where(sqlDb)
				}
			}
		}
		//团队
		{
			if nil != ct.Teams && len(ct.Teams) > 0 {
				depInfo, result := teamDb.FindAllByNoIn(ct.Teams)
				if result {
					sqlDb := r.Db()
					for i, obj := range depInfo {
						if 0 == i {
							sqlDb = sqlDb.Or("os->'teams' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
						} else {
							sqlDb = sqlDb.Or("os->'teams' @> ? ", strPg2.StrToArrayJsonExpr(obj.No))
						}
					}
					p.Condition.Where(sqlDb)
				}
			}
		}
		//注册时间 区间
		{
			if nil != ct.RegisterTimeRange {
				count := len(ct.RegisterTimeRange)
				if count == 2 && nil != ct.RegisterTimeRange[0] && nil != ct.RegisterTimeRange[1] {
					p.Condition.Where("register_time between ? and ?", ct.RegisterTimeRange[0], ct.RegisterTimeRange[1])
				} else if count == 1 && nil != ct.RegisterTimeRange[0] {
					p.Condition.Where("register_time >= ?", ct.RegisterTimeRange[0])
				}
			}
		}
		//登陆时间 区间
		{
			if nil != ct.LoginTimeRange {
				count := len(ct.LoginTimeRange)
				if count == 2 && nil != ct.LoginTimeRange[0] && nil != ct.LoginTimeRange[1] {
					p.Condition.Where("login_time between ? and ?", ct.LoginTimeRange[0], ct.LoginTimeRange[1])
				} else if count == 1 && nil != ct.LoginTimeRange[0] {
					p.Condition.Where("login_time >= ?", ct.LoginTimeRange[0])
				}
			}
		}
		//生日 区间
		{
			if nil != ct.BirthdayRange {
				count := len(ct.BirthdayRange)
				if count == 2 && nil != ct.BirthdayRange[0] && nil != ct.BirthdayRange[1] {
					p.Condition.Where("birthday between ? and ?", ct.BirthdayRange[0], ct.BirthdayRange[1])
				} else if count == 1 && nil != ct.BirthdayRange[0] {
					p.Condition.Where("birthday >= ?", ct.BirthdayRange[0])
				}
			}
		}
	})
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamAccount.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//
		mapDep := make(map[string]*entityRam.RamDepartmentEntity)
		mapRole := make(map[string]*entityRam.RamRoleEntity)
		mapLevel := make(map[string]*entityRam.RamLevelEntity)
		mapGroup := make(map[string]*entityRam.RamGroupEntity)
		mapTeam := make(map[string]*entityRam.RamTeamEntity)
		mapPosition := make(map[string]*entityRam.RamPositionEntity)
		mapPost := make(map[string]*entityRam.RamPostEntity)
		mapTenant := make(map[string]*entityTc.TcTenantEntity)
		idsDep := make([]string, 0)
		idsRole := make([]string, 0)
		idsLevel := make([]string, 0)
		idsGroup := make([]string, 0)
		idsTeam := make([]string, 0)
		idsPosition := make([]string, 0)
		idsPost := make([]string, 0)
		idsTenant := make([]string, 0)
		for _, item := range page.Data {
			//部门
			if nil != item.Os.Data().Departments && len(item.Os.Data().Departments) > 0 {
				for _, obj := range item.Os.Data().Departments {
					idsDep = append(idsDep, obj)
				}
			}
			if strPg.IsNotBlank(item.DepartmentNo) {
				idsDep = append(idsDep, item.DepartmentNo)
			}
			//角色
			if nil != item.Os.Data().Roles && len(item.Os.Data().Roles) > 0 {
				for _, obj := range item.Os.Data().Roles {
					idsRole = append(idsRole, obj)
				}
			}
			if strPg.IsNotBlank(item.RoleNo) {
				idsRole = append(idsRole, item.RoleNo)
			}
			//级别
			if nil != item.Os.Data().Levels && len(item.Os.Data().Levels) > 0 {
				for _, obj := range item.Os.Data().Levels {
					idsLevel = append(idsLevel, obj)
				}
			}
			if strPg.IsNotBlank(item.LevelNo) {
				idsLevel = append(idsLevel, item.LevelNo)
			}
			//分组
			if nil != item.Os.Data().Groups && len(item.Os.Data().Groups) > 0 {
				for _, obj := range item.Os.Data().Groups {
					idsGroup = append(idsGroup, obj)
				}
			}
			if strPg.IsNotBlank(item.GroupNo) {
				idsGroup = append(idsGroup, item.GroupNo)
			}
			//团队
			if nil != item.Os.Data().Teams && len(item.Os.Data().Teams) > 0 {
				for _, obj := range item.Os.Data().Teams {
					idsTeam = append(idsTeam, obj)
				}
			}
			//
			if strPg.IsNotBlank(item.Position) {
				idsPosition = append(idsPosition, item.Position)
			}
			//
			if strPg.IsNotBlank(item.Job) {
				idsPost = append(idsPost, item.Job)
			}
			//租户
			if nil != item.Os.Data().Tenants && len(item.Os.Data().Tenants) > 0 {
				for _, obj := range item.Os.Data().Tenants {
					idsTenant = append(idsTenant, obj)
				}
			}
			if strPg.IsNotBlank(item.TenantNo) {
				idsTenant = append(idsTenant, item.TenantNo)
			}

		}
		//部门
		{
			if len(idsDep) > 0 {
				infos, result := depDb.FindAllByNoIn(idsDep)
				if result {
					mapDep = slicePg.ToMap(infos, func(t *entityRam.RamDepartmentEntity) (string, *entityRam.RamDepartmentEntity) {
						return t.No, t
					})
				}
			}
		}
		//角色
		{
			if len(idsRole) > 0 {
				infos, result := roleDb.FindAllByNoIn(idsRole)
				if result {
					mapRole = slicePg.ToMap(infos, func(t *entityRam.RamRoleEntity) (string, *entityRam.RamRoleEntity) {
						return t.No, t
					})
				}
			}
		}
		//级别
		{
			if len(idsLevel) > 0 {
				infos, result := levelDb.FindAllByNoIn(idsLevel)
				if result {
					mapLevel = slicePg.ToMap(infos, func(t *entityRam.RamLevelEntity) (string, *entityRam.RamLevelEntity) {
						return t.No, t
					})
				}
			}
		}
		//分组
		{
			if len(idsGroup) > 0 {
				infos, result := groupDb.FindAllByNoIn(idsGroup)
				if result {
					mapGroup = slicePg.ToMap(infos, func(t *entityRam.RamGroupEntity) (string, *entityRam.RamGroupEntity) {
						return t.No, t
					})
				}
			}
		}
		//分组
		{
			if len(idsTeam) > 0 {
				infos, result := teamDb.FindAllByNoIn(idsTeam)
				if result {
					mapTeam = slicePg.ToMap(infos, func(t *entityRam.RamTeamEntity) (string, *entityRam.RamTeamEntity) {
						return t.No, t
					})
				}
			}
		}
		//职位
		{
			if len(idsPosition) > 0 {
				infos, result := c.positionDb.FindAllByNoIn(idsPosition)
				if result {
					mapPosition = slicePg.ToMap(infos, func(t *entityRam.RamPositionEntity) (string, *entityRam.RamPositionEntity) {
						return t.No, t
					})
				}
			}
		}
		//职位
		{
			if len(idsPost) > 0 {
				infos, result := c.postDb.FindAllByNoIn(idsPost)
				if result {
					mapPost = slicePg.ToMap(infos, func(t *entityRam.RamPostEntity) (string, *entityRam.RamPostEntity) {
						return t.No, t
					})
				}
			}
		}
		//租户
		{
			if len(idsTenant) > 0 {
				infos, result := c.tenDb.FindAllByNoIn(idsTenant)
				if result {
					mapTenant = slicePg.ToMap(infos, func(t *entityTc.TcTenantEntity) (string, *entityTc.TcTenantEntity) {
						return t.No, t
					})
				}
			}
		}
		//字段赋值
		for _, item := range page.Data {
			var vo modRamAccount.Vo
			copier.Copy(&vo, &item)
			//vo.Os.No = item.Os.Data()
			//部门
			if nil != item.Os.Data().Departments && len(item.Os.Data().Departments) > 0 {
				vo.Departments = item.Os.Data().Departments
				vo.Os.No.Departments = item.Os.Data().Departments
				vo.Os.NoName.Departments = make([]string, 0)
				for _, obj := range item.Os.Data().Departments {
					if get, ok := mapDep[obj]; ok {
						vo.Os.NoName.Departments = append(vo.Os.NoName.Departments, get.Name)
					}
				}
			}
			if strPg.IsNotBlank(item.DepartmentNo) {
				if obj, ok := mapDep[item.DepartmentNo]; ok {
					vo.DepartmentNoName = obj.Name
				}
			}
			//角色
			if nil != item.Os.Data().Roles && len(item.Os.Data().Roles) > 0 {
				vo.Roles = item.Os.Data().Roles
				vo.Os.No.Roles = item.Os.Data().Roles
				vo.Os.NoName.Roles = make([]string, 0)
				for _, obj := range item.Os.Data().Roles {
					if get, ok := mapRole[obj]; ok {
						vo.Os.NoName.Roles = append(vo.Os.NoName.Roles, get.Name)
					}
				}
			}
			//级别
			if nil != item.Os.Data().Levels && len(item.Os.Data().Levels) > 0 {
				vo.Os.No.Levels = item.Os.Data().Levels
				vo.Os.NoName.Levels = make([]string, 0)
				for _, obj := range item.Os.Data().Levels {
					if get, ok := mapLevel[obj]; ok {
						vo.Os.NoName.Levels = append(vo.Os.NoName.Levels, get.Name)
					}
				}
			}
			if strPg.IsNotBlank(item.LevelNo) {
				if obj, ok := mapLevel[item.LevelNo]; ok {
					vo.LevelNoName = obj.Name
				}
			}
			//分组
			if nil != item.Os.Data().Groups && len(item.Os.Data().Groups) > 0 {
				vo.Os.No.Groups = item.Os.Data().Groups
				vo.Os.NoName.Groups = make([]string, 0)
				for _, obj := range item.Os.Data().Groups {
					if get, ok := mapGroup[obj]; ok {
						vo.Os.NoName.Groups = append(vo.Os.NoName.Levels, get.Name)
					}
				}
			}
			if strPg.IsNotBlank(item.GroupNo) {
				if obj, ok := mapLevel[item.GroupNo]; ok {
					vo.GroupNoName = obj.Name
				}
			}
			//团队
			if nil != item.Os.Data().Teams && len(item.Os.Data().Teams) > 0 {
				vo.Teams = item.Os.Data().Teams
				vo.Os.No.Teams = item.Os.Data().Teams
				vo.Os.NoName.Teams = make([]string, 0)
				for _, obj := range item.Os.Data().Teams {
					if get, ok := mapTeam[obj]; ok {
						vo.Os.NoName.Teams = append(vo.Os.NoName.Teams, get.Name)
					}
				}
			}
			// 职位
			if strPg.IsNotBlank(item.Position) {
				if get, ok := mapPosition[item.Position]; ok {
					vo.PositionName = get.Name
				}
			}
			//职位
			if strPg.IsNotBlank(item.Job) {
				if get, ok := mapPost[item.Job]; ok {
					vo.JobName = get.Name
				}
			}
			//租户
			if nil != item.Os.Data().Tenants && len(item.Os.Data().Tenants) > 0 {
				vo.Os.No.Tenants = item.Os.Data().Tenants
				vo.Os.NoName.Tenants = make([]string, 0)
				for _, obj := range item.Os.Data().Tenants {
					if get, ok := mapTenant[obj]; ok {
						vo.Os.NoName.Tenants = append(vo.Os.NoName.Tenants, get.Name)
					}
				}
			}
			if strPg.IsNotBlank(item.TenantNo) {
				if get, ok := mapTenant[item.TenantNo]; ok {
					vo.TenantNoName = get.Name
				}
			}

			slice = append(slice, vo)
		}
		pg.Data = slice
		pg.Pageable = page.Pageable
		rt.Data = pg
		return rt.Ok()
	}
	return rt.Ok()
}

// Create 创建
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) Create(ctx *gin.Context, ct modRamAccount.CreateCt, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	return tcAccount.NewCreate(c.log,
		c.sp, ctx).Process(ctx, ct, tp)
}

// CreateAccount 创建
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) CreateAccount(ctx *gin.Context, ct modRamAccount.CreateAccountCt, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	return tcAccount.NewCreate(c.log,
		c.sp, ctx).CreateAccount(ct, tp)
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) Update(ctx *gin.Context, ct modRamAccount.UpdateCt, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	return tcAccount.NewUpdate(c.log,
		c.sp, ctx).Process(ctx, ct, tp)
}

// UpdateAccount 更新
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) UpdateAccount(ctx *gin.Context, ct modRamAccount.UpdateAccountCt, tp appModulePg.AppModule) (rt rg.Rs[string]) {
	return tcAccount.NewUpdate(c.log,
		c.sp, ctx).UpdateAccount(ct, tp)
}

// ExistAccount 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) ExistAccount(ctx *gin.Context, ct model.BaseExistWdCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
	}
	_, result := c.sv.FindByAccountAndTypeDomainAndIdNot(ct.Wd, tp.ToTypeDomain().String(), ct.Id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

// ExistPhone 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) ExistPhone(ctx *gin.Context, ct model.BaseExistWdCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
	}
	_, result := c.sv.FindByPhoneAndTypeDomainAndIdNot(ct.Wd, tp.ToTypeDomain().String(), ct.Id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

// ExistMail 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) ExistMail(ctx *gin.Context, ct model.BaseExistWdCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
	}
	_, result := c.sv.FindByMailAndTypeDomainAndIdNot(ct.Wd, tp.ToTypeDomain().String(), ct.Id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

// ExistIdentityCode 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) ExistIdentityCode(ctx *gin.Context, ct model.BaseExistWdCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
	}
	_, result := c.sv.FindByIdentityCodeAndTypeDomainAndIdNot(ct.Wd, tp.ToTypeDomain().String(), ct.Id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

// ExistRealName 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *TcTenantAccountService) ExistRealName(ctx *gin.Context, ct model.BaseExistWdCt[string], tp appModulePg.AppModule) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
	}
	_, result := c.sv.FindByRealNameAndTypeDomainAndIdNot(ct.Wd, tp.ToTypeDomain().String(), ct.Id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}
