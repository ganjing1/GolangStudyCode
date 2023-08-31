package main

import "fmt"



func main(){
	test()
	fmt.Println()
	test(1,2,3,4)
}

func test(args...int) {
	for i:=0;i<len(args);i++{
	fmt.Printf("%d ",args[i])
	}
}
