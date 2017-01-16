package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "List",
			Router: `/admin_role`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/admin_role/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/admin_role`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/admin_role/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/admin_role/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "CheckName",
			Router: `/admin_role/check_name`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminRole"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/admin_role/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "CheckTitle",
			Router: `/admin/check_title`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "List",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/admin/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/admin/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/admin`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/admin/detail/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/admin/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AdminUser"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/admin/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AppCsdn"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AppCsdn"],
		beego.ControllerComments{
			Method: "List",
			Router: `/auth_csdn`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:AppCsdn"] = append(beego.GlobalControllerRouter["blog/controllers/admin:AppCsdn"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/auth_token`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Area"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Area"],
		beego.ControllerComments{
			Method: "List",
			Router: `/area`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Attachment"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Attachment"],
		beego.ControllerComments{
			Method: "List",
			Router: `/attachment`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "CheckTitle",
			Router: `/blog/check_title`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "List",
			Router: `/blog`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/blog/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/blog/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/blog`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/blog/detail/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/blog/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Blog"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Blog"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/blog/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "List",
			Router: `/blog/cat`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/blog/cat/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/blog/cat/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/blog/cat`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/blog/cat/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogCat"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/blog/cat/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"],
		beego.ControllerComments{
			Method: "List",
			Router: `/auth_csdn`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"],
		beego.ControllerComments{
			Method: "Go",
			Router: `/blog_sync/go`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/blog_sync/go`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogSync"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/blog_sync/go`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:BlogTag"] = append(beego.GlobalControllerRouter["blog/controllers/admin:BlogTag"],
		beego.ControllerComments{
			Method: "List",
			Router: `/blog/tag`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Index"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Index"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/index`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Index"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Index"],
		beego.ControllerComments{
			Method: "V2",
			Router: `/index/v2`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Index"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Index"],
		beego.ControllerComments{
			Method: "Default",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "CheckTitle",
			Router: `/member/check_title`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "List",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/member/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/member/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/member`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/member/detail/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/member/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Member"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Member"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/member/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "List",
			Router: `/member_group`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/member_group/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/member_group`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/member_group/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/member_group/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "CheckName",
			Router: `/member_group/check_name`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"] = append(beego.GlobalControllerRouter["blog/controllers/admin:MemberGroup"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/member_group/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Oauth"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Oauth"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/oauth`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Oauth"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Oauth"],
		beego.ControllerComments{
			Method: "Csdn",
			Router: `/oauth_csdn`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Select"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Select"],
		beego.ControllerComments{
			Method: "Type",
			Router: `/select/type`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Select"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Select"],
		beego.ControllerComments{
			Method: "Type",
			Router: `/select/type/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Site"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Site"],
		beego.ControllerComments{
			Method: "List",
			Router: `/site`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Site"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Site"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/site`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "List",
			Router: `/type`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "ListChild",
			Router: `/type/list_child/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/type/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/type/add/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/type`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/type/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/type/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Type"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Type"],
		beego.ControllerComments{
			Method: "CheckName",
			Router: `/type/check_name`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Upload"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Upload"],
		beego.ControllerComments{
			Method: "Image",
			Router: `/upload/image`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Upload"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Upload"],
		beego.ControllerComments{
			Method: "File",
			Router: `/upload/file`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers/admin:Upload"] = append(beego.GlobalControllerRouter["blog/controllers/admin:Upload"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/upload/image`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
