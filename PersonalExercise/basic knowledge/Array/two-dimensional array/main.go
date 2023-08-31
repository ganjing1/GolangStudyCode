package main

import "fmt"

func main(){
	var arr[3][3]int=[3][3]int{{1,4,7},{2,5,8},{3,6,9}}
	for key,value:=range arr{
		for k,v:=range value{
			fmt.Printf("arr[%v][%v]=%v ",key,k,v)
	}
	fmt.Println()
	}
}
