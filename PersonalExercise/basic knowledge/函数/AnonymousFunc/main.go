package main

import "fmt"

func main(){
	//定义匿名函数：定义的同时调用
	result:=func (num1 int,num2 int)int{
		return num1+num2 }(10,20)
	fmt.Println(result)
	sub:=func (a int,b int) int{return a-b}
	fmt.Println(sub(29,30))
}

