package pg

// Domain
// @Description: 域模块
type Domain struct {
	System bool `value:"${system:=false}" label:"系统后台启用"`
}
