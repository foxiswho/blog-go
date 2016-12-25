package blog

import (
	"github.com/astaxie/beego/orm"
	"fox/models"
	"fmt"
	"fox/util"
	"time"
	"strings"
	"fox/util/array"
)

type BlogTag struct {

}
func (c *BlogTag)Query(str string) (data []interface{}, err error) {
	query := map[string]string{}
	if str!=""{
		query["name"] = str
	}
	var fields []string
	sortby := []string{"Id"}
	order := []string{"desc"}
	var offset int64
	var limit int64
	offset = 0
	limit = 100
	data, err = models.GetAllBlogTag(query, fields, sortby, order, offset, limit)
	//fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//创建
func (c *BlogTag)Create(m *models.BlogTag) (int64, error) {

	fmt.Println("DATA:", m)
	if len(m.Name) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	id, err := models.AddBlogTag(m)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//删除
func (c *BlogTag)DeleteByName(id int,str string) (bool, error) {
	if str == "" {
		return false, &util.Error{Msg:"名称 不能为空"}
	}
	o := orm.NewOrm()
	v := models.BlogTag{Name:str,BlogId:id}
	if num, err := o.Delete(&v, "name","blog_id"); err == nil {
		fmt.Println("Number of records deleted in database:", num)
		return true, nil
	}
	return false, nil
}
//根据自定义伪静态查询
func (c *BlogTag)GetBlogTagCheckName(str string) (models.BlogTag, error) {
	o := orm.NewOrm()
	v := models.BlogTag{Name:str}
	err := o.Read(&v, "name")
	if err == nil {
		return v, nil
	}
	return models.BlogTag{}, err
}
//创建 和删除
func (c *BlogTag)CreateFromTags(id int,tag, old string) (bool, error) {
	fmt.Println("CreateFromTags:")
	//if tag == "" {
	//	return false, nil
	//}
	fmt.Println("DATA:", tag)
	var olds, tags []string
	check := make(map[string]bool)
	if old != "" {
		olds = strings.Split(old, ",")
	}
	if tag != "" {
		//拆分成数组
		tags = strings.Split(tag, ",")
		fmt.Println(tags)
		//创建
		for _, v := range tags {
			if v == "" {
				continue
			}
			//fmt.Println(k,v)
			if old == "" {
				model := &models.BlogTag{Name:v}
				model.BlogId=id
				_, _ = c.Create(model)
			} else {
				check[v] = false
				if array.SliceContains(olds, v) {
					check[v] = true
					continue
				}
				model := &models.BlogTag{Name:v}
				model.BlogId=id
				_, _ = c.Create(model)
			}
		}
	}
	//旧 tag 检测
	if old != "" {
		for _, val := range olds {
			if tag != "" {
				if !check[val] {
					//没有，从数据库里删除
					if !array.SliceContains(tags, val) {
						ok, err := c.DeleteByName(id,val)
						fmt.Println(ok)
						fmt.Println(err)
					}
				}
			} else {
				//删除所有
				ok, err := c.DeleteByName(id,val)
				fmt.Println(ok)
				fmt.Println(err)
			}

		}
	}

	return false, nil
}