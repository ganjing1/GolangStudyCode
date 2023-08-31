package main

import "fmt"

func main(){
	//1
	var arr1[3] int=[3]int{3,6,9}
	fmt.Println(arr1)
	//2
	var arr2 =[3]int{1,4,7}
	fmt.Println(arr2)
	//3
	var arr3 =[...]int{4,5,6,7}
	fmt.Println(arr3)
	//4
	var arr4=[...]int{2:66,0:1,1:99}
	fmt.Println(arr4)
}
