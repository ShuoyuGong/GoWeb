package controllers

import (
	"github.com/astaxie/beego/context"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["IsHome"] = checkAccountLogin()
	this.TplName = "home/index/index.html"
}

// 检查是否在登陆中
func checkAccountLogin() bool {
	cookie_uanme := context.Context.GetCookie("uname")
	cookie_upwd := context.Context.GetCookie("upwd")
	conf_uname, _ := beego.AppConfig.String("uanme")
	conf_pwd, _ := beego.AppConfig.String("upwd")
	if cookie_uanme == conf_uname && cookie_upwd == conf_pwd {
		return true
	}
	return false
}
