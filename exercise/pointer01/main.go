package main

import "fmt"

func main() {
	// //基本数据类型在内存中的布局
	// var i int = 10
	// fmt.Println("i的地址",&i)
	// //下面的var ptr *int =&i
	// //1.ptr 是一个指针变量
	// //2.ptr 的类型*int
	// //3.ptr 本身的值&i
	// var ptr *int=&i
	// fmt.Printf("ptr=%v\n",ptr)
	// fmt.Printf("ptr的地址=%v\n",&ptr)//这里是ptr的地址
	// fmt.Printf("ptr指向的值=%v\n",*ptr)//这里是ptr指向的值 
	


	//作业
	var num int =9
	var ptr1 *int
	ptr1=&num
	fmt.Printf("num=%v\n",&num)
	*ptr1=10//相当于把之前的9修改为10
	fmt.Println("a=\n",num)

// 	var a int=300
// 	var ptr *int=a//错误不能把值给指针变量要把地址给指针变量修改为&a
     

    // var a int=30
	// var ptr *float=&a//给定的是整型指针不能同浮点型指针，float改为int
    
}