package controllers

import (
	"DataCertPlatform/blockchain"
	"fmt"
	"github.com/astaxie/beego"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get() {
	//1.解析和接收前端页面传递的数据
	cert_id:=c.GetString("cert_id")
	//2.到区块链上查询区块数据
	block,err:=blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if err != nil {
		c.Ctx.WriteString("抱歉，查询链上数据遇到错误，请重试")
		return
	}
	if block == nil {
     c.Ctx.WriteString("抱歉，未查到链上数据")
		return
	}
	fmt.Println("查询到的区块高度：",block.Height)
	c.Data["CertId"]=string(block.Data)
	//3.跳转证书详情页面
	c.TplName="cert_detail.html"
}