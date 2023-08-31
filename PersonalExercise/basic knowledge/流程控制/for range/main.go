package main

import "fmt"

func main() {
	var str string = "hello world你好"

	for i, value := range str {
		fmt.Printf("索引：%d,值：%c\n", i, value)
		//对str进行遍历，遍历的每个结果的索引值被i接收，每个值被value接收
		//对字符进行遍历

	}
}
