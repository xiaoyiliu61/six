package models

import "DataCertPlatform/db_mysql"

type SmsRecord struct {
	BizId string
    Phone string
	Code string
	Status string
	Message string
	TimeStamp int64
}

func QuerySmsRecord(bizId string, phone string, code string)  {
	db_mysql.Db.QueryRow("")
}

func (s *SmsRecord) SaveSmsRecord() {
	db_mysql.Db.Exec("insert into sms_record(biz_id,phone,code,status,message,timestamp )")
}
