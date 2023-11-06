package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义结构体：
type User struct {
	Age  int
	Name string
}
type UserInfo struct {
	Age  int
	Name string
}
type DBUserInfo struct {
	Age  int
	Name string
}
type MyUser struct {
	Age  int
	Name string
}

func (MyUser) TableName() string { //自定义表名
	return "test_my_user"
}
func (DBUserInfo) TableName() string { //自定义表名
	return "DB_username"
}
func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db03?charset=utf8&parseTime=True&loc=Local&parseTime=True")
	if err != nil {
		panic(err)
	}
	//创建表：通常情况下，数据库中新建的标的名字是结构体名字的复数形式，例如结构体User，表名 users
	db.CreateTable(&User{})
	db.CreateTable(&UserInfo{})
	db.CreateTable(&DBUserInfo{})
	db.CreateTable(&MyUser{})

}
