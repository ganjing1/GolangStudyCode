package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值为=%v\nintChan本身的地址为%p\n", intChan, &intChan)

	//给管道写入数据
	intChan <- 10
	intChan <- 100
	num := 222

	intChan <- num
	fmt.Printf("管道长度，管道容量:%v,%v\n", len(intChan), cap(intChan)) //,3,3
	//给管道写入数据时，不能超过管道的容量
	//intChan <- 20 //报错

	//从管道中读取数据
	n2 := <-intChan
	fmt.Println("从管道取的第一个值为", n2)
	fmt.Printf("管道长度，管道容量:%v,%v\n", len(intChan), cap(intChan)) //,2,3
	n3 := <-intChan
	n4 := <-intChan
	fmt.Println("从管道取的第2，3个值为", n3, n4)
	n5 := <-intChan //报错 么有东西可以取了
	fmt.Println("从管道取的第5个值为", n5)

}
