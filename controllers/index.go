package controllers

import (
	"beegodemo03/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
)

type IndexController struct {
	beego.Controller
}

// ShowIndex 显示列表页面内容
func (c *IndexController) ShowIndex() {
	//思路
	//1.获取前端传过来的id
	o := orm.NewOrm()
	id, _ := c.GetInt("select")
	//if err != nil {
	//	fmt.Println("从前端获取id失败")
	//	return
	//}
	//2.从数据库查询所有文章的信息
	var articles []models.Article
	_, err := o.QueryTable("Article").All(&articles)
	if err != nil {
		fmt.Println("查询所有文章失败")
		return
	}
	//3.进行分页处理
	//获得数据总数，总页数，当前页码
	count, err := o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__Id", id).Count()
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	pageSize := int64(2)                //每页显示数据条目
	index, err := c.GetInt("pageIndex") //当前页码
	if err != nil {
		index = 1
	}
	pageCount := math.Ceil(float64(count) / float64(pageSize)) //总页数=数据总数/每页显示数据条目

	if index <= 0 || index > int(pageCount) {
		index = 1
	}

	start := (int64(index) - 1) * pageSize

	//分页后当前页面的文章
	var artis []models.Article
	o.QueryTable("Article").Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__Id", id).All(&artis)

	//获取类型数据
	var artisTypes []models.ArticleType
	_, err = o.QueryTable("ArticleType").All(&artisTypes)
	if err != nil {
		fmt.Println("获取类型错误")
		return
	}

	c.Data["articleType"] = artisTypes //获取类型数据
	c.Data["pageCount"] = pageCount    //总页数
	c.Data["count"] = count            //数据总数
	c.Data["articles"] = artis         //分页后当前页面的文章
	c.Data["pageIndex"] = index        //当前页码
	c.Data["typeid"] = id              //文章类型ID

	c.TplName = "index.html"
}
