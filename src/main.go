package main

import (
	"github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController)Get()  {
	this.Ctx.WriteString("Hello,World!")
}
func main()  {
	beego.Router("/",&HomeController{})
	beego.Run()
}
 