package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql的驱动
)

type User struct {
	Age  int
	Name string
}

func main() { //连接数据库：
	//Open传入两个参数：
	//第一个参数：指定你要连接的数据库
	//第二个参数：指的是数据库的设置信息：用户名:密码@tcp(ip:port)/数据库名字?charset=utf8&parseTime=True&loc=Local
	//charset=utf8设置字符集
	//parseTime=True为了处理time.Time
	//loc=Local 时区设置，与本地时区保持一致
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db03?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) ////如果出错，后续代码没有必要执行，想让程序中断，panic来执行即可

	}
	defer db.Close()
	//创建表
	////创建表：通常情况下，数据库中新建的表的名字是结构体名字的复数形式，例如结构体User，表名 users
	db.CreateTable(&User{})

	///Table方法可以指定你要创建的数据库的表名
	db.Table("user").CreateTable(&User{})

	//删除表：
	db.DropTable(&User{}) //通过&User{}来删除users表
	db.DropTable("user")  //通过"user"删除user表
	//判断表是否存在:
	flag := db.HasTable("user")
	fmt.Println(flag)
}
