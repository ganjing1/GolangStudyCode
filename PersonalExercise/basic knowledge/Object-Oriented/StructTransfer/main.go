package main

import (
	"fmt"

)

type student struct {
	age int
}
type teacher struct {
	age int
}


type Student struct {
	age int
}

type Stu Student
func main(){
	var a Student
	var b teacher
	a=Student(b)
	fmt.Println(a)
	//结构体转换
	var s1 Student=Student{19}
	var s2 Stu=Stu{19}
	s1=Student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
