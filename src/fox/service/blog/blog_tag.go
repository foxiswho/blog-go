package blog

import (
	"github.com/astaxie/beego/orm"
	"fox/models"
	"fmt"
	"fox/util"
	"time"
	"strings"
	"fox/util/array"
	"fox/model"
	"fox/util/db"
)

type BlogTag struct {

}
func NewBlogTagService() *BlogTag{
	return new(BlogTag)
}
func (c *BlogTag)Query(str string) (data []interface{}, err error) {
	query := make(map[string]interface{})
	fields := []string{}
	if str != "" {
		query["name"] = str
	}
	mode := model.NewBlogTag()
	data, err = mode.GetAll(query, fields, "tag_id desc", 0, 999)
	//fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//创建
func (c *BlogTag)Create(m *model.BlogTag) (int64, error) {

	fmt.Println("DATA:", m)
	if len(m.Name) < 1 {
		return 0, &util.Error{Msg:"标题 不能为空"}
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	o := db.NewDb()
	id, err := o.Insert(m)
	if err != nil {
		return 0, &util.Error{Msg:"创建错误：" + err.Error()}
	}
	fmt.Println("DATA:", m)
	fmt.Println("Id:", id)
	return id, nil
}
//删除
func (c *BlogTag)DeleteByName(id int, str string) (bool, error) {
	if str == "" {
		return false, &util.Error{Msg:"名称 不能为空"}
	}
	mode := model.NewBlogTag()
	mode.BlogId = id
	mode.Name = str
	o := db.NewDb()

	if num, err := o.Delete(mode); err == nil {
		fmt.Println("Number of records deleted in database:", num)
		return true, nil
	}
	return false, nil
}
//根据
func (c *BlogTag)GetBlogTagCheckName(str string) (model.BlogTag, error) {
	mode := model.NewBlogTag()
	mode.Name = str
	o := db.NewDb()
	err := o.Find(mode, "name")
	if err == nil {
		return mode, nil
	}
	return nil, err
}
//创建 和删除
func (c *BlogTag)CreateFromTags(id int, tag, old string) (bool, error) {
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
	o := db.NewDb()
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
				mode := model.NewBlogTag()
				mode.Name = v
				mode.BlogId = id
				_, _ = o.Insert(mode)
			} else {
				check[v] = false
				if array.SliceContains(olds, v) {
					check[v] = true
					continue
				}
				mode := model.NewBlogTag()
				mode.Name = v
				mode.BlogId = id
				_, _ = o.Insert(mode)
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
						ok, err := c.DeleteByName(id, val)
						fmt.Println(ok)
						fmt.Println(err)
					}
				}
			} else {
				//删除所有
				ok, err := c.DeleteByName(id, val)
				fmt.Println(ok)
				fmt.Println(err)
			}

		}
	}

	return false, nil
}