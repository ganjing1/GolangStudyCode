package main

import "fmt"

type Cat struct {
	Name   string
	Age    int
	Weight float64
}

func main() {

	catChan := make(chan interface{}, 10)
	cat1 := "haha"
	cat2 := Cat{Name: "Hitty", Age: 20, Weight: 12.1}
	catChan <- cat1
	catChan <- cat2
	cat3 := <-catChan
	cat4 := <-catChan
	fmt.Println(cat3, cat4)
	/*
		fmt.Printf("cat4.name=%v", cat4.Name)是错的，编译不能通过
		cat4是一个interface{}类型的变量，
		因为cat4是一个interface{}类型的变量，它不知道自己的底层类型是什么，
		也不知道自己有哪些字段或方法。
		interface{}类型是一个空接口，它可以表示任何类型的值。
		但是，空接口没有任何方法，所以它不能调用任何方法。
	*/cat5 := cat4.(Cat)
	fmt.Printf("cat5.name=%v", cat5.Name)
}
