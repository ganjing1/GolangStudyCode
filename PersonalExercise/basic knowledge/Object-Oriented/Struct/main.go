package main

import (
	A "Struct/demo" //将路径命名为A
	"fmt"
)

type student struct {
	name  string
	age   int
	score float64
}

func main() {

	var st student
	st.name = "cheerio"
	st.score = 12.3
	st.age = 22
	fmt.Println(st)

	//2
	var s student = student{"ganjin", 18, 15.5}
	fmt.Println(s)

	//3
	var t *student = new(student)
	t.name = "daad"
	t.age = 20 //相当于(*t).age,go编译器底层自动转换了
	fmt.Println(*t)

	//4
	var n *student = &student{
		"hello",
		18,
		18.8,
	}
	fmt.Println(*n)

	b := A.NewStudent("zhangsan", 20)
	fmt.Println(*b)

	c := A.Student{} //创建一个Student类型的变量c，这里的Student是结构体指针类型
	c.St("张三", 1)
}
