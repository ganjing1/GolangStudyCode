package main

import (
	"fmt"
	"reflect"
)

// 定义一个结构体：
type Student struct {
	Name string
	Age  int
}

// 给结构体绑定方法：
func (s Student) CPrint() {
	fmt.Println("调用了Print()方法")
	fmt.Println("学生的名字是：", s.Name)
}
func (s Student) AGetSum(n1, n2 int) int {
	fmt.Println("调用了AGetSum方法")
	return n1 + n2
}
func (s Student) BSet(name string, age int) {
	s.Name = name
	s.Age = age
}

// 定义函数操作结构体进行反射操作：
func TestStudentStruct(a interface{}) {
	//a转成reflect.Value类型：
	val := reflect.ValueOf(a)
	fmt.Println(val)
	n := val.Elem().NumField()
	fmt.Println(n)
	//修改字段的值：
	val.Elem().Field(0).SetString("张三")
}
func main() {
	//定义结构体具体的实例：
	s := Student{
		Name: "丽丽",
		Age:  18,
	}
	//调用TestStudentStruct：
	TestStudentStruct(&s)
	fmt.Println(s)
}
