package {{.Model}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
	"fmt"
    "blog/fox"
    "blog/fox/db"
)
{{else}}
import (
    "fmt"
	"blog/fox"
    "blog/fox/db"
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

//初始化
func New{{Mapper .Name}}() *{{Mapper .Name}}{
	return new({{Mapper .Name}})
}
//初始化列表
func (c *{{Mapper .Name}})newMakeDataArr() ([]{{Mapper .Name}}){
	return make([]{{Mapper .Name}}, 0)
}
//列表查询
func (c *{{Mapper .Name}})GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	session := db.Filter(q)
	count, err := session.Count(c)
	if err != nil {
		fmt.Println("数据查询错误：",err)
		return nil,fox.NewError("数据查询错误:"+err.Error())
	}
	Query := db.Pagination(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}

	session = db.Filter(q)
	if orderBy != "" {
		session.OrderBy(orderBy)
	}
	session.Limit(limit, Query.Offset)
	if len(fields) == 0 {
		session.AllCols()
	}
	data := c.newMakeDataArr()
	err = session.Find(&data)
	if err != nil {
		fmt.Println("数据查询错误:",err)
		return nil,fox.NewError("数据查询错误:"+err.Error())
	}
	Query.Data = make([]interface{}, len(data))
	for y, x := range data {
		Query.Data[y] = x
	}
	return Query, nil
}
// 获取 单条记录
func (c *{{Mapper .Name}}) GetById(id int) (*{{Mapper .Name}}, error) {
    m:=New{{Mapper .Name}}()
	{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}
	{{if $col.IsPrimaryKey}}
	m.{{Mapper $col.Name}} = id
	{{end}}
    {{end}}
    o := db.NewDb()
	ok, err := o.Get(m)
    if err != nil {
        return nil, err
    }
    if !ok{
        return nil,fox.NewError("数据不存在:"+err.Error())
    }
    return m, nil
}
// 删除 单条记录
func (c *{{Mapper .Name}}) Delete(id int) (int64, error) {
	m:=New{{Mapper .Name}}()
	{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}
	{{if $col.IsPrimaryKey}}
	m.{{Mapper $col.Name}} = id
	{{end}}
    {{end}}
    o := db.NewDb()
	num, err := o.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}
{{end}}
