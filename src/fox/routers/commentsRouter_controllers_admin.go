package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "List",
			Router: `/blog/cat`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/blog/cat/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/blog/cat/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/blog/cat`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/blog/cat/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/blog/cat/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "CheckTitle",
			Router: `/blog/check_title`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/blog`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/blog/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/blog/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/blog`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/blog/detail/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/blog/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/blog/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogController"],
		beego.ControllerComments{
			Method: "Image",
			Router: `/blog/upload/image`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:BlogTag"] = append(beego.GlobalControllerRouter["fox/controllers/admin:BlogTag"],
		beego.ControllerComments{
			Method: "List",
			Router: `/blog/tag`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:Index"] = append(beego.GlobalControllerRouter["fox/controllers/admin:Index"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:Index"] = append(beego.GlobalControllerRouter["fox/controllers/admin:Index"],
		beego.ControllerComments{
			Method: "V2",
			Router: `/index/v2`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:Select"] = append(beego.GlobalControllerRouter["fox/controllers/admin:Select"],
		beego.ControllerComments{
			Method: "Type",
			Router: `/select/type`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:Select"] = append(beego.GlobalControllerRouter["fox/controllers/admin:Select"],
		beego.ControllerComments{
			Method: "Type",
			Router: `/select/type/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/type`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "ListChild",
			Router: `/type/list_child/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/type/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/type/add/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/type`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/type/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/type/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers/admin:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers/admin:TypeController"],
		beego.ControllerComments{
			Method: "CheckName",
			Router: `/type/check_name`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
