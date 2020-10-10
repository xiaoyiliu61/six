package main

import (
"DataCertPlatform/db_mysql"
_ "DataCertPlatform/routers"
"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("img", "./static/img")
	beego.SetStaticPath("css", "./static/css")

	beego.Run()
}
