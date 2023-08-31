package main

import "fmt"

func main(){
	//float 与 int之间的转换
	var a int64=11
	var a1 float64=float64(a)
	fmt.Println(a1)

	var b float64=11.1
	var b1 int64=int64(b)
	fmt.Println(b1)

	//bool 和 int
	//var c int64=1
	//var c1 bool=bool(c)
	/*数据类型和bool类型不能相互转换*/
	//println(c1)
	// int32 与int64
	var d int32=1
	var d1 int64=int64(d)
	fmt.Println(d1)
}
