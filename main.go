package main

import (
	_ "beegodemo03/models"
	_ "beegodemo03/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

func main() {
	//ignoreStaticPath()
	beego.AddFuncMap("showprepage", prepage)
	beego.AddFuncMap("shownextpage", shownextpage)
	beego.Run()
}

func ignoreStaticPath() {
	//pattern 路由规则，可以根据一定的规则进行路由，如果你全匹配可以用"*"
	// beego.InsertFilter("*",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url:", orpath)
	//如果请求url还有api字段，说明指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	if strings.Index(orpath, "test") >= 0 {

		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}

//试图函数，获取上一页页码

/*
1.在试图中定义视图函数函数名    | funcName

2.一般在main.go里面实现试图函数

3.在main函数里面把实现的函数和试图函关联起来   beego.AddFuncMap()
*/
func prepage(pageindex int) (preIndex int) {
	preIndex = pageindex - 1
	return
}

func shownextpage(pageindex int) (nextIndex int) {
	nextIndex = pageindex + 1
	return
}
