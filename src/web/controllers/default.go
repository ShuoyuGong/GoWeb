package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.TplName = "index.html"

}
