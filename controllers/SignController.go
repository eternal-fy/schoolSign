package controllers

import (
	"github.com/astaxie/beego"
	"sign/service"
)

type SignController struct {
	beego.Controller
}

//  @router /sign/?:key [get]
func (c *SignController) Sign() {
	fetch := service.FetchString("http://180.76.172.177/18nodone")
	classNames := []string{"软件工程1804"}
	data, _ := service.GetData(classNames, fetch)
	c.Ctx.WriteString(data)
}
