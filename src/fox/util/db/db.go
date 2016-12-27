package db

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/xormplus/xorm"
)

type Db struct {
	Db *xorm.Engine
}

var DB *Db
//
func dsn() string {
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_name := beego.AppConfig.String("db_name")
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8"
	return dsn
}
//初始化
func Init() {
	var err error
	DB=&Db{}
	DB.Db, err = xorm.NewEngine("mysql", dsn())
	if err != nil {
		fmt.Println("NewEngine", err)
	}
}
// NewDb create new db
func NewDb() *xorm.Engine {
	return DB.Db
}

type DbSql struct {

}

func (c *Db)Filter(where map[string]interface{})  {
	//db:=c.Db
	//str := ""
	//for k, v := range where {
	//	//fmt.Println(k, v, reflect.TypeOf(v))
	//	//fmt.Println("?号个数为", strings.Count(k, "?"))
	//	QuestionMarkCount := strings.Count(k, "?")
	//	isEmpty := false
	//	isMap := false
	//	switch v.(type) {
	//	case string:
	//		//是字符时做的事情
	//		isEmpty = v == ""
	//	case int:
	//	//是整数时做的事情
	//	case []string :
	//		isMap = true
	//		isEmpty = len(v) == 0
	//	case []int :
	//		isMap = true
	//		isEmpty = len(v) == 0
	//	}
	//	if QuestionMarkCount == 0 && isEmpty {
	//
	//		str += " AND " + k + " = '' "
	//	} else if QuestionMarkCount == 0 && !isEmpty {
	//		//是数组
	//		if (isMap) {
	//			str += " AND " + k + " = " +
	//		} else {
	//			//不是数组
	//			str += " AND " + k + " = " + v
	//		}
	//	} else if QuestionMarkCount == 1 && isEmpty {
	//		//值为空字符串,不是数组
	//		str += " AND " + k + " = " + v
	//	} else if QuestionMarkCount == 1 && !isEmpty {
	//		//是数组
	//		if isMap {
	//			str += " AND " + k + " = " + JsonEnCode(v)
	//		} else {
	//			//不是数组
	//			//不是数组，有值
	//			str += " AND " + k + " = '' "
	//		}
	//	} else if QuestionMarkCount > 1 && isEmpty {
	//		//不是数组，空值
	//		str += " AND " + k + " = ''"
	//	} else if QuestionMarkCount > 1 && !isEmpty & isMap {
	//		count := len(v)
	//		//问号 与  数组相同时
	//		if QuestionMarkCount == count {
	//			//不是数组，空值
	//			str += " AND " + k + " = ''"
	//		}else{
	//			//问号 与  数组不同时
	//			str += " AND " + k + " = ''"
	//		}
	//	}else {
	//		fmt.Println("其他还没有收录")
	//	}
	//}
}