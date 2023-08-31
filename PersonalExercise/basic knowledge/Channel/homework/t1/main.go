package main

import (
	"fmt"
	"sync"
	"time"
)

var intChan chan int
var wg sync.WaitGroup

func writeDate(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入的数据为：", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}
func readDate(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("读取的value=", v)
		time.Sleep(time.Second)
	}
}
func main() {
	intChan = make(chan int, 50)
	wg.Add(2)
	go writeDate(intChan)
	go readDate(intChan)
	//主线程一直在阻塞，什么时候wg减为0了，就停止
	wg.Wait()
}
