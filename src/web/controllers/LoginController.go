package controllers

import (
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
	uname := this.GetString("account")
	//upwd := this.GetString("password")
	//autoLogin := this.GetString("autoLogin") == "on"
	if uname == beego.AppConfig.String("appname") {

	}
	return
}
