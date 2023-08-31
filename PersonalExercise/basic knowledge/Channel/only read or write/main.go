package main

import "fmt"

func main() {
	//var intChan1 chan int	//可读可写

	//声明为只读:
	var intChan2 chan<- int //管道具备<- 只写性质
	intChan2 = make(chan int, 3)
	intChan2 <- 10
	intChan2 <- 20
	intChan2 <- 30
	//num1 := <-intChan2//报错
	fmt.Printf("%v", intChan2)
	//var intChan2 chan int //管道具备<- 只写性质

	//声明只读
	var intChan3 <-chan int
	var num1 int
	if intChan3 != nil {
		num1 = <-intChan3
	}
	fmt.Println(num1)

}
