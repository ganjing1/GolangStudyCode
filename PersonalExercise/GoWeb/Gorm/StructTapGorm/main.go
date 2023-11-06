package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	StuID   int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"not null"`
	Age     int    `gorm:"unique_index"`
	Email   string `gorm:"unique"`
	Sex     string `gorm:"column:gender;size:10"` //将列明sex在数据库中以gender显示
	Desc    string `gorm:"-"`                     //不映射这个字段，不在数据库中体现，只在struct中体现
	Classno string `gorm:"type:int"`              //将string类型转换成int类型
}

func main() {
	//连接数据库：
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) //如果出错，后续代码没有必要执行，想让程序中断，panic来执行即可
	}
	//数据库资源释放：
	defer db.Close()
	//创建表：
	db.CreateTable(&Student{})

}
