package pg

// 上传
type Upload struct {
	UploadFolder     string `json:"uploadFolder" value:"${uploadFolder}"`         // 根目录
	StaticAccessPath string `json:"staticAccessPath" value:"${staticAccessPath}"` // 相对路径
	SaveLocal        bool   `json:"saveLocal" value:"${saveLocal}"`               // 保存到本地
	SaveOss          bool   `json:"saveOss" value:"${saveOss}"`                   // 保存到云存储
}
