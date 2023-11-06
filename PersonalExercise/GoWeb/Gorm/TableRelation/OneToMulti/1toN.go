package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Author struct {
	AID  int `gorm:"primary_key;AUTO_INCREMENT"` // 修改为int类型
	Name string
	Age  int
	Sex  string
	//关联
	As []Article `gorm:"ForeignKey:Auid;AssociationForeignKey:AID"` // 修改为Auid和Aid
}
type Article struct {
	ArId  int `gorm:"primary_key;AUTO_INCREMENT"` // 修改为int类型
	Title string
	//设置外键
	Auid int
}

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//db.CreateTable(&Author{})
	//db.CreateTable(&Article{})
	//author2 := Author{
	//	AID:  2,
	//	Name: "莫言",
	//	Age:  55,
	//	Sex:  "女",
	//	// 创建一个Article切片，包含四个Article对象，分别为live，dead，world，common
	//	As: []Article{
	//		{Title: "live"},
	//		{Title: "dead"},
	//		{Title: "world"},
	//		{Title: "common"},
	//	},
	//}

	// 使用Create方法插入数据，同时插入关联的Article数据
	//db.Debug().Create(&author2)

	//查
	var moyan Author
	db.Where("name=?", "莫言").First(&moyan)
	// 给莫言添加两篇新的文章
	//moyan.As = append(moyan.As, Article{Title: "blue"})
	//moyan.As = append(moyan.As, Article{Title: "yellow"})
	//db.Save(&moyan)

	//改
	//db.Model(&moyan).Where("name=?", "莫言").Update("age", 20)

	//改关联表的数据
	//var moyan1 Author
	//db.Preload("As").Find(&moyan1, "name=?", "莫言")
	//db.Model(&Article{}).Where(moyan1.As).Where("title=?", "black").Update("title", "pink")

	//删
	var moyan3 Author
	db.Preload("As").Find(&moyan3, "name=?", "莫言")
	db.Where("title=?", "pink").Delete(&moyan3.As)
}
