package controllers

import (
    beego "github.com/beego/beego/v2/server/web"
)

type AccountController struct {
    beego.Controller
}

func (c *AccountController) Signup() {
    c.TplName = "account/signup.html"
}

func (c *AccountController) Login() {
    c.TplName = "account/login.html"
}
