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
	upwd := this.GetString("password")
	autoLogin := this.GetString("autoLogin") == "on"
	conf_uname, _ := beego.AppConfig.String("uanme")
	conf_pwd, _ := beego.AppConfig.String("upwd")
	if uname == conf_uname && upwd == conf_pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("upwd", upwd, maxAge, "/")
	}

	this.Redirect("/", 301)
	return
}
