package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口：
func testReflect(i interface{}) { //空接口没有任何方法,所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口。
	//1.调用TypeOf函数，返回reflect.Type类型数据：
	reType := reflect.TypeOf(i)
	fmt.Println("reType:", reType)
	fmt.Printf("reType的具体类型是：%T", reType) //*reflect.rtype

	//2.调用ValueOf函数，返回reflect.Value类型数据：
	reValue := reflect.ValueOf(i)
	fmt.Println("\nreValue:", reValue)
	fmt.Printf("reValue的具体类型是：%T\n", reValue)
	//num1 := 100
	//如果真想获取reValue的数值，要调用Int()方法：返回v持有的有符号整数
	num2 := 80 + reValue.Int()
	fmt.Println(num2)
	//reValue转成空接口：
	i2 := reValue.Interface()
	//类型断言：
	n := i2.(int) //一个接口类型的值 i2 转换为具体的 int 类型，并将结果赋值给变量 n
	n2 := n + 30
	fmt.Println(n2)
}
func main() {
	//对基本数据类型进行反射：
	//定义一个基本数据类型：
	var num int = 100
	testReflect(num)
}
