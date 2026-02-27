package modRamResourceMenu

import "github.com/foxiswho/blog-go/pkg/tools/typePg"

type UpdateByMenuCt struct {
	MenuId typePg.Int64String   `json:"menuId" validate:"required" label:"菜单id" `
	Data   []ResourceAndGroupVo `json:"data" validate:"required" label:"数据" `
}
