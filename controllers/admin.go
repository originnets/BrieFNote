package controllers

import (
	"BrieFNote/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type AdminController struct {
	BaseController
}

var Rtype struct{
	Id int
	Typename string
	Status	bool
}

// @router /admin [get]
func (c *ArticleController) GetAdminIndex() {
	c.LoginStatus()
	c.TplName = "admin/index.html"
}

// @router /admin/console [get]
func (c *ArticleController) GetConsole(){
	c.TplName = "admin/home/console.html"
}

// @router /admin/add [get]
func (c *ArticleController) GetAdd(){
	c.LoginStatus()
	o := orm.NewOrm()
	var atypes []*models.Type
	qs := o.QueryTable("type")
	_, err := qs.All(&atypes)
	if err != nil {
		beego.Info("查询失败", err)
	}
	c.Data["Atypes"] = atypes
	c.TplName = "admin/home/add.html"
}

// @router /admin/add [post]
func (c *AdminController) PostAdd(){
	c.LoginStatus()
	title := c.GetMUstString("title", "标题不能为空")
	if title == "" {
		return
	}
	num, _ := c.GetInt("type_num")
	types := make([]int, 0) // 定义一个切片用于接收type
	for i := 1;i <= num; i++ {
		str := "type_" + strconv.Itoa(i)
		atype := c.GetString(str)
		if atype == "on" {
			types = append(types, i)
		}
	}
	img := c.GetMUstString("img", "请先上传文章图片")
	if img == "" {
		return
	}
	content := c.GetMUstString("content", "内容不能为空")
	status := c.GetString("status")
	settop := c.GetString("settop")
	thickening := 	c.GetString("thickening")
	if len(types) == 0 {
		c.Revert(2002,"请至少选择一个标签")
		return
	}
	if content == "" {
		c.Revert(2003,"文章内容不能为空")
		return
	}
	o := orm.NewOrm()
	user := models.User{Name: c.User.Name}
	err := o.Read(&user, "Name")
	if err != nil {
		c.Revert(4003,"请先登录")
		return
	}
	article := models.Article{}
	article.Title = title
	article.Img = img
	article.User = &user
	article.Content = content
	if status == "on" {
		article.Status = 1
	}else {
		article.Status = 0
	}

	if settop == "on" {
		article.Settop = true
	}else {
		article.Settop = false
	}

	if thickening == "on" {
		article.Thickening = true
	}else {
		article.Thickening = false
	}
	_, err = o.Insert(&article)
	if err != nil {
		c.Revert(5001,"文章添加失败")
		return
	}
	//添加标签
	m2m := o.QueryM2M(&article, "Types")
	for _,value := range types {
		atype := models.Type{Id:value}
		err := o.Read(&atype)
		if err != nil {
			c.Revert(5001,"标签读取失败")
			continue
		}
		_, err = m2m.Add(atype)
		if err != nil {
			c.Revert(5002,"标签添加失败")
			continue
		}
	}

	//添加阅读数
	read := models.Read{}
	read.ReadNum = 0
	read.Article = &article
	_, err = o.Insert(&read)
	if err != nil {
		c.Revert(5001,"阅读库添加失败")
		return
	}
	c.Revert(0, "添加成功")
}

// @router /admin/typelist [get]
func (c *AdminController) GetTypeList() {
	c.LoginStatus()
	res := make(map[string]interface{})
	list :=make([]interface{},0)
	defer c.ServeJSON()
	o:= orm.NewOrm()
	var atypes []models.Type
	qs := o.QueryTable("type")
	num, err := qs.All(&atypes)
	if err != nil {
		res["code"] = 1
		res["msg"] = "查询失败"
		c.Data["json"] = res
		c.TplName = "admin/home/type.html"
		return
	}
	res["code"] = 0
	res["msg"] = "查询成功"
	res["count"] = num
	for _, atype := range atypes {
		list = append(list, map[string]interface{}{ "id": atype.Id, "typename": atype.Typename})
	}
	res["data"] = list
	c.Data["json"] = res
}

// @router /admin/type [get]
func (c *AdminController) GetType(){
	c.LoginStatus()
	c.TplName = "admin/home/type.html"
}

// @router /admin/addtype [post]
func (c *AdminController) PostAddType(){
	c.LoginStatus()
	add_type := c.GetMUstString("add_type", "添加标签为空")
	if add_type == "" {
		return
	}
	o := orm.NewOrm()
	atype := models.Type{Typename:add_type}
	err := o.Read(&atype)
	if err == nil {
		c.Revert(1002, "该标签已经存在")
		return
	}
	_,err = o.Insert(&atype)
	if err != nil {
		c.Revert(1003, "标签添加失败")
		return
	}
	c.Revert(0,"添加成功")
}

// @router /admin/list [get]
func (c *AdminController) GetList() {
	c.LoginStatus()
	//获取当前URL
	url := c.Ctx.Request.RequestURI
	c.Data["Url"] = url
	c.TplName = "admin/home/list.html"
}

// @router /admin/listdata [get]
func (c *AdminController) GetListData() {
	c.LoginStatus()
	res := make(map[string]interface{})
	list :=make([]interface{},0)
	defer c.ServeJSON()
	o:= orm.NewOrm()
	var articles []*models.Article
	qs := o.QueryTable("article")
	var (
		num int64
		err error
	)
	if c.User.Role == 0 {
		num, err = qs.OrderBy("-Updatetime").RelatedSel().All(&articles)
	} else {
		num, err = qs.Filter("User__Name", c.User.Name).OrderBy("-Updatetime").RelatedSel().All(&articles)
	}
	if err != nil {
		res["code"] = 1
		res["msg"] = "查询失败"
		c.Data["json"] = res
		return
	}
	res["count"] = num
	for _, article := range articles {
		_, err = o.LoadRelated(article, "Types")
	}
	if err != nil {
		c.GetError()
		return
	}
	res["code"] = 0
	res["msg"] = "查询成功"
	for _, article := range articles {
		Createtime := article.Createtime.Format("2006-01-02")
		Updatetime := article.Updatetime.Format("2006-01-02")
		var Status string
		var Settop string
		var Thickening string
		var atype string
		if article.Status == 1 {
			Status = "已发布"
		}else {
			Status = "未发布"
		}
		if article.Settop == true {
			Settop = "置顶"
		}else {
			Settop = "未置顶"
		}
		if article.Thickening == true {
			Thickening = "精帖"
		}else {
			Thickening = "非精帖"
		}
		l := make([]string,0)
		for i:=0;i<len(article.Types);i++{
			if article.Types[i].Typename != ""{
				l = append(l,article.Types[i].Typename)
			}
			atype = strings.Replace(strings.Trim(fmt.Sprint(l), "[]"), " ", ",", -1)
		}
		list = append(list, map[string]interface{}{ "id": article.Id, "title":article.Title, "createtime":Createtime, "updatetime":Updatetime, "status":Status, "settop":Settop, "thickening":Thickening ,"types":atype})
	}
	res["data"] = list
	c.Data["json"] = res
}

// @router /admin/modifytype [post]
func (c *AdminController) PostModifyType() {
	c.LoginStatus()
	id := c.GetMUstString("id", "标签ID为空")
	if id == "" {
		return
	}
	typename := c.GetMUstString("typename", "标签类型为空")
	if typename == "" {
		return
	}
	id1, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	atype := models.Type{Id:id1}
	err := o.Read(&atype)
	if err != nil {
		c.Revert(1002, "该标签不存在")
		return
	}
	atype.Typename = typename
	err = o.Read(&atype, "Typename")
	if err == nil {
		c.Revert(1002, "该标签已经存在")
		return
	}
	_, err = o.Update(&atype)
	if err != nil {
		c.Revert(1002, "标签修改失败")
		return
	}
	c.Revert(0,"修改成功")
}

// @router /admin/deltype [post]
func (c *AdminController) PostDelType() {
	c.LoginStatus()
	id := c.GetMUstString("id", "标签ID为空")
	if id == "" {
		return
	}
	id1, _ := strconv.Atoi(id)
	o:= orm.NewOrm()
	atype := models.Type{Id:id1}
	err := o.Read(&atype)
	if err != nil {
		c.Revert(1002,"该标签不存在")
		return
	}
	num, _ := o.LoadRelated(&atype,"Articles")
	if num != 0 {
		c.Revert(1002,"该标签已经绑定文章,请先清除该标签下的文章再删除")
		return
	}
	_, err = o.Delete(&atype)
	if err != nil{
		c.Revert(1002,"标签删除失败")
		return
	}
	c.Revert(0,"删除成功")
}

// @router /admin/edit/?:id [get]
func (c *AdminController) GetEdit() {
	//接收url
	//rel := c.GetString("rel")
	//beego.Info(rel)
	c.LoginStatus()
	id := c.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	var article models.Article
	var err error
	if c.User.Role == 0 {
		_ , err = qs.Filter("Id", id).RelatedSel().All(&article)
	}else {
		_ , err = qs.Filter("User",c.User).Filter("Id", id).RelatedSel().All(&article)
	}
	if err != nil {
		c.GetError()
		return
	}
	_, err = o.LoadRelated(&article, "Types")

	if err != nil {
		c.GetError()
		return
	}
	var atypes []*models.Type
	qs1 := o.QueryTable("type")
	_, err = qs1.All(&atypes)
	if err != nil {
		c.GetError()
		return
	}
	Types := make([]map[string]interface{}, 0)
	for _, type1 := range atypes {
		Rtype.Id = type1.Id
		Rtype.Typename = type1.Typename
		Rtype.Status = false
		Types = append(Types,map[string]interface{}{"Id":Rtype.Id, "Typename":Rtype.Typename, "Status":Rtype.Status})
	}
	for i:=0;i<len(article.Types);i++{
		for _, type2 := range Types {
			if article.Types[i].Id == type2["Id"] {
				type2["Status"] = true
			}
		}
	}
	c.Data["Article"] = article
	c.Data["Types"] = Types
	c.TplName = "admin/home/edit.html"
}

// @router /admin/edit/?:id [post]
func (c *AdminController) PostEdit(){
	c.LoginStatus()
	id := c.Ctx.Input.Param(":id")
	title := c.GetMUstString("title", "标题不能为空")
	if title == "" {
		return
	}
	num, _ := c.GetInt("type_num")
	types := make([]int, 0) // 定义一个切片用于接收type
	for i := 1;i <= num; i++ {
		str := "type_" + strconv.Itoa(i)
		atype := c.GetString(str)
		if atype == "on" {
			types = append(types, i)
		}
	}
	img := c.GetMUstString("img", "请先上传文章图片")
	if img == "" {
		return
	}
	content := c.GetMUstString("content", "内容不能为空")
	status := c.GetString("status")
	settop := c.GetString("settop")
	thickening := 	c.GetString("thickening")
	if len(types) == 0 {
		c.Revert(2002,"请至少选择一个标签")
		return
	}
	if content == "" {
		c.Revert(2003,"文章内容不能为空")
		return
	}
	o := orm.NewOrm()
	user := models.User{Name: c.User.Name}
	err := o.Read(&user, "Name")
	if err != nil {
		c.Revert(4003,"请先登录")
		return
	}
	article := models.Article{}
	qs := o.QueryTable(&article)
	if c.User.Role == 0 {
		_ , err = qs.Filter("Id", id).RelatedSel().All(&article)
	}else {
		_ , err = qs.Filter("User",user).Filter("Id", id).RelatedSel().All(&article)
	}
	if err != nil {
		c.Revert(4003,"无权修改本文章")
		return
	}
	article.Title = title
	article.Img = img
	article.User = &user
	article.Content = content
	if status == "on" {
		article.Status = 1
	}else {
		article.Status = 0
	}

	if settop == "on" {
		article.Settop = true
	}else {
		article.Settop = false
	}

	if thickening == "on" {
		article.Thickening = true
	}else {
		article.Thickening = false
	}
	_, err = o.Update(&article)
	if err != nil {
		c.Revert(5001,"文章更新失败")
		return
	}
	m2m := o.QueryM2M(&article, "Types")
	_, err = m2m.Clear()
	if err != nil {
		c.Revert(5001,"清除标签失败")
		return
	}
	for _,value := range types {
		atype := models.Type{Id:value}
		err := o.Read(&atype)
		if err != nil {
			c.Revert(5001,"标签读取失败")
			continue
		}
		_, err = m2m.Add(atype)
		if err != nil {
			c.Revert(5002,"标签修改失败")
			continue
		}
	}
	c.Revert(0, "修改成功")
}

// @router /admin/delarticle/?:id [post]
func (c *AdminController) PostDelArticle() {
	c.LoginStatus()
	id := c.GetString("id")
	if id == ""{
		c.Revert(5002,"id不能为空")
		return
	}
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	article :=models.Article{}
	var err error
	if c.User.Role == 0 {
		_ , err = qs.Filter("Id", id).RelatedSel().All(&article)
	}else {
		_ , err = qs.Filter("User",c.User).Filter("Id", id).RelatedSel().All(&article)
	}
	if err != nil {
		c.Revert(1003,"该文章不存在")
		return
	}
	_, err = o.LoadRelated(&article, "Types")
	if err != nil {
		c.Revert(1003,"该文章不存在")
		return
	}

	//移除标签库
	m2m := o.QueryM2M(&article, "Types")
	_, err = m2m.Remove(article.Types)
	if err != nil {
		c.Revert(1003,"移除标签失败")
		return
	}
	_, err = o.Delete(&article)
	if err != nil {
		c.Revert(1003,"删除文章失败")
		return
	}

	//移除阅读库
	read := models.Read{}
	read.Article = &article
	_, err = o.Delete(&read)
	if err != nil {
		c.Revert(1003,"移除阅读数失败")
		return
	}
	c.Revert(0,"删除成功")
}