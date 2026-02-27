package entityRam

import (
	"time"
)

// RamMenuRelationEntity 菜单权限关系
type RamMenuRelationEntity struct {
	ID           int64      `gorm:"column:id;type:bigserial;primaryKey;autoIncrement:true" json:"id" comment:"" `
	No           string     `gorm:"column:no;type:varchar(80);index;default:;comment:编号" json:"no" comment:"编号" `
	Name         string     `gorm:"column:name;type:varchar(255);comment:名称" json:"name" comment:"名称" `
	Code         string     `gorm:"column:code;type:varchar(80);comment:标志" json:"code" comment:"标志" `                                                              // 编号代号
	CreateAt     *time.Time `gorm:"column:create_at;type:timestamptz;index;autoCreateTime;default:current_timestamp;comment:创建时间" json:"create_at" comment:"创建时间" ` // 创建时间
	CreateBy     int64      `gorm:"column:create_by;type:bigint;not null;index;default:0;comment:创建人" json:"create_by" comment:"创建人" `
	Sort         int64      `gorm:"column:sort;type:bigint;not null;default:0;;comment:排序" json:"sort" comment:"排序" `               // 排序
	TenantNo     string     `gorm:"column:tenant_no;type:varchar(80);index;default:;comment:租户编号" json:"tenant_no" comment:"租户编号" ` // 租户
	OrgNo        string     `gorm:"column:org_no;type:varchar(80);index;default:;comment:组织编号" json:"org_no" comment:"组织编号" `
	StoreNo      string     `gorm:"column:store_no;type:varchar(80);index;default:;comment:店编号" json:"store_no" comment:"店编号" `
	DepartmentNo string     `gorm:"column:department_no;type:varchar(80);index;default:;comment:部门编号" json:"department_no" comment:"部门编号" `
	TypeSys      string     `gorm:"column:type_sys;type:varchar(100);index;default:;comment:类型|普通|系统|api" json:"type_sys" comment:"类型;普通;系统;api" `
	Type         string     `gorm:"column:type;type:varchar(100);index;default:;comment:类型|资源组|资源" json:"type" comment:"类型|资源组|资源" `
	TypeValue    string     `gorm:"column:type_value;type:varchar(80);index;default:;comment:对应类型id" json:"type_value" comment:"对应类型id" `
	MenuId       int64      `gorm:"column:menu_id;type:bigint;not null;index;default:0;comment:菜单id" json:"menu_id" comment:"菜单id" `
}

func (*RamMenuRelationEntity) TableName() string {
	return "ram_menu_relation"
}

func (*RamMenuRelationEntity) TableComment() string {
	return "菜单权限关系"
}
