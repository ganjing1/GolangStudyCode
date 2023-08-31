package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口：
func testReflect(i interface{}) { //空接口没有任何方法,所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口。
	//1.调用TypeOf函数，返回reflect.Type类型数据：
	reType := reflect.TypeOf(i)

	//2.调用ValueOf函数，返回reflect.Value类型数据：
	reValue := reflect.ValueOf(i)
	k1 := reType.Kind()
	k2 := reValue.Kind()
	fmt.Println(k1)
	fmt.Println(k2)
	i2 := reValue.Interface()
	n, flag := i2.(student)
	if flag == true {
		fmt.Printf("结构体的类型是:%T", n)
	}

}

type student struct {
	Name string
	Age  int
}

func main() {
	//对基本数据类型进行反射：
	//定义结构体类型：
	stu := student{
		Name: "ganjing",
		Age:  22,
	}
	testReflect(stu)
}
