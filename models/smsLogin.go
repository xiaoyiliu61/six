package models

type SmsLogin struct {
   Phone string `form:"phone"`
   code string `form:"code"`
}
