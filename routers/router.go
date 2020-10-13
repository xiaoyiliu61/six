package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//rootpath:路由
    beego.Router("/", &controllers.MainController{})

    beego.Router("/register",&controllers.RegisterController{})

    beego.Router("/login",&controllers.LoginController{})

    beego.Router("/login.html",&controllers.LoginController{})


}
