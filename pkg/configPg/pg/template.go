package pg

// Template
// 模版
type Template struct {
	Theme string `json:"theme" value:"${theme:=nisarg}" label:"主题"`
	//Engine   string `json:"engine"`   //模板引擎 PONGO2,TEMPLATE
	//Path string `json:"path"` //路径 ,读取模板相对路径 template/pongo2
	//FileData string `json:"fileData"` //读取文件位置,BIDATA 内,还是 当前文件夹
	Suffix string `json:"suffix" value:"${suffix:=html}" label:"后缀"`
}
