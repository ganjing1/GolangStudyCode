package main

import "fmt"

func main() {
	// var price float32 = 89.12
	// fmt.Println("price=",price)
	//.112相当于0.112
	num6:=5.12
	num7:=.112
	fmt.Println(num6)
	fmt.Println(num7)
	//科学计数法形式
	num8:=5.1234e2//5.1234*10的二次方
	num9:=5.1234E2//5.1234*10的二次方
	num10:=5.1234e-2//5.1234/10的二次方
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)
}