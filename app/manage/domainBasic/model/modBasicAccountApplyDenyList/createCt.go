package modBasicAccountApplyDenyList

type CreateCt struct {
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" `
	Description string `json:"description" label:"描述" `
	TypeSys     string `json:"typeSys" label:"系统类型" `
	TypeDomain  string `json:"typeDomain" label:"域类型" `
	TypeField   string `json:"typeField" label:"字段类型" `
	TypeExpr    string `json:"typeExpr" label:"表达式类型" `
	Expr        string `json:"expr" label:"表达式" `
	Message     string `json:"message" label:"错误时消息" `
}
