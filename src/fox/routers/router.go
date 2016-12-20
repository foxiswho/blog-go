// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"fox/controllers/admin"
)

func init() {
	beego.Router("/admin/login", &admin.LoginController{})
	//ns := beego.NewNamespace("/v1",
	//
	//	beego.NSNamespace("/admin",
	//		beego.NSInclude(
	//			&controllers.AdminController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/admin_menu",
	//		beego.NSInclude(
	//			&controllers.AdminMenuController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/admin_role_priv",
	//		beego.NSInclude(
	//			&controllers.AdminRolePrivController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/admin_status",
	//		beego.NSInclude(
	//			&controllers.AdminStatusController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/area",
	//		beego.NSInclude(
	//			&controllers.AreaController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/area_ext",
	//		beego.NSInclude(
	//			&controllers.AreaExtController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/attachment",
	//		beego.NSInclude(
	//			&controllers.AttachmentController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/connect",
	//		beego.NSInclude(
	//			&controllers.ConnectController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/log",
	//		beego.NSInclude(
	//			&controllers.LogController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/member",
	//		beego.NSInclude(
	//			&controllers.MemberController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/member_group_ext",
	//		beego.NSInclude(
	//			&controllers.MemberGroupExtController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/member_profile",
	//		beego.NSInclude(
	//			&controllers.MemberProfileController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/member_status",
	//		beego.NSInclude(
	//			&controllers.MemberStatusController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/session",
	//		beego.NSInclude(
	//			&controllers.SessionController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/template",
	//		beego.NSInclude(
	//			&controllers.TemplateController{},
	//		),
	//	),
	//
	//	beego.NSNamespace("/type",
	//		beego.NSInclude(
	//			&controllers.TypeController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
}
