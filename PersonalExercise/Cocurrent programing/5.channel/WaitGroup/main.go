package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Printf("协程 %d完成\n,", i)
		}(i)
		wg.Wait() //Wait()方法会解除阻塞，继续执行后续的操作
	}
	fmt.Println("所有协程完成")
}
