package main

import "fmt"

func main() {
	//切片的扩容
	//切片要初始化后才能使用
	// var a[]int//此时没有申请内存
	/* for i:=0 ;i<10;i++{
		 a=append(a,i)
		 fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	 */
	 /*a=append(a, 1,2,3,4,5,6,7,8,9)
	 fmt.Println(a)
	 b:=[]int {11,12,13,14,33,45,55,66,99,88,999,88}
	 a=append(a, b...)
	 fmt.Println(a)*/
	 //切片删除元素
	 a:=[]string{"北京","shanghai","wuhan","xiamen","guangzhou","fuzhou"}
	 a=append(a[0:2],a[3:]...)
	 fmt.Println(a)
	 //append(a[：index],a[index+1:]...)
	
	}
	
