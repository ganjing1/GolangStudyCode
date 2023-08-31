package main

import "fmt"

//判断输入的参数类型的函数

// 函数参数里加... 表示参数个数是可变的
func TypeJudge(items ...interface{}) {
	for key, value := range items {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型。值是%v\n", key, value)
		case int:
			fmt.Printf("第%v个参数是int类型。值是%v\n", key, value)
		case string:
			fmt.Printf("第%v个参数是strint类型。值是%v\n", key, value)
		case float32:
			fmt.Printf("第%v个参数是float32类型。值是%v\n", key, value)
		case float64:
			fmt.Printf("第%v个参数是float64类型。值是%v\n", key, value)
		default:
			fmt.Printf("第%v个参数类型不确定类型。值是%v\n", key, value)
		}
	}

}

func main() {
	var a float32 = 15.1
	var b float64 = 15.2
	var c string = "helloworld"
	var d int = 11
	e := "192.168.1.1"
	f := 15
	TypeJudge(a, b, c, d, e, f)
}
