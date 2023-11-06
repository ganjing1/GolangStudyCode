package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Sta struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string
	//指定外键：
	IID int
}
type Stb struct {
	InfoID  int `gorm:"primary_key;AUTO_INCREMENT"`
	Address string
	//关联关系
	Sta Sta `gorm:"ForeignKey:IID;AssociationForeignKey:InfoID"`
}

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) //如果出错，后续代码没有必要执行，想让程序中断，panic来执行即可
	}
	defer db.Close()
	db.CreateTable(&Sta{})
	db.CreateTable(&Stb{})
	userinfo := Stb{
		Address: "wuhan",
		Sta: Sta{
			Name: "cheerio",
		},
	}
	db.Create(&userinfo)
	var info1 Stb
	//如果只是执行下面这步操作，那么关联的User信息是查询不到的：
	//db.First(&userinfo, "info_id = ?", 1)
	//fmt.Println(userinfo) //{1 /upload/1.jpg 北京海淀区 124234@126.com {0 0  0}}
	//如果想要查询到User相关内容，必须执行如下操作：
	//Model参数：要查询的表数据，Association参数：关联到的具体的模型：模型名字User（字段名字）
	//Find参数：查询的数据要放在什么字段中&userinfo.User
	db.Model(&info1).Association("Sta").Find(&info1.Sta)
	fmt.Println(info1) //{1 /upload/1.jpg 北京海淀区 124234@126.com {1 19 丽丽 1}}
}
