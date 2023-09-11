package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 函数会返回两个值，一个是int64类型的written，一个是error类型的err
func Copyfile(dstFilename string, srcFilename string) (written int64, err error) {
	srcfile, err := os.Open(srcFilename)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer srcfile.Close()
	//通过srcfile,获取到Reader
	reader := bufio.NewReader(srcfile) //NewReader创建一个带缓冲的读取器
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFilename, os.O_WRONLY|os.O_CREATE, 066) //WRONLY是write only，只写
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//dstFile,获取writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader) //就reader拷贝到writer
}

func main() {
	srcFile := "F:\\Users\\gj\\gostudy\\PersonalExercise\\basic knowledge\\File\\homework\\work3\\girl.png"
	dstFile := "F:\\Users\\gj\\gostudy\\PersonalExercise\\basic knowledge\\File\\homework\\work3\\girl1.png"
	_, err := Copyfile(dstFile, srcFile) //自己定义的Copyfile函数调用后会返回两个参数,1.拷贝字节的个数，2.拷贝是否出错
	if err != nil {
		fmt.Printf("拷贝错误，err=%v", err)
		return
	} else {
		fmt.Println("拷贝完成")
	}
	fmt.Println("复制完成")
}
