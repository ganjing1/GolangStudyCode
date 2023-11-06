package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type User struct {
	Id      int
	Keyword string `gorm:"column:keywords"`
	City    string
}

func (User) TableName() string {
	return "user"
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func read(client *gorm.DB, city string) *User {
	var users []User
	client.Where("city=?", city).Find(&users)
	if len(users) > 0 {
		return &users[0]
	} else {
		return nil
	}
}
func main() {
	dataSourceName := "tester:123456@tcp(127.0.0.1:3306)/PythonWork?charset=utf8&parseTime=True"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil) //第二个参数是一些选项配置，一般不动
	/*
		数据库的用户名为tester，密码为123456
		@tcp代表协议，127.0.0.1是一个特殊的IP地址，用于访问本地主机
		:3306表示的是访问mysql的端口
		/test表示操作的是test数据库
		?charset=utf8&parseTime=True
		如果配置了parseTime=true，MySQL中的DATE、DATETIME等时间类型字段将自动转换为golang中的time.Time类型。 类似的0000-00-00 00:00:00 ，会被转为time.Time的零值。
		否则，如果没有配置或配置了parseTime=false, 只会转为 []byte / string 。
	*/
	checkError(err)
	user := read(client, "beijing")
	if user != nil {
		fmt.Printf("%+v\n", *user)
	} else {
		fmt.Println("无结果")
	}
}
