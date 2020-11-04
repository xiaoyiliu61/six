package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type LoginSmsController struct {
	beego.Controller
}

func (l *LoginSmsController) Get() {
	l.TplName="login_sms.html"
}
/*
短信验证码登录功能
*/
func (l *LoginSmsController) Post() {
  var smsLogin models.SmsLogin
  err:=l.ParseForm(&smsLogin)
	if err != nil {
		l.Ctx.WriteString("抱歉，验证码登录失败")
		return
	}
}