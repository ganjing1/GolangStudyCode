package main

import "github.com/jinzhu/gorm"

func main() {
	gorm.Open("mysql", "root:hsp&tcp(localhost:3306)/hsp_db03?charset=utf8&loc=Local")
}
