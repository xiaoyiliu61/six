package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

/*	block0:=blockchain.CreateGenesisBlock()//创建创世区块
	block1:=blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})
	fmt.Printf("block0的哈希：%x\n",block0.Hash)
	fmt.Printf("block1的哈希：%x\n",block1.Hash)
	fmt.Printf("block1的PrevHash：%x\n",block1.PrevHash)

	block0Bytes:=block0.Serialize()
	fmt.Println("创世区块gob序列化后：",block0Bytes)
    deBlock0,err:=blockchain.DeSerialize(block0Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("反序列化后的区块高度是：",deBlock0.Height)
    return

	  序列化：Marshal
	   将数据从内存中形式转换为可以持久化存储在硬盘上或者在网络上传输的形式，称为序列化
	  反序列化：Unmarshal

    blockJson,_:=json.Marshal(block0)
    fmt.Println("通过json序列化以后的block：",string(blockJson))*/

    bc:=blockchain.NewBlockChain()
    fmt.Printf("创世区块的哈希值：%x\n",bc.LastHash)
    bc.SaveData([]byte("用户要保存数据"))
	return


	db_mysql.Connect()
//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
