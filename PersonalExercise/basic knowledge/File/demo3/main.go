package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\Traditional PythonWork\\Test.txt")
	if err != nil {
		fmt.Println("can't read")
	} else {
		fmt.Println("Ok you got it")
	}
	//当函数退出时，让file关闭，防止内存泄露：
	defer file.Close()

	//创建一个流：reader
	reader := bufio.NewReader(file)
	//读取操作
	for {
		str, err := reader.ReadString('\n') //每次读取到\n后结束一次读取
		if err == io.EOF {
			//读到文件结尾跳去循环
			break
		}
		//没有到结尾
		fmt.Println(str)
	}
}
