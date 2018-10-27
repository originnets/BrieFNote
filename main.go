package main

import (
	"BrieFNote/models"
	_ "BriefNote/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"html/template"
)

//markdown支持
func MarkDown(raw string) (output template.HTML) {
	input := []byte(raw)
	bOutput := blackfriday.MarkdownBasic(input)
	output = template.HTML(string(bOutput))
	return
}

//开启session功能
func InitSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "briefnote"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func main() {
	//orm.Debug = true
	InitSession()
	//注册自定义函数
	beego.AddFuncMap("markdown",MarkDown)
	beego.Run()
}

