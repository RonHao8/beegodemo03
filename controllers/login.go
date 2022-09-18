package controllers

import (
	"beegodemo03/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

// ShowLogin 登陆页面
func (c *LoginController) ShowLogin() {
	c.TplName = "login.html"
}

// HandleLogin 登陆功能业务逻辑
func (c *LoginController) HandleLogin() {
	//resp := make(map[string]interface{})
	//1.获取前端传过来的数据
	username := c.GetString("userName")
	pwd := c.GetString("pwd")
	fmt.Println(username, pwd)
	//2.判断数据的合法性
	if username == "" || pwd == "" {
		fmt.Println(models.RECODE_DATAERR)
		c.TplName = "login.html"
		return
	}
	if len(pwd) < 6 || len(pwd) > 20 {
		fmt.Println(models.RECODE_DATAERR)
		c.TplName = "login.html"
		return
	}
	//3.根据用户名取数据库中查找
	user := models.User{Name: username}
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("name", username).All(&user)
	if err != nil {
		fmt.Println(models.RECODE_DBERR)
		c.TplName = "login.html"
		return
	}

	//4.校验密码是否正确
	if pwd != user.Pwd {
		fmt.Println(models.RECODE_PWDERR)
		c.TplName = "login.html"
		return
	}
	//5.跳转到首页
	c.Redirect("/index", 302)
}
