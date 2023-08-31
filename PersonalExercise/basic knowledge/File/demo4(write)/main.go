package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Openfile(路径,打开方式,权限对应的数字)
	file, err := os.OpenFile("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\demo4(write)\\test.txt",
		os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	//perm：文件权限，只有在创建新文件时才有效，可以使用 os.FileMode
	//类型的常量来指定文件权限。如果不想设置文件权限，可以使用 0。
	if err != nil { //文件打开失败
		fmt.Printf("打开文件失败", err)
		return
	}
	//及时将文件关闭：
	defer file.Close() //main函数最后一步执行

	//写入文件操作：---》IO流---》缓冲输出流(带缓冲区)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("你好 马士兵\n") //在write所指向的文件进行写入字符串操作
	}
	//流带缓冲区，刷新数据--->真正写入文件中：
	writer.Flush() //从缓冲区移到对应的文件
	s := os.FileMode(0666).String()
	fmt.Println(s) //打印s对应的权限
	/*
		0666 表示所有人都有读写权限，但没有执行权限
		0777：所有人都有读、写、执行权限；
		0755：所有人都有读、执行权限，只有所有者有写权限；
		0644：所有人都有读权限，只有所有者有写权限。
	*/
}
