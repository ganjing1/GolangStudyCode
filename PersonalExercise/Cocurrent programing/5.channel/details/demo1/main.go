package main

import "fmt"

func main() {
	//双向管道
	//var doubleChan chan int

	//只写int
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 3
	fmt.Println(chan2)

	//只读int
	//var chan3<-chan int
	//chan2=make( chan int,3)

}
