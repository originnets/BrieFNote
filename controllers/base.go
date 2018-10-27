package controllers

import (
	"BrieFNote/models"
	"BrieFNote/utils"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dchest/captcha"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	User models.User
	IsLogin	bool
}

//定义session
const SESSION_USER_KEY  = "SESSION_USER_KEY"

//获取当前用户是否处于登录状态
func (c *BaseController) LoginStatus() {
	c.IsLogin = false
	user, ok := c.GetSession(SESSION_USER_KEY).(models.User)
	if ok {
		c.User = user
		c.IsLogin = true
		c.Data["User"] = c.User
	}
	c.Data["LoginStatus"] = c.IsLogin
}

func (c *BaseController) LoginStatus1() {
	c.IsLogin = false
	user, ok := c.GetSession(SESSION_USER_KEY).(models.User)
	if ok {
		c.User = user
		c.IsLogin = true
		c.Data["User"] = c.User
	}
	c.Data["LoginStatus"] = c.IsLogin
	beego.Info(c.IsLogin)
}

func (c *BaseController) MustLogin() {
	if !c.IsLogin {
		c.Revert(100,"用户没有登录")
	}
}


//json返回
func (c *BaseController) Read(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

//失败返回
func (c *BaseController) Revert(code int, msg string) {
	res := map[string]interface{}{"code":code, "msg":msg }
	c.Data["json"] = res
	c.ServeJSON()
}

//成功返回
func (c *BaseController) RevertOk(code int, msg ,action string) {
	res := map[string]interface{}{"code":code, "msg":msg , "action":action}
	c.Data["json"] = res
	c.ServeJSON()
}

//markdown 上传返回
func (c *BaseController) RevertMd(success int, message ,url string) {
	res := map[string]interface{}{"success":success, "message":message , "url":url}
	c.Data["json"] = res
	c.ServeJSON()
}

//md5密码加密
func (c *BaseController) Md5(str string) (MD5PAW string){
	DATA := []byte(str)
	MD5PAW = fmt.Sprintf("%x",md5.Sum(DATA))
	return
}

//验证form提交的数据
func (c *BaseController)GetMUstString(key, msg string) string {
	resp := make(map[string]interface{})
	value := c.GetString(key)
	if len(value) == 0 {
		resp["code"] = 1001
		resp["msg"] = msg
		c.Read(resp)
		return value
	}
	value1 := strings.Replace(value, " ", "", -1)
	if  len(value1) == 0 {
		resp["code"] = 1001
		resp["msg"] = msg
		c.Read(resp)
		return value1
	}
	return value
}

//建立验证码ID
func (c *BaseController) NewCaptchaId(){
	Code := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}
	c.Data["CaptchaId"] = Code.CaptchaId
}

//获取随机数
func (c *BaseController) GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//发送邮件
//err := SendMail("开启查看邮箱用户是本地用户", "目标邮箱", "邮件标题", "邮件内容", "html")
func (c *BaseController) SendMail(eamil_exist bool, mail, subject, body, mtype string) error {
	//配置文件获取
	EMAIL_HOST_USER := beego.AppConfig.String("EMAIL_HOST_USER")
	EMAIL_HOST_PASSWORD := beego.AppConfig.String("EMAIL_HOST_PASSWORD")
	EMAIL_HOST := beego.AppConfig.String("EMAIL_HOST")
	EMAIL_PORT := beego.AppConfig.String("EMAIL_PORT")
	Display_Name := beego.AppConfig.String("Display_Name")

	addr := EMAIL_HOST + ":" + EMAIL_PORT
	auth := smtp.PlainAuth("", EMAIL_HOST_USER, EMAIL_HOST_PASSWORD, EMAIL_HOST)
	var c_type string
	if mtype == "html" {
		c_type = "Content-Type: text/" + mtype + "; charset=UTF-8"
	} else {
		c_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + mail + "\r\nFrom: " + Display_Name + "<" + EMAIL_HOST_USER + ">\r\nSubject: " + "重置邮箱验证码" + "\r\n" + c_type + "\r\n\r\n" + body)
	sto := strings.Split(mail, ";")
	if eamil_exist {
		o := orm.NewOrm()
		user := models.User{Email:mail}
		err := o.Read(&user, "Email")
		if err != nil {
			c.Revert(1002,"该用户不存在")
			return err
		}
		err = smtp.SendMail(addr, auth, EMAIL_HOST_USER, sto, msg)
		return err
	} else {
		o := orm.NewOrm()
		user := models.User{Email:mail}
		err := o.Read(&user, "Email")
		if err == nil {
			c.Revert(1002,"该用户已经存在")
			return err
		}
		//发送邮件
		err = smtp.SendMail(addr, auth, EMAIL_HOST_USER, sto, msg)
		return err
	}

}

//markdown图片上传
func (c *BaseController) MdUpload() {
	//f, h , err := c.GetFile("file")
	f, h , err := c.GetFile("editormd-image-file")
	defer f.Close()
	if err == nil {
		exStrArr := strings.Split(h.Filename, ".")
		exStr := strings.ToLower(exStrArr[len(exStrArr)-1])
		if exStr != "jpg" && exStr!="png" && exStr != "gif" && exStr != "jpeg" && exStr != "bmp的图片"{
			c.RevertMd(0, "上传格式只能是jpg|jpeg|gif|png|bmp的图片", "")
			return
		}
		img := "static/upload/" + (utils.TimeUUID()).String() + "." + exStr
		err = c.SaveToFile("editormd-image-file", img)
		if err != nil {
			c.RevertMd(0, "保存图片失败","")
			return
		}
		SERVER_NAME := beego.AppConfig.String("SERVER_NAME")
		url := SERVER_NAME + "/" +img
		c.RevertMd(1, "上传成功",url)
	}
}

//本地图片上传
func (c *BaseController) ImgUpload() {
	f, h , err := c.GetFile("file")
	defer f.Close()
	if err == nil {
		exStrArr := strings.Split(h.Filename, ".")
		exStr := strings.ToLower(exStrArr[len(exStrArr)-1])
		if exStr != "jpg" && exStr!="png" && exStr != "gif" && exStr != "jpeg" && exStr != "bmp的图片" {
			c.RevertMd(0, "上传格式只能是jpg|jpeg|gif|png|bmp的图片", "")
			return
		}
		img := "static/upload/" + (utils.TimeUUID()).String() + "." + exStr
		err = c.SaveToFile("file", img)
		if err != nil {
			c.Revert(1003, "保存图片失败")
			return
		}
		url := "/" + img
		c.RevertMd(0, "上传成功",url)
	}
}

//错误页面调用
func (c *BaseController) GetError(){
	c.TplName = "other/error.html"
}

//阅读统计量
func (c *BaseController) ReadingStatistics(id string)(key string){
	//获取用户ip
	addr := c.Ctx.Request.RemoteAddr
	key1 := addr + id
	value := c.GetSession(key1)
	if value == nil {
		// 总阅读数量 +1
		o :=orm.NewOrm()
		read := models.Read{}
		qs := o.QueryTable("read")
		err := qs.Filter("Article__Id", id).One(&read)
		if err != nil {
			c.Revert(6001, "查询失败")
		}
		read.ReadNum += 1
		_, err =o.Update(&read)
		if err != nil {
			c.Revert(5001, "阅读更新失败")
		}
	}
	return key1
}
