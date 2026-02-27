package pg

// app.template
// 模版
type Template struct {
	Type     string `json:"type"`     //模板类型 PONGO2
	Engine   string `json:"engine"`   //模板引擎 PONGO2,TEMPLATE
	Path     string `json:"path"`     //路径 ,读取模板相对路径 template/pongo2
	FileData string `json:"fileData"` //读取文件位置,BIDATA 内,还是 当前文件夹
	Suffix   string `json:"suffix"`   //后缀
}
