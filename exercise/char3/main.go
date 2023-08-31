package main

import (
	"fmt"
	"strconv"
)
func main() {

	//string 转化成基本数据类型
	var str string= "true"
	var b bool
	//b, _ =strconv.ParseBool(str)
	//说明
	//1.strconv.ParseBool(str)函数会返回两个值（value bool，err error）
	//2.应为我只想获取value bool，不想获取err 所以用“_”忽略
	b,_ =strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n",b,b)

	 var str2 string="1234590"
	 var n1 int64  
	 var n2 int 
	 n1,_= strconv.ParseInt(str2,10,64)
	 n2=int(n1) 
	 fmt.Printf("n1 type %T n1=%v\n",n1,n1)
	 fmt.Printf("n2 type %T n2=%v\n",n2,n2)



	var str3 string="123.5154165"
	var f1 float64
	f1,_=strconv.ParseFloat(str3,64)
	fmt.Printf("f1 type is %T f1=%v",f1,f1)


	//如果将字符串(非数字元素)转换成整型会直接转换成0
	var str6 string="hello"
	var k int64
	k,_=strconv.ParseInt(str6,10,64)
	fmt.Printf("k type is %T k=%v",k,k)

}


	