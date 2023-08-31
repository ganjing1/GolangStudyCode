package main

import "fmt"

func main(){
	var intarr [6]int64=[6]int64{1,2,3,4,5,6}
	//
	 a:=intarr[1:3]//不包含3索引对应的值
	fmt.Println(intarr)
	fmt.Println(a)
	fmt.Println(len(a))
	 //容量=数组长度-切片开始的位置
	fmt.Println(cap(a))//容量capacity，

	fmt.Printf("原数组下标为1位置的地址:%p",&intarr[1])//0xc00000e428
	fmt.Printf("切片下标为1位置的地址:%p",&a[0])//0xc00000e428
	a[1]=16
	fmt.Println(intarr)

}
