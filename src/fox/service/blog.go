package service

import (
	"fox/models"
	"fmt"
	"fox/util"
)

type Blog struct {

}

func (this *Blog)Get() (data []interface{}, err error) {
	var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var offset int64
	var limit int64
	offset = 0
	limit = 20
	data, err = models.GetAllBlog(query, fields, sortby, order, offset, limit)
	fmt.Println(data)
	fmt.Println(err)
	return data, err
}
//详情
func (this *Blog)Detail(id int) (data interface{}, err error) {
	if id < 1 {
		return nil, &util.Error{Msg:"ID 错误"}
	}
	data, err = models.GetBlogById(id)
	if err != nil {
		return nil, &util.Error{Msg:"数据不存在"}
	}
	//var Statistics *models.BlogStatistics
	//Statistics, err = models.GetBlogStatisticsById(id)
	//data["Comment"] = Statistics.Comment
	//data["Read"] = Statistics.Read
	return data, err
}