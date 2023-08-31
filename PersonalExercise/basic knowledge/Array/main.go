package main

import (
	"fmt"
)

type Hello interface {
	sayHello()
}

func greet(s Hello) {
	s.sayHello()
	//断言：
	var ch Chinese = s.(Chinese) //看s是否是chinese型的变量，如果是赋给ch变量
	ch.Dance()
}

type Chinese struct {
	name string
}

type American struct {
	name string
}

func (person Chinese) sayHello() {
	fmt.Println("你好")

}
func (person Chinese) Dance() {
	fmt.Println("跳舞")
}

func (person American) sayHello() {
	fmt.Println("hello")

}

func (person American) Dance() {
	fmt.Println("damce")
}

func main() {
	c := American{}
	greet(c)

}
