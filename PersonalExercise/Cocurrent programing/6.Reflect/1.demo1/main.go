package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取传入的变量的type,kind 值
	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	fmt.Printf("rVale=%v,rValue type=%T\n", rValue, rType)
	n2 := 2 + rValue.Float()
	fmt.Printf("n2=%v\n", n2)
	fmt.Println(rType, rValue)
	iV := rValue.Interface()
	fmt.Printf("iv=%T\n", iV)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b) //获取类别
	rValue := reflect.ValueOf(b)
	fmt.Printf("rType=%T,rValue=%v\n", rType, rValue)
	iV := rValue.Interface()
	fmt.Printf("iv=%v,type=%T\n", iV, iV)
	stu, ok := iV.(Student) //将IV断言程Student类型，判断是否能断言成功
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

	//获取类型(kind)
	rk := rValue.Kind()
	fmt.Println(rk)

}

type Student struct {
	Name string
	Age  int
}

func main() {
	//var num float64 = 99.9
	//reflectTest01(num)

	//Student的实例
	st := Student{Name: "cheerio", Age: 22}
	reflectTest02(st)
}
