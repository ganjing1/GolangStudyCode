package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\homework\\work1\\test1.txt")
	if err != nil {
		fmt.Println("创建失败")
	}
	for i := 0; i < 10; i++ {
		io.WriteString(file, "hello world\n")
	}
	defer file.Close()

	file2, err2 := os.Create("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\homework\\work1\\test2.txt")
	if err2 != nil {
		fmt.Println("创建失败")
	}
	defer file2.Close()
	content, err3 := ioutil.ReadFile("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\homework\\work1\\test1.txt")
	if err3 != nil {
		fmt.Println("读取错误")
	}
	_, err4 := file2.Write(content)
	if err4 != nil {
		fmt.Printf("写入错入\n", err4)

	}

}
