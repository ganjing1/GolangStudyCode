package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 100; i++ {
		fmt.Println("hello golang+" + strconv.Itoa(i))
		//阻塞一秒:
		time.Sleep((time.Second))
	}
}

func main() { //主线程
	//test()//只是调用函数
	go test() //开启协程，交替执行
	for i := 1; i <= 5; i++ {
		fmt.Println("java+" + strconv.Itoa(i))
		//阻塞一秒:
		time.Sleep((time.Second))
	}
}
