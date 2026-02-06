package routers

import (
	"demoApp/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/", &controllers.RootController{}, "get:Get")
	beego.Router("/", &controllers.RootController{}, "post:Post")
	beego.Router("/", &controllers.RootController{}, "put:Put")
	beego.Router("/", &controllers.RootController{}, "delete:Delete")

	beego.Router("/home", &controllers.HomeController{}, "get:Get")
	beego.Router("/home", &controllers.HomeController{}, "post:Post")
	beego.Router("/home", &controllers.HomeController{}, "put:Put")
	beego.Router("/home", &controllers.HomeController{}, "delete:Delete")

}
