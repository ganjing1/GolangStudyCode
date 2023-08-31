package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//启动五个协程
	for i := 1; i <= 5; i++ {
		wg.Add(1) //协程开始的时候+1操作
		go func(n int) {
			fmt.Println(n)
			wg.Done() //协程结束的时候-1操作
		}(i)
	}
	//主线程一直在阻塞，什么时候wg减为0了，就停止
	wg.Wait()
}
