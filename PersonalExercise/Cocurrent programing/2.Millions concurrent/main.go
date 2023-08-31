package main

import (
	"fmt"
	"runtime"
)

func main() {
	//for i := 1; i <= 10000000; i++ {
	//	go runTimes(1)
	//}
	runtime.GOMAXPROCS(16) //设置cpu最大核心数
	fmt.Println(runtime.NumCPU())
}

var num int = 1

func runTimes(times int) int {
	for i := 1; i <= times; i++ {
		fmt.Println("runTimes", i, "hello world", times-i)
		fmt.Println("num:", num)
	}
	num++
	return times
}
