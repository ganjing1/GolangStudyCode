package main

import "fmt"

func main() {

	n:=a(10,20)
	fmt.Println(n)
}
func a (num1 int,num2 int) int{
	//函数的语句被defer修饰时先不执行，而是把defer后的语句压入一个栈中
	//然后继续执行函数后面的语句
	defer fmt.Println(num1)
	defer fmt.Println(num2)
	//栈的特点：先进后出
	//在函数执行完毕以后，从栈中取出语句开始执行，按照先进后出的规则执行语句
	fmt.Println("hello")
	//
	return num1+num2
}