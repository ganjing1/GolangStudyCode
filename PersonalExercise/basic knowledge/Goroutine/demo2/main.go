package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		//启动协程，使用匿名函数
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
