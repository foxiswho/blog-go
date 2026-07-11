package modSysPub

import "github.com/hongmengzhu/xianfu-blog-go/pkg/model"

type AuthConfigVo struct {
	LoginEncrypt bool                  `json:"loginEncrypt" lable:"登陆加密"`
	AuthPubPrive model.AuthPubPriveDto `json:"pubPrive" label:"密钥"`
	Menu         RamMenuRouterVo       `json:"menuRouter" label:"菜单路由"`
}

func NewAuthConfig(dto model.AuthPubPriveDto) AuthConfigVo {
	return AuthConfigVo{
		AuthPubPrive: dto,
	}
}
func NewAuthConfigPublic(dto model.AuthPubPriveDto) AuthConfigVo {
	return AuthConfigVo{
		AuthPubPrive: model.AuthPubPriveDto{
			PublicKey: dto.PublicKey,
			Type:      dto.Type,
			No:        dto.No,
		},
	}
}
