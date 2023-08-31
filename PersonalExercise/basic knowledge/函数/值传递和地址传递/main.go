package main

import "fmt"

func main(){

	var a int=10
	var b int=20
	test(&a,&b)
	fmt.Println(a)
	fmt.Println(b)

}
func test(a *int,b *int){
	var t int
	t=*a
	*a=*b
	*b=t
}

