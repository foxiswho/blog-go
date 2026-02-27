package entityRam

import (
	"time"
)

// RamMenuEntity 菜单
type RamMenuEntity struct {
	ID          int64      `gorm:"column:id;type:bigserial;primaryKey;autoIncrement:true" json:"id" comment:"" `
	No          string     `gorm:"column:no;type:varchar(80);index;default:;comment:编号" json:"no" comment:"编号" `
	Name        string     `gorm:"column:name;type:varchar(255);comment:名称" json:"name" comment:"名称" `                                                             // 名称
	NameFl      string     `gorm:"column:name_fl;type:varchar(255);comment:名称外文" json:"name_fl" comment:"名称外文" `                                                   // 名称外文
	Code        string     `gorm:"column:code;type:varchar(80);comment:路由代号" json:"code" comment:"路由代号" `                                                          // 编号代号
	NameFull    string     `gorm:"column:name_full;type:varchar(255);comment:全称" json:"name_full" comment:"全称" `                                                   // 全称
	State       int8       `gorm:"column:state;type:int2;not null;index;default:1;comment:状态|1启用|2禁用" json:"state" comment:"状态:1启用;2禁用" `                          // 状态:1启用;2禁用
	Description string     `gorm:"column:description;type:varchar(255);comment:描述" json:"description" comment:"描述" `                                               // 描述
	CreateAt    *time.Time `gorm:"column:create_at;type:timestamptz;index;autoCreateTime;default:current_timestamp;comment:创建时间" json:"create_at" comment:"创建时间" ` // 创建时间
	UpdateAt    *time.Time `gorm:"column:update_at;type:timestamptz;autoUpdateTime;comment:更新时间" json:"update_at" comment:"更新时间" `                                 // 更新时间
	CreateBy    string     `gorm:"column:create_by;type:varchar(80);index;default:;comment:创建人" json:"create_by" comment:"创建人" `
	UpdateBy    string     `gorm:"column:update_by;type:varchar(80);default:;comment:更新人" json:"update_by" comment:"更新人" `
	Sort        int64      `gorm:"column:sort;type:bigint;not null;default:0;;comment:排序" json:"sort" comment:"排序" ` // 排序
	TypeSys     string     `gorm:"column:type_sys;type:varchar(100);index;default:;comment:类型|普通|系统|api" json:"type_sys" comment:"类型;普通;系统;api" `
	Path        string     `gorm:"column:path;type:varchar(5000);comment:路由路径" json:"path" comment:"路由路径" `
	Method      string     `gorm:"column:method;type:varchar(255);comment:方法" json:"method" comment:"方法" `
	Show        int64      `gorm:"column:show;type:int2;not null;default:2;comment:列表显示1是2否" json:"show" comment:"列表显示1是2否" `
	Component   string     `gorm:"column:component;type:varchar(2000);comment:对应前端文件路径" json:"component" comment:"对应前端文件路径" `
	ActiveName  string     `gorm:"column:active_name;type:varchar(2000);comment:高亮菜单" json:"active_name" comment:"高亮菜单" `
	KeepAlive   int8       `gorm:"column:keep_alive;type:int2;not null;default:2;comment:缓存1是2否" json:"keep_alive" comment:"缓存1是2否" `
	Icon        string     `gorm:"column:icon;type:varchar(2000);comment:菜单图标" json:"icon" comment:"菜单图标" `
	CloseTab    int8       `gorm:"column:close_tab;type:int2;not null;default:2;comment:关闭tab1是2否" json:"close_tab" comment:"关闭tab1是2否" `
	ParentId    string     `gorm:"column:parent_id;type:varchar(80);index;default:;comment:上级" json:"parent_id" comment:"上级" `
	IdLink      string     `gorm:"column:id_link;type:text;comment:上级" json:"parent_id_link" comment:"上级" `
	TypeAttr    string     `gorm:"column:type_attr;type:varchar(100);index;default:menu;comment:属性|菜单分类|资源" json:"type_attr" comment:"属性;菜单menu;按钮button;资源:resource;其他other" `
	TypeDomain  string     `gorm:"column:type_domain;type:varchar(80);index;default:'default';comment:域|平台platform|店铺shop|其他other" json:"type_domain" comment:"域;平台platform;店铺shop;其他other" `
	ParentNo    string     `gorm:"column:parent_no;type:varchar(80);index;default:;comment:上级编号" json:"parent_no" comment:"上级编号" `
	NoLink      string     `gorm:"column:no_link;type:text;comment:上级编号" json:"no_link" comment:"上级编号" `
}

func (*RamMenuEntity) TableName() string {
	return "ram_menu"
}

func (*RamMenuEntity) TableComment() string {
	return "菜单"
}
