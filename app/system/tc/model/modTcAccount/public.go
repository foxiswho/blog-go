package modRamAccount

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type AccountPub struct {
	TenantId      string      `json:"tenantId" comment:"租户" `            // 租户
	OrgId         string      `json:"orgId" comment:"组织id" `             // 组织id
	Name          string      `json:"name" comment:"名称" `                // 名称
	Account       string      `json:"account" comment:"账户" `             // 账户
	Cc            string      `json:"cc" comment:"国际区号"`                 // 国际区号
	Phone         string      `json:"phone" comment:"手机号" `              // 手机号
	Mail          string      `json:"mail" comment:"邮箱" `                // 邮箱
	Code          string      `json:"code" comment:"编码" `                // 编码
	MailVerify    typePg.Int8 `json:"mailVerify" comment:"邮箱验证1是2否" `    // 邮箱验证1是2否
	PhoneVerify   typePg.Int8 `json:"phoneVerify" comment:"手机验证1是2否" `   // 手机验证1是2否
	AccountVerify typePg.Int8 `json:"accountVerify" comment:"账户验证1是2否" ` // 账户验证1是2否
	State         typePg.Int8 `json:"state" comment:"启用1是2否" `           // 启用1是2否
	RegisterTime  *time.Time  `json:"registerTime" comment:"注册时间" `      // 注册时间
	RegisterIP    string      `json:"registerIp" comment:"注册ip" `        // 注册ip
	LoginTime     *time.Time  `json:"loginTime" comment:"登陆时间" `
	LevelNo       string      `json:"levelNo" label:"级别" `
	GroupNo       string      `json:"groupNo" label:"组" `
	RoleNo        string      `json:"roleNo" label:"角色" `
	DepartmentNo  string      `json:"departmentNo" label:"主部门id" `
	Departments   []string    `json:"departments" comment:"部门" `
	TeamIds       string      `json:"teamIds" comment:"团队" `
	TypeDomain    string      `json:"typeDomain" comment:"域类型" `
	Avatar        string      `json:"avatar"  `
	UserId        string      `json:"userId"  `
	Username      string      `json:"username"  `
	RealName      string      `json:"realName"`
}
