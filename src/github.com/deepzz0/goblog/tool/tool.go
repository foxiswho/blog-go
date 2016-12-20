package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type StatusItem struct {
	name string
	desc string
}

const (
	CODE_DIR = "../RS/resp.go"
	DESC_DIR = "../RS/desc.go"
)

func main() {
	workpath, _ := os.Getwd()
	workpath, _ = filepath.Abs(workpath)
	f, err := ioutil.ReadFile(CODE_DIR)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := map[int]StatusItem{}
	lines := strings.Split(string(f), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "RS_") {
			list0 := strings.Split(line, "=")
			name := list0[0]
			list1 := strings.Split(list0[1], "//")
			code, err := strconv.Atoi(strings.TrimSpace(list1[0]))
			if err != nil {
				fmt.Printf("On resp.go line:%d, %s", i, err.Error())
				return
			}
			desc := strings.TrimSpace(list1[1])
			result[code] = StatusItem{name: name, desc: desc}
		}
	}

	content := ""
	format := `package RS

// game server rpc status code and relevant description table
// Auto created by %s
//

var descDict = map[int]string{
%s}

func Desc(code int) string {
    desc, found := descDict[code]
    if !found {
		return "未定义状态"
    }
    return desc 
}
`
	sortedIdList := make([]int, 0, len(result))
	for id, _ := range result {
		sortedIdList = append(sortedIdList, id)
	}
	sort.Ints(sortedIdList)
	for _, id := range sortedIdList {
		name := result[id].name
		desc := result[id].desc
		content += fmt.Sprintf("\t%s: \"%s\",\n", name, desc)
	}
	output := fmt.Sprintf(format, workpath, content)
	if err := ioutil.WriteFile(DESC_DIR, []byte(output), 0644); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("->updated ", DESC_DIR+"!")
}
