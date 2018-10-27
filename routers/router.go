package routers

import (
	"BrieFNote/controllers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	beego.Handler("/authentication/*.png", captcha.Server(80, 34)) //图片验证码
	beego.Router( "/", &controllers.UserController{} , "get:GetIndex")
	beego.Router("/md/upload", &controllers.BaseController{}, "post:MdUpload")	//markdown内容图片上传
	beego.Router("/img/upload", &controllers.BaseController{}, "post:ImgUpload")	//图片上传
	beego.Include( &controllers.UserController{})	//用户
	beego.Include( &controllers.SendCodeController{}) //邮箱验证码
	beego.Include( &controllers.ArticleController{})	//文章
	beego.Include( &controllers.CommentController{})	//评论
	beego.Include( &controllers.AdminController{})		//后台
}
