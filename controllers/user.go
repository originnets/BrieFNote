package controllers

import (
	"BrieFNote/models"
	"github.com/astaxie/beego/orm"
	"github.com/dchest/captcha"
)


type UserController struct {
	BaseController
}

// @router /user/login [get]
func (c *UserController) GetLogin(){
	c.LoginStatus()
	//调用验证码ID
	c.NewCaptchaId()
	c.TplName = "user/login.html"
}

// @router /user/login [post]
func (c *UserController) PostLogin(){
	email := c.GetMUstString("email", "邮箱不能为空")
	if email == ""{
		return
	}
	pass := c.GetMUstString("pass", "密码不能为空")
	if pass == ""{
		return
	}
	captchaId := c.GetMUstString("vercodeId", "验证码ID不能为空")
	if captchaId == ""{
		return
	}
	captchaValue := c.GetMUstString("vercode", "验证码不能为空")
	if captchaValue == ""{
		return
	}
	//验证码验证
	if !captcha.VerifyString(captchaId, captchaValue) {
		c.Revert(3001, "验证码错误")
		return
	}

	o := orm.NewOrm()
	password := c.Md5(pass)
	user := models.User{Email:email, Password:password}
	err := o.Read(&user, "Email", "Password")
	if err != nil {
		c.Revert(2001, "邮箱或密码错误")
		return
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.RevertOk(0,"用户登录成功","/")
}

// @router /user/reg [get]
func (c *UserController) GetReg(){
	c.LoginStatus()

	//调用验证码ID
	c.NewCaptchaId()
	c.TplName = "user/reg.html"
}

// @router /user/reg [post]
func (c *UserController) PostRge(){

	email := c.GetMUstString("email", "邮箱不能为空")
	if email == ""{
		return
	}
	e_code := c.GetMUstString("e_code", "邮箱不能为空")
	if e_code == ""{
		return
	}
	username := c.GetMUstString("username", "昵称不能为空")
	if username == ""{
		return
	}
	pass := c.GetMUstString("pass", "密码不能为空")
	if pass == ""{
		return
	}
	repass := c.GetMUstString("repass", "确认密码不能为空")
	if repass == ""{
		return
	}
	captchaId := c.GetMUstString("vercodeId", "验证码ID不能为空")
	if captchaId == ""{
		return
	}
	captchaValue := c.GetMUstString("vercode", "验证码不能为空")
	if captchaValue == ""{
		return
	}
	if pass != repass {
		c.Revert(1002,"两次密码输入不一致")
		return
	}
	//验证码验证
	if !captcha.VerifyString(captchaId, captchaValue) {
		c.Revert(3001, "验证码错误")
		return
	}
	//从session中获取邮箱验证码
	se_code := c.GetSession("reg_code")
	if se_code == nil {
		c.Revert(1002,"请获取邮箱验证码")
		return
	}
	if se_code != e_code {
		c.Revert(1002,"邮箱验证码错误")
		return
	}

	o := orm.NewOrm()
	user := models.User{Email:email}
	err := o.Read(&user, "Email")
	if err == nil {
		c.Revert(1002,"该邮箱已经存在")
		return
	}
	user.Name = username
	err = o.Read(&user, "Name")
	if err == nil {
		c.Revert(1002,"该昵称已经存在")
		return
	}
	user.Password = c.Md5(pass)
	user.Role = 1
	user.Avatar = "/static/images/avatar/3.jpg"
	_, err = o.Insert(&user)
	if err != nil {
		c.Revert(1002,"注册失败")
		return
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.RevertOk(0,"用户注册成功","/")
}

// @router /user/forget [get]
func (c *UserController) GetForGet(){
	c.LoginStatus()
	//调用验证码ID
	c.NewCaptchaId()
	c.TplName = "user/forget.html"
}

// @router /user/forget [post]
func (c *UserController) PostForGet(){
	email := c.GetMUstString("email", "邮箱不能为空")
	if email == "" {
		return
	}
	e_code := c.GetMUstString("e_code", "邮箱不能为空")
	if e_code == "" {
		return
	}
	pass := c.GetMUstString("pass", "密码不能为空")
	if pass == "" {
		return
	}
	repass := c.GetMUstString("repass", "确认密码不能为空")
	if repass == "" {
		return
	}
	captchaId := c.GetMUstString("vercodeId", "验证码ID不能为空")
	if captchaId == "" {
		return
	}
	captchaValue := c.GetMUstString("vercode", "验证码不能为空")
	if captchaValue == "" {
		return
	}

	//从session中获取邮箱验证码
	se_code := c.GetSession("forget_code")
	if se_code == nil {
		c.Revert(1002,"请获取邮箱验证码")
		return
	}
	if se_code != e_code {
		c.Revert(1002,"邮箱验证码错误")
		return
	}

	if pass != repass {
		c.Revert(1002,"两次密码输入不一致")
		return
	}
	//验证码验证
	if !captcha.VerifyString(captchaId, captchaValue) {
		c.Revert(3001, "验证码错误")
		return
	}
	o := orm.NewOrm()
	user := models.User{Email:email}
	err := o.Read(&user,"Email")
	if err != nil {
		c.Revert(1002, "该邮箱用户不存在")
		return
	}
	user.Password = c.Md5(pass)
	_, err = o.Update(&user)
	if err != nil {
		c.Revert(4001, "密码修改失败")
		return
	}
	c.RevertOk(0, "密码修改成功", "/user/login")
}

// @router /user/logout [get]
func (c *UserController) GetLogout(){
	//c.MustLogin()
	c.DelSession(SESSION_USER_KEY)
	c.Redirect("/user/login", 302)
}