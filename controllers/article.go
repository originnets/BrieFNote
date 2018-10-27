package controllers

import (
	"BrieFNote/models"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	BaseController
}


func (c *BaseController) GetIndex(){
	c.LoginStatus()
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	var articles []*models.Article	//定义接收
	num, err:= qs.Filter("Settop", true).Filter("Status",1).OrderBy("-Updatetime").RelatedSel().Limit(5).All(&articles) //.RelatedSel()查出包括User中的数据

	//当articles为多个值是需要循环获取articles值
	for _, article := range articles {
		_, err = o.LoadRelated(article, "Types")
	}
	for _, article := range articles {
		_, err = o.LoadRelated(article, "Read")
	}
	if err != nil {
		c.GetError()
		return
	}
	if num == 0 {
		c.Data["Exist"] = false
	}
	var articles1 []*models.Article
	num, err= qs.Filter("Status",1).OrderBy("-Updatetime").RelatedSel().Limit(10).All(&articles1) //.RelatedSel()查出包括User中的数据

	//当articles为多个值是需要循环获取articles1值
	for _, article1 := range articles1 {
		_, err = o.LoadRelated(article1, "Types")
	}
	for _, article1 := range articles1 {
		_, err = o.LoadRelated(article1, "Read")
	}
	qs1 := o.QueryTable("read")
	var reads []*models.Read
	num1, err := qs1.OrderBy("-ReadNum").Limit(10).All(&reads)
	for _, read := range reads {
		_, err = o.LoadRelated(read, "Article")
	}

	c.Data["Exist"] = true
	c.Data["Articles"] = articles
	c.Data["Articles1"] = articles1
	c.Data["Reads"] = reads
	c.Data["Num"] = num1
	c.TplName = "index.html"
}

// @router /article/list [get]
func (c *ArticleController) GetList()  {
	c.LoginStatus()
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	var articles []*models.Article
	num, err:= qs.Filter("Status",1).OrderBy("-Updatetime").RelatedSel().All(&articles) //.RelatedSel()查出包括User中的数据
	if err != nil {
		c.GetError()
		return
	}
	if num == 0 {
		c.Data["Exist"] = false
	}
	//当articles为多个值是需要循环获取articles1值
	for _, article := range articles {
		_, err = o.LoadRelated(article, "Types")
	}
	for _, article := range articles {
		_, err = o.LoadRelated(article, "Read")
	}
	qs1 := o.QueryTable("read")
	var reads []*models.Read
	num1, err := qs1.OrderBy("-ReadNum").Limit(10).All(&reads)
	for _, read := range reads {
		_, err = o.LoadRelated(read, "Article")
	}
	c.Data["Exist"] = true
	c.Data["Articles"] = articles
	c.Data["Reads"] = reads
	c.Data["Num"] = num1
	c.TplName = "jie/list.html"
}

// @router /article/detail/?:id [get]
func (c *ArticleController) GetDetail() {
	c.LoginStatus()
	id := c.Ctx.Input.Param(":id")
	key := c.ReadingStatistics(id)
	//设置session 用于统计阅读数量
	c.SetSession(key, true)

	o := orm.NewOrm()
	var article models.Article
	qs := o.QueryTable("article")
	_ , err := qs.Filter("Id", id).Filter("Status",1).RelatedSel().All(&article)
	if err != nil {
		c.GetError()
		return
	}
	_, err = o.LoadRelated(&article, "Types")
	_, err = o.LoadRelated(&article, "Read")
	if err != nil {
		c.GetError()
		return
	}
	c.Data["Article"] = article

	var comments []models.Comment
	qs1 := o.QueryTable("comment")
	num, _ := qs1.Filter("Article", &article).OrderBy("-Createtime").RelatedSel().All(&comments)
	c.Data["CommentNum"] = num
	if num != 0 {
		c.Data["Comments"] = comments
	}
	qs2 := o.QueryTable("read")
	var reads []*models.Read
	num1, err := qs2.OrderBy("-ReadNum").Limit(10).All(&reads)
	for _, read := range reads {
		_, err = o.LoadRelated(read, "Article")
	}
	c.Data["Reads"] = reads
	c.Data["Num"] = num1
	c.TplName = "jie/detail.html"
}

//点赞统计
// @router /article/like/?:id [post]
//func (c *ArticleController) PostLike() {
//
//}