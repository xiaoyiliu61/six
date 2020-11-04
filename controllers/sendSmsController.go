package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
	"DataCertPlatform/utils"
	"time"
)

type SentSmsController struct {
	beego.Controller
}

func (s *SentSmsController) Post() {
	var smsLogin models.SmsLogin
	err:=s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("")
		return
	}
	phone:=smsLogin.Phone
	code:=utils.GenRandCode(6)
    result,err:=utils.SendSms(phone,code,utils.SMS_TLP_REGISTER)
	if err != nil {
		s.Ctx.WriteString("发送验证码失败，请重试！")
		return
	}
	if len(result.BizId)==0 {
		s.Ctx.WriteString(result.Message)
		return
	}
	smsRecord:=models.SmsRecord{
		BizId:     result.BizId,
		Phone:     phone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStamp: time.Now().Unix(),
	}
	_,err=smsRecord.SaveSmsRecord()
	if err != nil {
		s.Ctx.WriteString("")
		return
	}
	//保存成功 bizId
	s.Data["Phone"]=smsLogin.Phone
	s.Data["BizId"]=smsRecord.BizId
	//验证码登录
	s.TplName="login_sms.html"
}
