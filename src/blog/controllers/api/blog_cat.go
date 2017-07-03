package api

import (
	_ "github.com/go-sql-driver/mysql"
	"blog/controllers"
	"encoding/json"
	"fmt"
)

//
type BlogCat struct {
	controllers.BaseNoLogin
}

func (c *BlogCat) GetAll() {
	var result map[string]interface{}
	s := "{\"code\":1,\"info\":\"ok\",\"data\":[{\"cat_id\":1,\"name\":\"栏目一\"},{\"cat_id\":2,\"name\":\"栏目二\"}]}"
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	c.Data["json"] = result
	c.ServeJSON()
}
