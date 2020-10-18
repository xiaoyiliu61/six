package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName="login.html"
}

func (l *LoginController) Post()  {
	//1.解析客户端用户提交的登录数据
	var user models.User
	err:=l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录信息解析失败")
		return
	}
	//2.根据解析到的数据，执行数据库查询操作
    u,err:=user.QueryUser()

	//3.判断数据库查询结果
	if err != nil {
			fmt.Println(err.Error())
			l.Ctx.WriteString("抱歉，登录失败啊")
			return
	}
	//4.根据查询结果返回客户端相应的信息或页面跳转
	l.Data["Phone"]=u.Phone//动态数据
	l.TplName="home.html"//文件上传的界面
}