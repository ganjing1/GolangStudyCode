package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//读取文件
	in, err := ioutil.ReadFile("F:\\Users\\gj\\gostudy\\PersonalExercise\\File\\Traditional PythonWork\\Test.txt")
	//返回一个[]byte,err切片
	if err != nil {
		fmt.Println("fail read err is:", err)
	} else {
		//fmt.Printf("%v", in) //输出的是对应的asci码
		fmt.Printf("%v", string(in)) //将asci码切片以字符串的形式输出
	}

}
