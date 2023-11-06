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
【1】Association方式查询缺点：先First查询，再Association查询，费劲
【2】Preload方式查询：
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

	var info2 St2
	//Preload:预加载关联数据
	db.Preload("S1").Find(&info2, "Address = ?", "wuhan")
	//将查询到Address=wuhan的信息和S1关联的数据都放入info2中
	fmt.Println(info2)
	/* db.Preload 函数来预加载关联数据 S1，
	并使用 db.Find 函数来根据 Address 字段的值为 “wuhan” 来查询数据，
	并将结果赋值给info2。然后您使用 fmt.Println 函数来打印 info2 的值*/

}
