package routers

import (
	"github.com/astaxie/beego"
	. "sign/controllers"
)

func init() {
	beego.Router("/", &MainController{})
	beego.Include(&SignController{})
}
