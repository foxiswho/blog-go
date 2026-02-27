package excelPg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	DATE_PATTERN      = "2006-01-02"
	DATE_TIME_PATTERN = "2006-01-02 15:04:05"
)

// Field
// 字段
type ExcelFields struct {
	//name
	Name string //名称
	//Index starts at 0
	Index int //索引  从0 开始
	//JSON field name
	Field string //json 字段名称
	//Field type
	FieldType string //字段类型
	//Save all tags
	Tags map[string]string // 保存所有tags
}

// comprehensive
// 综合
type ExcelStruct struct {
	// index sort
	MapIndex map[int]string //按照 index 排序
	//max index
	IndexMax int // index 最大
	//All fields
	Fields map[string]ExcelFields //所有字段
	//The first few lines start with specific data
	StartRow int //第几行开始为具体数据
	//error
	Err error //错误
	//During type conversion, whether to directly prompt to report an error when an error occurs
	ConvertTypeErr bool //类型转换时候,产生错误时是否直接提示报错
}

// By default, it starts from the first row and the index starts from 0
// 默认 从第一行开始,索引从 0开始
func NewExcelStructDefault() *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = 1
	n.IndexMax = 10
	return n
}

// StartRow starts row, index starts from 0
// Indexmax indexes the maximum row. If the index in the structure is larger than the configured, the index in the structure is used
// StartRow 开始行,索引从 0开始
// IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
func NewExcelStruct(StartRow, IndexMax int) *ExcelStruct {
	n := new(ExcelStruct)
	n.StartRow = StartRow
	n.IndexMax = IndexMax
	return n
}

type Callback func(maps map[string]interface{}) error

// Struct pointer
// 结构体 指针
func (c *ExcelStruct) SetPointerStruct(ptr interface{}) *ExcelStruct {
	//Gets the type of the input parameter
	// 获取入参的类型
	t := reflect.TypeOf(ptr)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		//Argument should be a struct pointer
		c.Err = fmt.Errorf("参数应该为结构体指针")
		return c
	}
	//Take the structure variable that the pointer points to
	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()
	//Parsing fields
	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		//Take tag
		// 取tag
		fieldInfo := v.Type().Field(i)
		//
		fields := ExcelFields{}
		tag := fieldInfo.Tag
		//Parsing
		// 解析
		fields.Field = tag.Get("json")
		if fields.Field == "" {
			fields.Field = fieldInfo.Name
		}
		fields.Name = tag.Get("name")
		if fields.Name == "" {
			fields.Name = fieldInfo.Name
		}
		index := 0
		indexStr := tag.Get("index")
		if indexStr != "" {
			index, _ = strconv.Atoi(indexStr)
		}
		//If the index is large, the value is assigned
		//如果索引大 那么赋值
		if c.IndexMax < index {
			c.IndexMax = index
		}
		fields.Index = index
		fields.FieldType = fieldInfo.Type.String()
		m := make(map[string]string)
		m["json"] = fields.Field
		m["name"] = fields.Name
		m["index"] = strconv.Itoa(i)
		fields.Tags = m
		//
		if c.Fields == nil {
			c.Fields = make(map[string]ExcelFields)
		}
		c.Fields[fields.Field] = fields
		if c.MapIndex == nil {
			c.MapIndex = make(map[int]string)
		}
		c.MapIndex[index] = fields.Field
		// 解析uppercase tag
		//value := fmt.Sprintf("%v", v.Field(i))
		//if fieldInfo.TypeCategory.Kind() == reflect.String {
		//	uppercase := tag.Get("uppercase")
		//	if uppercase == "true" {
		//		value = strings.ToUpper(value)
		//	} else {
		//		value = strings.ToLower(value)
		//	}
		//}
	}
	return c
}

// process
// 处理
func (c *ExcelStruct) RowsProcess(rows [][]string, callback Callback) error {
	return c.RowsAllProcess(rows, callback)
}

// process
// 处理
func (c *ExcelStruct) RowsAllProcess(rows [][]string, callback Callback) error {
	if c.Fields == nil {
		//Please fill in the structure pointer
		return fmt.Errorf("请填写结构体指针")
	}
	if c.Err != nil {
		return c.Err
	}
	//data := []interface{}{}
	for index, row := range rows {
		//If the index is less than the set start row, skip
		//如果 索引 小于 已设置的 开始行,那么跳过
		if index < c.StartRow {
			continue
		}
		//单行处理
		maps, err := c.Row(row)
		if err != nil {
			return err
		}
		//json1, err := convUtil.ObjToJson(maps)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Println("MAP=>JSON")
		//fmt.Println(json1)
		//err = json.Unmarshal([]byte(json1), &tmp)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//fmt.Println("MAP=>JSON,JSON->STRUCT")
		//fmt.Println(tmp)
		//map 转为 struct
		//if err := mapstructure.Decode(maps, &ptr); err != nil {
		//	return nil, err
		//}
		//data = append(data, ptr)
		err2 := callback(maps)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

// process row
// 处理 单行
func (c *ExcelStruct) Row(row []string) (map[string]interface{}, error) {
	if c.Fields == nil {
		//Please fill in the structure pointer
		return nil, fmt.Errorf("请填写结构体指针")
	}
	if c.Err != nil {
		return nil, c.Err
	}
	maps := make(map[string]interface{})
	for i, colCell := range row {
		//Cannot judge null value, otherwise
		//不能判断空值,否则
		if len(colCell) < 1 {
			continue
		}
		//Determine whether the key name exists
		//判断键名是否存在
		if field, ok := c.MapIndex[i]; ok {
			maps[field] = ""
			//TypeCategory conversion
			//类型转换
			fields := c.Fields[field]
			//character
			//字符
			if fields.FieldType == "string" {
				maps[field] = colCell
				continue
			}
			//time
			//时间
			if fields.FieldType == "time.Time" && len(colCell) > 0 {
				t, err := time.ParseInLocation(DATE_TIME_PATTERN, colCell, time.Local)
				if err == nil {
					maps[field] = t
				} else {
					//During type conversion, whether to directly prompt to report an error when an error occurs
					//类型转换时候,产生错误时是否直接提示报错
					if c.ConvertTypeErr {
						return nil, err
					}
				}
			} else {
				//other
				//其他类型
				switch fields.FieldType {
				case "bool":
					lower := strings.ToLower(colCell)
					if lower == "true" {
						maps[field] = true
					} else {
						maps[field] = false
					}
				case "int":
					ret, err := strconv.Atoi(colCell)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "int8":
					ret, err := strconv.ParseInt(colCell, 10, 8)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "int16":
					ret, err := strconv.ParseInt(colCell, 10, 16)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "int32":
					ret, err := strconv.ParseInt(colCell, 10, 32)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "int64":
					ret, err := strconv.ParseInt(colCell, 10, 64)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
					//fmt.Println("int64=", int)
				case "uint":
					ret, err := strconv.Atoi(colCell)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = uint(ret)
					}
				case "uint8":
					ret, err := strconv.ParseUint(colCell, 10, 8)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "uint16":
					ret, err := strconv.ParseUint(colCell, 10, 16)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "uint32":
					ret, err := strconv.ParseUint(colCell, 10, 32)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "uint64":
					ret, err := strconv.ParseUint(colCell, 10, 64)
					if err != nil {
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "float32":
					ret, err := strconv.ParseFloat(colCell, 32)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "float64":
					ret, err := strconv.ParseFloat(colCell, 64)
					if err != nil {
						//During type conversion, whether to directly prompt to report an error when an error occurs
						//类型转换时候,产生错误时是否直接提示报错
						if c.ConvertTypeErr {
							return nil, err
						}
						maps[field] = 0
					} else {
						maps[field] = ret
					}
				case "string":
					maps[field] = colCell
				}
			}
		}
	}
	return maps, nil
}
