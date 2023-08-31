package main

import (
	"fmt"
	"sync"
)

var totalNum int
var wg sync.WaitGroup
var Lock sync.Mutex //互斥锁：

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Lock.Lock() //加锁
		totalNum += 1
		Lock.Unlock() //解锁
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		//加锁
		Lock.Lock()
		totalNum -= 1
		Lock.Unlock() //解锁
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(totalNum)
}
