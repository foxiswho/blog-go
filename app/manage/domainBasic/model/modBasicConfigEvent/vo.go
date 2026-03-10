package modBasicConfigEvent

import (
	"time"

	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"gorm.io/datatypes"
)

type Vo struct {
	ID       typePg.Int64String `json:"id" label:"id" `
	CreateAt *time.Time         `json:"createAt" label:"创建时间" `
	UpdateAt *time.Time         `json:"updateAt" label:"更新时间" `
	CreateBy string             `json:"createBy" label:"创建人" `
	UpdateBy string             `json:"updateBy" label:"更新人" `
	State    typePg.Int8        `json:"state" label:"状态" `
	Sort     typePg.Int64String ` json:"sort" comment:"排序" `
	//
	Description     string         `json:"description" label:"描述" ` // 描述
	Name            string         `json:"name" label:"名称" `
	LangCode        string         `json:"langCode" comment:"语言" `
	TypeSys         string         `json:"typeSys" comment:"类型|普通|系统|api" `
	Module          string         `json:"module" comment:"模块" `
	ModuleSub       string         `json:"moduleSub" comment:"子模块" `
	Owner           string         `json:"owner" comment:"所属/拥有者" `
	Model           string         `json:"model" comment:"模型" `
	Show            typePg.Int8    `json:"show" comment:"1显示2隐藏" `
	ExtraData       datatypes.JSON `json:"extraData" label:"额外数据" `
	Value           string         `json:"value" comment:"值" `
	Client          string         `json:"client" comment:"端" `
	ClientProgram   string         `json:"clientProgram" comment:"端内程序|隔开" `
	ClientSync      string         `json:"clientSync" comment:"端同步" `
	LoadingLocation string         `json:"loadingLocation" comment:"加载位置" `
	Cache           string         `json:"cache" comment:"缓存" `
	No              string         `json:"no" comment:"编号" `
}
