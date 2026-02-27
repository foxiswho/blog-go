package modRamCasbin

type UpdateCt struct {
	AuthorityId uint         `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
