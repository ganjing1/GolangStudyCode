package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file1path := "F:\\Users\\gj\\gostudy\\PersonalExercise\\basic knowledge\\File\\homework\\work2\\abc.txt"
	file2path := "F:\\Users\\gj\\gostudy\\PersonalExercise\\basic knowledge\\File\\homework\\work2\\kkk.txt"

	//将abc.txt的内容导入kkk.txt
	abccontent, err := ioutil.ReadFile(file1path)
	if err != nil {
		fmt.Printf("读取错误：%v", err)
		return
	}
	err = os.WriteFile(file2path, abccontent, 0666)
	if err != nil {
		fmt.Printf("写入错误:%v", err)
	}

}
