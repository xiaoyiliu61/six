package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}

/*
用于跳转到实名认证页面
*/
func (u *UserKycController) Get() {
	u.TplName="user_kyc_html"
}

/*
用于处理实名认证业务
*/
func (u *UserKycController) Post() {
	//1.解析前端的数据
	var user models.User
	err:=u.ParseForm(&user)
	if err != nil {
		u.Ctx.WriteString("抱歉，数据解析失败，请重试！")
		return
	}

	//2.把用户的实名认证的更新到数据库的用户表当中
   _,err=user.UpdateUser()
    //3.判断实名认证操作结果
	if err != nil {
		u.Ctx.WriteString("抱歉，实名认证失败，请重试！")
		return
	}
	u.TplName="home.html"
}
