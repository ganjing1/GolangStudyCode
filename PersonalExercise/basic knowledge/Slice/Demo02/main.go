package main

import "fmt"

func main(){
	//定义切片 make的三个参数；1.切片类型 2.切片长度 3.切片容量
	a:=make([]int,5,20)
	fmt.Println(a)
	b:=make([]int,6,6)
	b[0]=0
	b[1]=1
	b[2]=2
	b[3]=3
	b[4]=4
	b[5]=5
	for i,v:=range b{
		fmt.Printf("%v,%v\n",i,v)
	}

	//动态切片
	var intarr [6]int=[6]int{1,2,3,4,5,6}
	var s[]int=intarr[1:4]
	fmt.Println(len(s))
	s=append(s,88,55)
	fmt.Println(s)

	slice3:=[]int{99,44}
	s=append(s,slice3...)
	fmt.Println(s)
}