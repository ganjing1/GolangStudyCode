package main

import "fmt"

func main() {
/* 	var c1 byte = 'a'
	var c2 byte = '0' //字符0
	//当我们直接输出byte值，就是输出了对应的字符的asci码的值
	//'a'==>
	fmt.Println("c1=",c1)
	fmt.Println("c2=",c2)
	
	//var a byte='北'//overflow溢出
	var c3  int='北'//asci码值大于255时用int
	//fmt.Printf("a=%c",a)
	fmt.Printf("c3=%c c3对应码值=%d\n",c3,c3)
	//可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数对应的Unicode字符
	var c4 int =22269//22269->'国'120->'x'
	fmt.Printf("c4=%c\n",c4)
	//字符类型是可以进行运算的，相当于一个整数,运算时是安装码值进行的
	var n1=10+'a'//10+97=107
	fmt.Println("n1=",n1)
	//golang基本数据类型转换
		var i int32 =100
		//希望将i转换成float
		var n2 float32 = float32(i)
		var n3 int8=int8(i)
		fmt.Printf("i=%v n1=%v n2=%v\n",i,n2,n3)
		//被转换的是变量储存的数据（值），变量本身的数据类型并没有变化
		fmt.Printf("i type is %T\n",i)//int32
		// 在转换中，比如将int64转成int8【-128---127】，编译时不会报错，
		//只是转换的结果是按溢出处理，和我们希望的结果不一样
		var num int64= 99999
		var num2 int8= int8(num)//
		fmt.Println("num2=",num2)

 */


		//课堂练习
		var n1 int32 =12
		var n2 int64
		var n3 int8
		n2 = int64(n1) + 20 //int32--->int64 错误
		n3 = int8(n1) + 20//int32--->int8错误 
		fmt.Println("n2=",n2,"n3=",n3)

 

}