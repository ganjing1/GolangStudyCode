package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {
	reValue := reflect.ValueOf(i)
	//让reValue调用Elem()下的.SetInt(数据)函数来改变revalue的值
	reValue.Elem().SetInt(40)
}
func main() {
	A := 100
	testReflect(&A)
	fmt.Println(A)
}
