package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
//该方法用于处理用户注册的逻辑
func (r *RegisterController) Post() {
   //1.解析用户端提交的请求数据
	var user models.User
	err:=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据请求失败，请重试")
		return
	}

   //2.将解析的数据保存到数据库中
   _,err=user.AddUser()
	if err != nil {
		r.Ctx.WriteString("抱歉")
	}
    /*row,err:=db_mysql.AddUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户信息注册失败")
		r.TplName="err.html"
		return
	}
	fmt.Println(row)
    md5Hash:=md5.New()
    md5Hash.Write([]byte(user.Password))
    user.Password=hex.EncodeToString(md5Hash.Sum(nil))*/
	//3.将处理结果返回到客户端浏览器
	r.TplName="login.html"


}
