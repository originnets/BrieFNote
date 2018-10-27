package controllers

type SendCodeController struct {
	BaseController
}

// @router /sendcode/forget [post]
func (c *SendCodeController) PostForGetSendCode() {
	mail := c.GetMUstString("email", "邮箱为空")
	if mail == ""{
		return
	}
	subject := "重置密码"
	mtype := "html"
	code := c.GetRandomString(8)
	body := code

	out := make(chan error)
	go func() {
			err := c.SendMail(true, mail, subject, body, mtype)
			out <- err
		}()
	err, _ := <- out
	if err != nil {
		c.Revert(1002,"邮件发送失败")
		return
	}
	//设置session用于验证
	c.SetSession("forget_code",code)
	c.Revert(0,"邮件发送成功")
}

// @router /sendcode/reg [post]
func (c *SendCodeController) PostRegSendCode() {
	mail := c.GetMUstString("email", "邮箱为空")
	if mail == ""{
		return
	}
	subject := "邮箱注册"
	mtype := "html"
	code := c.GetRandomString(8)
	body := code

	out := make(chan error)
	go func() {
		err := c.SendMail( false, mail, subject, body, mtype)
		out <- err
	}()
	err, _ := <- out
	if err != nil {
		c.Revert(1002,"邮件发送失败")
		return
	}
	//设置session用于验证
	c.SetSession("reg_code",code)
	c.Revert(0,"邮件发送成功")
}