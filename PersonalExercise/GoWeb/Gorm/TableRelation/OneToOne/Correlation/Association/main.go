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

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.CreateTable(&St1{})
	db.CreateTable(&St2{})
	info := St2{
		Address: "wuhan",
		S1:      St1{Name: "ganjin"},
	}
	db.Create(&info)

	var info2 St2
	db.First(&info2, "Address=?", "wuhan") //这一句必须要，因为要关联
	db.Model(&info2).Association("S1").Find(&info2.S1)
	fmt.Println(info2)

}
