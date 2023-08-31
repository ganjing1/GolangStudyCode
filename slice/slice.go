package main

import (
	"fmt"
	"unsafe"
)

//切片
func main(){
	//基于数组得到切片
	a:=[5]int {55,56,57,58,59}
	b:=a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	//切片再次切片
	c:=b[:]
	fmt.Println(c)
	fmt.Printf("%T\n",c)
	//make函数构造切片
	d:=make([]int ,5,10)
	fmt.Println(d)
	fmt.Printf("%T\n",d)
	//通过len()函数获取切片刀的长度
	fmt.Println(len(d))
	//通过cap()函数获取切片的容量
	fmt.Println(cap(d))
	fmt.Println(unsafe.Sizeof(a))
	var n int 
	fmt.Scanf("%d",&n)
	fmt.Println(n)

}