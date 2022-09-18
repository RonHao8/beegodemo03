package controllers

import (
	"beegodemo03/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

func (c *MainController) Get() {
	c.TplName = "register.html"
}

// Register 注册功能业务逻辑
func (c *MainController) Register() {
	//resp := make(map[string]interface{})
	//1.从前端获取数据
	username := c.GetString("userName")
	pwd := c.GetString("pwd")
	//fmt.Println("username=", username, "pwd=", pwd)
	//json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	//2.判断数据的合法性
	if username == "" || pwd == "" {
		fmt.Println(models.RECODE_DATAERR)
		c.Redirect("/register", 302)
		return
	}
	if len(pwd) < 6 || len(pwd) > 20 {
		fmt.Println(models.RECODE_DATAERR)
		c.Redirect("/register", 302)
		return
	}
	//3.将数据插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println(models.RECODE_DBERR)
		c.Redirect("/register", 302)
		return
	}
	//4.返回登陆界面
	c.Redirect("/login", 302)

	//注意
	//注册界面如果注册的用户名存在，会出现错误，等待解决
}
