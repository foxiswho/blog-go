package modRamResource

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type UpdateCt struct {
	ID          typePg.Uint64String `json:"id" form:"id" validate:"required" label:"id" `
	OrgId       string              `json:"orgId" label:"组织id" `                                            // 组织id
	Name        string              `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" ` // 名称
	NameFl      string              `json:"nameFl" label:"名称外文" `                                           // 名称外文
	Code        string              `json:"code" form:"code"  label:"编号代号" `
	NameFull    string              `json:"nameFull" label:"全称" `    // 全称
	Description string              `json:"description" label:"描述" ` // 描述
	ParentId    string              `json:"parentId" label:"上级" `
	TypeSys     string              `json:"typeSys" validate:"required" label:"类型" `  //类型;普通;系统;api
	TypeAttr    string              `json:"typeAttr" validate:"required" label:"属性" ` //属性;菜单分类;资源
	Path        string              `json:"path" label:"路径" `
	Method      string              `json:"method" label:"方法" `
	MenuId      typePg.Int64String  `json:"menuId" label:"菜单id" `
	ParentNo    string              `json:"parentNo" label:"上级" `
}
