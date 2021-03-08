package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home/index/index.html"
}
