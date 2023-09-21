package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.253.158:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close() // 在发送完数据后关闭连接

	st := ""
	judge := ""
	reader := bufio.NewReader(os.Stdin)

	// 从终端读取用户输入，并发送给服务器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
			break
		}
		st += line
		fmt.Println("是否要继续输出，退出输入exit,否则按回车:")
		_, err = fmt.Scanln(judge)
		if err != nil {
			fmt.Println("Scanln err=", err)
			break
		}
		if judge == "exit" {
			break
		}
	}

	// 将line 发送给服务器
	n, err := conn.Write([]byte(st))
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	fmt.Printf("客户端发送了%d字节的数据，并退出\n", n)
	fmt.Println("conn success=", conn)
}
