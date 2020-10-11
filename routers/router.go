package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//rootpath:路由
    beego.Router("/", &controllers.MainController{})

}
