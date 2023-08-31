package main

import "fmt"

// 接口继承
type Binterface interface {
	b()
}
type Cinterface interface {
	c()
}

type Ainterface interface {
	Binterface
	Cinterface
	a()
}

type V interface { //空接口
}

type Stu struct {
}

func (s Stu) a() {
	fmt.Println("aaaaa")
}
func (s Stu) b() {
	fmt.Println("bbbbb")
}
func (c Stu) c() {
	fmt.Println("ccccc")
}
func main() {
	var s Stu
	var a Ainterface = s
	a.c()
	a.c()
	a.c()
	var v V = s
	fmt.Println(v)

	var v2 interface{} = s //第二种定义空接口
	fmt.Println(v2)
	var num float32 = 9.3
	var v3 interface{} = num
	fmt.Println(v3)
}
