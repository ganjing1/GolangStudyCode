package main

import "fmt"

type Person struct {
	name string
}

//这是方法和函数不一样
/*
Person 结构体有一个test()方法，PythonWork()和Person是绑定的
a为结构体Person调用的具体对象，用对象来调用test()
*/
func (a *Person) test() {
	(*a).name = "lisi"
	fmt.Println((*a).name)
}

func main() {
	var p Person
	p.name = "zhangsan"
	(&p).test()
	fmt.Println(p.name)
}
