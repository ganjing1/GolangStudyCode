package main

import "fmt"

func main(){
	var scores[5]int
	for i:=0;i<5;i++{
		fmt.Scanln(&scores[i])
	}

	//range 遍历
	for key,value:=range scores{
		fmt.Printf("%d,%d\n",key,value)
	}
}