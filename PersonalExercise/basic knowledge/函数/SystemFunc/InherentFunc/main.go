package main

import "fmt"

func main(){
	//new分配内存，new函数的参数是一个类型而不是具体数值，
	//new函数返回值是对应类型的指针 num:*int 相当于C语言中的malloc
	num:=new(int)
	fmt.Printf("%T,%v,%v,%v",num,num,*num,&num)
}