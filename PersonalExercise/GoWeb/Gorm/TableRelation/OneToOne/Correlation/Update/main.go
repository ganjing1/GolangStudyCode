package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
关联查询是先查询后更新
*/
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
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?loc=Local&parseTime=True&charset=utf8")
	if err != nil {
		panic(err)
	}
	//info := St2{
	//	Address: "shenzhen",
	//	S1:      St1{Name: "lili"},
	//}
	//db.Create(&info)

	//先查询
	var check1 St2
	db.Preload("S1").Find(&check1, "Address=?", "beijing") //将地址为北京和及其关联字段都输出
	fmt.Println(check1)
	//后跟新
	db.Model(&check1.S1).Update("name", "ls")
	fmt.Println(check1)
}
