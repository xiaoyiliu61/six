package main

import (
"DataCertPlatform/db_mysql"
_ "DataCertPlatform/routers"
"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()
//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
