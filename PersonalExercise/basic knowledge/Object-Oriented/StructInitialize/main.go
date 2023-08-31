package main

import "fmt"


type student struct {
	name string
	age int
}

func main(){
	//1
	var  a student=student{
		name:"zhangsan",
		age:18,
	}
	fmt.Println(a)
	//2
	var b student=student{"ganjing",
		18,
	}
	fmt.Println(b)

	//3
	var c *student=&student{name: "lisi",
		age: 19}
	fmt.Println(*c)

	//4
	var d *student=&student{name: "lisi",
		age: 19}
	fmt.Println(*d)



}
