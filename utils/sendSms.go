package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

type SmsCode struct {
	Code string `json:"code"`
}

type SmsResult struct {
	BizId string
	Code string
	Message string
	RequestId string
}

const SMS_TLP_REGISTER="SMS_205393604"
const SMS_TLO_LOGIN  ="SMS_205398654"
const SMS_TLO_KYC  =""
/*
该函数用于发送一条短消息
*/
func SendSms(phone string, code string,templateType string) (*SmsResult,error) {
	config:=beego.AppConfig
	accessKey:=config.String("sms_access_key")
	accessKeySecret := config.String("sms_access_secret")
	client,err:=dysmsapi.NewClientWithAccessKey("cn-hangzhou",":LTAI4FydRQswZLg94JhLGRLp","8Stma2gzzZaGYGo0bQckBMziHO3FWx")
	if err != nil {
		return err
	}
	request:=dysmsapi.CreateAddSmsSignRequest()
	request.PhoneNumbers = phone
	request.SignName="线上餐厅"
	request.TemplateCode = templateType
	//{"code":"xxxxxx";json格式}
	smsCode:=SmsCode{
		Code:code,
	}
	smsBytes,_:=json.Marshal(smsCode)
	request.TemplateParam = String(smsBytes)

	response,err:=client.SendSms(request)
	if err != nil {
		return nil,err
	}
	smsResult:=&SmsResult{
		BizId:response.BizId,
		Code:response.Code,
		Message:response.Message,
		RequestId:response.RequestId,
	}
	return smsResult, nil
}
func GenRandCode(width int) string {
	numeric:=[10]byte{0,1,2,3,4,5,6,7,8,9}
    r:=len(numeric)
    rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb,"%d",numeric[rand.Intn(r)])
	}
	return sb.String()
}