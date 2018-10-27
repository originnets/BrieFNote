package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostAdd",
			Router: `/admin/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostAddType",
			Router: `/admin/addtype`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostDelArticle",
			Router: `/admin/delarticle/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostDelType",
			Router: `/admin/deltype`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetEdit",
			Router: `/admin/edit/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostEdit",
			Router: `/admin/edit/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/admin/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetListData",
			Router: `/admin/listdata`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "PostModifyType",
			Router: `/admin/modifytype`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetType",
			Router: `/admin/type`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetTypeList",
			Router: `/admin/typelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAdminIndex",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAdd",
			Router: `/admin/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetConsole",
			Router: `/admin/console`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetDetail",
			Router: `/article/detail/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/article/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:CommentController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:CommentController"],
		beego.ControllerComments{
			Method: "PostAddComment",
			Router: `/comment/add/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:SendCodeController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:SendCodeController"],
		beego.ControllerComments{
			Method: "PostForGetSendCode",
			Router: `/sendcode/forget`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:SendCodeController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:SendCodeController"],
		beego.ControllerComments{
			Method: "PostRegSendCode",
			Router: `/sendcode/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetForGet",
			Router: `/user/forget`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostForGet",
			Router: `/user/forget`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetLogin",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostLogin",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetLogout",
			Router: `/user/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetReg",
			Router: `/user/reg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["BrieFNote/controllers:UserController"] = append(beego.GlobalControllerRouter["BrieFNote/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostRge",
			Router: `/user/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
