package controllers

import (
	"BrieFNote/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type CommentController struct {
	BaseController
}

// @router /comment/add/?:id [post]
func (c *CommentController) PostAddComment(){
	c.LoginStatus()
	if c.IsLogin == false {
		c.Revert(1002, "请先登入后评论" )
		return
	}
	id := c.Ctx.Input.Param(":id")
	content := c.GetMUstString("content", "评论内容为空")
	if content == "" {
		return
	}
	id1 , _:= strconv.Atoi(id)
	replyid,_ := c.GetInt("replyid")
	o := orm.NewOrm()
	user := models.User{}
	user = c.User
	err := o.Read(&user)
	if err != nil {
		c.Revert(3002, "用户不存在")
		return
	}

	article := models.Article{Id:id1}
	err = o.Read(&article)
	if err != nil {
		c.Revert(3002, "文章不存在")
		return
	}

	comment := models.Comment{}
	comment.User = &user
	comment.Article = &article
	comment.Content = content

	if replyid == 0 {
		comment.Top = nil
		comment.Upper = nil
	}else {
		comment1 := models.Comment{Id:replyid}
		err =o.Read(&comment1)
		if err != nil {
			c.Revert(3002, "回复评论Id不存在")
			return
		}
		if comment1.Top == nil {
			comment.Top = &comment1
		}else {
			comment2 := models.Comment{Id:comment1.Top.Id}
			err = o.Read(&comment2)
			if err != nil {
				c.Revert(3002, "回复评论Id不存在")
				return
			}
			comment.Top = &comment2
		}
		comment.Upper = &comment1
	}
	_, err = o.Insert(&comment)
	if err != nil {
		c.Revert(3002, "评论添加失败")
		return
	}
	c.Revert(0, "评论添加成功")
}