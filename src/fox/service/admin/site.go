package admin

import (
	"fox/util/db"
	"fox/service"
	"net/url"
	"fox/util"
	"strings"
	"fox/model"
	"fmt"
)

type Site struct {

}

func NewSiteService() *Site {
	return new(Site)
}
//列表
func (t *Site)Query() (*db.Paginator, error) {
	return NewTypeService().Query(service.SITE_ID)
}
//更新
func (t *Site)Update(form url.Values) (bool, error) {
	if len(form) < 1 {
		return false, &util.Error{Msg:"站点信息 不能为空"}
	}

	o := db.NewDb()
	for key, v := range form {
		key = strings.TrimSpace(key)
		val := strings.TrimSpace(v[0])
		if key == "SITE_NAME" && len(val) < 1 {
			return false, &util.Error{Msg:"站点名称 不能为空"}
		}
		if val != "" {
			mod := model.NewType()
			mod.Content=val
			num, err := o.Where("type_id=? and mark=? ", service.SITE_ID, key).Update(mod)
			if err != nil {
				return false, err
			}
			fmt.Println("更新 " + key + "=>" + val,num)
		}else{
			fmt.Println(key + "值为空 跳过更新")
		}

	}
	return true, nil
}