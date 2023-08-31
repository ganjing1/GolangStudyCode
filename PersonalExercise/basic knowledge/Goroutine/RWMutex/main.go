package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var Lock sync.RWMutex //读写锁：

func read() {
	defer wg.Done()
	//如果只是读数据，那么这个锁不产生影响，但是读写同时发生的时候，就会有影响
	Lock.RLock()
	fmt.Println("开始读取")
	time.Sleep(time.Second)
	fmt.Println("读取成功!!")
	Lock.RUnlock()
}

func write() {
	defer wg.Done()
	Lock.Lock()
	fmt.Println("开始修改数据..")
	time.Sleep(time.Second * 2)
	fmt.Println("修改完成!!")
	Lock.Unlock()
}

func main() {
	wg.Add(6)
	//场景：读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()

}
