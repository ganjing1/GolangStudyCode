package main

import "fmt"

type Hello interface {
	//首先再Hello结构中先定义一个SayHello方法
	SayHello()
}

// 定义American和 Chinese两个结构体
type Chinese struct {
	name string
}
type American struct {
	name string
}

//分别定义American和Chinese结构体对应的SayHello方法

func (person Chinese) SayHello() {
	fmt.Println("你好")
}
func (person American) SayHello() {
	fmt.Println("hello")
}

// 定义一个接口函数
func greet(h Hello) {
	h.SayHello()
}

// 自定义数据类型
type integrate int

func (i integrate) SayHello() {
	fmt.Println("integrate's hello", i) //拼接用逗号
}

func main() {
	//定义一个SayHello接口数组,里面存放American，chinese结构体变量
	var arr [3]Hello
	arr[0] = American{"Adrian"}
	arr[1] = Chinese{"小明"}
	arr[2] = Chinese{"小红"}
	fmt.Println(arr[0])
	fmt.Println(arr[1])

	fmt.Println(arr[2])

	var i integrate = 10
	var s1 Hello = i //接口等于自定义实践类型
	s1.SayHello()
	a := American{}
	c := Chinese{}
	greet(a)
	greet(c)
	var s Hello = c
	s.SayHello()
}
