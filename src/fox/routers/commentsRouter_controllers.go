package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["fox/controllers:AdminController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminMenuController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminMenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminRolePrivController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminStatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminStatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminStatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AdminStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:AdminStatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaExtController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaExtController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaExtController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaExtController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaExtController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaExtController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaExtController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaExtController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AreaExtController"] = append(beego.GlobalControllerRouter["fox/controllers:AreaExtController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AttachmentController"] = append(beego.GlobalControllerRouter["fox/controllers:AttachmentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AttachmentController"] = append(beego.GlobalControllerRouter["fox/controllers:AttachmentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AttachmentController"] = append(beego.GlobalControllerRouter["fox/controllers:AttachmentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AttachmentController"] = append(beego.GlobalControllerRouter["fox/controllers:AttachmentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:AttachmentController"] = append(beego.GlobalControllerRouter["fox/controllers:AttachmentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:ConnectController"] = append(beego.GlobalControllerRouter["fox/controllers:ConnectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:ConnectController"] = append(beego.GlobalControllerRouter["fox/controllers:ConnectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:ConnectController"] = append(beego.GlobalControllerRouter["fox/controllers:ConnectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:ConnectController"] = append(beego.GlobalControllerRouter["fox/controllers:ConnectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:ConnectController"] = append(beego.GlobalControllerRouter["fox/controllers:ConnectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:LogController"] = append(beego.GlobalControllerRouter["fox/controllers:LogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:LogController"] = append(beego.GlobalControllerRouter["fox/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:LogController"] = append(beego.GlobalControllerRouter["fox/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:LogController"] = append(beego.GlobalControllerRouter["fox/controllers:LogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:LogController"] = append(beego.GlobalControllerRouter["fox/controllers:LogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberGroupExtController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberProfileController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberProfileController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberProfileController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberProfileController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberProfileController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberProfileController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberProfileController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberProfileController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberProfileController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberProfileController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberStatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberStatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberStatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:MemberStatusController"] = append(beego.GlobalControllerRouter["fox/controllers:MemberStatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:SessionController"] = append(beego.GlobalControllerRouter["fox/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:SessionController"] = append(beego.GlobalControllerRouter["fox/controllers:SessionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:SessionController"] = append(beego.GlobalControllerRouter["fox/controllers:SessionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:SessionController"] = append(beego.GlobalControllerRouter["fox/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:SessionController"] = append(beego.GlobalControllerRouter["fox/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TemplateController"] = append(beego.GlobalControllerRouter["fox/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TemplateController"] = append(beego.GlobalControllerRouter["fox/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TemplateController"] = append(beego.GlobalControllerRouter["fox/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TemplateController"] = append(beego.GlobalControllerRouter["fox/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TemplateController"] = append(beego.GlobalControllerRouter["fox/controllers:TemplateController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers:TypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers:TypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers:TypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers:TypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["fox/controllers:TypeController"] = append(beego.GlobalControllerRouter["fox/controllers:TypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
