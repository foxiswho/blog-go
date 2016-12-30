package db

import "math"

type Paginator struct {
	Page        int		//当前页
	Pages       []int	//页码数组
	PageSize    int		//每页条数
	TotalPage   int		//总页码
	TotalCount  int		//总条数
	FirstPage   int
	FirstPageIs bool
	LastPageIs  bool
	LastPage    int
	Data        []interface{}
	OtherData   map[string]interface{}
	Offset      int
}
// 分页  总数，当前页，每页条数
func Pagination(count int, page int, pageSize int) (*Paginator) {
	Page := new(Paginator)
	Page.PageSize = pageSize
	Page.TotalCount = count
	//Page.TotalPage =  int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if count % pageSize > 0 {
		Page.TotalPage = count / pageSize + 1
	}
	if page > Page.TotalPage {
		page = Page.TotalPage
	}
	if page <= 0 {
		Page.Page = 1
	}
	//当前页
	page = Page.Page
	Page.FirstPageIs = page == 1
	Page.LastPageIs = page == Page.TotalPage
	//读取起始条数
	Page.Offset = (page - 1) * pageSize
	Page.Pages = make([]int, 5)
	var pages []int
	switch {
	case page >= Page.TotalPage - 5 && Page.TotalPage > 5: //最后5页
		start := Page.TotalPage - 5 + 1
		Page.FirstPage = page - 1
		Page.LastPage = int(math.Min(float64(Page.TotalPage), float64(page + 1)))
		for i, _ := range pages {
			Page.Pages[i] = start + i
		}
	case page >= 3 && Page.TotalPage > 5:
		start := page - 3 + 1
		Page.FirstPage = page - 3
		for i, _ := range pages {
			Page.Pages[i] = start + i
		}
		Page.FirstPage = page - 1
		Page.LastPage = page + 1
	default:
		Page.Pages = make([]int, int(math.Min(5, float64(Page.TotalPage))))
		for i, _ := range pages {
			Page.Pages[i] = i + 1
		}
		Page.FirstPage = int(math.Max(float64(1), float64(page - 1)))
		Page.LastPage = page + 1
	}
	return Page
}