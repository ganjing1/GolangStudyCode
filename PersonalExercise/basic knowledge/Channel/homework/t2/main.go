package main

import "fmt"

func main() {
	var wchan chan<- int
	var rchan <-chan int
	wchan = make(chan int, 3)
	rchan = make(chan int, 3)
	wchan <- 10
	wchan <- 20
	wchan <- 30
	num1 := <-rchan
	fmt.Println(num1)
}
