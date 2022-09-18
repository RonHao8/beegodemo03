package routers

import (
	"beegodemo03/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Register")
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.IndexController{}, "get:ShowIndex")
	//beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
	beego.Router("/addType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/content", &controllers.ArticleController{}, "get:ShowContent")
	beego.Router("/update", &controllers.ArticleController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/delete", &controllers.ArticleController{}, "get:HandleDelete")
	beego.Router("/deleteType", &controllers.TypeController{}, "get:DeleteType")
}
