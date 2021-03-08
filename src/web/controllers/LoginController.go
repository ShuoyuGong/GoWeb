package controllers

import (
	"fmt"
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
	fmt.Sprint(this.Input())
	return
}
