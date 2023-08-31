package main

import (
	"fmt"
)

func main() {
	//定义管道-> 定义int类型的管道
	var Channel chan int

	//通过make初始化:管道可以存放3个int类型的数据
	Channel = make(chan int, 3)

	//证明管道是引用类型:
	fmt.Printf("int型的Channel的值:%v", Channel) //输出管道的地址
	fmt.Println()
	//向管道存放数据:
	Channel <- 10 //向int管道存放的第一个数据
	num := 20
	Channel <- num //管道存放的第二个数据
	//不能存放大于容量的数据:
	Channel <- 30
	//Channel <- 40  总共有三个容量，这是第四个容量所以报错

	//输出管道的长度
	fmt.Printf("lenth of Channel:%d,capacity of Channel:%d\n", len(Channel), cap(Channel)) //len求实际长度，cap求容量

	//读取管道中的数据：
	num1 := <-Channel
	num2 := <-Channel
	num3 := <-Channel
	fmt.Println(num1, num2, num3)

	//如果在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错
	num4 := <-Channel //报错大于管道的容量，已经取取完了
	fmt.Println(num4)

}
