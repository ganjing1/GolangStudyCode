package demo

import (
	"fmt"
)

type Student struct {
	Age  int
	Name string
}

func main() {
	//1.按照顺序赋值
	var s1 Student = Student{12, "ganjing"}
	fmt.Println(s1)

	//2.按照指定类型
	var s2 Student = Student{
		Name: "xiaoming",
		Age:  18,
	}
	fmt.Println(s2)

	//3.想要返回结构体的指针类型
	var s3 *Student = &Student{18, "zhansan"}
	fmt.Println(*s3)

	var s4 *Student = &Student{
		Name: "Wangwu",
		Age:  28,
	}
	fmt.Println(s4)

}
func Test() {
	fmt.Println("hello, this is PythonWork function!")
}

type student struct {
	name string
	age  int
}

func (s *Student) St(name string, age int) {
	s.Age = age
	s.Name = name
	fmt.Println("name:", s.Name)
	fmt.Println("name:", s.Age)
}

// NewStudent为构造函数，它返回一个指向 student 结构体的新实例的指针。
// 该函数接受两个参数，n string 和 a int，用于初始化结构体的字段。*student 是指向 student 结构体的指针
func NewStudent(n string, a int) *student {
	return &student{n, a}
}
