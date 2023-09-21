package main

import (
	"fmt"
	"reflect"
)

//通过反射修改，num int的值
//修改student的值

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())

	rVal.Elem().SetInt(20)
}
func main() {
	num := 10
	reflect01(&num)
	fmt.Println("num=", num)
	/*如何理解rVal.Elem()
	num := 9
	ptr * int = &num
	*ptr=11
	类似于rVal.Elem()
	*/
}
