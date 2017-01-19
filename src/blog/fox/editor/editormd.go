package editor
//markdown编辑器
type EditorMd struct {
	Success int      `json:"success"`     // 0 表示上传失败，1 表示上传成功
	Message string   `json:"message"`    // 提示的信息，上传成功或上传失败及错误信息等
	Url     string   `json:"url"`   // 图片地址       // 上传成功时才返回
}