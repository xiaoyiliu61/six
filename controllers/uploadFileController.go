package controllers

import (
	"DataCertPlatform/models"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)


//该控制器结构体用于处理文件上传的功能

type UploadFileController struct {
	beego.Controller
}



//该post方法用于处理客户在客户端提交的文件
func (u UploadFileController) Post() {
	phone:=u.Ctx.Request.PostFormValue("phone")
	title:=u.Ctx.Request.PostFormValue("upload_title")
    fmt.Println("电子数据标签",title)
	//用户上传的文件
	file,header,err:=u.GetFile("xiaoyiliu")
	if err != nil {//解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	defer file.Close()//延迟执行 空指针错误
    //使用io包
    saveFilePath:="static/upload/"+header.Filename
	saveFile,err:=os.OpenFile(saveFilePath,os.O_CREATE|os.O_RDWR,777)
	if err != nil {
		u.Ctx.WriteString("抱歉,电子数据认证失败，请重试！")
		return
	}
    _,err=io.Copy(saveFile,file)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重新尝试！！！！！")
		return
	}

	//2.计算文件的SHA256值
	hash256:=sha256.New()
	fileBytes,_:=ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes:=hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))

	//先查询用户id
	 user1,err:=models.User{Phone:phone}.QueryUserByPhone()
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重新再试！——————@2——————")
		return
	}
	//把上传的文件作为记录保存到数据库当中
	//1.计算md5值
	md5Hash:=md5.New()
	fileMd5Bytes,err:=ioutil.ReadAll(saveFile)
	md5Hash.Write(fileMd5Bytes)
	bytes:=md5Hash.Sum(nil)
	record:=models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  hex.EncodeToString(bytes),
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}
	_,err=record.SaveRecord()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重试")
		return
	}
	//上传文件保存到数据库中成功
     records,err:=models.QueryRecordsByUserId(user1.Id)
	if err != nil {
		u.Ctx.WriteString("抱歉尼玛呢，获取电子数据列表失败，请重新尝试！")
		return
	}
     u.Data["Records"]=records
     u.TplName="list_record.html"
	//把上传的文件
	/*u.Ctx.WriteString("恭喜，已接收到上传文件！")*/
}

/*
 该post方法用于处理用户在客户端提交的认证文件
*/
func (u *UploadFileController) Post1() {
	//1.解析数据上传的文件
	//用户上传的自定义的标题
	title:=u.Ctx.Request.PostFormValue("upload_title")

    //用户上传的文件
    file,header,err:=u.GetFile("xiaoyiliu")
    if err != nil {
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标题：",title)
	//获取上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
    //eg:支持jpg，png类型，不支持jpeg。gif类型
	fileNameSlice:=strings.Split(header.Filename,".")
    fileType:=fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":",fileType)
	isJpg := strings.HasSuffix(header.Filename,".jpg")
	isPng := strings.HasSuffix(header.Filename,".png")
	if !isJpg && !isPng{
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}
	//文件大小 200kb
    config:=beego.AppConfig
    fileSize,err:=config.Int64("file_size")
	if header.Size / 1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
	}
	fmt.Println("文件的大小：",header.Size)

    //perm：permission   权限
    //权限的组成 ; a+b+c
      //a :文件所有者对文件的操作权限：读4、写2、执行1
      //b: 文件所有者所在组的用户的操作权限，读4、写2、执行1
      //c: 其他用户的操作权限，读4、写2、执行1
    saveDir:="static/upload"
    //1.先尝试打开文件夹
    _,err=os.Open(saveDir)
 //   os.OpenFile("文件名",os.,)

	if err != nil {
		//2.自己动手，打开文件夹
		err=os.Mkdir(saveDir,777)
		if err != nil {
			u.Ctx.WriteString("抱歉，文件认证失败，请重试")
			return
		}
	}

    saveName:="static/upload"+header.Filename
    fmt.Println("要保存的文件名",saveName)

    err = u.SaveToFile("xiaoyiliu",saveName)
	if err != nil {
		u.Ctx.WriteString("抱歉文件认证失败，请重试")
		return
	}

	// u.SaveToFile()
    fmt.Println("上传的文件:",file)
	u.Ctx.WriteString("已获取到上传文件。")

}