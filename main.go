package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0:=blockchain.CreateGenesisBlock()
	block1:=blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	fmt.Println(block0,block1)


	db_mysql.Connect()
//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
