package entity
//分页
type Page struct {
	Page  int `json:"page" 当前页`
	Count int `json:"count" 总记录数`
	Size  int `json:"size" 每页记录数`
	List  []Article        `json:"-" `
}
