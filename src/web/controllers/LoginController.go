package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// 返回登陆页面
func (this *LoginController) Get() {
	this.TplName = "home/login/login.html"
}

// 处理登陆请求
func (this *LoginController) Post() {
	//var uname string
	//var upwd string
	//uname = this.GetString("account")
	//upwd = this.GetString("password")
	//o := orm.NewOrm()
	//res := o.Raw("SELECT * FROM user where user_name = ? AND password = ?", uname, upwd)
	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable("user")
	//if beego.AppConfig.String("uname") == uname {
	//}
	//this.Ctx.WriteString(uname)
	//this.Ctx.WriteString(upwd)
	this.Ctx.WriteString(fmt.Sprint(qs))
	//var conf_uname string
	//conf_uname = beego.AppConfig.String("uname")
	//fmt.Println(beego.AppConfig.String("uname"))

	//autoLogin := this.GetString("autoLogin") == "on"
	//if uname == beego.AppConfig.String("appname") {

	//}
	return
}
