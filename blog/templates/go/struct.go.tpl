package {{.Model}}

import (
	{{range .Imports}}"{{.}}"{{end}}
)

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .Columns}}	{{Mapper .Name}}	{{Type .}}
{{end}}
}

{{end}}