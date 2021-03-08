package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"web/controllers"
)

func init() {
	// 首页
	beego.Router("/", &controllers.IndexController{})
	// 登陆
	beego.Router("/login", &controllers.LoginController{})

	beego.Router(
		"/test",
		&controllers.TestController{},
	)
}
