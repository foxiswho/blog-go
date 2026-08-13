package modSysPub

// RamMenuRouterMetaVo 菜单路由元数据
type RamMenuRouterMetaVo struct {
	// 标题名称
	Title string `json:"title"`
	// 用于路由->菜单排序
	Order int `json:"order"`
	// 激活图标（菜单/tab）
	ActiveIcon string `json:"activeIcon"`
	// 当前激活的菜单，有时候不想激活现有菜单，需要激活父级菜单时使用
	ActivePath string `json:"activePath"`
	// 是否固定标签页
	AffixTab bool `json:"affixTab"`
	// 固定标签页的顺序
	AffixTabOrder int `json:"affixTabOrder"`
	// 需要特定的角色标识才可以访问
	Authority []string `json:"authority"`
	// 徽标
	Badge string `json:"badge"`
	// 徽标类型 dot normal
	BadgeType string `json:"badgeType"`
	// 徽标颜色
	BadgeVariants string `json:"badgeVariants"`
	// 当前路由的子级在菜单中不展现
	HideChildrenInMenu bool `json:"hideChildrenInMenu"`
	// 当前路由在面包屑中不展现
	HideInBreadcrumb bool `json:"hideInBreadcrumb"`
	// 当前路由在菜单中不展现
	HideInMenu bool `json:"hideInMenu"`
	// 当前路由在标签页不展现
	HideInTab bool `json:"hideInTab"`
	// 图标（菜单/tab）
	Icon string `json:"icon"`
	// iframe 地址
	IframeSrc string `json:"iframeSrc"`
	// 外链-跳转路径
	Link string `json:"link"`
	// 忽略权限，直接可以访问
	IgnoreAccess bool `json:"ignoreAccess"`
	// 开启KeepAlive缓存
	KeepAlive bool `json:"keepAlive"`
	// 菜单可以看到，但是访问会被重定向到403
	MenuVisibleWithForbidden bool `json:"menuVisibleWithForbidden"`
	// 在新窗口打开
	OpenInNewWindow bool `json:"openInNewWindow"`
	// 后端接口
	Api string `json:"api"`
	// 方法
	Method string `json:"method"`
}
