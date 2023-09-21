package main

import "fmt"

func main() {
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	/*
		遍历管道不能使用普通的for循环
		for i := 0; i < len(intChan2); i++ {

		}
	*/
	//如果channel管道没有关闭会产生死锁
	//入伙channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	/*
		1.遍历管道不能使用普通的for循环，因为管道的长度是不固定的，
		它会随着发送和接收而变化。普通的for循环需要一个固定的范围，
		否则会出现越界或漏读的情况。
		2.如果channel管道没有关闭会产生死锁，
		因为遍历管道使用的是range关键字，它会一直等待从管道中读取数据，
		直到管道被关闭。如果管道没有关闭，那么range就会陷入无限等待，
		导致程序无法继续执行。
		3.如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
		这是因为关闭的channel可以读取到零值和一个false标志，表示没有更多的数据了。
		range就会根据这个标志来判断是否结束循环。
	*/
}
