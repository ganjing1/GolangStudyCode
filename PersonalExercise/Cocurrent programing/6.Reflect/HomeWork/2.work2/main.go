package main

import (
	"fmt"
	"reflect"
)

func floatRef(f interface{}) {
	rVale := reflect.ValueOf(f)
	rType := reflect.TypeOf(f)
	fmt.Printf("f对应的Fvalue为%v，Type为%T,对应的Kind为%v\n", rVale, rType, rVale.Kind())
	iValue := rVale.Interface()

	fmt.Printf("iValue的Type为:%T\n", iValue)
	fValue, ok := iValue.(float64)
	if !ok {
		fmt.Println("类型断言失败")
		return
	}
	fmt.Printf("fValue的Type为:%T\n", fValue)
}

func main() {
	f := 19.9
	floatRef(f)
}
