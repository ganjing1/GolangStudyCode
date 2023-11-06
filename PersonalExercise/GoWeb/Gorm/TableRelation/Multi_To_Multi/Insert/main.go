package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	SId  int `gorm:"primary_key"`
	SNo  int
	Name string
	Sex  string
	Age  int
	//创建一个Course切片，名称也叫Course，表示一个学生的多门课程
	Course []Course `gorm:"many2many:Student2Course"`

	/*
		多对多的关系必要要创建一个关联表：
		例如在这里我们STudent和course的关联表是Student2Course
	*/
}
type Course struct {
	CId         int `gorm:"primary_key"`
	CName       string
	TeacherName string
	Room        string
}

func main() {
	db, err := gorm.Open("mysql", "root:hsp@tcp(localhost:3306)/hsp_db02?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//info := Student{
	//
	//	SNo:  2,
	//	Name: "zhansan",
	//	Sex:  "man",
	//	Age:  22,
	//	Course: []Course{
	//		{
	//			CName:       "Java",
	//			TeacherName: "jack",
	//			Room:        "j-408",
	//		},
	//	},
	//}
	//db.Create(&info)

	var stu Student
	db.Preload("Course").Where("name=?", "甘进").First(&stu)
	/*
		从数据库中查询名字为"甘进"的学生记录，并预加载学生的课程信息。
		具体来说，db.Preload("Course")表示在查询学生记录时，同时预加载学生的课程信息。
		Where("name=?", "甘进")表示查询条件，即只查询名字为"甘进"的学生记录。
		First(&stu)表示将查询结果的第一条记录保存到stu变量中。
	*/
	fmt.Println(stu.Course)
	var stu2 Student
	db.Preload("Course").Where("name=?", "zhansan").First(&stu2)
	fmt.Println(stu2)

	//给zhansan再添加一些课程，先查询再操作
	var stu1 Student
	db.Where("name=?", "zhansan").First(&stu1)
	var newCourses []Course
	db.Create(&[]Course{
		{
			CName:       "net",
			TeacherName: "marry",
			Room:        "s-111",
		},
		{
			CName:       "linux",
			TeacherName: "robot",
			Room:        "s-222",
		},
	}).Find(&newCourses)

	//使用Association方法来给zhansan添加新的课程信息，并且将它们同步到course表中。
	db.Model(&stu1).Association("Course").Append(newCourses)

	var info3 Student
	db.Preload("Course").Where("name=?", "zhansan").First(&info3)
	fmt.Println(info3)
}
