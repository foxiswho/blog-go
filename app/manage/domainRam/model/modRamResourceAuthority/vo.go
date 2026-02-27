package modRamResourceAuthority

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID           typePg.Uint64String `json:"id" label:"id" `
	Name         string              `json:"name" label:"名称" `     // 名称
	NameFl       string              `json:"nameFl" label:"名称外文" ` // 名称外文
	Code         string              `json:"code" label:"标志" `
	NameFull     string              `json:"nameFull" label:"全称" `                                 // 全称
	State        typePg.Int8         `json:"state" label:"1有效2停用11取消(对应有效)12弃置(对应停用)13批量删除(无状态)" ` // 1有效2停用11取消(对应有效)12弃置(对应停用)13批量删除(无状态)
	Description  string              `json:"description" label:"描述" `                              // 描述
	CreateAt     *time.Time          `json:"createAt" label:"创建时间" `                               // 创建时间
	UpdateAt     *time.Time          `json:"updateAt" label:"更新时间" `                               // 更新时间
	CreateBy     typePg.Int64String  `json:"createBy" label:"创建人" `                                // 创建人
	UpdateBy     typePg.Int64String  `json:"updateBy" label:"更新人" `                                // 更新人
	OrgId        string              `json:"orgId" label:"组织id" `                                  // 组织id
	Mark         string              `json:"mark" label:"标记" `
	TypeCategory string              `json:"typeCategory" label:"类型" `
	TypeValue    string              `json:"typeValue" label:"对应类型id" `
	TypeSys      string              `json:"typeSys" label:"类型" `  //类型;普通;系统;api
	TypeAttr     string              `json:"typeAttr" label:"属性" ` //属性;菜单分类;资源
	Path         string              `json:"path" label:"路径" `
	Method       string              `json:"method" label:"方法" `
	MenuId       typePg.Int64String  `json:"menuId" label:"菜单id" `
	ResourceId   typePg.Int64String  `json:"resourceId" label:"资源id" `
}
