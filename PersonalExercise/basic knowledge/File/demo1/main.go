package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\Traditional test\\Test.txt")
	if err != nil { //如果有错误说明错误不为空
		fmt.Println("have err:", err)
	} else {
		fmt.Println("right!!")
	}

	//返回的file是指针类型所以输出的是地址
	fmt.Printf("FILE:=%v", file)
	Fc := file.Close()
	if Fc != nil { //如果不能关闭成功就是file.close()的返回值不为0
		fmt.Println("can't close")
	} else {
		fmt.Println("achieve close")
	}
}
