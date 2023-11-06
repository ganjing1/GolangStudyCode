package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type St1 struct {
	S1ID int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	//外键
	FK int
}
type St2 struct {
	S2ID    int `gorm:"primary_key;AUTO_INCREMENT"`
	Address string
	//关联
	S1 St1 `gorm:"Foreignkey:FK;AssociationForeignKey:S2ID"`
}

/*
 */
func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//db.CreateTable(&St1{})
	//db.CreateTable(&St2{})
	//info := St2{
	//	Address: "wuhan",
	//	S1:      St1{Name: "ganjin"},
	//}
	//db.Create(&info)

	//关联查询操作：（关联关系在UserInfo表中，所以从UserInfo入手）
	var userinfo St2
	db.First(&userinfo, "Address=?", "wuhan") //只打印主表
	fmt.Println(userinfo)

	var userinfo2 St1
	db.Model(&userinfo).Related(&userinfo2, "S1") //只查询关联的数据
	fmt.Println(userinfo2)

}
