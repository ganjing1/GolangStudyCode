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
	//看s是否是chinese型的变量，如果是赋给ch变量
	//flag判断是否转成功
	//1.第一种断言判断方法
	//if ch, flag := s.(Chinese); flag {
	//	//逗号前面的对应ch，逗号后面的对应flag
	//	ch.Dance()
	//} else {
	//	fmt.Println("other country's people  can't dance")
	//}
	/*
		其中，if语句中的逗号前面的ch对应转换后的Chinese类型变量，
		逗号后面的flag则表示转换是否成功。如果转换成功
		，则flag为true；否则，flag为false。1
	*/

	//2.第二种断言判断方法
	//ah, flag1 := s.(American)
	//if flag1 == true {
	//	ah.disco()
	//} else {
	//	fmt.Println("American people  can't dance")
	//}

	//用switch判断，判断类型,和c语言的是一样的
	switch s.(type) { //type是go中的一个关键字
	case Chinese:
		ch := s.(Chinese)
		ch.Dance()
	case American:
		am := s.(American)
		am.disco()
	}
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
func (person American) disco() {
	fmt.Println("disco")
}
func main() {
	c := Chinese{}
	greet(c)
	//a := American{}
	//greet(a)
}
