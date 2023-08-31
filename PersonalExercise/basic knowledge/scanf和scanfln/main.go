package main

import "fmt"

func main(){
	//var age int
	//fmt.Println("输入学生的年龄：")
	//fmt.Scanln(&age)
	//
	//var name string
	//fmt.Println("input name of student:")
	//fmt.Scanln(&name)
	//
	//var score float32
	//fmt.Println("input score of student:")
	//fmt.Scanln(&score)
	//
	//var isVIP bool
	//fmt.Println("请输入true或者false:")
	//fmt.Scanln(&isVIP)
	//
	//fmt.Printf("%v,%v,%v,%v",age,name,score,isVIP)


	//Scanf
	fmt.Println("请输入学生的年龄，姓名，成绩，是否是VIP,使用空格分隔：")
	var age int
	var score float64
	var name string
	var VIP bool
	fmt.Scanf("%d %s %f %t",&age,&name,&score,&VIP)
	fmt.Printf("%v ,%v,%v,%v",age,score,name,VIP)

}
