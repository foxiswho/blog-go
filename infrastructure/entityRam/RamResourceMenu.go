package entityRam

import (
	"time"
)

// RamResourceMenuEntity 资源菜单关系
type RamResourceMenuEntity struct {
	ID           int64      `gorm:"column:id;type:bigserial;primaryKey;autoIncrement:true" json:"id" comment:"" `
	No           string     `gorm:"column:no;type:varchar(80);index;default:;comment:编号" json:"no" comment:"编号" `
	Code         string     `gorm:"column:code;type:varchar(80);comment:标志" json:"code" comment:"标志" `                                                              // 编号代号
	CreateAt     *time.Time `gorm:"column:create_at;type:timestamptz;index;autoCreateTime;default:current_timestamp;comment:创建时间" json:"create_at" comment:"创建时间" ` // 创建时间
	CreateBy     string     `gorm:"column:create_by;type:varchar(80);index;default:;comment:创建人" json:"create_by" comment:"创建人" `
	TenantNo     string     `gorm:"column:tenant_no;type:varchar(80);index;default:;comment:租户编号" json:"tenant_no" comment:"租户编号" ` // 租户
	TypeSys      string     `gorm:"column:type_sys;type:varchar(100);index;default:;comment:类型|普通|系统|api" json:"type_sys" comment:"类型;普通;系统;api" `
	TypeCategory string     `gorm:"column:type_category;type:varchar(100);index;default:;comment:类型\\;角色\\;资源组\\;部门" json:"type_category" comment:"类型;角色;资源组;部门" `
	TypeValue    string     `gorm:"column:type_value;type:varchar(80);index;default:;comment:对应类型id" json:"type_value" comment:"对应类型id" `
	TypeDomain   string     `gorm:"column:type_domain;type:varchar(80);index;default:'default';comment:域|平台platform|店铺shop|其他other" json:"type_domain" comment:"域;平台platform;店铺shop;其他other" `
	TypeAttr     string     `gorm:"column:type_attr;type:varchar(100);index;default:;comment:属性|菜单分类|资源" json:"type_attr" comment:"属性;菜单menu;按钮button;资源:resource;其他other" `
	Mark         string     `gorm:"column:mark;type:varchar(80);index;default:;comment:标记" json:"mark" comment:"标记" `
	ParentId     string     `gorm:"column:parent_id;type:varchar(80);index;default:;comment:上级" json:"parent_id" comment:"上级" `
	IdLink       string     `gorm:"column:id_link;type:text;comment:上级" json:"parent_id_link" comment:"上级" `
	ResourceId   int64      `gorm:"column:resource_id;type:bigint;not null;index;default:0;comment:资源id" json:"resource_id" comment:"资源id" `
	GroupId      int64      `gorm:"column:group_id;type:bigint;not null;index;default:0;comment:资源组id" json:"group_id" comment:"资源组id" `
	MenuId       int64      `gorm:"column:menu_id;type:bigint;not null;index;default:0;comment:菜单id" json:"menu_id" comment:"菜单id" `
}

func (*RamResourceMenuEntity) TableName() string {
	return "ram_resource_menu"
}

func (*RamResourceMenuEntity) TableComment() string {
	return "资源菜单关系"
}
