package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//源文件
	text1path := "F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\demo5(move_file)\\text1.txt"
	//目标文件
	text2path := "F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\demo5(move_file)\\text2.txt"
	content1, err1 := ioutil.ReadFile(text1path)
	if err1 != nil {
		fmt.Println("cant open text1.txt")
		return
	}

	err1 = os.WriteFile(text2path, content1, 0666) //将第一个文件的内容写到第二个文件
	if err1 != nil {
		fmt.Println("Fail write")
	}

}
