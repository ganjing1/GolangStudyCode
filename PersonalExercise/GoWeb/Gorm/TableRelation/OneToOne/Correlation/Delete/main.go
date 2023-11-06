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
	S2ID    int `gorm:"primary_key;AUTO_INCREMENT"` //S2ID和S1ID要保持一致
	Address string
	//关联
	S1 St1 `gorm:"Foreignkey:FK;AssociationForeignKey:S2ID"`
}

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//db.CreateTable(&St1{})
	//db.CreateTable(&St2{})
	//info := St2{
	//	S2ID:    0,
	//	Address: "wuhan",
	//	S1: St1{
	//		Name: "ganjin",
	//	},
	//}
	//db.Create(&info)
	//a := St2{
	//	S2ID:    0,
	//	Address: "shengzheng",
	//	S1: St1{
	//		Name: "wangwu",
	//	},
	//}
	//
	//db.Create(&a)

	//先查询
	var info2 St2
	//Preload:预加载关联数据
	//db.Preload("S1").Find(&info2, "s2_id= ?", 4)
	//将查询到Address=wuhan的信息和S1关联的数据都放入info2中
	fmt.Println(info2)
	//删除借助info2模型删除S1记录
	db.Model(&St1{}).Where("name = ?", "wangwu").Update("s1_id", 19)
	db.Model(&St1{}).Where("name = ?", "ganjin").Update("s1_id", 18)

}
