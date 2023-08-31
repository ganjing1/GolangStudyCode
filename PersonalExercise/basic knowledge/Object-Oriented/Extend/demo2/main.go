package main

import "fmt"

type A struct {
	a int
	b string
}

type B struct {
	c int
	d string
	a int
}

type C struct {
	*A
	*B
	int
}

type D struct {
	a int //结构体的名称属于一种变量类型，让a为C结构体型
	b string
	c B
}

func main() {

	c := C{&A{a: 10, b: "zhangsan"}, &B{c: 20, d: "lisi", a: 50}, 101}
	fmt.Println(c)
	fmt.Println(*c.A)
	fmt.Println(*c.B)
	d := D{10, "ganjing", B{1, "wangwu", 2}}
	fmt.Println(d)
	fmt.Println(d.c.d)
}
