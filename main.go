package main

import (
	"github.com/astaxie/beego"
	_ "sign/routers"
	_ "sign/service"
)

func main() {
	beego.Run()
}
