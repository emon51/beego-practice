package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Home() {
	c.Data["name"] = "Emon"
	c.Data["email"] = "emon@gmail.com"
	c.TplName = "home.html"
}
