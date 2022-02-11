package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["sign/controllers:SignController"] = append(beego.GlobalControllerRouter["sign/controllers:SignController"],
        beego.ControllerComments{
            Method: "Sign",
            Router: "/sign/?:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
