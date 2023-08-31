package main

import "fmt"

func main(){
	//int64 -> string
	var a int64=888
	var s1 string=fmt.Sprintf("%d",a)
	fmt.Println(s1)

	//float64 -> string
	var b float64=888.5
	var s2 string=fmt.Sprintf("%f",b)
	fmt.Println(s2)

	//bool -> string
	var c bool=true
	var s3 string=fmt.Sprintf("%T",c)
	fmt.Println(s3)

	//byte -> string
	var d byte='d'
	var s4 string=fmt.Sprintf("%c",d)
	fmt.Println(s4)

}

