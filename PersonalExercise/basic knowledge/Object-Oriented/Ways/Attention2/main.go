package main

import "fmt"

type integer int

func (i *integer) print() {
	*i = 30
	fmt.Println("i= ", *i)
}
func (i1 integer) pirnt1() {
	i1 = 11
	fmt.Println("i1=", i1)
}

func main() {
	var i integer = 20
	var b integer = 12
	b.pirnt1()
	i.print()
	fmt.Println(i) //还是30，地址传递不改变值
	fmt.Println(b)
}
