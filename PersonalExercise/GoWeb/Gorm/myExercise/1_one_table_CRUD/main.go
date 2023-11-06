package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	ID   int `gorm:"AUTO_INCREMENT"`
	Name string
	Age  int `gorm:"AUTO_INCREMENT"`
	Sex  string
}

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/gorm?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		panic(err)
	}
	db.CreateTable(&Customer{})

	c := Customer{
		Name: "zhangsan",
		Age:  20,
		Sex:  "男"}
	db.Create(&c)

}
