package utils

import (
	"fmt"
	"io"
	"os"
)

func SaveFile(fileName string,file io.Reader) (int64,error) {
	saveFile,err:=os.OpenFile(fileName,os.O_CREATE|os.O_RDWR,777)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	length,err:=io.Copy(saveFile,file)
	if err != nil {
		return -1,err
	}
	return length,err
}
