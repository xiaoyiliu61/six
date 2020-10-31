package models

import (
	"bytes"
	"encoding/gob"
)

/*
该结构体用于定义链上数据保存的信息
*/

type CertRecord struct {
	CertId []byte//认证id，本质是一个md5值
	CertHash []byte //存证文件的hash值，本质是一个sha256值
	CertName string //认证人的姓名
	Phone string //联系方式
	CertCard string //身份证号码
	FileName string //认证文件的名称
	FileSize int64 //文件的大小
	CertTime int64 //认证时间
}

/*
序列化操作
*/

func (c CertRecord) Serialize() ([]byte,error){
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(),err
}

/*
反序列化生成一个CertRecord结构体实例
*/
func DeserializerCertRecord(data []byte) (*CertRecord,error) {
	var certRecord *CertRecord
	err:=gob.NewDecoder(bytes.NewReader(data)).Decode(&certRecord)
	return certRecord,err
}