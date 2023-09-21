package main

import (
	"fmt"
	"net" //所需的网络编程全部都在net包下
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		for {
			n, err3 := conn.Read(buf)
			if err3 != nil {
				return
			}
			//fmt.Println((buf[0:n])) //因为以字节数组（[]byte）形式存储并打印的，所以打印的是ascii码
			fmt.Println(string(buf[0:n]))
		}
	}
}
func main() {
	//打印：
	fmt.Println("服务器端启动了。。")
	//进行监听：需要指定服务器端TCP协议，服务器端的IP+PORT
	listen, err := net.Listen("tcp", "192.168.31.122:8888")
	if err != nil { //监听失败
		fmt.Println("监听失败，err:", err)
		return
	}
	//监听成功以后：
	//循环等待客户端的链接：
	for {
		conn, err2 := listen.Accept()
		if err2 != nil { //客户端的等待失败
			fmt.Println("客户端的等待失败,err2:", err2)
		} else {
			//连接成功：
			fmt.Printf("等待链接成功，con=%v ，接收到的客户端信息：%v \n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程，协程处理客户端服务请求：
		go process(conn) //不同的客户端的请求，连接conn不一样的
	}
}
