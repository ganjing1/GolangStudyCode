package main

import "fmt"

func main() {
	var intChan chan int

	intChan = make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan <- i
	}
	close(intChan)
	//在遍历前，如果没有关闭管道，就会出现deadlock的错误
	//所以我们再遍历前要进行管道的关闭

	//遍历for-range
	for v := range intChan {
		fmt.Println("value=", v)
	}

}
