package controllers

import (
	"beetwo2/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	user:=c.Ctx.Input.Query("user")
	psd:=c.Ctx.Input.Query("psd")
	fmt.Println(user,psd)
	if user!="yuanshen"||psd!="123456" {
		c.Ctx.ResponseWriter.Write([]byte("请重试，谢谢"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据正确"))
	c.TplName="index.tpl"
}
func (c *MainController)post{
	dataByes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据加载失败，请重试")
		return
	}
	//json包解析
	var person models.Person
	err=json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名",person.User,"年龄",person.Age)
	c.Ctx.WriteString("用户名是："+person.User)

	user:=c.Ctx.Request.FormValue("user")
	psd:=c.Ctx.Request.FormValue("psd")
	fmt.Println("用户名是"+user+"，密码是"+psd)

	if user!="原神"||psd!="123456" {
		
	}
}
