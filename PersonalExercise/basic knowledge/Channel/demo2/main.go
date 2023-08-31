package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 2)
	//在管道中存放数据
	intChan <- 10
	intChan <- 20

	//关闭管道:关闭管道后不能再向管道写入数据
	close(intChan)
	//intChan <- 30//报错,关闭管道后不能再向管道写入数据

	//关闭管道后读取数据
	num := intChan    //num代表管道的地址
	num1 := <-intChan //取出管道中对应的数
	num2 := <-intChan
	num3 := <-intChan //管道关闭后读取不会报错，也可以越容量读取，但是结果为0
	fmt.Println(num, num1, num2, num3)
}
