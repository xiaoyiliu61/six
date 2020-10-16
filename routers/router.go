package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//rootpath:路由
    beego.Router("/", &controllers.MainController{})
   //用户注册接口
    beego.Router("/register",&controllers.RegisterController{})
    //用户登录的接口
    beego.Router("/login",&controllers.LoginController{})

    beego.Router("/login.html",&controllers.LoginController{})
    //用户上传文件的功能
    beego.Router("/upload",&controllers.UploadFileController{})
}
