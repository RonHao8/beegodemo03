package controllers

import (
	"beegodemo03/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// ShowAddType 显示添加类型界面
func (c *ArticleController) ShowAddType() {
	//fmt.Println("此处仅作为测试git")
	//思路
	//1.从数据库查询文章类型
	o := orm.NewOrm()
	var artiTypes []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&artiTypes)
	if err != nil {
		fmt.Println("没有获取到类型数据")
		return
	}
	//2.将查询到的数据返回给前端
	c.Data["articleType"] = artiTypes
	c.TplName = "addType.html"
}

// HandleAddType 处理添加类型传输的信息
func (c *ArticleController) HandleAddType() {
	//思路
	//1.从前端获取数据typeName: 初中
	typeName := c.GetString("typeName")
	//2.判断数据是否合法
	if typeName == "" {
		fmt.Println("获取文章类型数据错误")
		return
	}
	//3.将数据插入数据库
	o := orm.NewOrm()
	artiType := models.ArticleType{}
	artiType.Tname = typeName
	_, err := o.Insert(&artiType)
	if err != nil {
		fmt.Println("插入文章类型错误")
		return
	}
	//4.返回文章类型界面
	c.Redirect("/addType", 302)

}

// ShowAdd 显示添加文章界面
func (c *ArticleController) ShowAdd() {
	//思路
	//1.从数据库查询文章类型数据
	o := orm.NewOrm()
	var artiTypes []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&artiTypes)
	if err != nil {
		fmt.Println("文章类型查询失败")
		return
	}
	//2.将数据返回给前端
	c.Data["articleType"] = artiTypes
	c.TplName = "add.html"
}

// HandleAdd 处理添加文章界面数据
func (c *ArticleController) HandleAdd() {
	//articleName:
	//select: 1
	//content:
	//uploadname: (binary)
	//思路
	//1.从前端获取数据
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	id, err := c.GetInt("select")
	if err != nil {
		fmt.Println("获取前端select错误")
		return
	}
	f, h, err := c.GetFile("uploadname")
	defer f.Close()

	//2.对图片uploadname进行处理
	//1.要限定格式
	fileExt := path.Ext(h.Filename)
	if fileExt != ".jpg" && fileExt != ".png" {
		fmt.Println("上传文件格式错误")
		return
	}
	//2.限制大小
	if h.Size > 50000000 {
		fmt.Println("上传文件过大")
		return
	}
	//3.需要对文件重命名，防止文件名重复
	//fileName := time.Now().Format("2006-01-02 15:04:05") + fileExt
	fileName := time.Now().Format("2006-01-02 15:04:05") + fileExt
	if err != nil {
		fmt.Println("上传文件失败")
		return
	} else {
		//将上传的文件保存到新路径
		//c.SaveToFile("uploadname", "./static/img/"+fileName)
		c.SaveToFile("uploadname", "./static/img/"+fileName)
	}
	//3.判断数据的合法性
	if artiContent == "" || artiName == "" {
		fmt.Println("添加文章数据错误")
		return
	}
	//4.将数据插入数据库中
	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	//arti.Aimg = "./static/img/" + fileName
	arti.Aimg = "./static/img/" + fileName
	//查找type对象
	artiType := models.ArticleType{Id: id}
	o.Read(&artiType)
	arti.ArticleType = &artiType

	if _, err := o.Insert(&arti); err != nil {
		fmt.Println("插入数据库错误")
		return
	}
	//5.返回文章界面
	c.Redirect("/index", 302)

}

// ShowContent 显示内容详情页面
func (c *ArticleController) ShowContent() {
	//思路
	//1.从前端获取文章id
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("获取文章ID错误", err)
		return
	}
	//2.从数据库中查询数据
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	if err := o.Read(&arti); err != nil {
		fmt.Println("数据库查询错误", err)
		return
	}
	//3.将数据传给前端
	c.Data["article"] = arti

	c.TplName = "content.html"
}

// ShowUpdate 显示编辑界面
func (c *ArticleController) ShowUpdate() {
	//思路
	//1.从前端获取id
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("获取文章ID错误", err)
		return
	}
	//2.从数据库查询数据
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	if err := o.Read(&arti); err != nil {
		fmt.Println("查询错误", err)
		return
	}
	//3.将数据返回给前端
	c.Data["article"] = arti
	c.TplName = "update.html"
}

// HandleUpdate 处理更新业务数据
func (c *ArticleController) HandleUpdate() {
	//id: 2
	//articleName: shuxue
	//content: 1212121212xiugai
	//uploadname: (binary)

	//思路
	//1.从前端获取数据
	id, _ := c.GetInt("id")
	artiName := c.GetString("articleName")
	content := c.GetString("content")

	//2.对图片uploadname进行处理
	f, h, err := c.GetFile("uploadname")
	var fileName string
	if err != nil {
		fmt.Println("上传文件失败")
		return
	} else {
		defer f.Close()

		//1.要限定格式
		fileExt := path.Ext(h.Filename)
		if fileExt != ".jpg" && fileExt != ".png" {
			fmt.Println("上传文件格式错误")
			return
		}
		//2.限制大小
		if h.Size > 50000000 {
			fmt.Println("上传文件过大")
			return
		}
		//3.需要对文件重命名，防止文件名重复
		fileName = time.Now().Format("2006-01-02 15:04:05") + fileExt
		c.SaveToFile("uploadname", "./static/img/"+fileName)
	}

	//3.判断数据的合法性
	if artiName == "" || content == "" {
		fmt.Println("更新数据获取失败")
		return
	}
	//4.对数据进行更新，将数据插入数据库中
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	if err := o.Read(&arti); err != nil {
		fmt.Println("查询数据错误")
		return
	}
	arti.ArtiName = artiName
	arti.Acontent = content
	arti.Aimg = "./static/img/" + fileName
	if _, err := o.Update(&arti, "ArtiName", "Acontent", "Aimg"); err != nil {
		fmt.Println("更新数据显示错误")
		return
	}
	//5.返回列表页面
	c.Redirect("/index", 302)

}

func (c *ArticleController) HandleDelete() {
	//思路
	//1.从前端获取文章id
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("获取id数据错误")
		return
	}
	//2.在数据库进行删除
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	//查询数据库中是否有为该id的文章
	if err := o.Read(&arti); err != nil {
		fmt.Println("数据库查询失败")
		return
	}
	if _, err := o.Delete(&arti); err != nil {
		fmt.Println("数据库删除失败")
		return
	}
	//3.返回列表页面
	c.Redirect("/index", 302)
}
