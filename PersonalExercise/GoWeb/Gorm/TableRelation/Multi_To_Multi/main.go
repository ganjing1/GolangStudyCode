package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	SID    int `gorm:"primary_key"`
	Sno    int
	Name   string
	Sex    string
	Age    int
	Course []Course `gorm:"many2many:Student2Course"`
}

type Course struct {
	CId         int `gorm:"primary_key"`
	CName       string
	TeacherName string
	Room        string
}

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		panic(err)
	}
	db.Table("course").CreateTable(&Course{})
	db.Table("student").CreateTable(&Student{})
}
