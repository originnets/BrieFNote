package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type User struct {
	Id			int
	Email 		string		`orm:"unique"`
	Name 		string		`orm:"unique"`
	Password	string
	Avatar		string		`orm:"null"`
	Comment 	[]*Comment 	`orm:"reverse(many)"`
	Article		[]*Article 	`orm:"reverse(many)"`
	Role 		int			`orm:"default(1)"`	//0代表管理员, 1 代表正常用户
	Createtime	time.Time	`orm:"auto_now_add;type(datetime)"`
	Updatetime	time.Time	`orm:"auto_now;type(datetime)"`
}

type Type struct {
	Id			int
	Typename 	string
	Articles 	[]*Article 	`orm:"reverse(many)"`
}

type Article struct {
	Id 			int
	Title		string
	Img 		string		`orm:"null"`
	User  		*User  		`orm:"rel(fk)"`
	Content  	string		`orm:"type(text)"`
	Types		[]*Type		`orm:"rel(m2m)"`
	//Like		*Like		`orm:"reverse(one)"`
	Read		*Read		`orm:"reverse(one)"`
	Comment 	[]*Comment 	`orm:"reverse(many)"`
	Createtime	time.Time	`orm:"auto_now_add;type(datetime)"`
	Updatetime	time.Time	`orm:"auto_now;type(datetime)"`
	Status 		int			//0草稿,1发布
	Settop 		bool		//false不置顶, true置顶
	Thickening 	bool 		//false不精帖, true精帖
}

//type Like struct {
//	Id 			int
//	Article 	*Article 	`orm:"rel(one)"`
//	LikeNum 	int 		`orm:"default(0)"`
//}

type Read struct {
	Id 			int
	Article 	*Article 	`orm:"rel(one)"`
	ReadNum 	int 		`orm:"default(0)"`
}

type Comment struct {
	Id 			int
	User 		*User 		`orm:"rel(fk)"`
	Article  	*Article 	`orm:"rel(fk)"`
	Content		string
	Createtime	time.Time	`orm:"auto_now_add;type(datetime)"`
	Top 		*Comment	`orm:"null;rel(fk)"`	//顶级评论
	Upper		*Comment	`orm:"null;rel(fk)"`	//上级评论
}

func init(){

	//读取app.conf文件中的配置
	username := beego.AppConfig.String("username")
	hostname := beego.AppConfig.String("hostname")
	password := beego.AppConfig.String("password")
	dbname := beego.AppConfig.String("dbname")
	port := beego.AppConfig.String("port")

	//当port 为空指定默认端口3306
	if port == "" {
		port = "3306"
	}

	//设置连接串
	conn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn, 30) //连接数据库
	orm.RegisterModel(new(User), new(Article), new(Type), new(Read), new(Comment)) //注册表
//	orm.RunSyncdb("default", false, true) //force 是否同步表结构, verbone 是否创建表
}