package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Ctx.WriteString("This is the (Get) page..from home.go")
}



func (c *HomeController) Post() {
	c.Ctx.WriteString("This is the (Post) page..from home.go")
}

func (c *HomeController) Put() {
	c.Ctx.WriteString("This is the (Put) page..from home.go")
}



func (c *HomeController) Delete() {
	c.Ctx.WriteString("This is the (Delete) page..from home.go")
}

