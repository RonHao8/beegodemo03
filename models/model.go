package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// User 表的设计
type User struct {
	Id      int
	Name    string `orm:"unique"`
	Pwd     string
	Article []*Article `orm:"rel(m2m)"`
}

// Article 文章结构体
type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`
	Atime    time.Time `orm:"auto_now"`
	Acount   int       `orm:"default(0);null"`
	Acontent string
	Aimg     string
	Atype    string

	ArticleType *ArticleType `orm:"rel(fk)"`
	User        []*User      `orm:"reverse(many)"`
}

// ArticleType 类型表
type ArticleType struct {
	Id      int
	Tname   string
	Article []*Article `orm:"reverse(many)"`
}

func init() {
	// set default database  设置数据库
	orm.RegisterDataBase("default", "mysql", "root:guanhao@tcp(127.0.0.1:3306)/beegodemo?charset=utf8", 30)
	//把上面的username改成你的数据库用户名，password改成数据库密码,db_name改成数据库名。

	// register model  注册表，就是把结构体创建为数据库里的表
	orm.RegisterModel(new(User), new(Article), new(ArticleType))

	// create table
	orm.RunSyncdb("default", false, true)
}
