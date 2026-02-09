package routers

import (
	"demoApp1/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.HomeController{}, "get:Home")

	beego.Router("signup/", &controllers.AccountController{}, "get:Signup")
	beego.Router("login/", &controllers.AccountController{}, "get:Login")

}
