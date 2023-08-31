package main

import (
	"fmt"
	"time"
)

func main() {
	//定义int管道：
	intchan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intchan <- 10
	}()
	//string管道
	stringchan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 5)
		stringchan <- "hello world"
	}()

	fmt.Println(<-intchan)
	//fmt.Println(<-intchan)本身取数据就是阻塞的
	select {
	case v := <-intchan:
		fmt.Println("intchan", v)
	case v := <-stringchan:
		fmt.Println("stringchan", v)
	default:
		fmt.Println("防止select被阻塞")
	}

}
