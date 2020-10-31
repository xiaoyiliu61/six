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
	var user models.User
	err:=u.ParseForm(&user)
	if err != nil {
		u.Ctx.WriteString("抱歉，数据解析失败，请重试！")
		return
	}

   _,err=user.UpdateUser()

	if err != nil {
		u.Ctx.WriteString("抱歉，实名认证失败，请重试！")
		return
	}
	u.TplName="home.html"
}
