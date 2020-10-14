package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)


//该控制器结构体用于处理文件上传的功能

type UploadFileController struct {
	beego.Controller
}

/*
 该post方法用于处理用户在客户端提交的认证文件
*/
func (u *UploadFileController) Post() {
	//1.解析数据上传的文件
	//用户上传的自定义的标题
	title:=u.Ctx.Request.PostFormValue("uoload_title")

    //用户上传的文件
    file,header,err:=u.GetFile("yixiaoliu")
	if err != nil {
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}

	fmt.Println("自定义的标题：",title)
	//获取上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
    //eg:支持jpg，png类型，不支持jpeg。gif类型
	fileNameSlice:=strings.Split(header.Filename,".")
    fileType:=fileNameSlice[1]
	if  fileType!="jpg"&&fileType!="png"{
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传符合格式的文件")
		return
	}
    fmt.Println("文件的大小：",header.Size)

    fmt.Println("上传的文件:",file)
    u.Ctx.WriteString("已获取到上传文件。")
    //文件大小 200kb
    config:=beego.AppConfig
    fileSize,err:=config.Int64("file.size")
	if header.Size / 1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
	}

}