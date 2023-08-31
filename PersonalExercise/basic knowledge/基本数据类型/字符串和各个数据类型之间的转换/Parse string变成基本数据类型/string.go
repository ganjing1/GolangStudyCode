package main

import (
	"fmt"
	"strconv"
)

func main(){

	//string -> float64
	var  s1 string="-1234.55"
	var s float64
	//ParseFLoat:返回字符串表示的浮点型值，接受正负号
	s,_=strconv.ParseFloat(s1,64)
	fmt.Printf("%T,%v",s,s)
	fmt.Println()


	//string -> int64
	var  s2 string="1234"
	var a int64
	//ParseInt:返回字符串表示的整数值，接受正负号
	a,_=strconv.ParseInt(s2,10,64)
//上面代码的意思是将s2 以十进制的形式接收，存储在int64 的a中
	fmt.Printf("%T,%v",a,a)

	//string -> bool
	var s3 string="true"
	var b bool
	//ParseBool只接受一个字符串
	b,_=strconv.ParseBool(s3)
	fmt.Printf("%T,%v",b,b)

}

