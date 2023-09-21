package main

import (
	"fmt"
	"sync"
)

func a(a string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%v\n", a)
	}

}
func b(b string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%v\n", b)
	}

}
func main() {
	var WG sync.WaitGroup
	WG.Add(2)
	/*
		当你传递一个 sync.WaitGroup 类型的变量时，如果使用值传递，函数会在内部创建一个 sync.WaitGroup 的副本。
		这样，当函数执行 Done() 方法时，
		它只会对副本进行操作，而不会对原始的 sync.WaitGroup 变量进行操作。
		这就导致了 WaitGroup 的计数器无法正确减少，最终可能会导致程序无法正常结束。
	*/
	go a("hello", &WG)
	go b("world", &WG)
	WG.Wait()
}
