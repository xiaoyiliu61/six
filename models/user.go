package models

import (
	"DataCertPlatform/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

//将用户的信息保存到数据库中
func (u User) AddUser()(int64,error){
    //脱敏
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)
	u.Password=hex.EncodeToString(passwordBytes)

	rs,err:=db_mysql.Db.Exec("insert into user(phone,password) values(?,?)",
		u.Phone,u.Password)
	if err != nil {
		return -1,err
	}
	//id值代表的是次数据操作影响的行数，id是一个整数int64类型
	id,err:=rs.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}

func (u User) QueryUser() (*User,error) {
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)
	u.Password=hex.EncodeToString(passwordBytes)
	row:=db_mysql.Db.QueryRow("select phone from user where phone =? and  password = ?",
		u.Phone,u.Password)
	err:=row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u,nil
}