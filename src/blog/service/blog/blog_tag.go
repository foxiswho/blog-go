package blog

import (
	"fmt"
	"blog/fox"
	"time"
	"strings"
	"blog/fox/array"
	"blog/model"
	"blog/fox/db"
	"blog/fox/editor"
)
//标签
type BlogTag struct {

}
//快速初始化
func NewBlogTagService() *BlogTag {
	return new(BlogTag)
}
//创建
func (c *BlogTag)Create(m *model.BlogTag) (int, error) {

	fmt.Println("DATA:", m)
	if len(m.Name) < 1 {
		return 0,fox.NewError("标题 不能为空")
	}
	//时间
	if m.TimeAdd.IsZero() {
		m.TimeAdd = time.Now()
	}
	o := db.NewDb()
	affected, err := o.Insert(m)
	if err != nil {
		return 0,fox.NewError("创建错误" + err.Error())
	}
	fmt.Println("affected:", affected)
	fmt.Println("DATA:", m)
	fmt.Println("Id:", m.TagId)
	return m.TagId, nil
}
//删除
func (c *BlogTag)DeleteByName(id int, str string) (bool, error) {
	if str == "" {
		return false,fox.NewError("名称 不能为空")
	}
	mode := model.NewBlogTag()
	//mode.BlogId = id
	//mode.Name = str
	o := db.NewDb()

	if num, err := o.Where("blog_id=? and name=?", id, str).Delete(mode); err == nil {
		fmt.Println("Number of records deleted in database:", num)
		return true, nil
	}
	return false, nil
}
//根据名字获取数据
func (c *BlogTag)GetBlogTagCheckName(str string) (*model.BlogTag, error) {
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
		fmt.Println("拆分数组", tags)
		//创建
		for i, v := range tags {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			tags[i] = v
			//fmt.Println(k,v)
			if old == "" {
				mode := model.NewBlogTag()
				mode.Name = v
				mode.BlogId = id
				_, err := o.Insert(mode)
				fmt.Println("err", err)
			} else {
				check[v] = false
				if array.SliceContains(olds, v) {
					check[v] = true
					continue
				}
				mode := model.NewBlogTag()
				mode.Name = v
				mode.BlogId = id
				_, err := o.Insert(mode)
				fmt.Println("err", err)

			}
		}
	}
	//旧 tag 检测
	if old != "" {
		for i, val := range olds {
			val = strings.TrimSpace(val)
			tags[i] = val
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
//列表
func (c *BlogTag)GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {

	mode := model.NewBlogTag()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	ids := make([]int, data.TotalCount)
	for i, x := range data.Data {
		r := x.(model.BlogTag)
		ids[i] = r.BlogId
	}
	o := db.NewDb()
	blogs := make([]model.Blog, 0)
	err = o.Id(ids).Find(&blogs)
	if err != nil {
		blogs = []model.Blog{}
		fmt.Println(err)
	}
	//fmt.Println(blogs)
	stat := make([]model.BlogStatistics, 0)
	err = o.In("blog_id", ids).Find(&stat)
	if err != nil {
		stat = []model.BlogStatistics{}
		fmt.Println(err)
	}
	for i, x := range data.Data {
		tmp := x.(model.BlogTag)
		row := &Blog{}
		for _, r := range blogs {
			if tmp.BlogId == r.BlogId {
				//内容转换
				r.Content = string(editor.Markdown([]byte(r.Content)))
				row.Blog = &r
				row.Tags = []string{}
				if row.Tag != "" {
					row.Tags = strings.Split(row.Tag, ",")
				}
			}
		}
		row.BlogStatistics = &model.BlogStatistics{}
		for _, v := range stat {
			//fmt.Println(v)
			if (v.BlogId == tmp.BlogId) {
				row.Comment = v.Comment
				row.BlogStatistics.Read = v.Read
				row.SeoDescription = v.SeoDescription
				row.SeoKeyword = v.SeoKeyword
				row.SeoTitle = v.SeoTitle
				//fmt.Println(">>>>",row.BlogStatistics)
			}
		}
		fmt.Println("===", row.Blog)
		data.Data[i] = &row
	}

	return data, nil
}
//根据博客ID删除数据
func (c *BlogTag)DeleteByBlogId(id int) (int64, error) {
	mod := model.NewBlogTag()
	o := db.NewDb()
	num, err := o.Where("blog_id=?",id).Delete(mod)
	if err != nil {
		return 0, err
	}
	fmt.Println("删除条数", num)
	return num, nil
}