package db

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	Data       []interface{}
	Offset     int
}
// 分页  总数，当前页，每页条数
func Pagination(count int, pageNo int, pageSize int) (*Page) {
	page := &Page{PageNo: pageNo, PageSize: pageSize, TotalCount: count}
	page.TotalPage = count / pageSize
	if count % pageSize > 0 {
		page.TotalPage = count / pageSize + 1
	}
	page.FirstPage = pageNo == 1
	page.LastPage = pageNo == page.TotalPage
	page.Offset = (pageNo - 1) * pageSize
	return page
}