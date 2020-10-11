package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//连接mysql数据库
func Connect() {
	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//fmt.Println(dbDriver, dbUser, dbPassword)
	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db,_ := sql.Open(dbDriver, connUrl)
	/*if err != nil {
		panic("数据连接失败，请检查配置")
	}*/
	Db = db
	//fmt.Println(db)
}
//将用户信息保存到数据库中去的函数
/*func AddUser(u models.User)(int64,error) {
	//将密码进行hash计算，得到密码hash值，然后在存
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)
	u.Password=hex.EncodeToString(passwordBytes)
	//execute 执行, .exe
	result,err :=Db.Exec("insert into user(name ,birthday,address,password)"+
		"values(?,?,?,?)",u.Name,u.Birthday,u.Address,u.Password)
	if err != nil {
		return -1,err
	}
	row,err:=result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return row,nil
}
*/

