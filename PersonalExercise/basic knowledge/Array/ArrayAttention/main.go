package main

import "fmt"

func main(){
	var arr1=[3]int{3,6,9}
	test1(&arr1)
	//fmt.Printf("%T",arr1)//[3]int,数组的长度属于数组的类型的一部分
 	for i:=0;i<3;i++{
 		fmt.Println(arr1[i])
	}

}
func test1(arr1 *[3]int){
	(*arr1)[0]=7
}
