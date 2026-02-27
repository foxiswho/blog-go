package modRamAccount

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

// UpdateDataCt
// @Description:  更新数据
type UpdateDataCt struct {
	Name         string           `json:"name" label:"名称" `       // 名称
	RealName     string           `json:"realName" label:"真实姓名" ` // 名称
	LevelNo      string           `json:"levelNo" label:"级别" `
	GroupNo      string           `json:"groupNo" label:"组" `
	RoleNo       string           `json:"roleNo" label:"角色" `
	DepartmentNo string           `json:"departmentNo" label:"主部门id" `
	Departments  []string         `json:"departments" label:"部门" `              // 部门
	Roles        []string         `json:"roles" label:"角色" `                    // 角色
	TeamIds      []string         `json:"teamIds" label:"团队" `                  // 团队
	TypeDomain   string           `json:"typeDomain" label:"域类型" `              // 域类型
	TypeIdentity string           `json:"typeIdentity" label:"身份类型;普通;经理;副经理" ` // 身份类型;普通;经理;副经理
	Description  string           `json:"description" label:"描述" `              // 描述
	Position     string           `json:"position" label:"岗位" `
	Job          string           `json:"job" label:"职位" `
	JobTitle     string           `json:"jobTitle" label:"职衔" `
	JobRank      string           `json:"jobRank" label:"职级" `
	Avatar       string           `json:"avatar" label:"头像" `
	Birthday     *typePg.DateOnly `json:"birthday" label:"生日" `
	Sex          string           `json:"sex" label:"性别1男2女3未知" `
	IdentityCode string           `json:"identityCode" label:"身份编号 "`

	//MailVerify       jsonPg.Int8        `json:"mailVerify" label:"邮箱验证1是2否" `                           // 邮箱验证1是2否
	//PhoneVerify      jsonPg.Int8        `json:"phoneVerify" label:"手机验证1是2否" `                          // 手机验证1是2否
	//AccountVerify    jsonPg.Int8        `json:"accountVerify" label:"账户验证1是2否" `                        // 账户验证1是2否
}
