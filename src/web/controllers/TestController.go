package controllers

import (
beego "github.com/beego/beego/v2/server/web"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.TplName = "test.html"
}

