package utils

import (
	"bytes"
	"encoding/binary"
)

/*
讲一个int64转化为[]byte字节切片
*/

func In64ToByte(num int64) ([]byte,error) {

	//Buffer：缓冲区。
	buff:=new(bytes.Buffer)//通过new实例化一个缓冲区
	//buff.Write()//通过一系列的Write方法向缓冲区写入数据
	//buff.Bytes()//通过bytes方法从缓冲区中获取数据
	err:=binary.Write(buff,binary.BigEndian,num)//binary：二进制
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}

func StringToBytes(data string) []byte {
	return []byte(data)
}
