package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData", i)
	}
	close(intChan) //关闭了还可以继续读
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan //要从管道读出v表示值，ok表示是否读出成功
		if ok == false {
			break //把五十个数字读完了才退出
		}
		fmt.Printf("readData =%v\n", v) //读取成功一次打印一次
	} //readData 读取完数据后，向exitchan中写入一个true，并关闭exitchan管道
	exitChan <- true
	close(exitChan)
}
func main() {
	IntChan := make(chan int, 100)
	exitChan := make(chan bool, 1) //bool类型的管道长度为1
	go writeData(IntChan)          //读和写同时进行
	go readData(IntChan, exitChan)
	for {
		_, ok := <-exitChan //只要exitChan没有退出来就一直 读，否则就退出for的
		/*这是因为在Go语言中，从一个已经关闭的管道中接收数据，
		会返回该管道元素类型的零值和一个false标志，表示没有更多的数据了*/
		if !ok {
			break
		}
	}

}
