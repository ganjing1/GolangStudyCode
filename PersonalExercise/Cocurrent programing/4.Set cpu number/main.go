package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu numbers:", cpuNum)
	//设置自己设置使用几个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
}
