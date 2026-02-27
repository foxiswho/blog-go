package pg

type Multi struct {
	Tenant   MultiItem `json:"tenant" value:"${tenant:=}" toml:"tenant"`       //租户表
	Org      MultiItem `json:"org" value:"${org:=}" toml:"org"`                //组织机构表
	Merchant MultiItem `json:"merchant" value:"${merchant:=}" toml:"merchant"` //商户表
	Store    MultiItem `json:"store" value:"${store:=}" toml:"store"`          //店铺表
	Owner    MultiItem `json:"owner" value:"${owner:=}" toml:"owner"`          //所有者
}
