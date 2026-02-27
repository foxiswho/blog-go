package cacheDiplPg

type DiplCo struct {
	Name     string `json:"name"`
	No       string `json:"no"`
	TenantNo string `json:"tenantNo"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	HashSha  string `json:"hashSha"`
}
