package main

import "fmt"

func  main (){
	//var f int
	var i float64
	_,i=Add(10,20)//_的作用是忽略第二个返回值
	fmt.Println(i)
}

func  Add(a int,b int)(int,float64){
	var sum int=0;
	var sum1 float64=float64(a)+float64(b)+0.1
	sum+=a+b
	return sum,sum1
}


