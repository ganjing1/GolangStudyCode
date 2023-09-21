package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 3
	intChan <- 2
	close(intChan)
	//intChan <- 1管道关闭后不能写入但是可以读取
	a := <-intChan
	fmt.Println(a)
}
