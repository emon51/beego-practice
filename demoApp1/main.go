package main

import (
	_ "demoApp1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

