package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.31.122:8888")
	if err != nil {
		fmt.Printf("conn have an err:=%v\n", err)
		return
	}
	reader := bufio.NewReader(os.Stdin) //创建一个终端读取器
	for {
		content, err := reader.ReadString('\n') //读取数据读到回车为止
		if err != nil {
			fmt.Printf("read have an err:=%v\n", err)
			return
		}
		content = strings.Trim(content, "\r\n")
		if content == "exit" {
			fmt.Println("客户端退出。。")
			break
		}
		_, err1 := conn.Write([]byte(content + "\n")) //向服务器传输字段，并返回n为字段的长度，err为异常
		if err != nil {
			fmt.Println("conn.write err=", err1)
		}
	}

}
