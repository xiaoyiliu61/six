package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
_ "DataCertPlatform/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0:=blockchain.CreateGenesisBlock()//创建创世区块
	fmt.Printf("block0的哈希：%x\n",block0.Hash)
	block1:=blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})
	fmt.Printf("block0的哈希：%x\n",block0.Hash)
	fmt.Printf("block1的哈希：%x\n",block1.Hash)
	fmt.Printf("block1的PrevHash：%x\n",block1.PrevHash)

	/* 序列化：Marshal
	   将数据从内存中形式转换为可以持久化存储在硬盘上或者在网络上传输的形式，称为序列化
	  反序列化：
	*/
    blockJson,_:=json.Marshal(block0)
    fmt.Println("通过json序列化以后的block：",string(blockJson))


	db_mysql.Connect()
//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
