package main

import "fmt"

//函数名：getSum 参数为空
//getSum函数返回值为一个函数，这个函数的参数是一个int型的参数，
//返回值也是int型
func getSum() func(int)int{
 var sum int =0
 return func(num int) int{
 	sum=sum+num
 	return sum
 }
}

func main(){
	f:=getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

