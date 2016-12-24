package routers

import (
	"github.com/astaxie/beego"
)

func init() {

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

}
