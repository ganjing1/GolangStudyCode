package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MyUser2 struct {
	//增加一个匿名字段：
	gorm.Model
	Age  int
	Name string
}

func main() {
	//连接数据库：
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db03?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) //如果出错，后续代码没有必要执行，想让程序中断，panic来执行即可
	}
	//数据库资源释放：
	defer db.Close()
	//创建表：
	db.CreateTable(&MyUser2{})

}
