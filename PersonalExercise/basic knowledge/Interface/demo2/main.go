package main

import "fmt"

type Ainterface interface {
	A()
}
type Binterface interface {
	B()
}

type student struct {
}

func (s student) A() {
	fmt.Println("AAAAAAAA")
}

func (s student) B() {
	fmt.Println("BBBBBBBB")
}

func main() {

	var s student
	var a Ainterface = s
	var b Binterface = s
	a.A()
	b.B()
}