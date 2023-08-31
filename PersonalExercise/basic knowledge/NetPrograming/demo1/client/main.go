package main

import (
	"bufio"
	"fmt"
	"net" //所需的网络编程全部都在net包下
	"os"
)

func main() {
	//打印：
	fmt.Println("客服端启动。。")
	//调用Dial函数：参数需要指定tcp协议，需要指定服务器端的IP+PORT
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	//Conn：表示与指定网络地址建立的连接，实现了io.Reader和io.Writer接口，可以进行读写操作。
	if err != nil { //连接失败
		fmt.Println("客户端连接失败：err:", err)
		return
	}
	fmt.Println("连接成功，conn:", conn)
	//通过客户端发送单行数据，然后退出：
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表终端标准输入
	//从终端读取一行用户输入的信息：
	str, err := reader.ReadString('\n')
	/*
		ReadString('\n') 方法会从读取器对象中读取字符，直到遇到换行符 \n 为止，并将读取到的字符以字符串的形式返回给变量 str。同时，它还会返回一个错误对象 err，用于表示读取过程中是否发生了错误。
		这个操作通常用于从终端或其他输入源中逐行读取用户输入的内容。
	*/
	if err != nil {
		fmt.Println("终端输入失败，err:", err)
	}
	//将str数据发送给服务器：
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败，err:", err)
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节的数据,并退出", n)
}
