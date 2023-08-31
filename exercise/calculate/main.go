package main

import (
	"fmt"

)
func main(){
	var 
	(
		n int = 10
		a int =4
	)
	fmt.Printf("两数之差为=%d\n",n-a)
	fmt.Printf("两数之余为=%d\n",n%a)
	fmt.Printf("两数之积为=%d\n",n*a)
	fmt.Printf("两数之和为=%d\n",n+a)
	var c float32=10.0/4
	fmt.Println(c)//如果要保留小数则需要浮点数运算

	//演示%的使用
	fmt.Println("10%3=",10%3)//1
	fmt.Println("-10%3=",-10%3)//-1
	fmt.Println("10%-3=",10%-3)//1
	fmt.Println("-10%-3=",-10%-3)//-1
	fmt.Printf("hello world")

	//i++
	/* var i int=1
	i++
	fmt.Println("i的值为",i) */

	var j int =99.
	j--
	fmt.Println("j的值为",j)

	var n1,n2,n3= "man",11,90.1
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	 n4,n5,n6:= "man",11,90.1
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
}
