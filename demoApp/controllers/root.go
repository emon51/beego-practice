package controllers 
import (
	beego "github.com/beego/beego/v2/server/web"
)

type RootController struct {
	beego.Controller
}




func (c *RootController) Get() {
	c.Ctx.WriteString("This is the (Get) page..from root.go")
}


func (c *RootController) Post() {
	c.Ctx.WriteString("This is the (Post) page..from root.go")
}


func (c *RootController) Put() {
	c.Ctx.WriteString("This is the (put) page..from root.go")
}


func (c *RootController) Delete() {
	c.Ctx.WriteString("This is the (Delete) page..from root.go")
}